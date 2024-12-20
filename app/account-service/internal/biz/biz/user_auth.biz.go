package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	errorv1 "github.com/go-micro-saas/account-service/api/account-service/v1/errors"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
	datarepos "github.com/go-micro-saas/account-service/app/account-service/internal/data/repo"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// userAuthBiz ...
type userAuthBiz struct {
	log                     *log.Helper
	authRepo                authpkg.AuthRepo
	userDataRepo            datarepos.UserDataRepo
	userRegEmailDataRepo    datarepos.UserRegEmailDataRepo
	userRegPhoneDataRepo    datarepos.UserRegPhoneDataRepo
	userConfirmCodeDataRepo datarepos.UserConfirmCodeDataRepo
}

// NewUserAuthBiz ...
func NewUserAuthBiz(
	logger log.Logger,
	authRepo authpkg.AuthRepo,
	userDataRepo datarepos.UserDataRepo,
	userRegEmailDataRepo datarepos.UserRegEmailDataRepo,
	userRegPhoneDataRepo datarepos.UserRegPhoneDataRepo,
	userConfirmCodeDataRepo datarepos.UserConfirmCodeDataRepo,
) bizrepos.UserAuthBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "account-service/biz/user_auth"))
	return &userAuthBiz{
		log:                     logHelper,
		authRepo:                authRepo,
		userDataRepo:            userDataRepo,
		userRegEmailDataRepo:    userRegEmailDataRepo,
		userRegPhoneDataRepo:    userRegPhoneDataRepo,
		userConfirmCodeDataRepo: userConfirmCodeDataRepo,
	}
}

// LoginByEmail ...
func (s *userAuthBiz) LoginByEmail(ctx context.Context, in *resourcev1.LoginByEmailReq) (*po.User, *bo.SignTokenResp, error) {
	// 注册邮箱
	regEmailModel, err := s.CheckAndGetByRegisterEmail(ctx, in.Email)
	if err != nil {
		return nil, nil, err
	}
	loginParam := &bo.LoginParam{
		Password: in.Password,
	}
	return s.LoginByUserID(ctx, regEmailModel.UserId, loginParam)
}

// LoginByPhone ...
func (s *userAuthBiz) LoginByPhone(ctx context.Context, in *resourcev1.LoginByPhoneReq) (*po.User, *bo.SignTokenResp, error) {
	// 注册邮箱
	regPhoneModel, err := s.CheckAndGetByRegisterPhone(ctx, in.Phone)
	if err != nil {
		return nil, nil, err
	}
	loginParam := &bo.LoginParam{
		Password: in.Password,
	}
	return s.LoginByUserID(ctx, regPhoneModel.UserId, loginParam)
}

// LoginByUserID ...
func (s *userAuthBiz) LoginByUserID(ctx context.Context, userID uint64, loginParam *bo.LoginParam) (*po.User, *bo.SignTokenResp, error) {
	// user
	userModel, err := s.CheckAndGetUserByUserId(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	signResp, err := s.LoginByUser(ctx, userModel, loginParam)
	if err != nil {
		return nil, nil, err
	}
	return userModel, signResp, nil
}

// LoginByUser ...
func (s *userAuthBiz) LoginByUser(ctx context.Context, userModel *po.User, loginParam *bo.LoginParam) (*bo.SignTokenResp, error) {
	// 验证用户
	err := s.ValidateLoginUser(userModel, loginParam.Password)
	if err != nil {
		return nil, err
	}
	// 签证
	signReq, err := s.GenSignTokenRequestByUserModel(ctx, userModel)
	if err != nil {
		return nil, err
	}
	signResp, err := s.SignToken(ctx, signReq)
	if err != nil {
		return nil, err
	}
	return signResp, nil
}

// CheckAndGetByRegisterEmail 邮箱是否存在
func (s *userAuthBiz) CheckAndGetByRegisterEmail(ctx context.Context, email string) (*po.UserRegEmail, error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.userRegEmailDataRepo.QueryOneByUserEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return regEmailModel, nil
}

// CheckAndGetByRegisterPhone 手机是否存在
func (s *userAuthBiz) CheckAndGetByRegisterPhone(ctx context.Context, phone string) (*po.UserRegPhone, error) {
	// 注册手机
	regEmailModel, isNotFound, err := s.userRegPhoneDataRepo.QueryOneByUserPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return regEmailModel, nil
}

// CheckAndGetUserByUserId 用户是否存在
func (s *userAuthBiz) CheckAndGetUserByUserId(ctx context.Context, userId uint64) (*po.User, error) {
	userModel, isNotFound, err := s.userDataRepo.QueryOneByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return userModel, nil
}

// ValidateLoginUser 对比密码
func (s *userAuthBiz) ValidateLoginUser(userModel *po.User, plaintextPassword string) error {
	err := s.ComparePassword(userModel.PasswordHash, plaintextPassword)
	if err != nil {
		return err
	}
	if !userModel.IsValidStatus() {
		e := errorv1.ErrorS103UserStatusNotAllow("无效的登录状态")
		return errorpkg.WithStack(e)
	}
	return nil
}

// ComparePassword 对比密码
// plaintextPassword plaintext || md5
func (s *userAuthBiz) ComparePassword(hashPassword, plaintextPassword string) error {
	// 验证密码
	err := passwordutil.Compare(hashPassword, plaintextPassword)
	if err != nil {
		e := errorv1.ErrorS103UserPasswordIncorrect("密码不正确")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

// GenSignTokenRequestByUserModel auth claims
func (s *userAuthBiz) GenSignTokenRequestByUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error) {
	payload := &authpkg.Payload{
		TokenID:       uuidpkg.New(),
		UserID:        uint64(userModel.UserId),
		UserUuid:      "",
		LoginPlatform: authpkg.LoginPlatformEnum_UNSPECIFIED,
		LoginType:     authpkg.LoginTypeEnum_UNSPECIFIED,
		LoginLimit:    authpkg.LoginLimitEnum_UNLIMITED,
		TokenType:     authpkg.TokenTypeEnum_USER,
	}
	res := &bo.SignTokenReq{
		Claims: *authpkg.GenAuthClaimsByAuthPayload(payload, authpkg.AccessTokenExpire),
	}
	return res, nil
}

// GenSignTokenRequestByAdminUserModel auth claims
func (s *userAuthBiz) GenSignTokenRequestByAdminUserModel(ctx context.Context, userModel *po.User) (*bo.SignTokenReq, error) {
	payload := &authpkg.Payload{
		TokenID:       uuidpkg.New(),
		UserID:        uint64(userModel.UserId),
		UserUuid:      "",
		LoginPlatform: authpkg.LoginPlatformEnum_UNSPECIFIED,
		LoginType:     authpkg.LoginTypeEnum_UNSPECIFIED,
		LoginLimit:    authpkg.LoginLimitEnum_UNLIMITED,
		TokenType:     authpkg.TokenTypeEnum_ADMIN,
	}
	res := &bo.SignTokenReq{
		Claims: *authpkg.GenAuthClaimsByAuthPayload(payload, authpkg.AccessTokenExpire),
	}
	return res, nil
}

// SignToken token
func (s *userAuthBiz) SignToken(ctx context.Context, req *bo.SignTokenReq) (*bo.SignTokenResp, error) {
	tokenResp, err := s.authRepo.SignToken(ctx, &req.Claims)
	if err != nil {
		return nil, err
	}

	res := &bo.SignTokenResp{}
	res.SetByAuthTokenResponse(tokenResp)
	return res, nil
}

// RefreshToken token
func (s *userAuthBiz) RefreshToken(ctx context.Context, refreshToken string) (*bo.SignTokenResp, error) {
	if refreshToken == "" {
		e := errorpkg.ErrorInvalidParameter("refresh token is empty")
		return nil, errorpkg.WithStack(e)
	}
	authClaims, err := s.authRepo.DecodeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if err = s.authRepo.VerifyRefreshToken(ctx, authClaims); err != nil {
		return nil, err
	}

	tokenResp, err := s.authRepo.RefreshToken(ctx, authClaims)
	if err != nil {
		return nil, err
	}

	res := &bo.SignTokenResp{}
	res.SetByAuthTokenResponse(tokenResp)
	return res, nil
}
