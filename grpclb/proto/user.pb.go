// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: user.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type PasswordCheckInof struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Password        string                 `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	EncryptPassword string                 `protobuf:"bytes,2,opt,name=encryptPassword,proto3" json:"encryptPassword,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PasswordCheckInof) Reset() {
	*x = PasswordCheckInof{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordCheckInof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordCheckInof) ProtoMessage() {}

func (x *PasswordCheckInof) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PasswordCheckInof.ProtoReflect.Descriptor instead.
func (*PasswordCheckInof) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordCheckInof) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PasswordCheckInof) GetEncryptPassword() string {
	if x != nil {
		return x.EncryptPassword
	}
	return ""
}

type CheckReponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckReponse) Reset() {
	*x = CheckReponse{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReponse) ProtoMessage() {}

func (x *CheckReponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CheckReponse.ProtoReflect.Descriptor instead.
func (*CheckReponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *CheckReponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PageInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          uint32                 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      uint32                 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *PageInfo) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfo) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type MobileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mobile        string                 `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MobileRequest) Reset() {
	*x = MobileRequest{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileRequest) ProtoMessage() {}

func (x *MobileRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MobileRequest.ProtoReflect.Descriptor instead.
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *MobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *IdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateUserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Password      string                 `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Mobile        string                 `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Nickname      string                 `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Username      string                 `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Birthday      uint64                 `protobuf:"varint,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserInfo) Reset() {
	*x = CreateUserInfo{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserInfo) ProtoMessage() {}

func (x *CreateUserInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateUserInfo.ProtoReflect.Descriptor instead.
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CreateUserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateUserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserInfo) GetBirthday() uint64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

type UpdateUserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile        string                 `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Nickname      string                 `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Birthday      uint64                 `protobuf:"varint,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserInfo) Reset() {
	*x = UpdateUserInfo{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfo) ProtoMessage() {}

func (x *UpdateUserInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateUserInfo.ProtoReflect.Descriptor instead.
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UpdateUserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UpdateUserInfo) GetBirthday() uint64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile        string                 `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Nickname      string                 `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Birthday      uint64                 `protobuf:"varint,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender        string                 `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Avatar        string                 `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Role          int32                  `protobuf:"varint,9,opt,name=role,proto3" json:"role,omitempty"`
	Username      string                 `protobuf:"bytes,10,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfoResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfoResponse) GetBirthday() uint64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *UserInfoResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserInfoResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfoResponse) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *UserInfoResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         uint32                 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data          []*UserInfoResponse    `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserListResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserListResponse) GetData() []*UserInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x1a\x1bgoogle/protobuf/empty.proto\"Y\n" +
	"\x11PasswordCheckInof\x12\x1a\n" +
	"\bpassword\x18\x01 \x01(\tR\bpassword\x12(\n" +
	"\x0fencryptPassword\x18\x02 \x01(\tR\x0fencryptPassword\"(\n" +
	"\fCheckReponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\":\n" +
	"\bPageInfo\x12\x12\n" +
	"\x04page\x18\x01 \x01(\rR\x04page\x12\x1a\n" +
	"\bpageSize\x18\x02 \x01(\rR\bpageSize\"'\n" +
	"\rMobileRequest\x12\x16\n" +
	"\x06mobile\x18\x01 \x01(\tR\x06mobile\"\x1b\n" +
	"\tIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"\xae\x01\n" +
	"\x0eCreateUserInfo\x12\x1a\n" +
	"\bpassword\x18\x01 \x01(\tR\bpassword\x12\x16\n" +
	"\x06mobile\x18\x02 \x01(\tR\x06mobile\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\bnickname\x18\x04 \x01(\tR\bnickname\x12\x1a\n" +
	"\busername\x18\x05 \x01(\tR\busername\x12\x1a\n" +
	"\bbirthday\x18\x06 \x01(\x04R\bbirthday\"\xa2\x01\n" +
	"\x0eUpdateUserInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x16\n" +
	"\x06mobile\x18\x03 \x01(\tR\x06mobile\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12\x1a\n" +
	"\bnickname\x18\x05 \x01(\tR\bnickname\x12\x1a\n" +
	"\bbirthday\x18\x06 \x01(\x04R\bbirthday\"\x84\x02\n" +
	"\x10UserInfoResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x16\n" +
	"\x06mobile\x18\x03 \x01(\tR\x06mobile\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12\x1a\n" +
	"\bnickname\x18\x05 \x01(\tR\bnickname\x12\x1a\n" +
	"\bbirthday\x18\x06 \x01(\x04R\bbirthday\x12\x16\n" +
	"\x06gender\x18\a \x01(\tR\x06gender\x12\x16\n" +
	"\x06avatar\x18\b \x01(\tR\x06avatar\x12\x12\n" +
	"\x04role\x18\t \x01(\x05R\x04role\x12\x1a\n" +
	"\busername\x18\n" +
	" \x01(\tR\busername\"O\n" +
	"\x10UserListResponse\x12\x14\n" +
	"\x05total\x18\x01 \x01(\rR\x05total\x12%\n" +
	"\x04data\x18\x02 \x03(\v2\x11.UserInfoResponseR\x04data2\xf4\x02\n" +
	"\x04User\x12-\n" +
	"\vGetUserList\x12\t.PageInfo\x1a\x11.UserListResponse\"\x00\x126\n" +
	"\x0fGetUserByMobile\x12\x0e.MobileRequest\x1a\x11.UserInfoResponse\"\x00\x12.\n" +
	"\vGetUserById\x12\n" +
	".IdRequest\x1a\x11.UserInfoResponse\"\x00\x122\n" +
	"\n" +
	"CreateUser\x12\x0f.CreateUserInfo\x1a\x11.UserInfoResponse\"\x00\x127\n" +
	"\n" +
	"UpdateUser\x12\x0f.UpdateUserInfo\x1a\x16.google.protobuf.Empty\"\x00\x122\n" +
	"\n" +
	"DeleteUser\x12\n" +
	".IdRequest\x1a\x16.google.protobuf.Empty\"\x00\x124\n" +
	"\rCheckPassword\x12\x12.PasswordCheckInof\x1a\r.CheckReponse\"\x00B\tZ\a.;protob\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_proto_goTypes = []any{
	(*PasswordCheckInof)(nil), // 0: PasswordCheckInof
	(*CheckReponse)(nil),      // 1: CheckReponse
	(*PageInfo)(nil),          // 2: PageInfo
	(*MobileRequest)(nil),     // 3: MobileRequest
	(*IdRequest)(nil),         // 4: IdRequest
	(*CreateUserInfo)(nil),    // 5: CreateUserInfo
	(*UpdateUserInfo)(nil),    // 6: UpdateUserInfo
	(*UserInfoResponse)(nil),  // 7: UserInfoResponse
	(*UserListResponse)(nil),  // 8: UserListResponse
	(*emptypb.Empty)(nil),     // 9: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	7, // 0: UserListResponse.data:type_name -> UserInfoResponse
	2, // 1: User.GetUserList:input_type -> PageInfo
	3, // 2: User.GetUserByMobile:input_type -> MobileRequest
	4, // 3: User.GetUserById:input_type -> IdRequest
	5, // 4: User.CreateUser:input_type -> CreateUserInfo
	6, // 5: User.UpdateUser:input_type -> UpdateUserInfo
	4, // 6: User.DeleteUser:input_type -> IdRequest
	0, // 7: User.CheckPassword:input_type -> PasswordCheckInof
	8, // 8: User.GetUserList:output_type -> UserListResponse
	7, // 9: User.GetUserByMobile:output_type -> UserInfoResponse
	7, // 10: User.GetUserById:output_type -> UserInfoResponse
	7, // 11: User.CreateUser:output_type -> UserInfoResponse
	9, // 12: User.UpdateUser:output_type -> google.protobuf.Empty
	9, // 13: User.DeleteUser:output_type -> google.protobuf.Empty
	1, // 14: User.CheckPassword:output_type -> CheckReponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
