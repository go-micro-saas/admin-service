package accountapi

import (
	servicev1 "github.com/go-micro-saas/account-service/api/account-service/v1/services"
	_ "github.com/gorilla/websocket"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[servicev1.OperationSrvUserAuthV1Ping] = middlewareutil.TransportServiceKindALL

	whiteList[servicev1.OperationSrvUserAuthV1SendEmailSignupCode] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvUserAuthV1SendPhoneSignupCode] = middlewareutil.TransportServiceKindALL

	whiteList[servicev1.OperationSrvUserAuthV1SignupByEmail] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvUserAuthV1SignupByPhone] = middlewareutil.TransportServiceKindALL

	whiteList[servicev1.OperationSrvUserAuthV1LoginOrSignupByPhone] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvUserAuthV1LoginOrSignupByEmail] = middlewareutil.TransportServiceKindALL

	whiteList[servicev1.OperationSrvUserAuthV1LoginByEmail] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvUserAuthV1LoginByOpenApi] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvUserAuthV1LoginByPhone] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvUserAuthV1RefreshToken] = middlewareutil.TransportServiceKindALL

	return whiteList
}
