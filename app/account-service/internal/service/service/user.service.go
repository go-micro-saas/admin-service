package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	servicev1 "github.com/go-micro-saas/account-service/api/account-service/v1/services"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
)

type accountService struct {
	servicev1.UnimplementedSrvAccountV1Server

	log                    *log.Helper
	userAuthBizRepo        bizrepos.UserAuthBizRepo
	sendEmailCodeBizRepo   bizrepos.SendEmailCodeBizRepo
	sendEmailCodeEventRepo bizrepos.SendEmailCodeEventRepo
}

func (a accountService) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) GetUserInfo(ctx context.Context, req *resourcev1.GetUserInfoReq) (*resourcev1.GetUserInfoResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) GetUserInfoList(ctx context.Context, req *resourcev1.GetUserListReq) (*resourcev1.GetUserListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) CreateUser(ctx context.Context, req *resourcev1.CreateUserReq) (*resourcev1.CreateUserResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) CreateUserByPhone(ctx context.Context, req *resourcev1.CreateUserByPhoneReq) (*resourcev1.CreateUserResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) CreateUserByEmail(ctx context.Context, req *resourcev1.CreateUserByEmailReq) (*resourcev1.CreateUserResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) mustEmbedUnimplementedSrvAccountV1Server() {
	//TODO implement me
	panic("implement me")
}
