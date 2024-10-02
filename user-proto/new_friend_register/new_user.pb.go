// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: new_user.proto

package new_friend_register

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

type NewFriendRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Uuid  string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *NewFriendRegister) Reset() {
	*x = NewFriendRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFriendRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFriendRegister) ProtoMessage() {}

func (x *NewFriendRegister) ProtoReflect() protoreflect.Message {
	mi := &file_new_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFriendRegister.ProtoReflect.Descriptor instead.
func (*NewFriendRegister) Descriptor() ([]byte, []int) {
	return file_new_user_proto_rawDescGZIP(), []int{0}
}

func (x *NewFriendRegister) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewFriendRegister) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_new_user_proto protoreflect.FileDescriptor

var file_new_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x42,
	0x20, 0x5a, 0x1e, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65,
	0x77, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_new_user_proto_rawDescOnce sync.Once
	file_new_user_proto_rawDescData = file_new_user_proto_rawDesc
)

func file_new_user_proto_rawDescGZIP() []byte {
	file_new_user_proto_rawDescOnce.Do(func() {
		file_new_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_new_user_proto_rawDescData)
	})
	return file_new_user_proto_rawDescData
}

var file_new_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_new_user_proto_goTypes = []any{
	(*NewFriendRegister)(nil), // 0: NewFriendRegister
}
var file_new_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_new_user_proto_init() }
func file_new_user_proto_init() {
	if File_new_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_new_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NewFriendRegister); i {
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
			RawDescriptor: file_new_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_new_user_proto_goTypes,
		DependencyIndexes: file_new_user_proto_depIdxs,
		MessageInfos:      file_new_user_proto_msgTypes,
	}.Build()
	File_new_user_proto = out.File
	file_new_user_proto_rawDesc = nil
	file_new_user_proto_goTypes = nil
	file_new_user_proto_depIdxs = nil
}