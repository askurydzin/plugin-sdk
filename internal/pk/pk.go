package pk

import (
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func String(table *schema.Table, resource []any) string {
	parts := make([]string, 0, len(table.PrimaryKeys()))
	for i, col := range table.Columns {
		if !col.CreationOptions.PrimaryKey {
			continue
		}
		parts = append(parts, fmt.Sprint(resource[i]))
	}

	return "(" + strings.Join(parts, ",") + ")"
}