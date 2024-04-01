// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: proto/authentication_management_api.proto

package auth

import (
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

type LoginMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginMessage) Reset() {
	*x = LoginMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authentication_management_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMessage) ProtoMessage() {}

func (x *LoginMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authentication_management_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMessage.ProtoReflect.Descriptor instead.
func (*LoginMessage) Descriptor() ([]byte, []int) {
	return file_proto_authentication_management_api_proto_rawDescGZIP(), []int{0}
}

func (x *LoginMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (x *LoginResponseMessage) Reset() {
	*x = LoginResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authentication_management_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponseMessage) ProtoMessage() {}

func (x *LoginResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authentication_management_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponseMessage.ProtoReflect.Descriptor instead.
func (*LoginResponseMessage) Descriptor() ([]byte, []int) {
	return file_proto_authentication_management_api_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponseMessage) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type ExtensionAccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldAccessToken string `protobuf:"bytes,1,opt,name=oldAccessToken,proto3" json:"oldAccessToken,omitempty"`
}

func (x *ExtensionAccessToken) Reset() {
	*x = ExtensionAccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authentication_management_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionAccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionAccessToken) ProtoMessage() {}

func (x *ExtensionAccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authentication_management_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionAccessToken.ProtoReflect.Descriptor instead.
func (*ExtensionAccessToken) Descriptor() ([]byte, []int) {
	return file_proto_authentication_management_api_proto_rawDescGZIP(), []int{2}
}

func (x *ExtensionAccessToken) GetOldAccessToken() string {
	if x != nil {
		return x.OldAccessToken
	}
	return ""
}

type ReplyExtensionAccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewAccessToken string `protobuf:"bytes,1,opt,name=newAccessToken,proto3" json:"newAccessToken,omitempty"`
}

func (x *ReplyExtensionAccessToken) Reset() {
	*x = ReplyExtensionAccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authentication_management_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyExtensionAccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyExtensionAccessToken) ProtoMessage() {}

func (x *ReplyExtensionAccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authentication_management_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyExtensionAccessToken.ProtoReflect.Descriptor instead.
func (*ReplyExtensionAccessToken) Descriptor() ([]byte, []int) {
	return file_proto_authentication_management_api_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyExtensionAccessToken) GetNewAccessToken() string {
	if x != nil {
		return x.NewAccessToken
	}
	return ""
}

var File_proto_authentication_management_api_proto protoreflect.FileDescriptor

var file_proto_authentication_management_api_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x50, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x38, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x14,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x0e, 0x6f, 0x6c, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0e, 0x6f, 0x6c, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x43, 0x0a, 0x19, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc5, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x6a, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0b,
	0x5a, 0x09, 0x2f, 0x61, 0x70, 0x75, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_authentication_management_api_proto_rawDescOnce sync.Once
	file_proto_authentication_management_api_proto_rawDescData = file_proto_authentication_management_api_proto_rawDesc
)

func file_proto_authentication_management_api_proto_rawDescGZIP() []byte {
	file_proto_authentication_management_api_proto_rawDescOnce.Do(func() {
		file_proto_authentication_management_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_authentication_management_api_proto_rawDescData)
	})
	return file_proto_authentication_management_api_proto_rawDescData
}

var file_proto_authentication_management_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_authentication_management_api_proto_goTypes = []interface{}{
	(*LoginMessage)(nil),              // 0: auth.LoginMessage
	(*LoginResponseMessage)(nil),      // 1: auth.LoginResponseMessage
	(*ExtensionAccessToken)(nil),      // 2: auth.ExtensionAccessToken
	(*ReplyExtensionAccessToken)(nil), // 3: auth.ReplyExtensionAccessToken
}
var file_proto_authentication_management_api_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.Login:input_type -> auth.LoginMessage
	2, // 1: auth.AuthService.ExtensionToken:input_type -> auth.ExtensionAccessToken
	1, // 2: auth.AuthService.Login:output_type -> auth.LoginResponseMessage
	3, // 3: auth.AuthService.ExtensionToken:output_type -> auth.ReplyExtensionAccessToken
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_authentication_management_api_proto_init() }
func file_proto_authentication_management_api_proto_init() {
	if File_proto_authentication_management_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_authentication_management_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMessage); i {
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
		file_proto_authentication_management_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponseMessage); i {
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
		file_proto_authentication_management_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionAccessToken); i {
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
		file_proto_authentication_management_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyExtensionAccessToken); i {
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
			RawDescriptor: file_proto_authentication_management_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_authentication_management_api_proto_goTypes,
		DependencyIndexes: file_proto_authentication_management_api_proto_depIdxs,
		MessageInfos:      file_proto_authentication_management_api_proto_msgTypes,
	}.Build()
	File_proto_authentication_management_api_proto = out.File
	file_proto_authentication_management_api_proto_rawDesc = nil
	file_proto_authentication_management_api_proto_goTypes = nil
	file_proto_authentication_management_api_proto_depIdxs = nil
}
