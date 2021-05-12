// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: pkg/ingress/get-instance.proto

package ingress

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetWorkflowInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *GetWorkflowInstanceRequest) Reset() {
	*x = GetWorkflowInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkflowInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowInstanceRequest) ProtoMessage() {}

func (x *GetWorkflowInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowInstanceRequest.ProtoReflect.Descriptor instead.
func (*GetWorkflowInstanceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_instance_proto_rawDescGZIP(), []int{0}
}

func (x *GetWorkflowInstanceRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type GetWorkflowInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *string              `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Status       *string              `protobuf:"bytes,2,opt,name=status,proto3,oneof" json:"status,omitempty"`
	InvokedBy    *string              `protobuf:"bytes,3,opt,name=invokedBy,proto3,oneof" json:"invokedBy,omitempty"`
	Revision     *int32               `protobuf:"varint,4,opt,name=revision,proto3,oneof" json:"revision,omitempty"`
	BeginTime    *timestamp.Timestamp `protobuf:"bytes,5,opt,name=beginTime,proto3,oneof" json:"beginTime,omitempty"`
	EndTime      *timestamp.Timestamp `protobuf:"bytes,6,opt,name=endTime,proto3,oneof" json:"endTime,omitempty"`
	Flow         []string             `protobuf:"bytes,7,rep,name=flow,proto3" json:"flow,omitempty"`
	Input        []byte               `protobuf:"bytes,8,opt,name=input,proto3,oneof" json:"input,omitempty"`
	Output       []byte               `protobuf:"bytes,9,opt,name=output,proto3,oneof" json:"output,omitempty"`
	ErrorCode    *string              `protobuf:"bytes,10,opt,name=errorCode,proto3,oneof" json:"errorCode,omitempty"`
	ErrorMessage *string              `protobuf:"bytes,11,opt,name=errorMessage,proto3,oneof" json:"errorMessage,omitempty"`
}

func (x *GetWorkflowInstanceResponse) Reset() {
	*x = GetWorkflowInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_instance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkflowInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowInstanceResponse) ProtoMessage() {}

func (x *GetWorkflowInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_instance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowInstanceResponse.ProtoReflect.Descriptor instead.
func (*GetWorkflowInstanceResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_instance_proto_rawDescGZIP(), []int{1}
}

func (x *GetWorkflowInstanceResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *GetWorkflowInstanceResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *GetWorkflowInstanceResponse) GetInvokedBy() string {
	if x != nil && x.InvokedBy != nil {
		return *x.InvokedBy
	}
	return ""
}

func (x *GetWorkflowInstanceResponse) GetRevision() int32 {
	if x != nil && x.Revision != nil {
		return *x.Revision
	}
	return 0
}

func (x *GetWorkflowInstanceResponse) GetBeginTime() *timestamp.Timestamp {
	if x != nil {
		return x.BeginTime
	}
	return nil
}

func (x *GetWorkflowInstanceResponse) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *GetWorkflowInstanceResponse) GetFlow() []string {
	if x != nil {
		return x.Flow
	}
	return nil
}

func (x *GetWorkflowInstanceResponse) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *GetWorkflowInstanceResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *GetWorkflowInstanceResponse) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

func (x *GetWorkflowInstanceResponse) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

var File_pkg_ingress_get_instance_proto protoreflect.FileDescriptor

var file_pkg_ingress_get_instance_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x22, 0xa0, 0x04, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x69, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x09, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x06, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x07, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_ingress_get_instance_proto_rawDescOnce sync.Once
	file_pkg_ingress_get_instance_proto_rawDescData = file_pkg_ingress_get_instance_proto_rawDesc
)

func file_pkg_ingress_get_instance_proto_rawDescGZIP() []byte {
	file_pkg_ingress_get_instance_proto_rawDescOnce.Do(func() {
		file_pkg_ingress_get_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ingress_get_instance_proto_rawDescData)
	})
	return file_pkg_ingress_get_instance_proto_rawDescData
}

var file_pkg_ingress_get_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_ingress_get_instance_proto_goTypes = []interface{}{
	(*GetWorkflowInstanceRequest)(nil),  // 0: ingress.GetWorkflowInstanceRequest
	(*GetWorkflowInstanceResponse)(nil), // 1: ingress.GetWorkflowInstanceResponse
	(*timestamp.Timestamp)(nil),         // 2: google.protobuf.Timestamp
}
var file_pkg_ingress_get_instance_proto_depIdxs = []int32{
	2, // 0: ingress.GetWorkflowInstanceResponse.beginTime:type_name -> google.protobuf.Timestamp
	2, // 1: ingress.GetWorkflowInstanceResponse.endTime:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_ingress_get_instance_proto_init() }
func file_pkg_ingress_get_instance_proto_init() {
	if File_pkg_ingress_get_instance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ingress_get_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkflowInstanceRequest); i {
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
		file_pkg_ingress_get_instance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkflowInstanceResponse); i {
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
	file_pkg_ingress_get_instance_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_ingress_get_instance_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_get_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_ingress_get_instance_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_get_instance_proto_depIdxs,
		MessageInfos:      file_pkg_ingress_get_instance_proto_msgTypes,
	}.Build()
	File_pkg_ingress_get_instance_proto = out.File
	file_pkg_ingress_get_instance_proto_rawDesc = nil
	file_pkg_ingress_get_instance_proto_goTypes = nil
	file_pkg_ingress_get_instance_proto_depIdxs = nil
}
