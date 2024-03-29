// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: api/v1/quiz/quiz.proto

package quiz

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

type Quiz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId int64  `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Quiz) Reset() {
	*x = Quiz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_quiz_quiz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quiz) ProtoMessage() {}

func (x *Quiz) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_quiz_quiz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quiz.ProtoReflect.Descriptor instead.
func (*Quiz) Descriptor() ([]byte, []int) {
	return file_api_v1_quiz_quiz_proto_rawDescGZIP(), []int{0}
}

func (x *Quiz) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Quiz) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Quiz) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateQuizRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64  `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateQuizRequest) Reset() {
	*x = CreateQuizRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_quiz_quiz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuizRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuizRequest) ProtoMessage() {}

func (x *CreateQuizRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_quiz_quiz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuizRequest.ProtoReflect.Descriptor instead.
func (*CreateQuizRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_quiz_quiz_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuizRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *CreateQuizRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateQuizResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateQuizResponse) Reset() {
	*x = CreateQuizResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_quiz_quiz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuizResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuizResponse) ProtoMessage() {}

func (x *CreateQuizResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_quiz_quiz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuizResponse.ProtoReflect.Descriptor instead.
func (*CreateQuizResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_quiz_quiz_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQuizResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_v1_quiz_quiz_proto protoreflect.FileDescriptor

var file_api_v1_quiz_quiz_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2f, 0x71, 0x75,
	0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a,
	0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x22, 0x49, 0x0a,
	0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x6e, 0x0a, 0x0b, 0x51, 0x75, 0x69, 0x7a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x69, 0x7a, 0x12, 0x26, 0x2e, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x7a, 0x68,
	0x6f, 0x70, 0x69, 0x6a, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x73,
	0x6f, 0x6c, 0x6f, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x2f, 0x7a, 0x68, 0x6f, 0x70, 0x69, 0x6a, 0x2f,
	0x71, 0x75, 0x69, 0x7a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_quiz_quiz_proto_rawDescOnce sync.Once
	file_api_v1_quiz_quiz_proto_rawDescData = file_api_v1_quiz_quiz_proto_rawDesc
)

func file_api_v1_quiz_quiz_proto_rawDescGZIP() []byte {
	file_api_v1_quiz_quiz_proto_rawDescOnce.Do(func() {
		file_api_v1_quiz_quiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_quiz_quiz_proto_rawDescData)
	})
	return file_api_v1_quiz_quiz_proto_rawDescData
}

var file_api_v1_quiz_quiz_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_quiz_quiz_proto_goTypes = []interface{}{
	(*Quiz)(nil),               // 0: zhopij.quiz.v1.quiz.Quiz
	(*CreateQuizRequest)(nil),  // 1: zhopij.quiz.v1.quiz.CreateQuizRequest
	(*CreateQuizResponse)(nil), // 2: zhopij.quiz.v1.quiz.CreateQuizResponse
}
var file_api_v1_quiz_quiz_proto_depIdxs = []int32{
	1, // 0: zhopij.quiz.v1.quiz.QuizService.CreateQuiz:input_type -> zhopij.quiz.v1.quiz.CreateQuizRequest
	2, // 1: zhopij.quiz.v1.quiz.QuizService.CreateQuiz:output_type -> zhopij.quiz.v1.quiz.CreateQuizResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_quiz_quiz_proto_init() }
func file_api_v1_quiz_quiz_proto_init() {
	if File_api_v1_quiz_quiz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_quiz_quiz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quiz); i {
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
		file_api_v1_quiz_quiz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuizRequest); i {
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
		file_api_v1_quiz_quiz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuizResponse); i {
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
			RawDescriptor: file_api_v1_quiz_quiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_quiz_quiz_proto_goTypes,
		DependencyIndexes: file_api_v1_quiz_quiz_proto_depIdxs,
		MessageInfos:      file_api_v1_quiz_quiz_proto_msgTypes,
	}.Build()
	File_api_v1_quiz_quiz_proto = out.File
	file_api_v1_quiz_quiz_proto_rawDesc = nil
	file_api_v1_quiz_quiz_proto_goTypes = nil
	file_api_v1_quiz_quiz_proto_depIdxs = nil
}
