// Package datarepos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package datarepos

import (
	context "context"
	"database/sql"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
)

// UserRegPhoneDataRepo repo
type UserRegPhoneDataRepo interface {
	NewTransaction(ctx context.Context, opts ...*sql.TxOptions) gormpkg.TransactionInterface
	Create(ctx context.Context, dataModel *po.UserRegPhone) error
	CreateWithTransaction(ctx context.Context, tx gormpkg.TransactionInterface, dataModel *po.UserRegPhone) (err error)
	ExistCreate(ctx context.Context, dataModel *po.UserRegPhone) (anotherModel *po.UserRegPhone, isNotFound bool, err error)
	CreateInBatches(ctx context.Context, dataModels []*po.UserRegPhone, batchSize int) error
	Insert(ctx context.Context, dataModels []*po.UserRegPhone) error
	Update(ctx context.Context, dataModel *po.UserRegPhone) error
	ExistUpdate(ctx context.Context, dataModel *po.UserRegPhone) (anotherModel *po.UserRegPhone, isNotFound bool, err error)
	QueryOneById(ctx context.Context, id interface{}) (dataModel *po.UserRegPhone, isNotFound bool, err error)
	QueryOneByUserPhone(ctx context.Context, userPhone string) (dataModel *po.UserRegPhone, isNotFound bool, err error)
	QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *po.UserRegPhone, isNotFound bool, err error)
	QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) (dataModels []*po.UserRegPhone, err error)
	List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) (dataModels []*po.UserRegPhone, totalNumber int64, err error)
	Delete(ctx context.Context, dataModel *po.UserRegPhone) error
	DeleteByIds(ctx context.Context, ids interface{}) error
}
