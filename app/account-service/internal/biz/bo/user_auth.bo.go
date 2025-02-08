package bo

import (
	"encoding/json"
	enumv1 "github.com/go-micro-saas/account-service/api/account-service/v1/enums"
	errorv1 "github.com/go-micro-saas/account-service/api/account-service/v1/errors"
	passwordpkg "github.com/ikaiguang/go-srv-kit/kit/password"
	regexpkg "github.com/ikaiguang/go-srv-kit/kit/regex"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"strings"
)

type LoginParam struct {
	SkipValidatePassword bool
	PlaintextPassword    string // md5(plaintext)
}

type SignTokenReq struct {
	authpkg.Claims
}

func (s *SignTokenReq) SetByAuthClaims(authClaims *authpkg.Claims) {
	s.Claims = *authClaims
}

type SignTokenResp struct {
	authpkg.TokenResponse
}

func (s *SignTokenResp) SetByAuthTokenResponse(authResp *authpkg.TokenResponse) {
	s.TokenResponse = *authResp
}

type SignupByPhoneParam struct {
	Phone           string
	Password        string
	PasswordConfirm string
	Code            string
	SkipVerifyCode  bool
}

type SignupByEmailParam struct {
	Email           string
	Password        string
	PasswordConfirm string
	Code            string
	SkipVerifyCode  bool
}

type LoginByPhoneParam struct {
	Phone          string
	Password       string
	Code           string
	SkipVerifyCode bool
}

type LoginByEmailParam struct {
	Email          string
	Password       string
	Code           string
	SkipVerifyCode bool
}

type PasswordParam struct {
	Password        string
	PasswordConfirm string
}

func (s *PasswordParam) ValidateAndEncrypt() (string, error) {
	if s.Password != s.PasswordConfirm {
		e := errorv1.DefaultErrorS103PasswordNotMatchConfirm()
		return "", errorpkg.WithStack(e)
	}
	if len(s.Password) < 6 {
		e := errorv1.DefaultErrorS103PasswordNotMatchRule()
		return "", errorpkg.WithStack(e)
	}
	hash, err := passwordpkg.Encrypt(s.Password)
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return "", errorpkg.WithStack(e)
	}
	return string(hash), nil
}

type SendVerifyCodeParam struct {
	VerifyAccount string // 用户标识；手机、邮箱、。。。
	VerifyType    enumv1.UserVerifyTypeEnum_UserVerifyType
}

func (s *SendVerifyCodeParam) Validate() error {
	s.VerifyAccount = strings.TrimSpace(s.VerifyAccount)
	switch s.VerifyType {
	default:
		if s.VerifyAccount == "" {
			e := errorpkg.ErrorInvalidParameter("account is empty")
			return errorpkg.WithStack(e)
		}
	case enumv1.UserVerifyTypeEnum_SIGNUP_BY_PHONE:
		if !regexpkg.IsPhone(s.VerifyAccount) {
			e := errorv1.DefaultErrorS103InvalidPhone()
			return errorpkg.WithStack(e)
		}
	case enumv1.UserVerifyTypeEnum_SIGNUP_BY_EMAIL:
		if !regexpkg.IsEmail(s.VerifyAccount) {
			e := errorv1.DefaultErrorS103InvalidEmail()
			return errorpkg.WithStack(e)
		}
	}
	return nil
}

type SendVerifyCodeReply struct {
	IsSendToMQ    bool   // 是否发送队列
	VerifyAccount string // 用户标识；手机、邮箱、。。。
	VerifyType    enumv1.UserVerifyTypeEnum_UserVerifyType
	VerifyCode    string // code
}

type SendEmailCodeParam struct {
	VerifyAccount string                                   `json:"verify_account"` // 用户标识；手机、邮箱、。。。
	VerifyType    enumv1.UserVerifyTypeEnum_UserVerifyType `json:"verify_type"`    // 验证类型
	VerifyCode    string                                   `json:"verify_code"`    // 验证码
}

func (s *SendEmailCodeParam) MarshalToJSON() ([]byte, error) {
	buf, err := json.Marshal(s)
	if err != nil {
		return nil, errorpkg.WithStack(errorpkg.ErrorInternalServer(err.Error()))
	}
	return buf, nil
}

func (s *SendEmailCodeParam) UnmarshalFromJSON(buf []byte) error {
	err := json.Unmarshal(buf, s)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalServer(err.Error()))
	}
	return nil
}

type SendEmailCodeReply struct {
	IsSendToServer bool
	Code           string
}

type ConfirmVerifyCodeParam struct {
	VerifyAccount string                                   // 用户标识；手机、邮箱、。。。
	VerifyType    enumv1.UserVerifyTypeEnum_UserVerifyType //
	VerifyCode    string                                   // 验证码
}

func (s *ConfirmVerifyCodeParam) Validate() error {
	s.VerifyAccount = strings.TrimSpace(s.VerifyAccount)
	s.VerifyCode = strings.TrimSpace(s.VerifyCode)
	if s.VerifyAccount == "" {
		e := errorpkg.ErrorInvalidParameter("account is empty")
		return errorpkg.WithStack(e)
	}
	if s.VerifyCode == "" {
		e := errorpkg.ErrorInvalidParameter("code is empty")
		return errorpkg.WithStack(e)
	}
	return nil

}
