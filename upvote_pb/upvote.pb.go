// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: upvote.proto

package upvote_pb

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

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId int32 `protobuf:"varint,1,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upvote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_upvote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_upvote_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetResourceId() int32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

type VotePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VoteIncrement int32 `protobuf:"varint,2,opt,name=voteIncrement,proto3" json:"voteIncrement,omitempty"`
}

func (x *VotePayload) Reset() {
	*x = VotePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upvote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VotePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VotePayload) ProtoMessage() {}

func (x *VotePayload) ProtoReflect() protoreflect.Message {
	mi := &file_upvote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VotePayload.ProtoReflect.Descriptor instead.
func (*VotePayload) Descriptor() ([]byte, []int) {
	return file_upvote_proto_rawDescGZIP(), []int{1}
}

func (x *VotePayload) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VotePayload) GetVoteIncrement() int32 {
	if x != nil {
		return x.VoteIncrement
	}
	return 0
}

type VoteCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VotesCount int32 `protobuf:"varint,2,opt,name=votesCount,proto3" json:"votesCount,omitempty"`
}

func (x *VoteCount) Reset() {
	*x = VoteCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upvote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteCount) ProtoMessage() {}

func (x *VoteCount) ProtoReflect() protoreflect.Message {
	mi := &file_upvote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteCount.ProtoReflect.Descriptor instead.
func (*VoteCount) Descriptor() ([]byte, []int) {
	return file_upvote_proto_rawDescGZIP(), []int{2}
}

func (x *VoteCount) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VoteCount) GetVotesCount() int32 {
	if x != nil {
		return x.VotesCount
	}
	return 0
}

var File_upvote_proto protoreflect.FileDescriptor

var file_upvote_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x43, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0x75, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x10, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x1a, 0x11, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x56, 0x6f, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x11, 0x2e, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x2e,
	0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2f,
	0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x62, 0x3b, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upvote_proto_rawDescOnce sync.Once
	file_upvote_proto_rawDescData = file_upvote_proto_rawDesc
)

func file_upvote_proto_rawDescGZIP() []byte {
	file_upvote_proto_rawDescOnce.Do(func() {
		file_upvote_proto_rawDescData = protoimpl.X.CompressGZIP(file_upvote_proto_rawDescData)
	})
	return file_upvote_proto_rawDescData
}

var file_upvote_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_upvote_proto_goTypes = []interface{}{
	(*Resource)(nil),    // 0: upvote.Resource
	(*VotePayload)(nil), // 1: upvote.VotePayload
	(*VoteCount)(nil),   // 2: upvote.VoteCount
}
var file_upvote_proto_depIdxs = []int32{
	0, // 0: upvote.UserManagement.GetVotes:input_type -> upvote.Resource
	1, // 1: upvote.UserManagement.Vote:input_type -> upvote.VotePayload
	2, // 2: upvote.UserManagement.GetVotes:output_type -> upvote.VoteCount
	2, // 3: upvote.UserManagement.Vote:output_type -> upvote.VoteCount
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upvote_proto_init() }
func file_upvote_proto_init() {
	if File_upvote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upvote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_upvote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VotePayload); i {
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
		file_upvote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteCount); i {
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
			RawDescriptor: file_upvote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upvote_proto_goTypes,
		DependencyIndexes: file_upvote_proto_depIdxs,
		MessageInfos:      file_upvote_proto_msgTypes,
	}.Build()
	File_upvote_proto = out.File
	file_upvote_proto_rawDesc = nil
	file_upvote_proto_goTypes = nil
	file_upvote_proto_depIdxs = nil
}
