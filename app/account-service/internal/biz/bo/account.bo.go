package bo

import enums "github.com/go-micro-saas/account-service/api/account-service/v1/enums"

type CreateUserParam struct {
	UserPhone    string                          // 电话
	UserEmail    string                          // 邮箱
	UserNickname string                          // 昵称
	UserAvatar   string                          // 头像
	UserGender   enums.UserGenderEnum_UserGender // 性别；0：INIT，1：MALE，2：FEMALE，3：SECRET
	UserStatus   enums.UserStatusEnum_UserStatus // 状态；0：INIT，1：ENABLE，2：DISABLE，3：WHITELIST，4：BLACKLIST，5：DELETED
	Password     string                          // md5(密码)

}
