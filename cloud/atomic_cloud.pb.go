// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: atomic_cloud.proto

package cloud

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CallResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CallResult) Reset() {
	*x = CallResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_atomic_cloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResult) ProtoMessage() {}

func (x *CallResult) ProtoReflect() protoreflect.Message {
	mi := &file_atomic_cloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResult.ProtoReflect.Descriptor instead.
func (*CallResult) Descriptor() ([]byte, []int) {
	return file_atomic_cloud_proto_rawDescGZIP(), []int{0}
}

func (x *CallResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CallResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CallResult) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_atomic_cloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_atomic_cloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_atomic_cloud_proto_rawDescGZIP(), []int{1}
}

func (x *Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Service) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_atomic_cloud_proto protoreflect.FileDescriptor

var file_atomic_cloud_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0a, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x2d, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x32, 0x62, 0x0a, 0x12, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x12, 0x0e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x01, 0x2a, 0x42, 0x20, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x0b, 0x41,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x00, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_atomic_cloud_proto_rawDescOnce sync.Once
	file_atomic_cloud_proto_rawDescData = file_atomic_cloud_proto_rawDesc
)

func file_atomic_cloud_proto_rawDescGZIP() []byte {
	file_atomic_cloud_proto_rawDescOnce.Do(func() {
		file_atomic_cloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_atomic_cloud_proto_rawDescData)
	})
	return file_atomic_cloud_proto_rawDescData
}

var file_atomic_cloud_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_atomic_cloud_proto_goTypes = []interface{}{
	(*CallResult)(nil), // 0: cloud.CallResult
	(*Service)(nil),    // 1: cloud.Service
}
var file_atomic_cloud_proto_depIdxs = []int32{
	1, // 0: cloud.AtomicCloudService.CallCloud:input_type -> cloud.Service
	0, // 1: cloud.AtomicCloudService.CallCloud:output_type -> cloud.CallResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_atomic_cloud_proto_init() }
func file_atomic_cloud_proto_init() {
	if File_atomic_cloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_atomic_cloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResult); i {
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
		file_atomic_cloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
			RawDescriptor: file_atomic_cloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_atomic_cloud_proto_goTypes,
		DependencyIndexes: file_atomic_cloud_proto_depIdxs,
		MessageInfos:      file_atomic_cloud_proto_msgTypes,
	}.Build()
	File_atomic_cloud_proto = out.File
	file_atomic_cloud_proto_rawDesc = nil
	file_atomic_cloud_proto_goTypes = nil
	file_atomic_cloud_proto_depIdxs = nil
}
