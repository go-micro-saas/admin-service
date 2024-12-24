//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	servicev1 "github.com/go-micro-saas/account-service/api/account-service/v1/services"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/biz"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/conf"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/data"
	datarepos "github.com/go-micro-saas/account-service/app/account-service/internal/data/repo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/service/dto"
	"github.com/go-micro-saas/account-service/app/account-service/internal/service/service"
	snowflakeapi "github.com/go-micro-saas/service-api/app/snowflake-service"
	"github.com/google/wire"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportUserDataRepo(launcherManager setuputil.LauncherManager) (datarepos.UserDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		setuputil.GetRecommendDBConn, data.NewUserDataRepo,
	))
	return nil, nil
}

func exportUserRegPhoneDataRepo(launcherManager setuputil.LauncherManager) (datarepos.UserRegPhoneDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		setuputil.GetRecommendDBConn, data.NewUserRegPhoneDataRepo,
	))
	return nil, nil
}

func exportUserRegEmailDataRepo(launcherManager setuputil.LauncherManager) (datarepos.UserRegEmailDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		setuputil.GetRecommendDBConn, data.NewUserRegEmailDataRepo,
	))
	return nil, nil
}

func exportUserConfirmCodeDataRepo(launcherManager setuputil.LauncherManager) (datarepos.UserConfirmCodeDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		setuputil.GetRecommendDBConn, data.NewUserConfirmCodeDataRepo,
	))
	return nil, nil
}

func exportIdGenerator(launcherManager setuputil.LauncherManager) (idpkg.Snowflake, func(), error) {
	panic(wire.Build(
		setuputil.GetLogger,
		setuputil.GetServiceAPIManager,
		// req
		conf.GetServiceConfig,
		dto.ToPbGetNodeIdReq,
		// opts
		snowflakeapi.DefaultOptions,
		// snowflake
		snowflakeapi.GetSingletonIDGeneratorByHTTPAPI,
	))
	return nil, nil, nil
}

func exportUserBizRepo(launcherManager setuputil.LauncherManager) (bizrepos.UserAuthBizRepo, func(), error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// authRepo
		setuputil.GetAuthManager,
		// idGenerator
		exportIdGenerator,
		// data
		exportUserDataRepo, exportUserRegPhoneDataRepo, exportUserRegEmailDataRepo, exportUserConfirmCodeDataRepo,
		// biz
		biz.NewUserAuthBiz,
	))
	return nil, nil, nil
}

func exportUserAuthV1Service(launcherManager setuputil.LauncherManager) (servicev1.SrvUserAuthV1Server, func(), error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// biz
		exportUserBizRepo,
		// service
		service.NewUserAuthService,
	))
	return nil, nil, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, func(), error) {
	panic(wire.Build(
		// service
		exportUserAuthV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil, nil
}
