// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: user.proto

package user_proto

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

// Data in request or getting uuid by login. If User doesnt exist - user will be creating
type GetUserByLoginIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Email of target user
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *GetUserByLoginIn) Reset() {
	*x = GetUserByLoginIn{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByLoginIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByLoginIn) ProtoMessage() {}

func (x *GetUserByLoginIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByLoginIn.ProtoReflect.Descriptor instead.
func (*GetUserByLoginIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByLoginIn) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

// Message for response
type GetUserByLoginOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID of user
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Flag for indicate of new user
	IsNewUser bool `protobuf:"varint,2,opt,name=isNewUser,proto3" json:"isNewUser,omitempty"`
}

func (x *GetUserByLoginOut) Reset() {
	*x = GetUserByLoginOut{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByLoginOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByLoginOut) ProtoMessage() {}

func (x *GetUserByLoginOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByLoginOut.ProtoReflect.Descriptor instead.
func (*GetUserByLoginOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByLoginOut) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetUserByLoginOut) GetIsNewUser() bool {
	if x != nil {
		return x.IsNewUser
	}
	return false
}

// Message for request
type IsUserExistByUUIDIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID for target user
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *IsUserExistByUUIDIn) Reset() {
	*x = IsUserExistByUUIDIn{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsUserExistByUUIDIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExistByUUIDIn) ProtoMessage() {}

func (x *IsUserExistByUUIDIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExistByUUIDIn.ProtoReflect.Descriptor instead.
func (*IsUserExistByUUIDIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *IsUserExistByUUIDIn) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// Message for response
type IsUserExistByUUIDOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Flag of indicate user exist
	IsExist bool `protobuf:"varint,1,opt,name=isExist,proto3" json:"isExist,omitempty"`
}

func (x *IsUserExistByUUIDOut) Reset() {
	*x = IsUserExistByUUIDOut{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsUserExistByUUIDOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExistByUUIDOut) ProtoMessage() {}

func (x *IsUserExistByUUIDOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExistByUUIDOut.ProtoReflect.Descriptor instead.
func (*IsUserExistByUUIDOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *IsUserExistByUUIDOut) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

// Request data fo getting user info (for initiator page)
type GetUserInfoByUUIDIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID for target user
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetUserInfoByUUIDIn) Reset() {
	*x = GetUserInfoByUUIDIn{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserInfoByUUIDIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoByUUIDIn) ProtoMessage() {}

func (x *GetUserInfoByUUIDIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoByUUIDIn.ProtoReflect.Descriptor instead.
func (*GetUserInfoByUUIDIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserInfoByUUIDIn) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetOs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *GetOs) Reset() {
	*x = GetOs{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOs) ProtoMessage() {}

func (x *GetOs) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOs.ProtoReflect.Descriptor instead.
func (*GetOs) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetOs) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetOs) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// Response data for initiator page
type GetUserInfoByUUIDOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname   string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar     string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Name       *string  `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Surname    *string  `protobuf:"bytes,4,opt,name=surname,proto3,oneof" json:"surname,omitempty"`
	Birthdate  *string  `protobuf:"bytes,5,opt,name=birthdate,proto3,oneof" json:"birthdate,omitempty"`
	Phone      *string  `protobuf:"bytes,6,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	City       *string  `protobuf:"bytes,7,opt,name=city,proto3,oneof" json:"city,omitempty"`
	Telegram   *string  `protobuf:"bytes,8,opt,name=telegram,proto3,oneof" json:"telegram,omitempty"`
	Git        *string  `protobuf:"bytes,9,opt,name=git,proto3,oneof" json:"git,omitempty"`
	Os         *GetOs   `protobuf:"bytes,10,opt,name=os,proto3,oneof" json:"os,omitempty"`
	Work       *string  `protobuf:"bytes,11,opt,name=work,proto3,oneof" json:"work,omitempty"`
	University *string  `protobuf:"bytes,12,opt,name=university,proto3,oneof" json:"university,omitempty"`
	Skills     []string `protobuf:"bytes,13,rep,name=skills,proto3" json:"skills,omitempty"`
	Hobbies    []string `protobuf:"bytes,14,rep,name=hobbies,proto3" json:"hobbies,omitempty"`
}

func (x *GetUserInfoByUUIDOut) Reset() {
	*x = GetUserInfoByUUIDOut{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserInfoByUUIDOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoByUUIDOut) ProtoMessage() {}

func (x *GetUserInfoByUUIDOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoByUUIDOut.ProtoReflect.Descriptor instead.
func (*GetUserInfoByUUIDOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserInfoByUUIDOut) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetSurname() string {
	if x != nil && x.Surname != nil {
		return *x.Surname
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetBirthdate() string {
	if x != nil && x.Birthdate != nil {
		return *x.Birthdate
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetTelegram() string {
	if x != nil && x.Telegram != nil {
		return *x.Telegram
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetGit() string {
	if x != nil && x.Git != nil {
		return *x.Git
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetOs() *GetOs {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *GetUserInfoByUUIDOut) GetWork() string {
	if x != nil && x.Work != nil {
		return *x.Work
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetUniversity() string {
	if x != nil && x.University != nil {
		return *x.University
	}
	return ""
}

func (x *GetUserInfoByUUIDOut) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *GetUserInfoByUUIDOut) GetHobbies() []string {
	if x != nil {
		return x.Hobbies
	}
	return nil
}

type GetLoginByUUIDIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetLoginByUUIDIn) Reset() {
	*x = GetLoginByUUIDIn{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLoginByUUIDIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginByUUIDIn) ProtoMessage() {}

func (x *GetLoginByUUIDIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginByUUIDIn.ProtoReflect.Descriptor instead.
func (*GetLoginByUUIDIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetLoginByUUIDIn) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetLoginByUUIDOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *GetLoginByUUIDOut) Reset() {
	*x = GetLoginByUUIDOut{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLoginByUUIDOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginByUUIDOut) ProtoMessage() {}

func (x *GetLoginByUUIDOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginByUUIDOut.ProtoReflect.Descriptor instead.
func (*GetLoginByUUIDOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetLoginByUUIDOut) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type GetUserWithOffsetIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit    int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *GetUserWithOffsetIn) Reset() {
	*x = GetUserWithOffsetIn{}
	mi := &file_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserWithOffsetIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWithOffsetIn) ProtoMessage() {}

func (x *GetUserWithOffsetIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWithOffsetIn.ProtoReflect.Descriptor instead.
func (*GetUserWithOffsetIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserWithOffsetIn) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUserWithOffsetIn) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUserWithOffsetIn) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type GetUserWithOffsetOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetUserWithOffsetOut) Reset() {
	*x = GetUserWithOffsetOut{}
	mi := &file_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserWithOffsetOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWithOffsetOut) ProtoMessage() {}

func (x *GetUserWithOffsetOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWithOffsetOut.ProtoReflect.Descriptor instead.
func (*GetUserWithOffsetOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserWithOffsetOut) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetUserWithOffsetOut) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname   string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Uuid       string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AvatarLink string `protobuf:"bytes,3,opt,name=avatar_link,json=avatarLink,proto3" json:"avatar_link,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string `protobuf:"bytes,5,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if x != nil {
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
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetAvatarLink() string {
	if x != nil {
		return x.AvatarLink
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type UpdateProfileIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Birthday string `protobuf:"bytes,2,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Telegram string `protobuf:"bytes,3,opt,name=telegram,proto3" json:"telegram,omitempty"`
	Github   string `protobuf:"bytes,4,opt,name=github,proto3" json:"github,omitempty"`
	OsId     int64  `protobuf:"varint,5,opt,name=os_id,json=osId,proto3" json:"os_id,omitempty"`
}

func (x *UpdateProfileIn) Reset() {
	*x = UpdateProfileIn{}
	mi := &file_user_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileIn) ProtoMessage() {}

func (x *UpdateProfileIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileIn.ProtoReflect.Descriptor instead.
func (*UpdateProfileIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateProfileIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProfileIn) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UpdateProfileIn) GetTelegram() string {
	if x != nil {
		return x.Telegram
	}
	return ""
}

func (x *UpdateProfileIn) GetGithub() string {
	if x != nil {
		return x.Github
	}
	return ""
}

func (x *UpdateProfileIn) GetOsId() int64 {
	if x != nil {
		return x.OsId
	}
	return 0
}

type UpdateProfileOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateProfileOut) Reset() {
	*x = UpdateProfileOut{}
	mi := &file_user_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileOut) ProtoMessage() {}

func (x *UpdateProfileOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileOut.ProtoReflect.Descriptor instead.
func (*UpdateProfileOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateProfileOut) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a,
	0x13, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x55,
	0x49, 0x44, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x49, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x4f, 0x75, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x49,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4f, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x22, 0x88, 0x04, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x4f, 0x75, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x73,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x88, 0x01,
	0x01, 0x12, 0x15, 0x0a, 0x03, 0x67, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06,
	0x52, 0x03, 0x67, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x73, 0x48, 0x07, 0x52, 0x02,
	0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x09, 0x52, 0x0a, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f,
	0x62, 0x62, 0x69, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x67, 0x69, 0x74,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x6f, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x22,
	0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x55, 0x55, 0x49,
	0x44, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x4f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x22, 0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74,
	0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69,
	0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x85, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x13, 0x0a, 0x05,
	0x6f, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6f, 0x73, 0x49,
	0x64, 0x22, 0x2a, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x87, 0x03,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x11, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x49, 0x6e, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x11, 0x49, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x2e,
	0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x55, 0x49,
	0x44, 0x49, 0x6e, 0x1a, 0x15, 0x2e, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x55, 0x49,
	0x44, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x79, 0x55, 0x55, 0x49, 0x44, 0x49, 0x6e, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x4f, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x55,
	0x55, 0x49, 0x44, 0x49, 0x6e, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x10, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_user_proto_goTypes = []any{
	(*GetUserByLoginIn)(nil),     // 0: GetUserByLoginIn
	(*GetUserByLoginOut)(nil),    // 1: GetUserByLoginOut
	(*IsUserExistByUUIDIn)(nil),  // 2: IsUserExistByUUIDIn
	(*IsUserExistByUUIDOut)(nil), // 3: IsUserExistByUUIDOut
	(*GetUserInfoByUUIDIn)(nil),  // 4: GetUserInfoByUUIDIn
	(*GetOs)(nil),                // 5: GetOs
	(*GetUserInfoByUUIDOut)(nil), // 6: GetUserInfoByUUIDOut
	(*GetLoginByUUIDIn)(nil),     // 7: GetLoginByUUIDIn
	(*GetLoginByUUIDOut)(nil),    // 8: GetLoginByUUIDOut
	(*GetUserWithOffsetIn)(nil),  // 9: GetUserWithOffsetIn
	(*GetUserWithOffsetOut)(nil), // 10: GetUserWithOffsetOut
	(*User)(nil),                 // 11: User
	(*UpdateProfileIn)(nil),      // 12: UpdateProfileIn
	(*UpdateProfileOut)(nil),     // 13: UpdateProfileOut
}
var file_user_proto_depIdxs = []int32{
	5,  // 0: GetUserInfoByUUIDOut.os:type_name -> GetOs
	11, // 1: GetUserWithOffsetOut.user:type_name -> User
	0,  // 2: UserService.GetUserByLogin:input_type -> GetUserByLoginIn
	2,  // 3: UserService.IsUserExistByUUID:input_type -> IsUserExistByUUIDIn
	4,  // 4: UserService.GetUserInfoByUUID:input_type -> GetUserInfoByUUIDIn
	7,  // 5: UserService.GetLoginByUUID:input_type -> GetLoginByUUIDIn
	9,  // 6: UserService.GetUserWithOffset:input_type -> GetUserWithOffsetIn
	12, // 7: UserService.UpdateProfile:input_type -> UpdateProfileIn
	1,  // 8: UserService.GetUserByLogin:output_type -> GetUserByLoginOut
	3,  // 9: UserService.IsUserExistByUUID:output_type -> IsUserExistByUUIDOut
	6,  // 10: UserService.GetUserInfoByUUID:output_type -> GetUserInfoByUUIDOut
	8,  // 11: UserService.GetLoginByUUID:output_type -> GetLoginByUUIDOut
	10, // 12: UserService.GetUserWithOffset:output_type -> GetUserWithOffsetOut
	13, // 13: UserService.UpdateProfile:output_type -> UpdateProfileOut
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_user_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
