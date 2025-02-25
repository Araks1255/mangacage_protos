// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/service_notifications/service_notifications.proto

package service_notifications

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TitleOnModeration struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TitleName     string                 `protobuf:"bytes,1,opt,name=TitleName,proto3" json:"TitleName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TitleOnModeration) Reset() {
	*x = TitleOnModeration{}
	mi := &file_proto_service_notifications_service_notifications_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TitleOnModeration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitleOnModeration) ProtoMessage() {}

func (x *TitleOnModeration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_notifications_service_notifications_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitleOnModeration.ProtoReflect.Descriptor instead.
func (*TitleOnModeration) Descriptor() ([]byte, []int) {
	return file_proto_service_notifications_service_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *TitleOnModeration) GetTitleName() string {
	if x != nil {
		return x.TitleName
	}
	return ""
}

type ChapterOnModeration struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TitleName     string                 `protobuf:"bytes,1,opt,name=TitleName,proto3" json:"TitleName,omitempty"`
	ChapterName   string                 `protobuf:"bytes,2,opt,name=ChapterName,proto3" json:"ChapterName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChapterOnModeration) Reset() {
	*x = ChapterOnModeration{}
	mi := &file_proto_service_notifications_service_notifications_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChapterOnModeration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterOnModeration) ProtoMessage() {}

func (x *ChapterOnModeration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_notifications_service_notifications_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterOnModeration.ProtoReflect.Descriptor instead.
func (*ChapterOnModeration) Descriptor() ([]byte, []int) {
	return file_proto_service_notifications_service_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *ChapterOnModeration) GetTitleName() string {
	if x != nil {
		return x.TitleName
	}
	return ""
}

func (x *ChapterOnModeration) GetChapterName() string {
	if x != nil {
		return x.ChapterName
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_service_notifications_service_notifications_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_notifications_service_notifications_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_service_notifications_service_notifications_proto_rawDescGZIP(), []int{2}
}

var File_proto_service_notifications_service_notifications_proto protoreflect.FileDescriptor

var file_proto_service_notifications_service_notifications_proto_rawDesc = string([]byte{
	0x0a, 0x37, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x65, 0x72, 0x76, 0x6f,
	0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x31, 0x0a, 0x11, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x4f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0xf0, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x69, 0x0a, 0x1f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x4e, 0x65, 0x77, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x4f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x4f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x6f, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6d, 0x0a, 0x21, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x4f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x6f, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x6f,
	0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x28, 0x5a, 0x26, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x63,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_service_notifications_service_notifications_proto_rawDescOnce sync.Once
	file_proto_service_notifications_service_notifications_proto_rawDescData []byte
)

func file_proto_service_notifications_service_notifications_proto_rawDescGZIP() []byte {
	file_proto_service_notifications_service_notifications_proto_rawDescOnce.Do(func() {
		file_proto_service_notifications_service_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_service_notifications_service_notifications_proto_rawDesc), len(file_proto_service_notifications_service_notifications_proto_rawDesc)))
	})
	return file_proto_service_notifications_service_notifications_proto_rawDescData
}

var file_proto_service_notifications_service_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_service_notifications_service_notifications_proto_goTypes = []any{
	(*TitleOnModeration)(nil),   // 0: servoce_notifications.TitleOnModeration
	(*ChapterOnModeration)(nil), // 1: servoce_notifications.ChapterOnModeration
	(*Empty)(nil),               // 2: servoce_notifications.Empty
}
var file_proto_service_notifications_service_notifications_proto_depIdxs = []int32{
	0, // 0: servoce_notifications.ServiceNotifications.NotifyAboutNewTitleOnModeration:input_type -> servoce_notifications.TitleOnModeration
	1, // 1: servoce_notifications.ServiceNotifications.NotifyAboutNewChapterOnModeration:input_type -> servoce_notifications.ChapterOnModeration
	2, // 2: servoce_notifications.ServiceNotifications.NotifyAboutNewTitleOnModeration:output_type -> servoce_notifications.Empty
	2, // 3: servoce_notifications.ServiceNotifications.NotifyAboutNewChapterOnModeration:output_type -> servoce_notifications.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_notifications_service_notifications_proto_init() }
func file_proto_service_notifications_service_notifications_proto_init() {
	if File_proto_service_notifications_service_notifications_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_service_notifications_service_notifications_proto_rawDesc), len(file_proto_service_notifications_service_notifications_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_notifications_service_notifications_proto_goTypes,
		DependencyIndexes: file_proto_service_notifications_service_notifications_proto_depIdxs,
		MessageInfos:      file_proto_service_notifications_service_notifications_proto_msgTypes,
	}.Build()
	File_proto_service_notifications_service_notifications_proto = out.File
	file_proto_service_notifications_service_notifications_proto_goTypes = nil
	file_proto_service_notifications_service_notifications_proto_depIdxs = nil
}
