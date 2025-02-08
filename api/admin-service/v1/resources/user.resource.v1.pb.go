// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.6
// source: api/admin-service/v1/resources/user.resource.v1.proto

package resourcev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	enums "github.com/go-micro-saas/admin-service/api/admin-service/v1/enums"
	page "github.com/ikaiguang/go-srv-kit/kit/page"
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

// User 用户
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// deleted_time 删除时间
	DeletedTime uint64 `protobuf:"varint,4,opt,name=deleted_time,json=deletedTime,proto3" json:"deleted_time,omitempty"`
	// user_id uid
	UserId uint64 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// user_phone 手机
	UserPhone string `protobuf:"bytes,6,opt,name=user_phone,json=userPhone,proto3" json:"user_phone,omitempty"`
	// user_email 邮箱
	UserEmail string `protobuf:"bytes,7,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// user_nickname 昵称
	UserNickname string `protobuf:"bytes,8,opt,name=user_nickname,json=userNickname,proto3" json:"user_nickname,omitempty"`
	// user_avatar 头像
	UserAvatar string `protobuf:"bytes,9,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	// user_gender 性别；0：INIT，1：MALE，2：FEMALE，3：SECRET
	UserGender enums.UserGenderEnum_UserGender `protobuf:"varint,10,opt,name=user_gender,json=userGender,proto3,enum=saas.api.admin.enumv1.UserGenderEnum_UserGender" json:"user_gender,omitempty"`
	// user_status 状态；0：INIT，1：ENABLE，2：DISABLE，3：WHITELIST，4：BLACKLIST，5：DELETED
	UserStatus enums.UserStatusEnum_UserStatus `protobuf:"varint,12,opt,name=user_status,json=userStatus,proto3,enum=saas.api.admin.enumv1.UserStatusEnum_UserStatus" json:"user_status,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[0]
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
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *User) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *User) GetDeletedTime() uint64 {
	if x != nil {
		return x.DeletedTime
	}
	return 0
}

func (x *User) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

func (x *User) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *User) GetUserNickname() string {
	if x != nil {
		return x.UserNickname
	}
	return ""
}

func (x *User) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *User) GetUserGender() enums.UserGenderEnum_UserGender {
	if x != nil {
		return x.UserGender
	}
	return enums.UserGenderEnum_UserGender(0)
}

func (x *User) GetUserStatus() enums.UserStatusEnum_UserStatus {
	if x != nil {
		return x.UserStatus
	}
	return enums.UserStatusEnum_UserStatus(0)
}

// UserIdReq id
type UserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserIdReq) Reset() {
	*x = UserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReq) ProtoMessage() {}

func (x *UserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdReq.ProtoReflect.Descriptor instead.
func (*UserIdReq) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *UserIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// UserIdsReq ids
type UserIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []int64 `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *UserIdsReq) Reset() {
	*x = UserIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdsReq) ProtoMessage() {}

func (x *UserIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdsReq.ProtoReflect.Descriptor instead.
func (*UserIdsReq) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *UserIdsReq) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// UserListReq list
type UserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageRequest *page.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (x *UserListReq) Reset() {
	*x = UserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReq) ProtoMessage() {}

func (x *UserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReq.ProtoReflect.Descriptor instead.
func (*UserListReq) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{3}
}

func (x *UserListReq) GetPageRequest() *page.PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

type UserListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason   string            `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message  string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data     *UserListRespData `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserListResp) Reset() {
	*x = UserListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResp) ProtoMessage() {}

func (x *UserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResp.ProtoReflect.Descriptor instead.
func (*UserListResp) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{4}
}

func (x *UserListResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserListResp) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *UserListResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserListResp) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UserListResp) GetData() *UserListRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserListRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List     []*User            `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	PageInfo *page.PageResponse `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *UserListRespData) Reset() {
	*x = UserListRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRespData) ProtoMessage() {}

func (x *UserListRespData) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRespData.ProtoReflect.Descriptor instead.
func (*UserListRespData) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{5}
}

func (x *UserListRespData) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *UserListRespData) GetPageInfo() *page.PageResponse {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

// UserProcessResult process result
type UserProcessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *UserProcessResult) Reset() {
	*x = UserProcessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProcessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProcessResult) ProtoMessage() {}

func (x *UserProcessResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProcessResult.ProtoReflect.Descriptor instead.
func (*UserProcessResult) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{6}
}

func (x *UserProcessResult) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_api_admin_service_v1_resources_user_resource_v1_proto protoreflect.FileDescriptor

var file_api_admin_service_v1_resources_user_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x76, 0x31, 0x1a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x51, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e,
	0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x4f, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x40, 0x0a, 0x0c,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x70, 0x6b, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa5,
	0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x84, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x61, 0x61, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x70, 0x6b, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x87, 0x01, 0x0a, 0x19, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x42,
	0x16, 0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73,
	0x61, 0x61, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescOnce sync.Once
	file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescData = file_api_admin_service_v1_resources_user_resource_v1_proto_rawDesc
)

func file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescGZIP() []byte {
	file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescData)
	})
	return file_api_admin_service_v1_resources_user_resource_v1_proto_rawDescData
}

var file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_admin_service_v1_resources_user_resource_v1_proto_goTypes = []interface{}{
	(*User)(nil),                         // 0: saas.api.admin.resourcev1.User
	(*UserIdReq)(nil),                    // 1: saas.api.admin.resourcev1.UserIdReq
	(*UserIdsReq)(nil),                   // 2: saas.api.admin.resourcev1.UserIdsReq
	(*UserListReq)(nil),                  // 3: saas.api.admin.resourcev1.UserListReq
	(*UserListResp)(nil),                 // 4: saas.api.admin.resourcev1.UserListResp
	(*UserListRespData)(nil),             // 5: saas.api.admin.resourcev1.UserListRespData
	(*UserProcessResult)(nil),            // 6: saas.api.admin.resourcev1.UserProcessResult
	nil,                                  // 7: saas.api.admin.resourcev1.UserListResp.MetadataEntry
	(enums.UserGenderEnum_UserGender)(0), // 8: saas.api.admin.enumv1.UserGenderEnum.UserGender
	(enums.UserStatusEnum_UserStatus)(0), // 9: saas.api.admin.enumv1.UserStatusEnum.UserStatus
	(*page.PageRequest)(nil),             // 10: kit.page.pagepkg.PageRequest
	(*page.PageResponse)(nil),            // 11: kit.page.pagepkg.PageResponse
}
var file_api_admin_service_v1_resources_user_resource_v1_proto_depIdxs = []int32{
	8,  // 0: saas.api.admin.resourcev1.User.user_gender:type_name -> saas.api.admin.enumv1.UserGenderEnum.UserGender
	9,  // 1: saas.api.admin.resourcev1.User.user_status:type_name -> saas.api.admin.enumv1.UserStatusEnum.UserStatus
	10, // 2: saas.api.admin.resourcev1.UserListReq.page_request:type_name -> kit.page.pagepkg.PageRequest
	7,  // 3: saas.api.admin.resourcev1.UserListResp.metadata:type_name -> saas.api.admin.resourcev1.UserListResp.MetadataEntry
	5,  // 4: saas.api.admin.resourcev1.UserListResp.data:type_name -> saas.api.admin.resourcev1.UserListRespData
	0,  // 5: saas.api.admin.resourcev1.UserListRespData.list:type_name -> saas.api.admin.resourcev1.User
	11, // 6: saas.api.admin.resourcev1.UserListRespData.page_info:type_name -> kit.page.pagepkg.PageResponse
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_admin_service_v1_resources_user_resource_v1_proto_init() }
func file_api_admin_service_v1_resources_user_resource_v1_proto_init() {
	if File_api_admin_service_v1_resources_user_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdReq); i {
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
		file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdsReq); i {
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
		file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReq); i {
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
		file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResp); i {
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
		file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListRespData); i {
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
		file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProcessResult); i {
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
			RawDescriptor: file_api_admin_service_v1_resources_user_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_admin_service_v1_resources_user_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_admin_service_v1_resources_user_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_admin_service_v1_resources_user_resource_v1_proto_msgTypes,
	}.Build()
	File_api_admin_service_v1_resources_user_resource_v1_proto = out.File
	file_api_admin_service_v1_resources_user_resource_v1_proto_rawDesc = nil
	file_api_admin_service_v1_resources_user_resource_v1_proto_goTypes = nil
	file_api_admin_service_v1_resources_user_resource_v1_proto_depIdxs = nil
}
