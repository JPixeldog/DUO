// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: friend.proto

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

type FriendState int32

const (
	FriendState_offline FriendState = 0
	FriendState_inGame  FriendState = 1
	FriendState_inLobby FriendState = 2
	FriendState_online  FriendState = 3
)

// Enum value maps for FriendState.
var (
	FriendState_name = map[int32]string{
		0: "offline",
		1: "inGame",
		2: "inLobby",
		3: "online",
	}
	FriendState_value = map[string]int32{
		"offline": 0,
		"inGame":  1,
		"inLobby": 2,
		"online":  3,
	}
)

func (x FriendState) Enum() *FriendState {
	p := new(FriendState)
	*p = x
	return p
}

func (x FriendState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FriendState) Descriptor() protoreflect.EnumDescriptor {
	return file_friend_proto_enumTypes[0].Descriptor()
}

func (FriendState) Type() protoreflect.EnumType {
	return &file_friend_proto_enumTypes[0]
}

func (x FriendState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FriendState.Descriptor instead.
func (FriendState) EnumDescriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{0}
}

type Friend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` //user uuid
	Name  string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Score int32       `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	State FriendState `protobuf:"varint,4,opt,name=state,proto3,enum=pb.FriendState" json:"state,omitempty"`
}

func (x *Friend) Reset() {
	*x = Friend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Friend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Friend) ProtoMessage() {}

func (x *Friend) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Friend.ProtoReflect.Descriptor instead.
func (*Friend) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{0}
}

func (x *Friend) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Friend) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Friend) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Friend) GetState() FriendState {
	if x != nil {
		return x.State
	}
	return FriendState_offline
}

type FriendList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string token = 1;
	Friends []*Friend `protobuf:"bytes,2,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *FriendList) Reset() {
	*x = FriendList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendList) ProtoMessage() {}

func (x *FriendList) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendList.ProtoReflect.Descriptor instead.
func (*FriendList) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{1}
}

func (x *FriendList) GetFriends() []*Friend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type FriendRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` //token of requester
	TargetId      string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	RequesterName string `protobuf:"bytes,3,opt,name=requester_name,json=requesterName,proto3" json:"requester_name,omitempty"`
}

func (x *FriendRequestRequest) Reset() {
	*x = FriendRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestRequest) ProtoMessage() {}

func (x *FriendRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestRequest.ProtoReflect.Descriptor instead.
func (*FriendRequestRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{2}
}

func (x *FriendRequestRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FriendRequestRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *FriendRequestRequest) GetRequesterName() string {
	if x != nil {
		return x.RequesterName
	}
	return ""
}

type FriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterUuid string `protobuf:"bytes,1,opt,name=requester_uuid,json=requesterUuid,proto3" json:"requester_uuid,omitempty"`
	RequesterName string `protobuf:"bytes,2,opt,name=requester_name,json=requesterName,proto3" json:"requester_name,omitempty"`
	TargetUuid    string `protobuf:"bytes,3,opt,name=target_uuid,json=targetUuid,proto3" json:"target_uuid,omitempty"`
}

func (x *FriendRequest) Reset() {
	*x = FriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequest) ProtoMessage() {}

func (x *FriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequest.ProtoReflect.Descriptor instead.
func (*FriendRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{3}
}

func (x *FriendRequest) GetRequesterUuid() string {
	if x != nil {
		return x.RequesterUuid
	}
	return ""
}

func (x *FriendRequest) GetRequesterName() string {
	if x != nil {
		return x.RequesterName
	}
	return ""
}

func (x *FriendRequest) GetTargetUuid() string {
	if x != nil {
		return x.TargetUuid
	}
	return ""
}

type FriendRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*FriendRequest `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *FriendRequestList) Reset() {
	*x = FriendRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestList) ProtoMessage() {}

func (x *FriendRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestList.ProtoReflect.Descriptor instead.
func (*FriendRequestList) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{4}
}

func (x *FriendRequestList) GetRequests() []*FriendRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type FriendRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` //token of target
	RequesterId string `protobuf:"bytes,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	Accept      bool   `protobuf:"varint,3,opt,name=accept,proto3" json:"accept,omitempty"`
}

func (x *FriendRequestResponse) Reset() {
	*x = FriendRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestResponse) ProtoMessage() {}

func (x *FriendRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestResponse.ProtoReflect.Descriptor instead.
func (*FriendRequestResponse) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{5}
}

func (x *FriendRequestResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FriendRequestResponse) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *FriendRequestResponse) GetAccept() bool {
	if x != nil {
		return x.Accept
	}
	return false
}

type DeleteFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TargetId string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
}

func (x *DeleteFriendRequest) Reset() {
	*x = DeleteFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendRequest) ProtoMessage() {}

func (x *DeleteFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendRequest.ProtoReflect.Descriptor instead.
func (*DeleteFriendRequest) Descriptor() ([]byte, []int) {
	return file_friend_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFriendRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteFriendRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

var File_friend_proto protoreflect.FileDescriptor

var file_friend_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x6d, 0x0a, 0x06, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x32, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x70, 0x0a, 0x14, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x68, 0x0a, 0x15, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x22, 0x48, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x2a,
	0x3f, 0x0a, 0x0b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x69,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x69, 0x6e, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x03,
	0x42, 0x13, 0x5a, 0x11, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x75, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_proto_rawDescOnce sync.Once
	file_friend_proto_rawDescData = file_friend_proto_rawDesc
)

func file_friend_proto_rawDescGZIP() []byte {
	file_friend_proto_rawDescOnce.Do(func() {
		file_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_proto_rawDescData)
	})
	return file_friend_proto_rawDescData
}

var file_friend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_friend_proto_goTypes = []interface{}{
	(FriendState)(0),              // 0: pb.FriendState
	(*Friend)(nil),                // 1: pb.Friend
	(*FriendList)(nil),            // 2: pb.FriendList
	(*FriendRequestRequest)(nil),  // 3: pb.FriendRequestRequest
	(*FriendRequest)(nil),         // 4: pb.FriendRequest
	(*FriendRequestList)(nil),     // 5: pb.FriendRequestList
	(*FriendRequestResponse)(nil), // 6: pb.FriendRequestResponse
	(*DeleteFriendRequest)(nil),   // 7: pb.DeleteFriendRequest
}
var file_friend_proto_depIdxs = []int32{
	0, // 0: pb.Friend.state:type_name -> pb.FriendState
	1, // 1: pb.FriendList.friends:type_name -> pb.Friend
	4, // 2: pb.FriendRequestList.requests:type_name -> pb.FriendRequest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_friend_proto_init() }
func file_friend_proto_init() {
	if File_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Friend); i {
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
		file_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendList); i {
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
		file_friend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestRequest); i {
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
		file_friend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequest); i {
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
		file_friend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestList); i {
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
		file_friend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestResponse); i {
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
		file_friend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendRequest); i {
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
			RawDescriptor: file_friend_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_friend_proto_goTypes,
		DependencyIndexes: file_friend_proto_depIdxs,
		EnumInfos:         file_friend_proto_enumTypes,
		MessageInfos:      file_friend_proto_msgTypes,
	}.Build()
	File_friend_proto = out.File
	file_friend_proto_rawDesc = nil
	file_friend_proto_goTypes = nil
	file_friend_proto_depIdxs = nil
}
