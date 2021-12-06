// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: envoy/api/v2/core/backoff.proto

package core

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Configuration defining a jittered exponential back off strategy.
type BackoffStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base interval to be used for the next back off computation. It should
	// be greater than zero and less than or equal to :ref:`max_interval
	// <envoy_api_field_core.BackoffStrategy.max_interval>`.
	BaseInterval *duration.Duration `protobuf:"bytes,1,opt,name=base_interval,json=baseInterval,proto3" json:"base_interval,omitempty"`
	// Specifies the maximum interval between retries. This parameter is optional,
	// but must be greater than or equal to the :ref:`base_interval
	// <envoy_api_field_core.BackoffStrategy.base_interval>` if set. The default
	// is 10 times the :ref:`base_interval
	// <envoy_api_field_core.BackoffStrategy.base_interval>`.
	MaxInterval *duration.Duration `protobuf:"bytes,2,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
}

func (x *BackoffStrategy) Reset() {
	*x = BackoffStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_backoff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackoffStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackoffStrategy) ProtoMessage() {}

func (x *BackoffStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_backoff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackoffStrategy.ProtoReflect.Descriptor instead.
func (*BackoffStrategy) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_backoff_proto_rawDescGZIP(), []int{0}
}

func (x *BackoffStrategy) GetBaseInterval() *duration.Duration {
	if x != nil {
		return x.BaseInterval
	}
	return nil
}

func (x *BackoffStrategy) GetMaxInterval() *duration.Duration {
	if x != nil {
		return x.MaxInterval
	}
	return nil
}

var File_envoy_api_v2_core_backoff_proto protoreflect.FileDescriptor

var file_envoy_api_v2_core_backoff_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a,
	0x0f, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x12, 0x4e, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0xaa, 0x01, 0x08, 0x08, 0x01, 0x32, 0x04, 0x10, 0xc0,
	0x84, 0x3d, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x46, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x0b, 0x6d, 0x61, 0x78,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x8f, 0x01, 0x0a, 0x1f, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0c, 0x42, 0x61,
	0x63, 0x6b, 0x6f, 0x66, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x16, 0x12, 0x14, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_api_v2_core_backoff_proto_rawDescOnce sync.Once
	file_envoy_api_v2_core_backoff_proto_rawDescData = file_envoy_api_v2_core_backoff_proto_rawDesc
)

func file_envoy_api_v2_core_backoff_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_core_backoff_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_core_backoff_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_core_backoff_proto_rawDescData)
	})
	return file_envoy_api_v2_core_backoff_proto_rawDescData
}

var file_envoy_api_v2_core_backoff_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_api_v2_core_backoff_proto_goTypes = []interface{}{
	(*BackoffStrategy)(nil),   // 0: envoy.api.v2.core.BackoffStrategy
	(*duration.Duration)(nil), // 1: google.protobuf.Duration
}
var file_envoy_api_v2_core_backoff_proto_depIdxs = []int32{
	1, // 0: envoy.api.v2.core.BackoffStrategy.base_interval:type_name -> google.protobuf.Duration
	1, // 1: envoy.api.v2.core.BackoffStrategy.max_interval:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_core_backoff_proto_init() }
func file_envoy_api_v2_core_backoff_proto_init() {
	if File_envoy_api_v2_core_backoff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_core_backoff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackoffStrategy); i {
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
			RawDescriptor: file_envoy_api_v2_core_backoff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_core_backoff_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_core_backoff_proto_depIdxs,
		MessageInfos:      file_envoy_api_v2_core_backoff_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_core_backoff_proto = out.File
	file_envoy_api_v2_core_backoff_proto_rawDesc = nil
	file_envoy_api_v2_core_backoff_proto_goTypes = nil
	file_envoy_api_v2_core_backoff_proto_depIdxs = nil
}
