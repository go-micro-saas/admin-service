// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errorv1

import (
	errors "github.com/go-kratos/kratos/v2/errors"
	strconv "strconv"
)

var ERROR_http_code = map[string]int{

	"UNKNOWN":                         500,
	"S104_USER_NOT_EXIST":             400,
	"S104_USER_EXIST":                 400,
	"S104_USER_NAME_INVALID":          400,
	"S104_USER_NAME_EXIST":            400,
	"S104_USER_PASSWORD_INVALID":      400,
	"S104_USER_PASSWORD_INCORRECT":    400,
	"S104_USER_TOKEN_INVALID":         400,
	"S104_USER_ACCOUNT_EXPIRE":        400,
	"S104_USER_STATUS_NOT_ALLOW":      400,
	"S104_USER_ACTIVE_TIME_INVALID":   400,
	"S104_INVALID_PHONE":              400,
	"S104_INVALID_EMAIL":              400,
	"S104_PASSWORD_NOT_MATCH_CONFIRM": 400,
	"S104_PASSWORD_NOT_MATCH_RULE":    400,
	"S104_PASSWORD_INCORRECT":         400,
	"S104_VERIFY_CODE_INCORRECT":      400,
}

func (x ERROR) HTTPCode() int {
	if v, ok := ERROR_http_code[x.String()]; ok {
		return v
	}
	return 500
}

// 未知错误
func DefaultErrorUnknown() *errors.Error {
	e := errors.New(500, ERROR_UNKNOWN.String(), "未知错误")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_UNKNOWN.Number()))}
	return e
}

// 用户不存在
func DefaultErrorS104UserNotExist() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_NOT_EXIST.String(), "用户不存在")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_NOT_EXIST.Number()))}
	return e
}

// 用户已存在
func DefaultErrorS104UserExist() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_EXIST.String(), "用户已存在")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_EXIST.Number()))}
	return e
}

// 用户名不合法
func DefaultErrorS104UserNameInvalid() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_NAME_INVALID.String(), "用户名不合法")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_NAME_INVALID.Number()))}
	return e
}

// 用户名已存在
func DefaultErrorS104UserNameExist() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_NAME_EXIST.String(), "用户名已存在")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_NAME_EXIST.Number()))}
	return e
}

// 用户密码不合法
func DefaultErrorS104UserPasswordInvalid() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_PASSWORD_INVALID.String(), "用户密码不合法")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_PASSWORD_INVALID.Number()))}
	return e
}

// 用户密码不正确
func DefaultErrorS104UserPasswordIncorrect() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_PASSWORD_INCORRECT.String(), "用户密码不正确")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_PASSWORD_INCORRECT.Number()))}
	return e
}

// 令牌已失效
func DefaultErrorS104UserTokenInvalid() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_TOKEN_INVALID.String(), "令牌已失效")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_TOKEN_INVALID.Number()))}
	return e
}

// 账户已过期
func DefaultErrorS104UserAccountExpire() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_ACCOUNT_EXPIRE.String(), "账户已过期")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_ACCOUNT_EXPIRE.Number()))}
	return e
}

// 无效的登录状态
func DefaultErrorS104UserStatusNotAllow() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_STATUS_NOT_ALLOW.String(), "无效的登录状态")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_STATUS_NOT_ALLOW.Number()))}
	return e
}

// 不在有效的激活期间
func DefaultErrorS104UserActiveTimeInvalid() *errors.Error {
	e := errors.New(400, ERROR_S104_USER_ACTIVE_TIME_INVALID.String(), "不在有效的激活期间")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_USER_ACTIVE_TIME_INVALID.Number()))}
	return e
}

// 无效的手机号
func DefaultErrorS104InvalidPhone() *errors.Error {
	e := errors.New(400, ERROR_S104_INVALID_PHONE.String(), "无效的手机号")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_INVALID_PHONE.Number()))}
	return e
}

// 无效的邮箱
func DefaultErrorS104InvalidEmail() *errors.Error {
	e := errors.New(400, ERROR_S104_INVALID_EMAIL.String(), "无效的邮箱")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_INVALID_EMAIL.Number()))}
	return e
}

// 密码不匹配
func DefaultErrorS104PasswordNotMatchConfirm() *errors.Error {
	e := errors.New(400, ERROR_S104_PASSWORD_NOT_MATCH_CONFIRM.String(), "密码不匹配")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_PASSWORD_NOT_MATCH_CONFIRM.Number()))}
	return e
}

// 密码不符合规则
func DefaultErrorS104PasswordNotMatchRule() *errors.Error {
	e := errors.New(400, ERROR_S104_PASSWORD_NOT_MATCH_RULE.String(), "密码不符合规则")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_PASSWORD_NOT_MATCH_RULE.Number()))}
	return e
}

// 密码不正确
func DefaultErrorS104PasswordIncorrect() *errors.Error {
	e := errors.New(400, ERROR_S104_PASSWORD_INCORRECT.String(), "密码不正确")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_PASSWORD_INCORRECT.Number()))}
	return e
}

// 验证码不正确
func DefaultErrorS104VerifyCodeIncorrect() *errors.Error {
	e := errors.New(400, ERROR_S104_VERIFY_CODE_INCORRECT.String(), "验证码不正确")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S104_VERIFY_CODE_INCORRECT.Number()))}
	return e
}
