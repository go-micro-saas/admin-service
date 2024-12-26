package bizrepos

import (
	"context"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
)

type UserAuthBizRepo interface {
	SendVerifyCode(ctx context.Context, param *bo.SendVerifyCodeParam) (*bo.SendVerifyCodeReply, error)
	ConfirmVerifyCode(ctx context.Context, param *bo.ConfirmVerifyCodeParam) error

	SignupByPhone(ctx context.Context, in *bo.SignupByPhoneParam) (*po.User, *bo.SignTokenResp, error)
	SignupByEmail(ctx context.Context, in *bo.SignupByEmailParam) (*po.User, *bo.SignTokenResp, error)
	LoginByEmail(ctx context.Context, in *bo.LoginByEmailParam) (*po.User, *bo.SignTokenResp, error)
	LoginByPhone(ctx context.Context, in *bo.LoginByPhoneParam) (*po.User, *bo.SignTokenResp, error)
	LoginByUserID(ctx context.Context, userID uint64, loginParam *bo.LoginParam) (*po.User, *bo.SignTokenResp, error)
	LoginByUser(ctx context.Context, userModel *po.User, loginParam *bo.LoginParam) (*bo.SignTokenResp, error)

	GenSignTokenRequestByUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error)
	GenSignTokenRequestByAdminUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error)
	SignToken(ctx context.Context, req *bo.SignTokenReq) (*bo.SignTokenResp, error)
	RefreshToken(ctx context.Context, refreshToken string) (*bo.SignTokenResp, error)

	CheckAndGetByRegisterEmail(ctx context.Context, email string) (*po.UserRegEmail, error)
	CheckAndGetByRegisterPhone(ctx context.Context, phone string) (*po.UserRegPhone, error)
	CheckAndGetUserByUserId(ctx context.Context, userId uint64) (*po.User, error)
	IsExistRegisterEmail(ctx context.Context, email string) (*po.UserRegEmail, bool, error)
	IsExistRegisterPhone(ctx context.Context, phone string) (*po.UserRegPhone, bool, error)
	ValidateLoginUser(userModel *po.User, loginParam *bo.LoginParam) error
	ComparePassword(hashPassword, plaintextPassword string) error
}
