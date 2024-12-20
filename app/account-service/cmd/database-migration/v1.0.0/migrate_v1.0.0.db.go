package dbv1_0_0

import (
	"context"
	dbv1_0_0_user "github.com/go-micro-saas/account-service/app/account-service/cmd/database-migration/v1.0.0/user"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	"gorm.io/gorm"
)

const (
	Version = "v1.0.0"
)

// Upgrade update
func Upgrade(ctx context.Context, dbConn *gorm.DB, migrateRepo migrationpkg.MigrateRepo) (err error) {
	//var (
	//	serverVersion     = Version
	//	migrateIdentifier = serverVersion + ":migrate"
	//)
	//// 已进行数据库迁移
	//dataModel, _, err := migrateRepo.QueryRecord(ctx, migrateIdentifier)
	//if err != nil {
	//	return
	//}
	//if dataModel.Id > 0 {
	//	return
	//}
	//// 记录数据库迁移
	//defer func() {
	//	if err == nil {
	//		err = migrateRepo.CreateRecord(ctx, serverVersion, migrateIdentifier)
	//	}
	//}()

	// dbv1_0_0_nodeid
	err = dbv1_0_0_user.NewMigrateHandler(dbConn, migrateRepo).Upgrade(ctx)
	if err != nil {
		return err
	}

	return err
}
