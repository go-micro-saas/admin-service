package bo

import (
	errorv1 "github.com/go-micro-saas/account-service/api/account-service/v1/errors"
	passwordpkg "github.com/ikaiguang/go-srv-kit/kit/password"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
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
