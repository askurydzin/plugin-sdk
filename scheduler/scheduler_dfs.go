package scheduler

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"

	"github.com/cloudquery/plugin-sdk/v4/helpers"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/getsentry/sentry-go"
	"golang.org/x/sync/semaphore"
)

func (s *Scheduler) syncDfs(ctx context.Context, resolvedResources chan<- *schema.Resource, syncOpts *SyncOptions) {
	// This is very similar to the concurrent web crawler problem with some minor changes.
	// We are using DFS to make sure memory usage is capped at O(h) where h is the height of the tree.
	tableConcurrency := max(s.concurrency/minResourceConcurrency, minTableConcurrency)
	resourceConcurrency := tableConcurrency * minResourceConcurrency

	s.tableSems = make([]*semaphore.Weighted, s.maxDepth)
	for i := uint64(0); i < s.maxDepth; i++ {
		s.tableSems[i] = semaphore.NewWeighted(int64(tableConcurrency))
		// reduce table concurrency logarithmically for every depth level
		tableConcurrency = max(tableConcurrency/2, minTableConcurrency)
	}
	s.resourceSem = semaphore.NewWeighted(int64(resourceConcurrency))

	// we have this because plugins can return sometimes clients in a random way which will cause
	// differences between this run and the next one.
	preInitialisedClients := make([][]schema.ClientMeta, len(s.tables))
	for i, table := range s.tables {
		clients := []schema.ClientMeta{s.client}
		if table.Multiplex != nil {
			clients = table.Multiplex(s.client)
		}
		// Detect duplicate clients while multiplexing
		seenClients := make(map[string]bool)
		for _, c := range clients {
			if _, ok := seenClients[c.ID()]; !ok {
				seenClients[c.ID()] = true
			} else {
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("table", table.Name)
					sentry.CurrentHub().CaptureMessage("duplicate client ID in " + table.Name)
				})
				s.logger.Warn().Str("client", c.ID()).Str("table", table.Name).Msg("multiplex returned duplicate client")
			}
		}
		preInitialisedClients[i] = clients
		// we do this here to avoid locks so we initial the metrics structure once in the main goroutines
		// and then we can just read from it in the other goroutines concurrently given we are not writing to it.
		s.metrics.initWithClients(table, clients)
	}

	var wg sync.WaitGroup
	for i, table := range s.tables {
		table := table
		clients := preInitialisedClients[i]
		for _, client := range clients {
			client := client
			if err := s.tableSems[0].Acquire(ctx, 1); err != nil {
				// This means context was cancelled
				wg.Wait()
				return
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer s.tableSems[0].Release(1)
				// not checking for error here as nothing much todo.
				// the error is logged and this happens when context is cancelled
				s.resolveTableDfs(ctx, table, client, nil, resolvedResources, 1, syncOpts)
			}()
		}
	}

	// Wait for all the worker goroutines to finish
	wg.Wait()
}

func (s *Scheduler) resolveTableDfs(ctx context.Context, table *schema.Table, client schema.ClientMeta, parent *schema.Resource, resolvedResources chan<- *schema.Resource, depth int, syncOpts *SyncOptions) {
	var validationErr *schema.ValidationError
	clientName := client.ID()
	logger := s.logger.With().Str("table", table.Name).Str("client", clientName).Logger()

	if parent == nil { // Log only for root tables, otherwise we spam too much.
		logger.Info().Msg("top level table resolver started")
	}
	tableMetrics := s.metrics.TableClient[table.Name][clientName]

	res := make(chan any)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := fmt.Sprintf("%s\n%s", err, string(debug.Stack()))
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("table", table.Name)
					sentry.CurrentHub().CaptureMessage(stack)
				})
				logger.Error().Interface("error", err).Str("stack", stack).Msg("table resolver finished with panic")
				atomic.AddUint64(&tableMetrics.Panics, 1)
			}
			close(res)
		}()
		if err := table.Resolver(ctx, client, parent, res); err != nil {
			logger.Error().Err(err).Msg("table resolver finished with error")
			atomic.AddUint64(&tableMetrics.Errors, 1)
			if errors.As(err, &validationErr) {
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("table", table.Name)
					sentry.CurrentHub().CaptureMessage(validationErr.MaskedError())
				})
			}
			return
		}
	}()

	for r := range res {
		s.resolveResourcesDfs(ctx, table, client, parent, r, resolvedResources, depth, syncOpts)
	}

	// we don't need any waitgroups here because we are waiting for the channel to close
	if parent == nil { // Log only for root tables and relations only after resolving is done, otherwise we spam per object instead of per table.
		logger.Info().Uint64("resources", tableMetrics.Resources).Uint64("errors", tableMetrics.Errors).Msg("table sync finished")
		s.logTablesMetrics(table.Relations, client)
	}
}

func (s *Scheduler) resolveResourcesDfs(ctx context.Context, table *schema.Table, client schema.ClientMeta, parent *schema.Resource, resources any, resolvedResources chan<- *schema.Resource, depth int, syncOpts *SyncOptions) {
	resourcesSlice := helpers.InterfaceSlice(resources)
	if len(resourcesSlice) == 0 {
		return
	}
	resourcesChan := make(chan *schema.Resource, len(resourcesSlice))
	go func() {
		defer close(resourcesChan)
		var wg sync.WaitGroup
		sentValidationErrors := sync.Map{}
		for i := range resourcesSlice {
			i := i
			if err := s.resourceSem.Acquire(ctx, 1); err != nil {
				s.logger.Warn().Err(err).Msg("failed to acquire semaphore. context cancelled")
				wg.Wait()
				// we have to continue emptying the channel to exit gracefully
				return
			}
			wg.Add(1)
			go func() {
				defer s.resourceSem.Release(1)
				defer wg.Done()
				//nolint:all
				resolvedResource := s.resolveResource(ctx, table, client, parent, resourcesSlice[i])
				if resolvedResource == nil {
					return
				}

				if err := resolvedResource.CalculateCQID(syncOpts.DeterministicCQID); err != nil {
					tableMetrics := s.metrics.TableClient[table.Name][client.ID()]
					s.logger.Error().Err(err).Str("table", table.Name).Str("client", client.ID()).Msg("resource resolver finished with primary key calculation error")
					if _, found := sentValidationErrors.LoadOrStore(table.Name, struct{}{}); !found {
						// send resource validation errors to Sentry only once per table,
						// to avoid sending too many duplicate messages
						sentry.WithScope(func(scope *sentry.Scope) {
							scope.SetTag("table", table.Name)
							sentry.CurrentHub().CaptureMessage(err.Error())
						})
					}
					atomic.AddUint64(&tableMetrics.Errors, 1)
					return
				}
				if err := resolvedResource.Validate(); err != nil {
					tableMetrics := s.metrics.TableClient[table.Name][client.ID()]
					s.logger.Error().Err(err).Str("table", table.Name).Str("client", client.ID()).Msg("resource resolver finished with validation error")
					if _, found := sentValidationErrors.LoadOrStore(table.Name, struct{}{}); !found {
						// send resource validation errors to Sentry only once per table,
						// to avoid sending too many duplicate messages
						sentry.WithScope(func(scope *sentry.Scope) {
							scope.SetTag("table", table.Name)
							sentry.CurrentHub().CaptureMessage(err.Error())
						})
					}
					atomic.AddUint64(&tableMetrics.Errors, 1)
					return
				}
				resourcesChan <- resolvedResource
			}()
		}
		wg.Wait()
	}()

	var wg sync.WaitGroup
	for resource := range resourcesChan {
		resource := resource
		resolvedResources <- resource
		for _, relation := range resource.Table.Relations {
			relation := relation
			if err := s.tableSems[depth].Acquire(ctx, 1); err != nil {
				// This means context was cancelled
				wg.Wait()
				return
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer s.tableSems[depth].Release(1)
				s.resolveTableDfs(ctx, relation, client, resource, resolvedResources, depth+1, syncOpts)
			}()
		}
	}
	wg.Wait()
}