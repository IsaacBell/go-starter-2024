// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: proto/basic/user.proto

package basic

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                 string      `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username              string      `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	IsActive              bool        `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Password              []byte      `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirmation  []byte      `protobuf:"bytes,6,opt,name=password_confirmation,json=passwordConfirmation,proto3" json:"password_confirmation,omitempty"`
	Uuid                  string      `protobuf:"bytes,7,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CommunicationChannels []*any1.Any `protobuf:"bytes,16,rep,name=communication_channels,json=communicationChannels,proto3" json:"communication_channels,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *User) GetPasswordConfirmation() []byte {
	if x != nil {
		return x.PasswordConfirmation
	}
	return nil
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetCommunicationChannels() []*any1.Any {
	if x != nil {
		return x.CommunicationChannels
	}
	return nil
}

type PaperMailChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *PaperMailChannel) Reset() {
	*x = PaperMailChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaperMailChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaperMailChannel) ProtoMessage() {}

func (x *PaperMailChannel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaperMailChannel.ProtoReflect.Descriptor instead.
func (*PaperMailChannel) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{1}
}

func (x *PaperMailChannel) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type SocialMediaChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SocialMediaChannel) Reset() {
	*x = SocialMediaChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialMediaChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialMediaChannel) ProtoMessage() {}

func (x *SocialMediaChannel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialMediaChannel.ProtoReflect.Descriptor instead.
func (*SocialMediaChannel) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{2}
}

func (x *SocialMediaChannel) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *SocialMediaChannel) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type EmailChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailChannel) Reset() {
	*x = EmailChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailChannel) ProtoMessage() {}

func (x *EmailChannel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailChannel.ProtoReflect.Descriptor instead.
func (*EmailChannel) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{3}
}

func (x *EmailChannel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_proto_basic_user_proto protoreflect.FileDescriptor

var file_proto_basic_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x33,
	0x0a, 0x15, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x4b, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x50, 0x61, 0x70, 0x65, 0x72, 0x4d, 0x61, 0x69,
	0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x4c, 0x0a, 0x12, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x24, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_basic_user_proto_rawDescOnce sync.Once
	file_proto_basic_user_proto_rawDescData = file_proto_basic_user_proto_rawDesc
)

func file_proto_basic_user_proto_rawDescGZIP() []byte {
	file_proto_basic_user_proto_rawDescOnce.Do(func() {
		file_proto_basic_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_basic_user_proto_rawDescData)
	})
	return file_proto_basic_user_proto_rawDescData
}

var file_proto_basic_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_basic_user_proto_goTypes = []interface{}{
	(*User)(nil),               // 0: basic.User
	(*PaperMailChannel)(nil),   // 1: basic.PaperMailChannel
	(*SocialMediaChannel)(nil), // 2: basic.SocialMediaChannel
	(*EmailChannel)(nil),       // 3: basic.EmailChannel
	(*any1.Any)(nil),           // 4: google.protobuf.Any
}
var file_proto_basic_user_proto_depIdxs = []int32{
	4, // 0: basic.User.communication_channels:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_basic_user_proto_init() }
func file_proto_basic_user_proto_init() {
	if File_proto_basic_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_basic_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_basic_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaperMailChannel); i {
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
		file_proto_basic_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialMediaChannel); i {
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
		file_proto_basic_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailChannel); i {
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
			RawDescriptor: file_proto_basic_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_basic_user_proto_goTypes,
		DependencyIndexes: file_proto_basic_user_proto_depIdxs,
		MessageInfos:      file_proto_basic_user_proto_msgTypes,
	}.Build()
	File_proto_basic_user_proto = out.File
	file_proto_basic_user_proto_rawDesc = nil
	file_proto_basic_user_proto_goTypes = nil
	file_proto_basic_user_proto_depIdxs = nil
}
