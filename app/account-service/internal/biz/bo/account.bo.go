package bo

import (
	enums "github.com/go-micro-saas/account-service/api/account-service/v1/enums"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
)

type CreateUserParam struct {
	UserPhone    string                          // 电话
	UserEmail    string                          // 邮箱
	UserNickname string                          // 昵称
	UserAvatar   string                          // 头像
	UserGender   enums.UserGenderEnum_UserGender // 性别；0：INIT，1：MALE，2：FEMALE，3：SECRET
	UserStatus   enums.UserStatusEnum_UserStatus // 状态；0：INIT，1：ENABLE，2：DISABLE，3：WHITELIST，4：BLACKLIST，5：DELETED
	Password     string                          // md5(密码)
}

type UserListParam struct {
	UidList      []uint64 // 用户ID列表
	ContactPhone string   // 联系电话
	ContactEmail string   // 联系邮箱

	PaginatorArgs *gormpkg.PaginatorArgs
}
