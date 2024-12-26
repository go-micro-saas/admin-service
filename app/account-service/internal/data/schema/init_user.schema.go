package schemas

import (
	enumv1 "github.com/go-micro-saas/admin-service/api/admin-service/v1/enums"
	userschemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user"
	emailschemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user_reg_email"
	phoneschemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user_reg_phone"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	md5pkg "github.com/ikaiguang/go-srv-kit/kit/md5"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	"time"

	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	"gorm.io/gorm"
)

const (
	DefaultInitUserID       uint64 = 1
	DefaultInitUserNickname        = "user"
	DefaultInitUserPhone           = "13800138000"
	DefaultInitUserEmail           = "user@user.user"
	DefaultInitUserPassword        = "User.123456"
)

// UserDB ...
type UserDB interface {
	InitializeUser(dbConn *gorm.DB) migrationutil.MigrationInterface
}

// NewUserDB ...
func NewUserDB() UserDB {
	return &user{}
}

type user struct {
	userSchema         userschemas.User
	userRegEmailSchema emailschemas.UserRegEmail
	userRegPhoneSchema phoneschemas.UserRegPhone
}

// InitializeUser 初始化
// 账户密码: user@user.user : User.123456
func (s *user) InitializeUser(dbConn *gorm.DB) migrationutil.MigrationInterface {
	var (
		tableName           = s.userSchema.TableName()
		userPhone           = DefaultInitUserPhone
		userEmail           = DefaultInitUserEmail
		userID              = DefaultInitUserID
		migrationVersion    = migrationutil.Version
		migrationIdentifier = migrationVersion + ":" + tableName + ":initialize.user:" + userEmail
	)
	migrationUp := func() error {
		return s.initializeUserUp(dbConn, userID, userEmail, userPhone)
	}
	migrationDown := func() error {
		return s.initializeUserDown(dbConn, userID, userEmail, userPhone)
	}
	return migrationutil.NewAnyMigrator(
		migrationVersion,
		migrationIdentifier,
		migrationUp,
		migrationDown,
	)
}

// initializeUserDown 初始化
func (s *user) initializeUserDown(dbConn *gorm.DB, userID uint64, userEmail, userPhone string) (err error) {
	tx := dbConn.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback().Error
		}
	}()

	err = tx.Table(s.userRegEmailSchema.TableName()).
		Where(emailschemas.FieldUserEmail+" = ?", userEmail).
		Delete(emailschemas.UserRegEmail{}).Error
	if err != nil {
		return err
	}

	err = tx.Table(s.userRegPhoneSchema.TableName()).
		Where(phoneschemas.FieldUserPhone+" = ?", userPhone).
		Delete(phoneschemas.UserRegPhone{}).Error
	if err != nil {
		return err
	}
	err = tx.Table(s.userSchema.TableName()).
		Where(userschemas.FieldUserId+" = ?", userID).
		Delete(userschemas.User{}).Error
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return err
}

// initializeUserUp 初始化
func (s *user) initializeUserUp(dbConn *gorm.DB, userID uint64, userEmail, userPhone string) (err error) {
	// 检查是否存在
	var (
		existEmailModel = &emailschemas.UserRegEmail{}
	)
	err = dbConn.Table(s.userRegEmailSchema.TableName()).
		Where(emailschemas.FieldUserEmail+" = ?", userEmail).
		First(existEmailModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
		} else {
			return err
		}
	}
	if existEmailModel.Id > 0 {
		return err
	}

	// 新增
	var (
		now             = time.Now()
		userNickname    = DefaultInitUserNickname
		passwordMd5, _  = md5pkg.Md5([]byte(DefaultInitUserPassword))
		passwordHash, _ = passwordutil.Encrypt(passwordMd5)
	)
	regEmailModel := &emailschemas.UserRegEmail{
		CreatedTime: now,
		UpdatedTime: now,
		UserId:      userID,
		UserEmail:   userEmail,
	}
	regPhoneModel := &phoneschemas.UserRegPhone{
		CreatedTime: now,
		UpdatedTime: now,
		UserId:      userID,
		UserPhone:   userPhone,
	}
	userModel := &userschemas.User{
		UserId:       userID,
		CreatedTime:  now,
		UpdatedTime:  now,
		UserPhone:    regPhoneModel.UserPhone,
		UserEmail:    regEmailModel.UserEmail,
		UserNickname: userNickname,
		UserAvatar:   "",
		UserGender:   uint32(enumv1.UserGenderEnum_SECRET),
		UserStatus:   uint32(enumv1.UserStatusEnum_ENABLE),
		PasswordHash: string(passwordHash),
		RegisterType: uint32(enumv1.UserRegisterTypeEnum_EMAIL),
	}

	// 新增用户
	tx := dbConn.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback().Error
		}
	}()

	if err = tx.Table(s.userRegEmailSchema.TableName()).Create(regEmailModel).Error; err != nil {
		return err
	}
	if err = tx.Table(s.userRegPhoneSchema.TableName()).Create(regPhoneModel).Error; err != nil {
		return err
	}
	if err = tx.Table(s.userSchema.TableName()).Create(userModel).Error; err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return err
}
