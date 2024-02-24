// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: session_messages.proto

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

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{0}
}

type SessionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JoinedUserUuids       string `protobuf:"bytes,1,opt,name=joined_user_uuids,json=joinedUserUuids,proto3" json:"joined_user_uuids,omitempty"`
	DisconnectedUserUuids string `protobuf:"bytes,2,opt,name=disconnected_user_uuids,json=disconnectedUserUuids,proto3" json:"disconnected_user_uuids,omitempty"`
}

func (x *SessionState) Reset() {
	*x = SessionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionState) ProtoMessage() {}

func (x *SessionState) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionState.ProtoReflect.Descriptor instead.
func (*SessionState) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{1}
}

func (x *SessionState) GetJoinedUserUuids() string {
	if x != nil {
		return x.JoinedUserUuids
	}
	return ""
}

func (x *SessionState) GetDisconnectedUserUuids() string {
	if x != nil {
		return x.DisconnectedUserUuids
	}
	return ""
}

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Pin   int32  `protobuf:"varint,2,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSessionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateSessionRequest) GetPin() int32 {
	if x != nil {
		return x.Pin
	}
	return 0
}

type JoinSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SessionId int32  `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Pin       int32  `protobuf:"varint,3,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *JoinSessionRequest) Reset() {
	*x = JoinSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSessionRequest) ProtoMessage() {}

func (x *JoinSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSessionRequest.ProtoReflect.Descriptor instead.
func (*JoinSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{3}
}

func (x *JoinSessionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JoinSessionRequest) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *JoinSessionRequest) GetPin() int32 {
	if x != nil {
		return x.Pin
	}
	return 0
}

type SessionStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId    int32         `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	GameState    *GameState    `protobuf:"bytes,2,opt,name=game_state,json=gameState,proto3" json:"game_state,omitempty"`
	SessionState *SessionState `protobuf:"bytes,3,opt,name=session_state,json=sessionState,proto3" json:"session_state,omitempty"`
}

func (x *SessionStream) Reset() {
	*x = SessionStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStream) ProtoMessage() {}

func (x *SessionStream) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStream.ProtoReflect.Descriptor instead.
func (*SessionStream) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{4}
}

func (x *SessionStream) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SessionStream) GetGameState() *GameState {
	if x != nil {
		return x.GameState
	}
	return nil
}

func (x *SessionStream) GetSessionState() *SessionState {
	if x != nil {
		return x.SessionState
	}
	return nil
}

type DisconnectSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SessionId int32  `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *DisconnectSessionRequest) Reset() {
	*x = DisconnectSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectSessionRequest) ProtoMessage() {}

func (x *DisconnectSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectSessionRequest.ProtoReflect.Descriptor instead.
func (*DisconnectSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{5}
}

func (x *DisconnectSessionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DisconnectSessionRequest) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type DisconnectSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DisconnectSessionResponse) Reset() {
	*x = DisconnectSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectSessionResponse) ProtoMessage() {}

func (x *DisconnectSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectSessionResponse.ProtoReflect.Descriptor instead.
func (*DisconnectSessionResponse) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{6}
}

func (x *DisconnectSessionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SessionId int32  `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *DeleteSessionRequest) Reset() {
	*x = DeleteSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSessionRequest) ProtoMessage() {}

func (x *DeleteSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSessionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSessionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteSessionRequest) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type DeleteSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteSessionResponse) Reset() {
	*x = DeleteSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSessionResponse) ProtoMessage() {}

func (x *DeleteSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSessionResponse.ProtoReflect.Descriptor instead.
func (*DeleteSessionResponse) Descriptor() ([]byte, []int) {
	return file_session_messages_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSessionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_session_messages_proto protoreflect.FileDescriptor

var file_session_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x0b, 0x0a, 0x09,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x72, 0x0a, 0x0c, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6a, 0x6f, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x73, 0x22, 0x3e, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x22, 0x5b, 0x0a,
	0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09,
	0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x4f, 0x0a, 0x18, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_messages_proto_rawDescOnce sync.Once
	file_session_messages_proto_rawDescData = file_session_messages_proto_rawDesc
)

func file_session_messages_proto_rawDescGZIP() []byte {
	file_session_messages_proto_rawDescOnce.Do(func() {
		file_session_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_messages_proto_rawDescData)
	})
	return file_session_messages_proto_rawDescData
}

var file_session_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_session_messages_proto_goTypes = []interface{}{
	(*GameState)(nil),                 // 0: pb.GameState
	(*SessionState)(nil),              // 1: pb.SessionState
	(*CreateSessionRequest)(nil),      // 2: pb.CreateSessionRequest
	(*JoinSessionRequest)(nil),        // 3: pb.JoinSessionRequest
	(*SessionStream)(nil),             // 4: pb.SessionStream
	(*DisconnectSessionRequest)(nil),  // 5: pb.DisconnectSessionRequest
	(*DisconnectSessionResponse)(nil), // 6: pb.DisconnectSessionResponse
	(*DeleteSessionRequest)(nil),      // 7: pb.DeleteSessionRequest
	(*DeleteSessionResponse)(nil),     // 8: pb.DeleteSessionResponse
}
var file_session_messages_proto_depIdxs = []int32{
	0, // 0: pb.SessionStream.game_state:type_name -> pb.GameState
	1, // 1: pb.SessionStream.session_state:type_name -> pb.SessionState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_session_messages_proto_init() }
func file_session_messages_proto_init() {
	if File_session_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_session_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameState); i {
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
		file_session_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionState); i {
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
		file_session_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_session_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinSessionRequest); i {
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
		file_session_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStream); i {
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
		file_session_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectSessionRequest); i {
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
		file_session_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectSessionResponse); i {
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
		file_session_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSessionRequest); i {
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
		file_session_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSessionResponse); i {
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
			RawDescriptor: file_session_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_session_messages_proto_goTypes,
		DependencyIndexes: file_session_messages_proto_depIdxs,
		MessageInfos:      file_session_messages_proto_msgTypes,
	}.Build()
	File_session_messages_proto = out.File
	file_session_messages_proto_rawDesc = nil
	file_session_messages_proto_goTypes = nil
	file_session_messages_proto_depIdxs = nil
}