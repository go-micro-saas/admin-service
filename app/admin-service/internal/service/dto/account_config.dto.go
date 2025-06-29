package dto

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/conf"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	snowflakeapi "github.com/go-micro-saas/service-api/app/snowflake-service"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	emailpkg "github.com/ikaiguang/go-srv-kit/kit/email"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func GetNodeIDOptions(logger log.Logger, cfg *conf.ServiceConfig) []snowflakeapi.Option {
	opts := snowflakeapi.DefaultOptions(logger)
	if cfg.GetAdminService().GetSnowflake().GetEnable() {
		opts = append(opts, snowflakeapi.WithMustGetNodeIdFromAPI(true))
	}
	return opts
}

func ToPbGetNodeIdReq(cfg *conf.ServiceConfig) (*nodeidresourcev1.GetNodeIdReq, error) {
	snowflakeConf := cfg.GetAdminService().GetSnowflake()
	if snowflakeConf == nil {
		e := errorpkg.ErrorInvalidParameter("snowflake config is nil")
		return nil, errorpkg.WithStack(e)
	}
	if err := snowflakeConf.Validate(); err != nil {
		e := errorpkg.ErrorInvalidParameter(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	res := &nodeidresourcev1.GetNodeIdReq{
		InstanceId:   snowflakeConf.InstanceId,
		InstanceName: snowflakeConf.InstanceName,
		Metadata:     snowflakeConf.Metadata,
	}
	return res, nil
}

func ToBoSendEmailCodeConfig(cfg *conf.ServiceConfig) (*bo.SendEmailCodeConfig, error) {
	if !cfg.GetAdminService().GetSendEmailCode().GetEnable() {
		return &bo.SendEmailCodeConfig{}, nil
	}
	emailConf := cfg.GetAdminService().GetSendEmailCode()
	if emailConf == nil {
		e := errorpkg.ErrorInvalidParameter("send_email_code config is nil")
		return nil, errorpkg.WithStack(e)
	}
	if err := emailConf.Validate(); err != nil {
		e := errorpkg.ErrorInvalidParameter(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	res := &bo.SendEmailCodeConfig{
		Enable: emailConf.GetEnable(),
		Sender: emailpkg.Sender{
			Issuer:   emailConf.GetIssuer(),
			Host:     emailConf.GetHost(),
			Port:     int(emailConf.GetPort()),
			Username: emailConf.GetUsername(),
			Password: emailConf.GetPassword(),
		},
		Message: emailpkg.Message{
			From:    emailConf.GetFrom(),
			To:      nil,
			Cc:      "",
			Subject: emailConf.GetSubject(),
			Body:    "",
		},
	}
	return res, nil
}

func ToBoVerifyCodeConfig(cfg *configpb.Bootstrap) *bo.VerifyCodeConfig {
	res := &bo.VerifyCodeConfig{
		CaptchaLen: po.DefaultVerifyCodeLength,
		CaptchaTTL: po.DefaultVerifyCodeExpiredTime,
	}
	captchaConfig := cfg.GetSetting().GetCaptcha()
	if captchaConfig == nil {
		return res
	}
	if captchaConfig.GetCaptchaLen() > 0 {
		res.CaptchaLen = int(captchaConfig.GetCaptchaLen())
	}
	if captchaConfig.GetCaptchaTtl().AsDuration() > 0 {
		res.CaptchaTTL = captchaConfig.GetCaptchaTtl().AsDuration()
	}
	return res
}
