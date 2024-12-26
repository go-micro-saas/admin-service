// Package po
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package po

import (
	enumv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/enums"
	randompkg "github.com/ikaiguang/go-srv-kit/kit/random"
	datatypes "gorm.io/datatypes"
	"strconv"
	time "time"
)

var _ = time.Time{}

var _ = datatypes.JSON{}

// User ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户表'
type User struct {
	Id            uint64                                       `gorm:"column:id;primaryKey" json:"id"`              // ID
	CreatedTime   time.Time                                    `gorm:"column:created_time" json:"created_time"`     // 创建时间
	UpdatedTime   time.Time                                    `gorm:"column:updated_time" json:"updated_time"`     // 最后修改时间
	DeletedTime   uint64                                       `gorm:"column:deleted_time" json:"deleted_time"`     // 删除时间
	UserId        uint64                                       `gorm:"column:user_id" json:"user_id"`               // uid
	UserPhone     string                                       `gorm:"column:user_phone" json:"user_phone"`         // 手机
	UserEmail     string                                       `gorm:"column:user_email" json:"user_email"`         // 邮箱
	UserNickname  string                                       `gorm:"column:user_nickname" json:"user_nickname"`   // 昵称
	UserAvatar    string                                       `gorm:"column:user_avatar" json:"user_avatar"`       // 头像
	UserGender    enumv1.UserGenderEnum_UserGender             `gorm:"column:user_gender" json:"user_gender"`       // 性别；0：INIT，1：MALE，2：FEMALE，3：SECRET
	RegisterType  enumv1.UserRegisterTypeEnum_UserRegisterType `gorm:"column:register_type" json:"register_type"`   // 注册类型；0：INIT，1：EMAIL，2：MOBILE，3：。。。参考ENUM定义
	UserStatus    enumv1.UserStatusEnum_UserStatus             `gorm:"column:user_status" json:"user_status"`       // 状态；0：INIT，1：ENABLE，2：DISABLE，3：WHITELIST，4：BLACKLIST，5：DELETED
	DisableTime   uint64                                       `gorm:"column:disable_time" json:"disable_time"`     // 禁用时间
	BlacklistTime uint64                                       `gorm:"column:blacklist_time" json:"blacklist_time"` // 黑名单时间
	PasswordHash  string                                       `gorm:"column:password_hash" json:"password_hash"`   // 密码HASH
}

func (s *User) IsValidStatus() bool {
	switch s.UserStatus {
	default:
		return true
	case enumv1.UserStatusEnum_DISABLE, enumv1.UserStatusEnum_BLACKLIST, enumv1.UserStatusEnum_DELETED:
		return false
	}
}

func RandomNickname() string {
	return "u_" + strconv.FormatInt(time.Now().UnixNano(), 36)
}

func RandomPassword() string {
	return randompkg.AlphabetLower(32)
}

func NewUser() *User {
	var (
		now = time.Now()
	)
	userModel := &User{
		Id:            0,
		CreatedTime:   now,
		UpdatedTime:   now,
		DeletedTime:   0,
		UserId:        0,
		UserPhone:     "",
		UserEmail:     "",
		UserNickname:  RandomNickname(),
		UserAvatar:    "",
		UserGender:    enumv1.UserGenderEnum_UNSPECIFIED,
		RegisterType:  enumv1.UserRegisterTypeEnum_UNSPECIFIED,
		UserStatus:    enumv1.UserStatusEnum_ENABLE,
		DisableTime:   0,
		BlacklistTime: 0,
		PasswordHash:  "",
	}
	return userModel
}

func NewUserByPhone(phone, passwdHash string) *User {
	userModel := NewUser()
	userModel.UserPhone = phone
	userModel.PasswordHash = passwdHash
	userModel.RegisterType = enumv1.UserRegisterTypeEnum_PHONE
	return userModel
}

func NewUserByEmail(email, passwdHash string) *User {
	userModel := NewUser()
	userModel.UserEmail = email
	userModel.PasswordHash = passwdHash
	userModel.RegisterType = enumv1.UserRegisterTypeEnum_EMAIL
	return userModel
}
