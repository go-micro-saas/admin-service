package dbv1_0_0_user

import (
	"context"
	schemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema"
	userschemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user"
	emailschemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user_reg_email"
	phoneschemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user_reg_phone"
	coceschemas "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/schema/user_verify_code"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"gorm.io/gorm"
)

// Migrate 数据库迁移
type Migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationpkg.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationpkg.MigrateRepo) *Migrate {
	return &Migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// Upgrade ...
func (s *Migrate) Upgrade(ctx context.Context) error {
	var (
		mr       migrationpkg.MigrationInterface
		migrator = s.dbConn.WithContext(ctx).Migrator()
	)

	// 创建表
	mr = userschemas.UserSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = phoneschemas.UserRegPhoneSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = emailschemas.UserRegEmailSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = coceschemas.UserVerifyCodeSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 初始化用户
	mr = schemas.NewUserDB().InitializeUser(s.dbConn)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}
