package bo

import (
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
	VerifyType    enumv1.UserConfirmTypeEnum_UserConfirmType
}

func (s *SendVerifyCodeParam) Validate() error {
	s.VerifyAccount = strings.TrimSpace(s.VerifyAccount)
	switch s.VerifyType {
	default:
		if s.VerifyAccount == "" {
			e := errorpkg.ErrorInvalidParameter("account is empty")
			return errorpkg.WithStack(e)
		}
	case enumv1.UserConfirmTypeEnum_PHONE:
		if !regexpkg.IsPhone(s.VerifyAccount) {
			e := errorv1.DefaultErrorS103InvalidPhone()
			return errorpkg.WithStack(e)
		}
	case enumv1.UserConfirmTypeEnum_EMAIL:
		if !regexpkg.IsEmail(s.VerifyAccount) {
			e := errorv1.DefaultErrorS103InvalidEmail()
			return errorpkg.WithStack(e)
		}
	case enumv1.UserConfirmTypeEnum_PASSWORD:
		if s.VerifyAccount == "" {
			e := errorpkg.ErrorInvalidParameter("account is empty")
			return errorpkg.WithStack(e)
		}
	}
	return nil
}

type SendVerifyCodeReply struct {
	IsSendSuccess bool   // 是否发送成功
	Code          string // code
}
