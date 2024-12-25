package bizrepos

import (
	"context"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
)

type UserAuthBizRepo interface {
	SendVerifyCode(ctx context.Context, param *bo.SendVerifyCodeParam) (*bo.SendVerifyCodeReply, error)
	ConfirmVerifyCode(ctx context.Context, param *bo.ConfirmVerifyCodeParam) error

	SignupByPhone(ctx context.Context, in *resourcev1.SignupByPhoneReq) (*po.User, *bo.SignTokenResp, error)
	LoginByEmail(ctx context.Context, in *resourcev1.LoginByEmailReq) (*po.User, *bo.SignTokenResp, error)
	LoginByPhone(ctx context.Context, in *resourcev1.LoginByPhoneReq) (*po.User, *bo.SignTokenResp, error)
	LoginByUserID(ctx context.Context, userID uint64, loginParam *bo.LoginParam) (*po.User, *bo.SignTokenResp, error)
	LoginByUser(ctx context.Context, userModel *po.User, loginParam *bo.LoginParam) (*bo.SignTokenResp, error)

	GenSignTokenRequestByUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error)
	GenSignTokenRequestByAdminUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error)
	SignToken(ctx context.Context, req *bo.SignTokenReq) (*bo.SignTokenResp, error)
	RefreshToken(ctx context.Context, refreshToken string) (*bo.SignTokenResp, error)

	CheckAndGetByRegisterEmail(ctx context.Context, email string) (*po.UserRegEmail, error)
	CheckAndGetByRegisterPhone(ctx context.Context, phone string) (*po.UserRegPhone, error)
	CheckAndGetUserByUserId(ctx context.Context, userId uint64) (*po.User, error)
	ValidateLoginUser(userModel *po.User, loginParam *bo.LoginParam) error
	ComparePassword(hashPassword, plaintextPassword string) error
}
