// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: protobuffers/thread.proto

package thread_service

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Thread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	OwnerUserId   string `protobuf:"bytes,3,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	Archived      bool   `protobuf:"varint,4,opt,name=archived,proto3" json:"archived,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	OwnerUserName string `protobuf:"bytes,6,opt,name=ownerUserName,proto3" json:"ownerUserName,omitempty"`
	FirmId        string `protobuf:"bytes,7,opt,name=firmId,proto3" json:"firmId,omitempty"`
	FirmName      string `protobuf:"bytes,8,opt,name=firmName,proto3" json:"firmName,omitempty"`
}

func (x *Thread) Reset() {
	*x = Thread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_thread_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thread) ProtoMessage() {}

func (x *Thread) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_thread_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thread.ProtoReflect.Descriptor instead.
func (*Thread) Descriptor() ([]byte, []int) {
	return file_protobuffers_thread_proto_rawDescGZIP(), []int{0}
}

func (x *Thread) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Thread) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Thread) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *Thread) GetArchived() bool {
	if x != nil {
		return x.Archived
	}
	return false
}

func (x *Thread) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Thread) GetOwnerUserName() string {
	if x != nil {
		return x.OwnerUserName
	}
	return ""
}

func (x *Thread) GetFirmId() string {
	if x != nil {
		return x.FirmId
	}
	return ""
}

func (x *Thread) GetFirmName() string {
	if x != nil {
		return x.FirmName
	}
	return ""
}

type GetThreadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadId string `protobuf:"bytes,1,opt,name=threadId,proto3" json:"threadId,omitempty"`
}

func (x *GetThreadRequest) Reset() {
	*x = GetThreadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_thread_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThreadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThreadRequest) ProtoMessage() {}

func (x *GetThreadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_thread_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThreadRequest.ProtoReflect.Descriptor instead.
func (*GetThreadRequest) Descriptor() ([]byte, []int) {
	return file_protobuffers_thread_proto_rawDescGZIP(), []int{1}
}

func (x *GetThreadRequest) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

type GetThreadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thread *Thread `protobuf:"bytes,1,opt,name=thread,proto3" json:"thread,omitempty"`
}

func (x *GetThreadResponse) Reset() {
	*x = GetThreadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_thread_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThreadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThreadResponse) ProtoMessage() {}

func (x *GetThreadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_thread_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThreadResponse.ProtoReflect.Descriptor instead.
func (*GetThreadResponse) Descriptor() ([]byte, []int) {
	return file_protobuffers_thread_proto_rawDescGZIP(), []int{2}
}

func (x *GetThreadResponse) GetThread() *Thread {
	if x != nil {
		return x.Thread
	}
	return nil
}

type CreateThreadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thread *Thread `protobuf:"bytes,1,opt,name=thread,proto3" json:"thread,omitempty"`
}

func (x *CreateThreadRequest) Reset() {
	*x = CreateThreadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_thread_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateThreadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateThreadRequest) ProtoMessage() {}

func (x *CreateThreadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_thread_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateThreadRequest.ProtoReflect.Descriptor instead.
func (*CreateThreadRequest) Descriptor() ([]byte, []int) {
	return file_protobuffers_thread_proto_rawDescGZIP(), []int{3}
}

func (x *CreateThreadRequest) GetThread() *Thread {
	if x != nil {
		return x.Thread
	}
	return nil
}

type CreateThreadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadId string `protobuf:"bytes,1,opt,name=threadId,proto3" json:"threadId,omitempty"`
}

func (x *CreateThreadResponse) Reset() {
	*x = CreateThreadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffers_thread_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateThreadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateThreadResponse) ProtoMessage() {}

func (x *CreateThreadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffers_thread_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateThreadResponse.ProtoReflect.Descriptor instead.
func (*CreateThreadResponse) Descriptor() ([]byte, []int) {
	return file_protobuffers_thread_proto_rawDescGZIP(), []int{4}
}

func (x *CreateThreadResponse) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

var File_protobuffers_thread_proto protoreflect.FileDescriptor

var file_protobuffers_thread_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x06, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x05, 0x18, 0x96, 0x01,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42,
	0x12, 0x72, 0x10, 0x32, 0x0e, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x2a, 0x24, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xe8, 0x07, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x66, 0x69, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15,
	0xfa, 0x42, 0x12, 0x72, 0x10, 0x32, 0x0e, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x2a, 0x24, 0x52, 0x06, 0x66, 0x69, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x72, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x72, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x06,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x06, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x22, 0x32, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x32, 0xcc, 0x01, 0x0a, 0x0d, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2f, 0x7b,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x20, 0x5a, 0x1e, 0x2f, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protobuffers_thread_proto_rawDescOnce sync.Once
	file_protobuffers_thread_proto_rawDescData = file_protobuffers_thread_proto_rawDesc
)

func file_protobuffers_thread_proto_rawDescGZIP() []byte {
	file_protobuffers_thread_proto_rawDescOnce.Do(func() {
		file_protobuffers_thread_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuffers_thread_proto_rawDescData)
	})
	return file_protobuffers_thread_proto_rawDescData
}

var file_protobuffers_thread_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuffers_thread_proto_goTypes = []interface{}{
	(*Thread)(nil),               // 0: thread.Thread
	(*GetThreadRequest)(nil),     // 1: thread.GetThreadRequest
	(*GetThreadResponse)(nil),    // 2: thread.GetThreadResponse
	(*CreateThreadRequest)(nil),  // 3: thread.CreateThreadRequest
	(*CreateThreadResponse)(nil), // 4: thread.CreateThreadResponse
}
var file_protobuffers_thread_proto_depIdxs = []int32{
	0, // 0: thread.GetThreadResponse.thread:type_name -> thread.Thread
	0, // 1: thread.CreateThreadRequest.thread:type_name -> thread.Thread
	1, // 2: thread.ThreadService.GetThread:input_type -> thread.GetThreadRequest
	3, // 3: thread.ThreadService.CreateThread:input_type -> thread.CreateThreadRequest
	2, // 4: thread.ThreadService.GetThread:output_type -> thread.GetThreadResponse
	4, // 5: thread.ThreadService.CreateThread:output_type -> thread.CreateThreadResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuffers_thread_proto_init() }
func file_protobuffers_thread_proto_init() {
	if File_protobuffers_thread_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuffers_thread_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thread); i {
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
		file_protobuffers_thread_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThreadRequest); i {
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
		file_protobuffers_thread_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThreadResponse); i {
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
		file_protobuffers_thread_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateThreadRequest); i {
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
		file_protobuffers_thread_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateThreadResponse); i {
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
			RawDescriptor: file_protobuffers_thread_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuffers_thread_proto_goTypes,
		DependencyIndexes: file_protobuffers_thread_proto_depIdxs,
		MessageInfos:      file_protobuffers_thread_proto_msgTypes,
	}.Build()
	File_protobuffers_thread_proto = out.File
	file_protobuffers_thread_proto_rawDesc = nil
	file_protobuffers_thread_proto_goTypes = nil
	file_protobuffers_thread_proto_depIdxs = nil
}
