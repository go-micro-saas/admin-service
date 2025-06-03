package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	servicev1 "github.com/go-micro-saas/account-service/api/account-service/v1/services"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	stdlog "log"
)

// RegisterServices 注册服务
// @return Services 用于wire
// @return func() = cleanup 关闭资源
// @return error 错误
func RegisterServices(
	hs *http.Server, gs *grpc.Server,
	userAuthV1Service servicev1.SrvUserAuthV1Server,
	accountService servicev1.SrvAccountV1Server,
) (cleanuputil.CleanupManager, error) {
	// 先进后出
	var cleanupManager = cleanuputil.NewCleanupManager()
	// grpc
	if gs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：GRPC: userAuthV1Service")
		servicev1.RegisterSrvUserAuthV1Server(gs, userAuthV1Service)
		servicev1.RegisterSrvAccountV1Server(gs, accountService)

		//cleanupManager.Append(cleanup)
	}

	// http
	if hs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：HTTP: userAuthV1Service")
		servicev1.RegisterSrvUserAuthV1HTTPServer(hs, userAuthV1Service)
		servicev1.RegisterSrvAccountV1HTTPServer(hs, accountService)

		// special
		//RegisterSpecialRouters(hs, homeService, websocketService)

		//cleanupManager.Append(cleanup)
	}

	// event
	stdlog.Println("|*** REGISTER_EVENT：SUBSCRIBE : SendEmailCodeEvent")
	_, err := userAuthV1Service.SubscribeSendEmailCodeEvent(context.Background(), &resourcev1.SubscribeSendEmailCodeEventReq{})
	if err != nil {
		return nil, err
	}
	cleanupManager.Append(func() {
		logpkg.Infow("msg", "StopSendEmailCodedEvent ...")
		_, err := userAuthV1Service.StopSendEmailCodedEvent(context.Background(), &resourcev1.StopSendEmailCodeEventReq{})
		if err != nil {
			logpkg.Warnw("msg", "StopSendEmailCodedEvent failed", "err", err)
		}
	})

	return cleanupManager, nil
}

//func RegisterSpecialRouters(hs *http.Server, homeService *HomeService, websocketService *WebsocketService) {
//	// router
//	router := mux.NewRouter()
//
//	stdlog.Println("|*** REGISTER_ROUTER：Root(/)")
//	router.HandleFunc("/", homeService.Homepage)
//	hs.Handle("/", router)
//
//	stdlog.Println("|*** REGISTER_ROUTER：Websocket")
//	router.HandleFunc("/ws/v1/testdata/websocket", websocketService.TestWebsocket)
//
//	// router
//	hs.Handle("/ws/v1/websocket", router)
//}
