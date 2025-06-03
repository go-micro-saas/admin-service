//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/biz"
	events "github.com/go-micro-saas/account-service/app/account-service/internal/biz/event"
	"github.com/go-micro-saas/account-service/app/account-service/internal/conf"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/data"
	"github.com/go-micro-saas/account-service/app/account-service/internal/service/dto"
	"github.com/go-micro-saas/account-service/app/account-service/internal/service/service"
	snowflakeapi "github.com/go-micro-saas/service-api/app/snowflake-service"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, func(), error) {
	panic(wire.Build(
		setuputil.GetLogger,
		setuputil.GetRecommendDBConn,
		setuputil.GetAuthManager,
		setuputil.GetServiceAPIManager,
		setuputil.GetRabbitmqConn,
		// conf
		conf.GetServiceConfig,
		dto.ToBoSendEmailCodeConfig,
		// idGenerator
		dto.ToPbGetNodeIdReq, dto.GetNodeIDOptions, snowflakeapi.GetSingletonIDGeneratorByHTTPAPI,
		// data
		data.NewUserDataRepo,
		data.NewUserRegPhoneDataRepo,
		data.NewUserRegEmailDataRepo,
		data.NewUserVerifyCodeRepo,
		data.NewUserEventHistoryRepo,
		// biz
		biz.NewUserAuthBiz,
		biz.NewSendEmailCodeBiz,
		biz.NewAccountBiz,
		// event
		events.NewSendEmailCodeEventRepo,
		// service
		service.NewUserAuthService,
		service.NewAccountService,
		// register services
		service.RegisterServices,
	))
	return nil, nil, nil
}
