// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.6
// source: app/admin-service/internal/conf/config.conf.proto

package conf

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminService *ServiceConfig_AdminService `protobuf:"bytes,1,opt,name=admin_service,json=adminService,proto3" json:"admin_service,omitempty"`
}

func (x *ServiceConfig) Reset() {
	*x = ServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig) ProtoMessage() {}

func (x *ServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig.ProtoReflect.Descriptor instead.
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return file_app_admin_service_internal_conf_config_conf_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceConfig) GetAdminService() *ServiceConfig_AdminService {
	if x != nil {
		return x.AdminService
	}
	return nil
}

type ServiceConfig_AdminService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snowflake     *ServiceConfig_AdminService_Snowflake     `protobuf:"bytes,1,opt,name=snowflake,proto3" json:"snowflake,omitempty"`
	SendEmailCode *ServiceConfig_AdminService_SendEmailCode `protobuf:"bytes,2,opt,name=send_email_code,json=sendEmailCode,proto3" json:"send_email_code,omitempty"`
}

func (x *ServiceConfig_AdminService) Reset() {
	*x = ServiceConfig_AdminService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig_AdminService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig_AdminService) ProtoMessage() {}

func (x *ServiceConfig_AdminService) ProtoReflect() protoreflect.Message {
	mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig_AdminService.ProtoReflect.Descriptor instead.
func (*ServiceConfig_AdminService) Descriptor() ([]byte, []int) {
	return file_app_admin_service_internal_conf_config_conf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServiceConfig_AdminService) GetSnowflake() *ServiceConfig_AdminService_Snowflake {
	if x != nil {
		return x.Snowflake
	}
	return nil
}

func (x *ServiceConfig_AdminService) GetSendEmailCode() *ServiceConfig_AdminService_SendEmailCode {
	if x != nil {
		return x.SendEmailCode
	}
	return nil
}

// Snowflake nodeid for snowflake
type ServiceConfig_AdminService_Snowflake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId   string            `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstanceName string            `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Metadata     map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceConfig_AdminService_Snowflake) Reset() {
	*x = ServiceConfig_AdminService_Snowflake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig_AdminService_Snowflake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig_AdminService_Snowflake) ProtoMessage() {}

func (x *ServiceConfig_AdminService_Snowflake) ProtoReflect() protoreflect.Message {
	mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig_AdminService_Snowflake.ProtoReflect.Descriptor instead.
func (*ServiceConfig_AdminService_Snowflake) Descriptor() ([]byte, []int) {
	return file_app_admin_service_internal_conf_config_conf_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ServiceConfig_AdminService_Snowflake) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *ServiceConfig_AdminService_Snowflake) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *ServiceConfig_AdminService_Snowflake) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// SendEmailCode 发送邮件验证码
type ServiceConfig_AdminService_SendEmailCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable   bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Issuer   string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject  string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Host     string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	From     string `protobuf:"bytes,8,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *ServiceConfig_AdminService_SendEmailCode) Reset() {
	*x = ServiceConfig_AdminService_SendEmailCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig_AdminService_SendEmailCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig_AdminService_SendEmailCode) ProtoMessage() {}

func (x *ServiceConfig_AdminService_SendEmailCode) ProtoReflect() protoreflect.Message {
	mi := &file_app_admin_service_internal_conf_config_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig_AdminService_SendEmailCode.ProtoReflect.Descriptor instead.
func (*ServiceConfig_AdminService_SendEmailCode) Descriptor() ([]byte, []int) {
	return file_app_admin_service_internal_conf_config_conf_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ServiceConfig_AdminService_SendEmailCode) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

var File_app_admin_service_internal_conf_config_conf_proto protoreflect.FileDescriptor

var file_app_admin_service_internal_conf_config_conf_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb6, 0x06, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0c, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0xce, 0x05, 0x0a, 0x0c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x09, 0x73, 0x6e,
	0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x52, 0x09, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c,
	0x61, 0x6b, 0x65, 0x12, 0x65, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x73,
	0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0d, 0x73, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0xfc, 0x01, 0x0a, 0x09, 0x53,
	0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x63, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x73, 0x61, 0x61, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6e, 0x6f, 0x77, 0x66,
	0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xfe, 0x01, 0x0a, 0x0d, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x1a, 0x06, 0x18, 0xff, 0xff, 0x03, 0x28, 0x01, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x76, 0x0a, 0x13, 0x73, 0x61,
	0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x42, 0x10, 0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_admin_service_internal_conf_config_conf_proto_rawDescOnce sync.Once
	file_app_admin_service_internal_conf_config_conf_proto_rawDescData = file_app_admin_service_internal_conf_config_conf_proto_rawDesc
)

func file_app_admin_service_internal_conf_config_conf_proto_rawDescGZIP() []byte {
	file_app_admin_service_internal_conf_config_conf_proto_rawDescOnce.Do(func() {
		file_app_admin_service_internal_conf_config_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_admin_service_internal_conf_config_conf_proto_rawDescData)
	})
	return file_app_admin_service_internal_conf_config_conf_proto_rawDescData
}

var file_app_admin_service_internal_conf_config_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_admin_service_internal_conf_config_conf_proto_goTypes = []interface{}{
	(*ServiceConfig)(nil),                            // 0: saas.api.admin.conf.ServiceConfig
	(*ServiceConfig_AdminService)(nil),               // 1: saas.api.admin.conf.ServiceConfig.AdminService
	(*ServiceConfig_AdminService_Snowflake)(nil),     // 2: saas.api.admin.conf.ServiceConfig.AdminService.Snowflake
	(*ServiceConfig_AdminService_SendEmailCode)(nil), // 3: saas.api.admin.conf.ServiceConfig.AdminService.SendEmailCode
	nil, // 4: saas.api.admin.conf.ServiceConfig.AdminService.Snowflake.MetadataEntry
}
var file_app_admin_service_internal_conf_config_conf_proto_depIdxs = []int32{
	1, // 0: saas.api.admin.conf.ServiceConfig.admin_service:type_name -> saas.api.admin.conf.ServiceConfig.AdminService
	2, // 1: saas.api.admin.conf.ServiceConfig.AdminService.snowflake:type_name -> saas.api.admin.conf.ServiceConfig.AdminService.Snowflake
	3, // 2: saas.api.admin.conf.ServiceConfig.AdminService.send_email_code:type_name -> saas.api.admin.conf.ServiceConfig.AdminService.SendEmailCode
	4, // 3: saas.api.admin.conf.ServiceConfig.AdminService.Snowflake.metadata:type_name -> saas.api.admin.conf.ServiceConfig.AdminService.Snowflake.MetadataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_app_admin_service_internal_conf_config_conf_proto_init() }
func file_app_admin_service_internal_conf_config_conf_proto_init() {
	if File_app_admin_service_internal_conf_config_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_admin_service_internal_conf_config_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig); i {
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
		file_app_admin_service_internal_conf_config_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig_AdminService); i {
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
		file_app_admin_service_internal_conf_config_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig_AdminService_Snowflake); i {
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
		file_app_admin_service_internal_conf_config_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig_AdminService_SendEmailCode); i {
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
			RawDescriptor: file_app_admin_service_internal_conf_config_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_admin_service_internal_conf_config_conf_proto_goTypes,
		DependencyIndexes: file_app_admin_service_internal_conf_config_conf_proto_depIdxs,
		MessageInfos:      file_app_admin_service_internal_conf_config_conf_proto_msgTypes,
	}.Build()
	File_app_admin_service_internal_conf_config_conf_proto = out.File
	file_app_admin_service_internal_conf_config_conf_proto_rawDesc = nil
	file_app_admin_service_internal_conf_config_conf_proto_goTypes = nil
	file_app_admin_service_internal_conf_config_conf_proto_depIdxs = nil
}
