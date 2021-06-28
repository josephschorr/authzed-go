// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: authzed/api/v1alpha1/schema.proto

package v1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ReadSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of names of the Object Definitions that are being requested.
	//
	// These names must be fully qualified with their namespace (e.g.
	// myblog/post).
	ObjectDefinitionNames []string `protobuf:"bytes,1,rep,name=object_definition_names,json=objectDefinitionNames,proto3" json:"object_definition_names,omitempty"`
}

func (x *ReadSchemaRequest) Reset() {
	*x = ReadSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSchemaRequest) ProtoMessage() {}

func (x *ReadSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSchemaRequest.ProtoReflect.Descriptor instead.
func (*ReadSchemaRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *ReadSchemaRequest) GetObjectDefinitionNames() []string {
	if x != nil {
		return x.ObjectDefinitionNames
	}
	return nil
}

type ReadSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Object Definitions that were requested.
	ObjectDefinitions []string `protobuf:"bytes,1,rep,name=object_definitions,json=objectDefinitions,proto3" json:"object_definitions,omitempty"`
}

func (x *ReadSchemaResponse) Reset() {
	*x = ReadSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSchemaResponse) ProtoMessage() {}

func (x *ReadSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSchemaResponse.ProtoReflect.Descriptor instead.
func (*ReadSchemaResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *ReadSchemaResponse) GetObjectDefinitions() []string {
	if x != nil {
		return x.ObjectDefinitions
	}
	return nil
}

type WriteSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Object Definitions that will be written to the permission system.
	ObjectDefinitions []string `protobuf:"bytes,2,rep,name=object_definitions,json=objectDefinitions,proto3" json:"object_definitions,omitempty"`
}

func (x *WriteSchemaRequest) Reset() {
	*x = WriteSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSchemaRequest) ProtoMessage() {}

func (x *WriteSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSchemaRequest.ProtoReflect.Descriptor instead.
func (*WriteSchemaRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_schema_proto_rawDescGZIP(), []int{2}
}

func (x *WriteSchemaRequest) GetObjectDefinitions() []string {
	if x != nil {
		return x.ObjectDefinitions
	}
	return nil
}

type WriteSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteSchemaResponse) Reset() {
	*x = WriteSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSchemaResponse) ProtoMessage() {}

func (x *WriteSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1alpha1_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSchemaResponse.ProtoReflect.Descriptor instead.
func (*WriteSchemaResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1alpha1_schema_proto_rawDescGZIP(), []int{3}
}

var File_authzed_api_v1alpha1_schema_proto protoreflect.FileDescriptor

var file_authzed_api_v1alpha1_schema_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x85, 0x01, 0x0a, 0x17, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x4d, 0xfa, 0x42, 0x4a, 0x92,
	0x01, 0x47, 0x22, 0x45, 0x72, 0x43, 0x28, 0x80, 0x01, 0x32, 0x3e, 0x5e, 0x28, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2f, 0x29, 0x3f, 0x5b, 0x61, 0x2d, 0x7a,
	0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x15, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0x43, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x11, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x54, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x12, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92, 0x01, 0x09, 0x08,
	0x01, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x11, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xd8, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x28, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_authzed_api_v1alpha1_schema_proto_rawDescOnce sync.Once
	file_authzed_api_v1alpha1_schema_proto_rawDescData = file_authzed_api_v1alpha1_schema_proto_rawDesc
)

func file_authzed_api_v1alpha1_schema_proto_rawDescGZIP() []byte {
	file_authzed_api_v1alpha1_schema_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1alpha1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1alpha1_schema_proto_rawDescData)
	})
	return file_authzed_api_v1alpha1_schema_proto_rawDescData
}

var file_authzed_api_v1alpha1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_v1alpha1_schema_proto_goTypes = []interface{}{
	(*ReadSchemaRequest)(nil),   // 0: authzed.api.v1alpha1.ReadSchemaRequest
	(*ReadSchemaResponse)(nil),  // 1: authzed.api.v1alpha1.ReadSchemaResponse
	(*WriteSchemaRequest)(nil),  // 2: authzed.api.v1alpha1.WriteSchemaRequest
	(*WriteSchemaResponse)(nil), // 3: authzed.api.v1alpha1.WriteSchemaResponse
}
var file_authzed_api_v1alpha1_schema_proto_depIdxs = []int32{
	0, // 0: authzed.api.v1alpha1.SchemaService.ReadSchema:input_type -> authzed.api.v1alpha1.ReadSchemaRequest
	2, // 1: authzed.api.v1alpha1.SchemaService.WriteSchema:input_type -> authzed.api.v1alpha1.WriteSchemaRequest
	1, // 2: authzed.api.v1alpha1.SchemaService.ReadSchema:output_type -> authzed.api.v1alpha1.ReadSchemaResponse
	3, // 3: authzed.api.v1alpha1.SchemaService.WriteSchema:output_type -> authzed.api.v1alpha1.WriteSchemaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authzed_api_v1alpha1_schema_proto_init() }
func file_authzed_api_v1alpha1_schema_proto_init() {
	if File_authzed_api_v1alpha1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v1alpha1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSchemaRequest); i {
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
		file_authzed_api_v1alpha1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSchemaResponse); i {
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
		file_authzed_api_v1alpha1_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSchemaRequest); i {
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
		file_authzed_api_v1alpha1_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSchemaResponse); i {
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
			RawDescriptor: file_authzed_api_v1alpha1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authzed_api_v1alpha1_schema_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1alpha1_schema_proto_depIdxs,
		MessageInfos:      file_authzed_api_v1alpha1_schema_proto_msgTypes,
	}.Build()
	File_authzed_api_v1alpha1_schema_proto = out.File
	file_authzed_api_v1alpha1_schema_proto_rawDesc = nil
	file_authzed_api_v1alpha1_schema_proto_goTypes = nil
	file_authzed_api_v1alpha1_schema_proto_depIdxs = nil
}
