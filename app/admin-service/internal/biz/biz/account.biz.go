package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	errorv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/errors"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/repo"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
	datarepos "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/repo"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	regexpkg "github.com/ikaiguang/go-srv-kit/kit/regex"
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
	logHelper := log.NewHelper(log.With(logger, "module", "admin-service/biz/account"))
	return &accountBiz{
		log:         logHelper,
		idGenerator: idGenerator,

		userDataRepo:           userDataRepo,
		userRegEmailDataRepo:   userRegEmailDataRepo,
		userRegPhoneDataRepo:   userRegPhoneDataRepo,
		userVerifyCodeDataRepo: userVerifyCodeDataRepo,
	}
}

func (s *accountBiz) GetUserByUid(ctx context.Context, uid uint64) (*po.User, error) {
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

func (s *accountBiz) GetUsersByUidList(ctx context.Context, uidList []uint64) ([]*po.User, error) {
	userModels, err := s.userDataRepo.QueryByUserIdList(ctx, uidList)
	if err != nil {
		return nil, err
	}
	return userModels, nil
}

func (s *accountBiz) ListUser(ctx context.Context, param *bo.UserListParam) ([]*po.User, int64, error) {
	queryParam := &po.QueryUserParam{
		UidList:      param.UidList,
		ContactPhone: param.ContactPhone,
		ContactEmail: param.ContactEmail,

		PaginatorArgs: param.PaginatorArgs,
	}
	dataModels, counter, err := s.userDataRepo.ListUsers(ctx, queryParam, param.PaginatorArgs)
	if err != nil {
		return dataModels, counter, err
	}
	return dataModels, counter, err
}

func (s *accountBiz) CreateUser(ctx context.Context, param *bo.CreateUserParam) (*po.User, error) {
	if regexpkg.IsPhone(param.UserPhone) == false {
		e := errorv1.ErrorS103InvalidPhone("无效的手机号")
		return nil, errorpkg.WithStack(e)
	}
	if regexpkg.IsEmail(param.UserEmail) == false {
		e := errorv1.ErrorS103InvalidEmail("无效的邮箱")
		return nil, errorpkg.WithStack(e)
	}
	// passwd
	passwdParam := &bo.PasswordParam{
		Password:        param.Password,
		PasswordConfirm: param.Password,
	}
	passwdHash, err := passwdParam.ValidateAndEncrypt()
	if err != nil {
		return nil, err
	}

	// exist email?
	_, isNotFound, err := s.userRegEmailDataRepo.QueryOneByUserEmail(ctx, param.UserEmail)
	if err != nil {
		return nil, err
	}
	if !isNotFound {
		e := errorv1.ErrorS103UserEmailExist("用户已存在")
		return nil, errorpkg.WithStack(e)
	}
	// exist phone?
	_, isNotFound, err = s.userRegPhoneDataRepo.QueryOneByUserPhone(ctx, param.UserPhone)
	if err != nil {
		return nil, err
	}
	if !isNotFound {
		e := errorv1.ErrorS103UserPhoneExist("用户已存在")
		return nil, errorpkg.WithStack(e)
	}

	// user
	var (
		dataModel = po.NewUserByEmail(param.UserEmail, passwdHash)
		regEmail  = po.NewUserRegEmail(param.UserEmail)
		regPhone  = po.NewUserRegPhone(param.UserPhone)
	)
	s.attachUserModelByCreateUserParam(dataModel, param)
	dataModel.UserId, err = s.idGenerator.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	regEmail.UserId = dataModel.UserId
	regPhone.UserId = dataModel.UserId

	// create
	createParam := &po.CreateAccountParam{
		UserModel:     dataModel,
		RegPhoneModel: regPhone,
		RegEmailModel: regEmail,
	}
	if err = s.CreateAccount(ctx, createParam); err != nil {
		return nil, err
	}
	return dataModel, nil
}

func (s *accountBiz) CreateOrGetUserByEmail(ctx context.Context, param *bo.CreateUserParam) (*po.User, bool, error) {
	return s.createUserByEmail(ctx, param, true)
}

func (s *accountBiz) CreateUserByEmail(ctx context.Context, param *bo.CreateUserParam) (*po.User, error) {
	dataModel, _, err := s.createUserByEmail(ctx, param, false)
	if err != nil {
		return nil, err
	}
	return dataModel, nil
}

func (s *accountBiz) createUserByEmail(ctx context.Context, param *bo.CreateUserParam, isCreateOrGet bool) (*po.User, bool, error) {
	var isCreate bool
	if regexpkg.IsEmail(param.UserEmail) == false {
		e := errorv1.ErrorS103InvalidEmail("无效的邮箱")
		return nil, isCreate, errorpkg.WithStack(e)
	}

	// exist?
	existModel, isNotFound, err := s.userRegEmailDataRepo.QueryOneByUserEmail(ctx, param.UserEmail)
	if err != nil {
		return nil, isCreate, err
	}
	if !isNotFound {
		if isCreateOrGet {
			existUser, err := s.GetUserByUid(ctx, existModel.UserId)
			if err != nil {
				return nil, isCreate, err
			}
			return existUser, isCreate, nil
		}
		e := errorv1.ErrorS103UserEmailExist("用户已存在")
		return nil, isCreate, errorpkg.WithStack(e)
	}
	isCreate = true

	// passwd
	passwdParam := &bo.PasswordParam{
		Password:        param.Password,
		PasswordConfirm: param.Password,
	}
	passwdHash, err := passwdParam.ValidateAndEncrypt()
	if err != nil {
		return nil, isCreate, err
	}

	// user
	var (
		dataModel = po.NewUserByEmail(param.UserEmail, passwdHash)
		regModel  = po.NewUserRegEmail(param.UserEmail)
	)
	s.attachUserModelByCreateUserParam(dataModel, param)
	dataModel.UserId, err = s.idGenerator.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, isCreate, errorpkg.WithStack(e)
	}
	regModel.UserId = dataModel.UserId

	// create
	createParam := &po.CreateAccountParam{
		UserModel:     dataModel,
		RegPhoneModel: nil,
		RegEmailModel: regModel,
	}
	if err = s.CreateAccount(ctx, createParam); err != nil {
		return nil, isCreate, err
	}
	return dataModel, isCreate, nil
}

func (s *accountBiz) CreateOrGetUserByPhone(ctx context.Context, param *bo.CreateUserParam) (*po.User, bool, error) {
	return s.createUserByPhone(ctx, param, true)
}

func (s *accountBiz) CreateUserByPhone(ctx context.Context, param *bo.CreateUserParam) (*po.User, error) {
	dataModel, _, err := s.createUserByPhone(ctx, param, false)
	if err != nil {
		return nil, err
	}
	return dataModel, nil
}

func (s *accountBiz) createUserByPhone(ctx context.Context, param *bo.CreateUserParam, isCreateOrGet bool) (*po.User, bool, error) {
	var isCreate bool
	if regexpkg.IsPhone(param.UserPhone) == false {
		e := errorv1.ErrorS103InvalidPhone("无效的手机号")
		return nil, isCreate, errorpkg.WithStack(e)
	}

	// exist?
	existModel, isNotFound, err := s.userRegPhoneDataRepo.QueryOneByUserPhone(ctx, param.UserPhone)
	if err != nil {
		return nil, isCreate, err
	}
	if !isNotFound {
		if isCreateOrGet {
			existUser, err := s.GetUserByUid(ctx, existModel.UserId)
			if err != nil {
				return nil, isCreate, err
			}
			return existUser, isCreate, nil
		}
		e := errorv1.ErrorS103UserPhoneExist("用户已存在")
		return nil, isCreate, errorpkg.WithStack(e)
	}
	isCreate = true

	// passwd
	passwdParam := &bo.PasswordParam{
		Password:        param.Password,
		PasswordConfirm: param.Password,
	}
	passwdHash, err := passwdParam.ValidateAndEncrypt()
	if err != nil {
		return nil, isCreate, err
	}

	// user
	var (
		dataModel = po.NewUserByPhone(param.UserPhone, passwdHash)
		regModel  = po.NewUserRegPhone(param.UserPhone)
	)
	s.attachUserModelByCreateUserParam(dataModel, param)
	dataModel.UserId, err = s.idGenerator.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, isCreate, errorpkg.WithStack(e)
	}
	regModel.UserId = dataModel.UserId

	// create
	createParam := &po.CreateAccountParam{
		UserModel:     dataModel,
		RegPhoneModel: regModel,
		RegEmailModel: nil,
	}
	if err = s.CreateAccount(ctx, createParam); err != nil {
		return nil, isCreate, err
	}
	return dataModel, isCreate, nil
}

func (s *accountBiz) attachUserModelByCreateUserParam(dataModel *po.User, param *bo.CreateUserParam) {
	dataModel.UserGender = param.UserGender
	dataModel.UserStatus = param.UserStatus
	if param.UserPhone != "" {
		dataModel.UserPhone = param.UserPhone
	}
	if param.UserEmail != "" {
		dataModel.UserEmail = param.UserEmail
	}
	if param.UserNickname != "" {
		dataModel.UserNickname = param.UserNickname
	}
	if param.UserAvatar != "" {
		dataModel.UserAvatar = param.UserAvatar
	}
}

func (s *accountBiz) CreateAccount(ctx context.Context, param *po.CreateAccountParam) (err error) {
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
