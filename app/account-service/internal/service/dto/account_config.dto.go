package dto

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/conf"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	snowflakeapi "github.com/go-micro-saas/service-api/app/snowflake-service"
	emailpkg "github.com/ikaiguang/go-srv-kit/kit/email"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func GetNodeIDOptions(logger log.Logger, cfg *conf.ServiceConfig) []snowflakeapi.Option {
	opts := snowflakeapi.DefaultOptions(logger)
	if cfg.GetAccountService().GetSnowflake().GetEnable() {
		opts = append(opts, snowflakeapi.WithMustGetNodeIdFromAPI(true))
	}
	return opts
}

func ToPbGetNodeIdReq(cfg *conf.ServiceConfig) (*nodeidresourcev1.GetNodeIdReq, error) {
	snowflakeConf := cfg.GetAccountService().GetSnowflake()
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
	if !cfg.GetAccountService().GetSendEmailCode().GetEnable() {
		return &bo.SendEmailCodeConfig{}, nil
	}
	emailConf := cfg.GetAccountService().GetSendEmailCode()
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
