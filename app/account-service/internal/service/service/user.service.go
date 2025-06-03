package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	servicev1 "github.com/go-micro-saas/account-service/api/account-service/v1/services"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/service/dto"
)

type accountService struct {
	servicev1.UnimplementedSrvAccountV1Server

	log         *log.Helper
	userBizRepo bizrepos.UserBizRepo
}

func NewAccountService(
	logger log.Logger,
	userBizRepo bizrepos.UserBizRepo,
) servicev1.SrvAccountV1Server {
	return &accountService{
		log:         log.NewHelper(log.With(logger, "module", "account-service/service/account")),
		userBizRepo: userBizRepo,
	}
}

func (s *accountService) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	return &resourcev1.PingResp{
		Data: &resourcev1.PingRespData{
			Message: req.Message,
		},
	}, nil
}

func (s *accountService) GetUserInfo(ctx context.Context, req *resourcev1.GetUserInfoReq) (*resourcev1.GetUserInfoResp, error) {
	userModel, err := s.userBizRepo.GetUserByUid(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetUserInfoResp{
		Data: dto.AccountDto.ToPbAccountInfo(userModel),
	}, nil
}

func (s *accountService) GetUserInfoList(ctx context.Context, req *resourcev1.GetUserListReq) (*resourcev1.GetUserListResp, error) {
	dataModels, err := s.userBizRepo.GetUsersByUidList(ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetUserListResp{
		Data: dto.AccountDto.ToPbAccountInfoList(dataModels),
	}, nil
}

func (s *accountService) CreateUser(ctx context.Context, req *resourcev1.CreateUserReq) (*resourcev1.CreateUserResp, error) {
	param := dto.AccountDto.ToBoCreateUserParam(req)
	userModel, err := s.userBizRepo.CreateUserByPhone(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateUserResp{
		Data: dto.AccountDto.ToPbAccountInfo(userModel),
	}, nil
}

func (s *accountService) CreateUserByPhone(ctx context.Context, req *resourcev1.CreateUserByPhoneReq) (*resourcev1.CreateUserResp, error) {
	param := dto.AccountDto.ToBoCreateUserParam3(req)
	userModel, err := s.userBizRepo.CreateUserByPhone(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateUserResp{
		Data: dto.AccountDto.ToPbAccountInfo(userModel),
	}, nil
}

func (s *accountService) CreateUserByEmail(ctx context.Context, req *resourcev1.CreateUserByEmailReq) (*resourcev1.CreateUserResp, error) {
	param := dto.AccountDto.ToBoCreateUserParam2(req)
	userModel, err := s.userBizRepo.CreateUserByEmail(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateUserResp{
		Data: dto.AccountDto.ToPbAccountInfo(userModel),
	}, nil
}
