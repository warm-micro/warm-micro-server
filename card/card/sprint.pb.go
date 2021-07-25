// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: sprint.proto

package __

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetSprintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SprintId string `protobuf:"bytes,1,opt,name=sprint_id,json=sprintId,proto3" json:"sprint_id,omitempty"`
}

func (x *GetSprintRequest) Reset() {
	*x = GetSprintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sprint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSprintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSprintRequest) ProtoMessage() {}

func (x *GetSprintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sprint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSprintRequest.ProtoReflect.Descriptor instead.
func (*GetSprintRequest) Descriptor() ([]byte, []int) {
	return file_sprint_proto_rawDescGZIP(), []int{0}
}

func (x *GetSprintRequest) GetSprintId() string {
	if x != nil {
		return x.SprintId
	}
	return ""
}

type Sprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkspaceId string `protobuf:"bytes,2,opt,name=workspaceId,proto3" json:"workspaceId,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Sprint) Reset() {
	*x = Sprint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sprint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sprint) ProtoMessage() {}

func (x *Sprint) ProtoReflect() protoreflect.Message {
	mi := &file_sprint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sprint.ProtoReflect.Descriptor instead.
func (*Sprint) Descriptor() ([]byte, []int) {
	return file_sprint_proto_rawDescGZIP(), []int{1}
}

func (x *Sprint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sprint) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *Sprint) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_sprint_proto protoreflect.FileDescriptor

var file_sprint_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0xd5, 0x02, 0x0a, 0x10, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x53,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x09, 0x67, 0x65,
	0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x2f, 0x7b, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x51, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0e,
	0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2e,
	0x2f, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sprint_proto_rawDescOnce sync.Once
	file_sprint_proto_rawDescData = file_sprint_proto_rawDesc
)

func file_sprint_proto_rawDescGZIP() []byte {
	file_sprint_proto_rawDescOnce.Do(func() {
		file_sprint_proto_rawDescData = protoimpl.X.CompressGZIP(file_sprint_proto_rawDescData)
	})
	return file_sprint_proto_rawDescData
}

var file_sprint_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sprint_proto_goTypes = []interface{}{
	(*GetSprintRequest)(nil),       // 0: sprint.GetSprintRequest
	(*Sprint)(nil),                 // 1: sprint.Sprint
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 3: google.protobuf.BoolValue
}
var file_sprint_proto_depIdxs = []int32{
	1, // 0: sprint.SprintManagement.addSprint:input_type -> sprint.Sprint
	0, // 1: sprint.SprintManagement.getSprint:input_type -> sprint.GetSprintRequest
	2, // 2: sprint.SprintManagement.listSprints:input_type -> google.protobuf.StringValue
	0, // 3: sprint.SprintManagement.checkSprint:input_type -> sprint.GetSprintRequest
	2, // 4: sprint.SprintManagement.addSprint:output_type -> google.protobuf.StringValue
	1, // 5: sprint.SprintManagement.getSprint:output_type -> sprint.Sprint
	1, // 6: sprint.SprintManagement.listSprints:output_type -> sprint.Sprint
	3, // 7: sprint.SprintManagement.checkSprint:output_type -> google.protobuf.BoolValue
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sprint_proto_init() }
func file_sprint_proto_init() {
	if File_sprint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sprint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSprintRequest); i {
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
		file_sprint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sprint); i {
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
			RawDescriptor: file_sprint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sprint_proto_goTypes,
		DependencyIndexes: file_sprint_proto_depIdxs,
		MessageInfos:      file_sprint_proto_msgTypes,
	}.Build()
	File_sprint_proto = out.File
	file_sprint_proto_rawDesc = nil
	file_sprint_proto_goTypes = nil
	file_sprint_proto_depIdxs = nil
}