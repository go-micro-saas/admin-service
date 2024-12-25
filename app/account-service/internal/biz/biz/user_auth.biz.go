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
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	regexpkg "github.com/ikaiguang/go-srv-kit/kit/regex"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// userAuthBiz ...
type userAuthBiz struct {
	log         *log.Helper
	authRepo    authpkg.AuthRepo
	idGenerator idpkg.Snowflake

	userDataRepo           datarepos.UserDataRepo
	userRegEmailDataRepo   datarepos.UserRegEmailDataRepo
	userRegPhoneDataRepo   datarepos.UserRegPhoneDataRepo
	userVerifyCodeDataRepo datarepos.UserVerifyCodeDataRepo
}

// NewUserAuthBiz ...
func NewUserAuthBiz(
	logger log.Logger,
	authRepo authpkg.AuthRepo,
	idGenerator idpkg.Snowflake,

	userDataRepo datarepos.UserDataRepo,
	userRegEmailDataRepo datarepos.UserRegEmailDataRepo,
	userRegPhoneDataRepo datarepos.UserRegPhoneDataRepo,
	userVerifyCodeDataRepo datarepos.UserVerifyCodeDataRepo,
) bizrepos.UserAuthBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "account-service/biz/user_auth"))
	return &userAuthBiz{
		log:         logHelper,
		authRepo:    authRepo,
		idGenerator: idGenerator,

		userDataRepo:           userDataRepo,
		userRegEmailDataRepo:   userRegEmailDataRepo,
		userRegPhoneDataRepo:   userRegPhoneDataRepo,
		userVerifyCodeDataRepo: userVerifyCodeDataRepo,
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
		PlaintextPassword: in.Password,
	}
	return s.LoginByUserID(ctx, regEmailModel.UserId, loginParam)
}

// LoginByPhone ...
func (s *userAuthBiz) LoginByPhone(ctx context.Context, in *resourcev1.LoginByPhoneReq) (*po.User, *bo.SignTokenResp, error) {
	regPhoneModel, err := s.CheckAndGetByRegisterPhone(ctx, in.Phone)
	if err != nil {
		return nil, nil, err
	}
	loginParam := &bo.LoginParam{
		PlaintextPassword: in.Password,
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
	err := s.ValidateLoginUser(userModel, loginParam)
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
	regModel, isNotFound, err := s.userRegEmailDataRepo.QueryOneByUserEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return regModel, nil
}

// CheckAndGetByRegisterPhone 手机是否存在
func (s *userAuthBiz) CheckAndGetByRegisterPhone(ctx context.Context, phone string) (*po.UserRegPhone, error) {
	// 注册手机
	regModel, isNotFound, err := s.userRegPhoneDataRepo.QueryOneByUserPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.ErrorS103UserNotExist("用户不存在")
		return nil, errorpkg.WithStack(e)
	}
	return regModel, nil
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
func (s *userAuthBiz) ValidateLoginUser(userModel *po.User, loginParam *bo.LoginParam) error {
	if !loginParam.SkipValidatePassword {
		err := s.ComparePassword(userModel.PasswordHash, loginParam.PlaintextPassword)
		if err != nil {
			return err
		}
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

func (s *userAuthBiz) SignupByPhone(ctx context.Context, in *resourcev1.SignupByPhoneReq) (*po.User, *bo.SignTokenResp, error) {
	if regexpkg.IsPhone(in.Phone) == false {
		e := errorv1.ErrorS103InvalidPhone("无效的手机号")
		return nil, nil, errorpkg.WithStack(e)

	}
	_, isNotFound, err := s.userRegPhoneDataRepo.QueryOneByUserPhone(ctx, in.Phone)
	if err != nil {
		return nil, nil, err
	}
	if !isNotFound {
		e := errorv1.ErrorS103UserExist("用户已存在")
		return nil, nil, errorpkg.WithStack(e)
	}

	// passwd
	passwdParam := &bo.PasswordParam{
		Password:        in.Password,
		PasswordConfirm: in.PasswordConfirm,
	}
	passwdHash, err := passwdParam.ValidateAndEncrypt()
	if err != nil {
		return nil, nil, err
	}

	// user
	var (
		dataModel = po.NewUserByPhone(in.Phone, passwdHash)
		regModel  = po.NewUserRegPhone(in.Phone)
	)
	dataModel.UserId, err = s.idGenerator.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, nil, errorpkg.WithStack(e)
	}
	regModel.UserId = dataModel.UserId

	// 事务
	tx := s.userDataRepo.NewTransaction(ctx)
	defer func() {
		commitErr := tx.CommitAndErrRollback(ctx, err)
		if commitErr != nil {
			s.log.WithContext(ctx).Errorw(
				"mgs", "GetNodeId tx.CommitAndErrRollback failed!",
				"err", commitErr,
			)
		}
	}()
	err = s.userDataRepo.CreateWithTransaction(ctx, tx, dataModel)
	if err != nil {
		return nil, nil, err
	}
	err = s.userRegPhoneDataRepo.CreateWithTransaction(ctx, tx, regModel)
	if err != nil {
		return nil, nil, err
	}

	loginParam := &bo.LoginParam{
		SkipValidatePassword: true,
		PlaintextPassword:    in.Password,
	}
	return s.LoginByUserID(ctx, dataModel.UserId, loginParam)
}

func (s *userAuthBiz) SendVerifyCode(ctx context.Context, param *bo.SendVerifyCodeParam) (*bo.SignTokenResp, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	var (
		code      = po.NewVerifyCode()
		dataModel = po.NewUserVerifyCode(code)
	)
	dataModel.VerifyCode = param.VerifyAccount
	dataModel.VerifyType = param.VerifyType
	return nil, nil
}
