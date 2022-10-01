// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: internal/pb/source.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Sync) Reset() {
	*x = Sync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sync) ProtoMessage() {}

func (x *Sync) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sync.ProtoReflect.Descriptor instead.
func (*Sync) Descriptor() ([]byte, []int) {
	return file_internal_pb_source_proto_rawDescGZIP(), []int{0}
}

type GetTables struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTables) Reset() {
	*x = GetTables{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTables) ProtoMessage() {}

func (x *GetTables) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTables.ProtoReflect.Descriptor instead.
func (*GetTables) Descriptor() ([]byte, []int) {
	return file_internal_pb_source_proto_rawDescGZIP(), []int{1}
}

type Sync_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec []byte `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Sync_Request) Reset() {
	*x = Sync_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sync_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sync_Request) ProtoMessage() {}

func (x *Sync_Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sync_Request.ProtoReflect.Descriptor instead.
func (*Sync_Request) Descriptor() ([]byte, []int) {
	return file_internal_pb_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Sync_Request) GetSpec() []byte {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Sync_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// marshalled *schema.Resources
	Resource []byte `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Sync_Response) Reset() {
	*x = Sync_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sync_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sync_Response) ProtoMessage() {}

func (x *Sync_Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sync_Response.ProtoReflect.Descriptor instead.
func (*Sync_Response) Descriptor() ([]byte, []int) {
	return file_internal_pb_source_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Sync_Response) GetResource() []byte {
	if x != nil {
		return x.Resource
	}
	return nil
}

type GetTables_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTables_Request) Reset() {
	*x = GetTables_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTables_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTables_Request) ProtoMessage() {}

func (x *GetTables_Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTables_Request.ProtoReflect.Descriptor instead.
func (*GetTables_Request) Descriptor() ([]byte, []int) {
	return file_internal_pb_source_proto_rawDescGZIP(), []int{1, 0}
}

type GetTables_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Marshalled []*schema.Table
	Tables []byte `protobuf:"bytes,3,opt,name=tables,proto3" json:"tables,omitempty"`
}

func (x *GetTables_Response) Reset() {
	*x = GetTables_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTables_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTables_Response) ProtoMessage() {}

func (x *GetTables_Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTables_Response.ProtoReflect.Descriptor instead.
func (*GetTables_Response) Descriptor() ([]byte, []int) {
	return file_internal_pb_source_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetTables_Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTables_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetTables_Response) GetTables() []byte {
	if x != nil {
		return x.Tables
	}
	return nil
}

var File_internal_pb_source_proto protoreflect.FileDescriptor

var file_internal_pb_source_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x04, 0x53, 0x79, 0x6e,
	0x63, 0x1a, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x1a, 0x26, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x68, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x50, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x32, 0x80, 0x02, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_source_proto_rawDescOnce sync.Once
	file_internal_pb_source_proto_rawDescData = file_internal_pb_source_proto_rawDesc
)

func file_internal_pb_source_proto_rawDescGZIP() []byte {
	file_internal_pb_source_proto_rawDescOnce.Do(func() {
		file_internal_pb_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_source_proto_rawDescData)
	})
	return file_internal_pb_source_proto_rawDescData
}

var file_internal_pb_source_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_pb_source_proto_goTypes = []interface{}{
	(*Sync)(nil),                // 0: proto.Sync
	(*GetTables)(nil),           // 1: proto.GetTables
	(*Sync_Request)(nil),        // 2: proto.Sync.Request
	(*Sync_Response)(nil),       // 3: proto.Sync.Response
	(*GetTables_Request)(nil),   // 4: proto.GetTables.Request
	(*GetTables_Response)(nil),  // 5: proto.GetTables.Response
	(*GetName_Request)(nil),     // 6: proto.GetName.Request
	(*GetVersion_Request)(nil),  // 7: proto.GetVersion.Request
	(*GetName_Response)(nil),    // 8: proto.GetName.Response
	(*GetVersion_Response)(nil), // 9: proto.GetVersion.Response
}
var file_internal_pb_source_proto_depIdxs = []int32{
	6, // 0: proto.Source.GetName:input_type -> proto.GetName.Request
	7, // 1: proto.Source.GetVersion:input_type -> proto.GetVersion.Request
	4, // 2: proto.Source.GetTables:input_type -> proto.GetTables.Request
	2, // 3: proto.Source.Sync:input_type -> proto.Sync.Request
	8, // 4: proto.Source.GetName:output_type -> proto.GetName.Response
	9, // 5: proto.Source.GetVersion:output_type -> proto.GetVersion.Response
	5, // 6: proto.Source.GetTables:output_type -> proto.GetTables.Response
	3, // 7: proto.Source.Sync:output_type -> proto.Sync.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pb_source_proto_init() }
func file_internal_pb_source_proto_init() {
	if File_internal_pb_source_proto != nil {
		return
	}
	file_internal_pb_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sync); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTables); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sync_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sync_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTables_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pb_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTables_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pb_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pb_source_proto_goTypes,
		DependencyIndexes: file_internal_pb_source_proto_depIdxs,
		MessageInfos:      file_internal_pb_source_proto_msgTypes,
	}.Build()
	File_internal_pb_source_proto = out.File
	file_internal_pb_source_proto_rawDesc = nil
	file_internal_pb_source_proto_goTypes = nil
	file_internal_pb_source_proto_depIdxs = nil
}
