// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: protos/auth.proto

package v1

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

type UserRegisterInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password       string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	RepeatPassword string `protobuf:"bytes,3,opt,name=repeatPassword,proto3" json:"repeatPassword,omitempty"`
}

func (x *UserRegisterInput) Reset() {
	*x = UserRegisterInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterInput) ProtoMessage() {}

func (x *UserRegisterInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterInput.ProtoReflect.Descriptor instead.
func (*UserRegisterInput) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegisterInput) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRegisterInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRegisterInput) GetRepeatPassword() string {
	if x != nil {
		return x.RepeatPassword
	}
	return ""
}

type UserLoginInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginInput) Reset() {
	*x = UserLoginInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginInput) ProtoMessage() {}

func (x *UserLoginInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginInput.ProtoReflect.Descriptor instead.
func (*UserLoginInput) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginInput) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *UserEntity      `protobuf:"bytes,1,opt,name=user,proto3,oneof" json:"user,omitempty"`
	Token *TokenPayloadDto `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *RegisterPayload) Reset() {
	*x = RegisterPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPayload) ProtoMessage() {}

func (x *RegisterPayload) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPayload.ProtoReflect.Descriptor instead.
func (*RegisterPayload) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterPayload) GetUser() *UserEntity {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RegisterPayload) GetToken() *TokenPayloadDto {
	if x != nil {
		return x.Token
	}
	return nil
}

type TokenPayloadDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpiresIn   *int32  `protobuf:"varint,1,opt,name=expiresIn,proto3,oneof" json:"expiresIn,omitempty"`
	AccessToken *string `protobuf:"bytes,2,opt,name=accessToken,proto3,oneof" json:"accessToken,omitempty"`
}

func (x *TokenPayloadDto) Reset() {
	*x = TokenPayloadDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenPayloadDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPayloadDto) ProtoMessage() {}

func (x *TokenPayloadDto) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenPayloadDto.ProtoReflect.Descriptor instead.
func (*TokenPayloadDto) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{3}
}

func (x *TokenPayloadDto) GetExpiresIn() int32 {
	if x != nil && x.ExpiresIn != nil {
		return *x.ExpiresIn
	}
	return 0
}

func (x *TokenPayloadDto) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

var File_protos_auth_proto protoreflect.FileDescriptor

var file_protos_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x1a, 0x1d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x48, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x34,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x74, 0x6f, 0x48, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x79, 0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x49, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x32, 0x93, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_auth_proto_rawDescOnce sync.Once
	file_protos_auth_proto_rawDescData = file_protos_auth_proto_rawDesc
)

func file_protos_auth_proto_rawDescGZIP() []byte {
	file_protos_auth_proto_rawDescOnce.Do(func() {
		file_protos_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_auth_proto_rawDescData)
	})
	return file_protos_auth_proto_rawDescData
}

var file_protos_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_auth_proto_goTypes = []interface{}{
	(*UserRegisterInput)(nil), // 0: protogen.UserRegisterInput
	(*UserLoginInput)(nil),    // 1: protogen.UserLoginInput
	(*RegisterPayload)(nil),   // 2: protogen.RegisterPayload
	(*TokenPayloadDto)(nil),   // 3: protogen.TokenPayloadDto
	(*UserEntity)(nil),        // 4: protogen.UserEntity
}
var file_protos_auth_proto_depIdxs = []int32{
	4, // 0: protogen.RegisterPayload.user:type_name -> protogen.UserEntity
	3, // 1: protogen.RegisterPayload.token:type_name -> protogen.TokenPayloadDto
	0, // 2: protogen.AuthService.Register:input_type -> protogen.UserRegisterInput
	1, // 3: protogen.AuthService.Login:input_type -> protogen.UserLoginInput
	2, // 4: protogen.AuthService.Register:output_type -> protogen.RegisterPayload
	3, // 5: protogen.AuthService.Login:output_type -> protogen.TokenPayloadDto
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_auth_proto_init() }
func file_protos_auth_proto_init() {
	if File_protos_auth_proto != nil {
		return
	}
	file_protos_user_user_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterInput); i {
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
		file_protos_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginInput); i {
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
		file_protos_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPayload); i {
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
		file_protos_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenPayloadDto); i {
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
	file_protos_auth_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protos_auth_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_auth_proto_goTypes,
		DependencyIndexes: file_protos_auth_proto_depIdxs,
		MessageInfos:      file_protos_auth_proto_msgTypes,
	}.Build()
	File_protos_auth_proto = out.File
	file_protos_auth_proto_rawDesc = nil
	file_protos_auth_proto_goTypes = nil
	file_protos_auth_proto_depIdxs = nil
}
