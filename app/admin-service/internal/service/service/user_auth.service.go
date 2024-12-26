package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/enums"
	errorv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/errors"
	resourcev1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/resources"
	servicev1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/services"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/repo"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/service/dto"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
	"sync/atomic"
)

// userAuth ...
type userAuth struct {
	servicev1.UnimplementedSrvUserAuthV1Server

	log                    *log.Helper
	userAuthBizRepo        bizrepos.UserAuthBizRepo
	sendEmailCodeBizRepo   bizrepos.SendEmailCodeBizRepo
	sendEmailCodeEventRepo bizrepos.SendEmailCodeEventRepo

	isConsumingSendEmailCodeEvent atomic.Int64
}

// NewUserAuthService ...
func NewUserAuthService(
	logger log.Logger,
	userAuthBizRepo bizrepos.UserAuthBizRepo,
	sendEmailCodeBizRepo bizrepos.SendEmailCodeBizRepo,
	sendEmailCodeEventRepo bizrepos.SendEmailCodeEventRepo,
) servicev1.SrvUserAuthV1Server {
	return &userAuth{
		log:                    log.NewHelper(log.With(logger, "module", "admin-service/service/user_auth")),
		userAuthBizRepo:        userAuthBizRepo,
		sendEmailCodeBizRepo:   sendEmailCodeBizRepo,
		sendEmailCodeEventRepo: sendEmailCodeEventRepo,
	}
}

// Ping ping pong
func (s *userAuth) Ping(ctx context.Context, in *resourcev1.PingReq) (out *resourcev1.PingResp, err error) {
	// 可以解析
	authClaims, tokenClaimsOK := authpkg.GetAuthClaimsFromContext(ctx)
	s.log.WithContext(ctx).Infow(
		"tokenClaimsOK", tokenClaimsOK,
		"authClaims", authClaims,
	)

	out = &resourcev1.PingResp{
		Message: in.Message,
	}
	return out, err
}

// LoginByEmail Email登录
func (s *userAuth) LoginByEmail(ctx context.Context, in *resourcev1.LoginByEmailReq) (*resourcev1.LoginResp, error) {
	param := dto.AccountDto.ToBoLoginByEmailParam(in)
	param.SkipVerifyCode = in.Code == ""
	userModel, signResp, err := s.userAuthBizRepo.LoginByEmail(ctx, param)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

// LoginByPhone 手机登录
func (s *userAuth) LoginByPhone(ctx context.Context, in *resourcev1.LoginByPhoneReq) (*resourcev1.LoginResp, error) {
	param := dto.AccountDto.ToBoLoginByPhoneParam(in)
	param.SkipVerifyCode = in.Code == ""
	userModel, signResp, err := s.userAuthBizRepo.LoginByPhone(ctx, param)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

// RefreshToken 刷新token
func (s *userAuth) RefreshToken(ctx context.Context, in *resourcev1.RefreshTokenReq) (*resourcev1.LoginResp, error) {
	signResp, err := s.userAuthBizRepo.RefreshToken(ctx, in.RefreshToken)
	if err != nil {
		return nil, err
	}

	// user
	userId := signResp.TokenResponse.RefreshTokenItem.Payload.UserID
	userModel, err := s.userAuthBizRepo.CheckAndGetUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

func (s *userAuth) SendPhoneSignupCode(ctx context.Context, req *resourcev1.SendPhoneVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_SIGNUP_BY_PHONE
	return s.sendPhoneCode(ctx, param)
}

func (s *userAuth) SendEmailSignupCode(ctx context.Context, req *resourcev1.SendEmailVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam2(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_SIGNUP_BY_EMAIL
	return s.sendEmailCode(ctx, param)
}

// SignupByEmail 身份验证-Email注册
func (s *userAuth) SignupByEmail(ctx context.Context, req *resourcev1.SignupByEmailReq) (*resourcev1.LoginResp, error) {
	param := dto.AccountDto.ToBoSignupByEmailParam(req)
	userModel, signResp, err := s.userAuthBizRepo.SignupByEmail(ctx, param)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

// SignupByPhone 身份验证-手机注册
func (s *userAuth) SignupByPhone(ctx context.Context, req *resourcev1.SignupByPhoneReq) (*resourcev1.LoginResp, error) {
	param := dto.AccountDto.ToBoSignupByPhoneParam(req)
	userModel, signResp, err := s.userAuthBizRepo.SignupByPhone(ctx, param)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

func (s *userAuth) LoginOrSignupByPhone(ctx context.Context, req *resourcev1.LoginOrSignupByPhoneReq) (*resourcev1.LoginResp, error) {
	verifyParam := &bo.ConfirmVerifyCodeParam{
		VerifyAccount: req.Phone,
		VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_PHONE,
		VerifyCode:    req.Code,
	}
	err := s.userAuthBizRepo.ConfirmVerifyCode(ctx, verifyParam)
	if err != nil {
		return nil, err
	}

	// exist
	regModel, isExist, err := s.userAuthBizRepo.IsExistRegisterPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if isExist {
		return s.loginByUserID(ctx, regModel.UserId)
	}

	// signup
	passwd := po.RandomPassword()
	param := &bo.SignupByPhoneParam{
		Phone:           req.Phone,
		Password:        passwd,
		PasswordConfirm: passwd,
		Code:            req.Code,
		SkipVerifyCode:  true,
	}
	userModel, signResp, err := s.userAuthBizRepo.SignupByPhone(ctx, param)
	if err != nil {
		return nil, err
	}
	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

func (s *userAuth) LoginOrSignupByEmail(ctx context.Context, req *resourcev1.LoginOrSignupByEmailReq) (*resourcev1.LoginResp, error) {
	verifyParam := &bo.ConfirmVerifyCodeParam{
		VerifyAccount: req.Email,
		VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_EMAIL,
		VerifyCode:    req.Code,
	}
	err := s.userAuthBizRepo.ConfirmVerifyCode(ctx, verifyParam)
	if err != nil {
		return nil, err
	}

	// exist
	regModel, isExist, err := s.userAuthBizRepo.IsExistRegisterEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return s.loginByUserID(ctx, regModel.UserId)
	}

	// signup
	passwd := po.RandomPassword()
	param := &bo.SignupByEmailParam{
		Email:           req.Email,
		Password:        passwd,
		PasswordConfirm: passwd,
		Code:            req.Code,
		SkipVerifyCode:  true,
	}
	userModel, signResp, err := s.userAuthBizRepo.SignupByEmail(ctx, param)
	if err != nil {
		return nil, err
	}
	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

func (s *userAuth) loginByUserID(ctx context.Context, userID uint64) (*resourcev1.LoginResp, error) {
	param := &bo.LoginParam{
		SkipValidatePassword: true,
	}
	userModel, signResp, err := s.userAuthBizRepo.LoginByUserID(ctx, userID, param)
	if err != nil {
		return nil, err
	}

	out := &resourcev1.LoginResp{
		Data: dto.AccountDto.ToPbLoginRespData(userModel, signResp),
	}
	return out, nil
}

func (s *userAuth) LoginByOpenApi(ctx context.Context, req *resourcev1.LoginByOpenApiReq) (*resourcev1.LoginResp, error) {
	return s.UnimplementedSrvUserAuthV1Server.LoginByOpenApi(ctx, req)
}

func (s *userAuth) SubscribeSendEmailCodeEvent(ctx context.Context, req *resourcev1.SubscribeSendEmailCodeEventReq) (*resourcev1.SubscribeSendEmailCodeEventResp, error) {
	if s.isConsumingSendEmailCodeEvent.Load() < 1 {
		threadpkg.GoSafe(func() {
			s.isConsumingSendEmailCodeEvent.Add(1)
			defer func() { s.isConsumingSendEmailCodeEvent.Add(-1) }()
			err := s.sendEmailCodeEventRepo.Consume(ctx, s.sendEmailCodeBizRepo.SendEmailCode)
			if err != nil {
				s.log.WithContext(ctx).Errorw("msg", "SubscribeSendEmailCodeEvent failed!", "err", err)
			}
		})
	}
	return &resourcev1.SubscribeSendEmailCodeEventResp{
		Data: &resourcev1.SubscribeSendEmailCodeEventRespData{
			ConsumerCounter: s.isConsumingSendEmailCodeEvent.Load(),
		},
	}, nil
}

func (s *userAuth) StopSendEmailCodeEvent(ctx context.Context, req *resourcev1.StopSendEmailCodeEventReq) (*resourcev1.StopSendEmailCodeEventResp, error) {
	err := s.sendEmailCodeEventRepo.Close(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg", "run StopSendEmailCodeEvent failed!", "err", err)
		return nil, err
	}
	return &resourcev1.StopSendEmailCodeEventResp{
		Data: &resourcev1.StopSendEmailCodeEventRespData{
			ConsumerCounter: s.isConsumingSendEmailCodeEvent.Load(),
		},
	}, nil
}

func (s *userAuth) sendEmailCode(ctx context.Context, param *bo.SendVerifyCodeParam) (*resourcev1.SendVerifyCodeResp, error) {
	// send
	dataModel, err := s.userAuthBizRepo.SendVerifyCode(ctx, param)
	if err != nil {
		return nil, err
	}

	// mq
	sendEmailParam := dto.AccountDto.ToBoSendEmailCodeParam(dataModel)
	if err := s.sendEmailCodeEventRepo.Publish(ctx, sendEmailParam); err != nil {
		return nil, err
	}
	dataModel.IsSendToMQ = true

	return &resourcev1.SendVerifyCodeResp{
		Data: dto.AccountDto.ToPbSendSignupCodeRespData(dataModel),
	}, nil
}

func (s *userAuth) sendPhoneCode(ctx context.Context, param *bo.SendVerifyCodeParam) (*resourcev1.SendVerifyCodeResp, error) {
	return nil, errorpkg.WithStack(errorpkg.DefaultErrorMethodNotAllowed())

	dataModel, err := s.userAuthBizRepo.SendVerifyCode(ctx, param)
	if err != nil {
		return nil, err
	}

	return &resourcev1.SendVerifyCodeResp{
		Data: dto.AccountDto.ToPbSendSignupCodeRespData(dataModel),
	}, nil
}

func (s *userAuth) SendPhoneLoginCode(ctx context.Context, req *resourcev1.SendPhoneVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_LOGIN_BY_PHONE
	return s.sendPhoneCode(ctx, param)
}

func (s *userAuth) SendEmailLoginCode(ctx context.Context, req *resourcev1.SendEmailVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam2(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_LOGIN_BY_EMAIL
	return s.sendEmailCode(ctx, param)
}

func (s *userAuth) SendPhoneForgotPasswdCode(ctx context.Context, req *resourcev1.SendPhoneVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_FORGOT_PASSWORD_BY_PHONE
	return s.sendPhoneCode(ctx, param)
}

func (s *userAuth) SendEmailForgotPasswdCode(ctx context.Context, req *resourcev1.SendEmailVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam2(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_FORGOT_PASSWORD_BY_EMAIL
	return s.sendEmailCode(ctx, param)
}

func (s *userAuth) SendPhoneChangePasswdCode(ctx context.Context, req *resourcev1.SendPhoneVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_CHANGE_PASSWORD_BY_PHONE
	return s.sendPhoneCode(ctx, param)
}

func (s *userAuth) SendEmailChangePasswdCode(ctx context.Context, req *resourcev1.SendEmailVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam2(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_CHANGE_PASSWORD_BY_EMAIL
	return s.sendEmailCode(ctx, param)
}

func (s *userAuth) SendPhoneChangeCode(ctx context.Context, req *resourcev1.SendPhoneVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_CHANGE_PHONE
	return s.sendPhoneCode(ctx, param)
}

func (s *userAuth) SendEmailChangeCode(ctx context.Context, req *resourcev1.SendEmailVerifyCodeReq) (*resourcev1.SendVerifyCodeResp, error) {
	_, isExist, err := s.userAuthBizRepo.IsExistRegisterEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errorpkg.WithStack(errorv1.DefaultErrorS104UserNotExist())
	}

	// send
	param := dto.AccountDto.ToBoSendVerifyCodeParam2(req)
	param.VerifyType = enumv1.UserVerifyTypeEnum_CHANGE_EMAIL
	return s.sendEmailCode(ctx, param)
}
