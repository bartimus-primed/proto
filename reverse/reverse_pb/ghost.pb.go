// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ghost.proto

package reverse_pb

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

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd          string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Timeout      int32  `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	SendResp     bool   `protobuf:"varint,3,opt,name=send_resp,json=sendResp,proto3" json:"send_resp,omitempty"`
	StartHandsOn bool   `protobuf:"varint,4,opt,name=start_hands_on,json=startHandsOn,proto3" json:"start_hands_on,omitempty"`
	ShouldExit   bool   `protobuf:"varint,5,opt,name=should_exit,json=shouldExit,proto3" json:"should_exit,omitempty"`
	IsSession    bool   `protobuf:"varint,6,opt,name=is_session,json=isSession,proto3" json:"is_session,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *Command) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Command) GetSendResp() bool {
	if x != nil {
		return x.SendResp
	}
	return false
}

func (x *Command) GetStartHandsOn() bool {
	if x != nil {
		return x.StartHandsOn
	}
	return false
}

func (x *Command) GetShouldExit() bool {
	if x != nil {
		return x.ShouldExit
	}
	return false
}

func (x *Command) GetIsSession() bool {
	if x != nil {
		return x.IsSession
	}
	return false
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RanCommand string `protobuf:"bytes,1,opt,name=ranCommand,proto3" json:"ranCommand,omitempty"`
	Resp       string `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	Success    bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Ready      bool   `protobuf:"varint,4,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_ghost_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetRanCommand() string {
	if x != nil {
		return x.RanCommand
	}
	return ""
}

func (x *Response) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

var File_ghost_proto protoreflect.FileDescriptor

var file_ghost_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xb8, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x45, 0x78, 0x69, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x61,
	0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x32, 0x68, 0x0a, 0x0f,
	0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x29, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x07, 0x48, 0x61,
	0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2d, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ghost_proto_rawDescOnce sync.Once
	file_ghost_proto_rawDescData = file_ghost_proto_rawDesc
)

func file_ghost_proto_rawDescGZIP() []byte {
	file_ghost_proto_rawDescOnce.Do(func() {
		file_ghost_proto_rawDescData = protoimpl.X.CompressGZIP(file_ghost_proto_rawDescData)
	})
	return file_ghost_proto_rawDescData
}

var file_ghost_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ghost_proto_goTypes = []interface{}{
	(*Command)(nil),  // 0: pb.Command
	(*Response)(nil), // 1: pb.Response
}
var file_ghost_proto_depIdxs = []int32{
	1, // 0: pb.ReverseInteract.GetCommand:input_type -> pb.Response
	1, // 1: pb.ReverseInteract.HandsOn:input_type -> pb.Response
	0, // 2: pb.ReverseInteract.GetCommand:output_type -> pb.Command
	0, // 3: pb.ReverseInteract.HandsOn:output_type -> pb.Command
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ghost_proto_init() }
func file_ghost_proto_init() {
	if File_ghost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ghost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_ghost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_ghost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ghost_proto_goTypes,
		DependencyIndexes: file_ghost_proto_depIdxs,
		MessageInfos:      file_ghost_proto_msgTypes,
	}.Build()
	File_ghost_proto = out.File
	file_ghost_proto_rawDesc = nil
	file_ghost_proto_goTypes = nil
	file_ghost_proto_depIdxs = nil
}
