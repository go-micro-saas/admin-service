package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	errorv1 "github.com/go-micro-saas/account-service/api/account-service/v1/errors"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
	datarepos "github.com/go-micro-saas/account-service/app/account-service/internal/data/repo"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type accountBiz struct {
	log         *log.Helper
	idGenerator idpkg.Snowflake

	userDataRepo           datarepos.UserDataRepo
	userRegEmailDataRepo   datarepos.UserRegEmailDataRepo
	userRegPhoneDataRepo   datarepos.UserRegPhoneDataRepo
	userVerifyCodeDataRepo datarepos.UserVerifyCodeDataRepo
}

func NewAccountBiz(
	logger log.Logger,
	idGenerator idpkg.Snowflake,

	userDataRepo datarepos.UserDataRepo,
	userRegEmailDataRepo datarepos.UserRegEmailDataRepo,
	userRegPhoneDataRepo datarepos.UserRegPhoneDataRepo,
	userVerifyCodeDataRepo datarepos.UserVerifyCodeDataRepo,
) bizrepos.UserBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "account-service/biz/account"))
	return &accountBiz{
		log:         logHelper,
		idGenerator: idGenerator,

		userDataRepo:           userDataRepo,
		userRegEmailDataRepo:   userRegEmailDataRepo,
		userRegPhoneDataRepo:   userRegPhoneDataRepo,
		userVerifyCodeDataRepo: userVerifyCodeDataRepo,
	}
}

func (s accountBiz) GetUserByUid(ctx context.Context, uid uint64) (*po.User, error) {
	userModel, isNotFound, err := s.userDataRepo.QueryOneByUserId(ctx, uid)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return userModel, nil
}

func (s accountBiz) GetUsersByUidList(ctx context.Context, uidList []uint64) ([]*po.User, error) {
	userModels, err := s.userDataRepo.QueryByUserIdList(ctx, uidList)
	if err != nil {
		return nil, err
	}
	return userModels, nil
}
