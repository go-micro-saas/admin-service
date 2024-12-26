package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/enums"
	errorv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/errors"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/repo"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
	datarepos "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/repo"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	regexpkg "github.com/ikaiguang/go-srv-kit/kit/regex"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"time"
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
	logHelper := log.NewHelper(log.With(logger, "module", "admin-service/biz/user_auth"))
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
func (s *userAuthBiz) LoginByEmail(ctx context.Context, in *bo.LoginByEmailParam) (*po.User, *bo.SignTokenResp, error) {
	// code
	if !in.SkipVerifyCode {
		verifyParam := &bo.ConfirmVerifyCodeParam{
			VerifyAccount: in.Email,
			VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_EMAIL,
			VerifyCode:    in.Code,
		}
		err := s.ConfirmVerifyCode(ctx, verifyParam)
		if err != nil {
			return nil, nil, err
		}
	}

	// 登录
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
func (s *userAuthBiz) LoginByPhone(ctx context.Context, in *bo.LoginByPhoneParam) (*po.User, *bo.SignTokenResp, error) {
	// code
	if !in.SkipVerifyCode {
		verifyParam := &bo.ConfirmVerifyCodeParam{
			VerifyAccount: in.Phone,
			VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_PHONE,
			VerifyCode:    in.Code,
		}
		err := s.ConfirmVerifyCode(ctx, verifyParam)
		if err != nil {
			return nil, nil, err
		}
	}

	// 登录
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

func (s *userAuthBiz) IsExistRegisterEmail(ctx context.Context, email string) (*po.UserRegEmail, bool, error) {
	// 注册邮箱
	regModel, isNotFound, err := s.userRegEmailDataRepo.QueryOneByUserEmail(ctx, email)
	if err != nil {
		return nil, false, err
	}
	return regModel, !isNotFound, err
}

func (s *userAuthBiz) IsExistRegisterPhone(ctx context.Context, phone string) (*po.UserRegPhone, bool, error) {
	// 注册手机
	regModel, isNotFound, err := s.userRegPhoneDataRepo.QueryOneByUserPhone(ctx, phone)
	if err != nil {
		return nil, false, err
	}
	return regModel, !isNotFound, nil
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

func (s *userAuthBiz) SignupByPhone(ctx context.Context, in *bo.SignupByPhoneParam) (*po.User, *bo.SignTokenResp, error) {
	if regexpkg.IsPhone(in.Phone) == false {
		e := errorv1.ErrorS103InvalidPhone("无效的手机号")
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

	// code
	if !in.SkipVerifyCode {
		verifyParam := &bo.ConfirmVerifyCodeParam{
			VerifyAccount: in.Phone,
			VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_PHONE,
			VerifyCode:    in.Code,
		}
		err = s.ConfirmVerifyCode(ctx, verifyParam)
		if err != nil {
			return nil, nil, err
		}
	}

	// exist?
	_, isNotFound, err := s.userRegPhoneDataRepo.QueryOneByUserPhone(ctx, in.Phone)
	if err != nil {
		return nil, nil, err
	}
	if !isNotFound {
		e := errorv1.ErrorS103UserExist("用户已存在")
		return nil, nil, errorpkg.WithStack(e)
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

	// create
	createParam := &po.CreateAccountParam{
		UserModel:     dataModel,
		RegPhoneModel: regModel,
		RegEmailModel: nil,
	}
	if err = s.CreateAccount(ctx, createParam); err != nil {
		return nil, nil, err
	}

	loginParam := &bo.LoginParam{
		SkipValidatePassword: true,
		PlaintextPassword:    in.Password,
	}
	return s.LoginByUserID(ctx, dataModel.UserId, loginParam)
}

func (s *userAuthBiz) SignupByEmail(ctx context.Context, in *bo.SignupByEmailParam) (*po.User, *bo.SignTokenResp, error) {
	if regexpkg.IsEmail(in.Email) == false {
		e := errorv1.ErrorS103InvalidEmail("无效的邮箱")
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

	// code
	if !in.SkipVerifyCode {
		verifyParam := &bo.ConfirmVerifyCodeParam{
			VerifyAccount: in.Email,
			VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_EMAIL,
			VerifyCode:    in.Code,
		}
		err = s.ConfirmVerifyCode(ctx, verifyParam)
		if err != nil {
			return nil, nil, err
		}
	}

	// exist?
	_, isNotFound, err := s.userRegEmailDataRepo.QueryOneByUserEmail(ctx, in.Email)
	if err != nil {
		return nil, nil, err
	}
	if !isNotFound {
		e := errorv1.ErrorS103UserExist("用户已存在")
		return nil, nil, errorpkg.WithStack(e)
	}

	// user
	var (
		dataModel = po.NewUserByEmail(in.Email, passwdHash)
		regModel  = po.NewUserRegEmail(in.Email)
	)
	dataModel.UserId, err = s.idGenerator.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, nil, errorpkg.WithStack(e)
	}
	regModel.UserId = dataModel.UserId

	// create
	createParam := &po.CreateAccountParam{
		UserModel:     dataModel,
		RegPhoneModel: nil,
		RegEmailModel: regModel,
	}
	if err = s.CreateAccount(ctx, createParam); err != nil {
		return nil, nil, err
	}

	loginParam := &bo.LoginParam{
		SkipValidatePassword: true,
		PlaintextPassword:    in.Password,
	}
	return s.LoginByUserID(ctx, dataModel.UserId, loginParam)
}

func (s *userAuthBiz) CreateAccount(ctx context.Context, param *po.CreateAccountParam) (err error) {
	// 事务
	tx := s.userDataRepo.NewTransaction(ctx)
	defer func() {
		commitErr := tx.CommitAndErrRollback(ctx, err)
		if commitErr != nil {
			s.log.WithContext(ctx).Errorw(
				"mgs", "CreateAccount tx.CommitAndErrRollback failed!",
				"err", commitErr,
			)
		}
	}()
	err = s.userDataRepo.CreateWithTransaction(ctx, tx, param.UserModel)
	if err != nil {
		return err
	}
	if param.RegPhoneModel != nil {
		err = s.userRegPhoneDataRepo.CreateWithTransaction(ctx, tx, param.RegPhoneModel)
		if err != nil {
			return err
		}
	}
	if param.RegEmailModel != nil {
		err = s.userRegEmailDataRepo.CreateWithTransaction(ctx, tx, param.RegEmailModel)
		if err != nil {
			return err
		}
	}
	return err
}

func (s *userAuthBiz) SendVerifyCode(ctx context.Context, param *bo.SendVerifyCodeParam) (*bo.SendVerifyCodeReply, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	var (
		code      = po.NewVerifyCode()
		dataModel = po.NewUserVerifyCode(code)
	)
	dataModel.VerifyAccount = param.VerifyAccount
	dataModel.VerifyType = param.VerifyType

	err := s.userVerifyCodeDataRepo.Create(ctx, dataModel)
	if err != nil {
		return nil, err
	}

	resp := &bo.SendVerifyCodeReply{
		IsSendToMQ:    false,
		VerifyAccount: param.VerifyAccount,
		VerifyType:    param.VerifyType,
		VerifyCode:    dataModel.VerifyCode,
	}
	return resp, nil
}

func (s *userAuthBiz) ConfirmVerifyCode(ctx context.Context, param *bo.ConfirmVerifyCodeParam) error {
	if err := param.Validate(); err != nil {
		return nil
	}
	queryParam := &po.GetVerifyCodeParam{
		VerifyAccount: param.VerifyAccount,
		VerifyType:    param.VerifyType,
		VerifyCode:    param.VerifyCode,
		VerifyStatusSlice: []enumv1.UserVerifyStatusEnum_UserVerifyStatus{
			enumv1.UserVerifyStatusEnum_UNSPECIFIED,
			enumv1.UserVerifyStatusEnum_CONFIRMING,
		},
		GTCreateTime: time.Now().Add(-po.DefaultVerifyCodeExpiredTime),
	}
	dataModel, isNotFound, err := s.userVerifyCodeDataRepo.QueryOneVerifyCode(ctx, queryParam)
	if err != nil {
		return nil
	}
	if isNotFound || dataModel.VerifyCode != param.VerifyCode || !dataModel.CanVerification() {
		e := errorv1.ErrorS103VerifyCodeIncorrect("验证码不正确")
		return errorpkg.WithStack(e)
	}

	// used
	dataModel.VerifyStatus = enumv1.UserVerifyStatusEnum_CONFIRMED
	dataModel.ConfirmTime = uint64(time.Now().Unix())
	err = s.userVerifyCodeDataRepo.UpdateVerifyStatus(ctx, dataModel)
	if err != nil {
		return err
	}
	return nil
}
