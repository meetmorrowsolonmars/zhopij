// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: api/v1/message/answer.proto

package message

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

type EventType int32

const (
	EventType_EVENT_TYPE_UNKNOWN EventType = 0
	EventType_CREATE             EventType = 1
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNKNOWN",
		1: "CREATE",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNKNOWN": 0,
		"CREATE":             1,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_message_answer_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_api_v1_message_answer_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_message_answer_proto_rawDescGZIP(), []int{0}
}

type MessageVersion int32

const (
	MessageVersion_MESSAGE_VERSION_UNKNOWN MessageVersion = 0
	MessageVersion_V1_0_0                  MessageVersion = 1
)

// Enum value maps for MessageVersion.
var (
	MessageVersion_name = map[int32]string{
		0: "MESSAGE_VERSION_UNKNOWN",
		1: "V1_0_0",
	}
	MessageVersion_value = map[string]int32{
		"MESSAGE_VERSION_UNKNOWN": 0,
		"V1_0_0":                  1,
	}
)

func (x MessageVersion) Enum() *MessageVersion {
	p := new(MessageVersion)
	*p = x
	return p
}

func (x MessageVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_message_answer_proto_enumTypes[1].Descriptor()
}

func (MessageVersion) Type() protoreflect.EnumType {
	return &file_api_v1_message_answer_proto_enumTypes[1]
}

func (x MessageVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageVersion.Descriptor instead.
func (MessageVersion) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_message_answer_proto_rawDescGZIP(), []int{1}
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string         `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Version   MessageVersion `protobuf:"varint,2,opt,name=version,proto3,enum=zhopij.answer.v1.message.MessageVersion" json:"version,omitempty"`
	Event     EventType      `protobuf:"varint,3,opt,name=event,proto3,enum=zhopij.answer.v1.message.EventType" json:"event,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_message_answer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_message_answer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_api_v1_message_answer_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Meta) GetVersion() MessageVersion {
	if x != nil {
		return x.Version
	}
	return MessageVersion_MESSAGE_VERSION_UNKNOWN
}

func (x *Meta) GetEvent() EventType {
	if x != nil {
		return x.Event
	}
	return EventType_EVENT_TYPE_UNKNOWN
}

type BaseEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta  *Meta            `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Event *BaseEvent_Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *BaseEvent) Reset() {
	*x = BaseEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_message_answer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseEvent) ProtoMessage() {}

func (x *BaseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_message_answer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseEvent.ProtoReflect.Descriptor instead.
func (*BaseEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_message_answer_proto_rawDescGZIP(), []int{1}
}

func (x *BaseEvent) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *BaseEvent) GetEvent() *BaseEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type V1CreateAnswerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta  *Meta                      `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Event *V1CreateAnswerEvent_Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *V1CreateAnswerEvent) Reset() {
	*x = V1CreateAnswerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_message_answer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1CreateAnswerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1CreateAnswerEvent) ProtoMessage() {}

func (x *V1CreateAnswerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_message_answer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1CreateAnswerEvent.ProtoReflect.Descriptor instead.
func (*V1CreateAnswerEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_message_answer_proto_rawDescGZIP(), []int{2}
}

func (x *V1CreateAnswerEvent) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *V1CreateAnswerEvent) GetEvent() *V1CreateAnswerEvent_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type BaseEvent_Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BaseEvent_Event) Reset() {
	*x = BaseEvent_Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_message_answer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseEvent_Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseEvent_Event) ProtoMessage() {}

func (x *BaseEvent_Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_message_answer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseEvent_Event.ProtoReflect.Descriptor instead.
func (*BaseEvent_Event) Descriptor() ([]byte, []int) {
	return file_api_v1_message_answer_proto_rawDescGZIP(), []int{1, 0}
}

type V1CreateAnswerEvent_Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizId   int64 `protobuf:"varint,1,opt,name=quiz_id,json=quizId,proto3" json:"quiz_id,omitempty"`
	AuthorId int64 `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *V1CreateAnswerEvent_Event) Reset() {
	*x = V1CreateAnswerEvent_Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_message_answer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1CreateAnswerEvent_Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1CreateAnswerEvent_Event) ProtoMessage() {}

func (x *V1CreateAnswerEvent_Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_message_answer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1CreateAnswerEvent_Event.ProtoReflect.Descriptor instead.
func (*V1CreateAnswerEvent_Event) Descriptor() ([]byte, []int) {
	return file_api_v1_message_answer_proto_rawDescGZIP(), []int{2, 0}
}

func (x *V1CreateAnswerEvent_Event) GetQuizId() int64 {
	if x != nil {
		return x.QuizId
	}
	return 0
}

func (x *V1CreateAnswerEvent_Event) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

var File_api_v1_message_answer_proto protoreflect.FileDescriptor

var file_api_v1_message_answer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x7a,
	0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x42, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x89,
	0x01, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x68, 0x6f,
	0x70, 0x69, 0x6a, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x3f, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x1a, 0x07, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x13, 0x56,
	0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x56, 0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x1a, 0x3d, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x75,
	0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x71, 0x75, 0x69,
	0x7a, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x2a, 0x2f, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10,
	0x01, 0x2a, 0x39, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x56, 0x31, 0x5f, 0x30, 0x5f, 0x30, 0x10, 0x01, 0x42, 0x49, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x6d,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x73, 0x6f, 0x6c, 0x6f, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x2f, 0x7a,
	0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_message_answer_proto_rawDescOnce sync.Once
	file_api_v1_message_answer_proto_rawDescData = file_api_v1_message_answer_proto_rawDesc
)

func file_api_v1_message_answer_proto_rawDescGZIP() []byte {
	file_api_v1_message_answer_proto_rawDescOnce.Do(func() {
		file_api_v1_message_answer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_message_answer_proto_rawDescData)
	})
	return file_api_v1_message_answer_proto_rawDescData
}

var file_api_v1_message_answer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_message_answer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_message_answer_proto_goTypes = []interface{}{
	(EventType)(0),                    // 0: zhopij.answer.v1.message.EventType
	(MessageVersion)(0),               // 1: zhopij.answer.v1.message.MessageVersion
	(*Meta)(nil),                      // 2: zhopij.answer.v1.message.Meta
	(*BaseEvent)(nil),                 // 3: zhopij.answer.v1.message.BaseEvent
	(*V1CreateAnswerEvent)(nil),       // 4: zhopij.answer.v1.message.V1CreateAnswerEvent
	(*BaseEvent_Event)(nil),           // 5: zhopij.answer.v1.message.BaseEvent.Event
	(*V1CreateAnswerEvent_Event)(nil), // 6: zhopij.answer.v1.message.V1CreateAnswerEvent.Event
}
var file_api_v1_message_answer_proto_depIdxs = []int32{
	1, // 0: zhopij.answer.v1.message.Meta.version:type_name -> zhopij.answer.v1.message.MessageVersion
	0, // 1: zhopij.answer.v1.message.Meta.event:type_name -> zhopij.answer.v1.message.EventType
	2, // 2: zhopij.answer.v1.message.BaseEvent.meta:type_name -> zhopij.answer.v1.message.Meta
	5, // 3: zhopij.answer.v1.message.BaseEvent.event:type_name -> zhopij.answer.v1.message.BaseEvent.Event
	2, // 4: zhopij.answer.v1.message.V1CreateAnswerEvent.meta:type_name -> zhopij.answer.v1.message.Meta
	6, // 5: zhopij.answer.v1.message.V1CreateAnswerEvent.event:type_name -> zhopij.answer.v1.message.V1CreateAnswerEvent.Event
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1_message_answer_proto_init() }
func file_api_v1_message_answer_proto_init() {
	if File_api_v1_message_answer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_message_answer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_api_v1_message_answer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseEvent); i {
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
		file_api_v1_message_answer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1CreateAnswerEvent); i {
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
		file_api_v1_message_answer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseEvent_Event); i {
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
		file_api_v1_message_answer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1CreateAnswerEvent_Event); i {
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
			RawDescriptor: file_api_v1_message_answer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_message_answer_proto_goTypes,
		DependencyIndexes: file_api_v1_message_answer_proto_depIdxs,
		EnumInfos:         file_api_v1_message_answer_proto_enumTypes,
		MessageInfos:      file_api_v1_message_answer_proto_msgTypes,
	}.Build()
	File_api_v1_message_answer_proto = out.File
	file_api_v1_message_answer_proto_rawDesc = nil
	file_api_v1_message_answer_proto_goTypes = nil
	file_api_v1_message_answer_proto_depIdxs = nil
}
