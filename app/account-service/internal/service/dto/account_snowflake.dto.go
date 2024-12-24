package dto

import (
	"github.com/go-micro-saas/account-service/app/account-service/internal/conf"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

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
