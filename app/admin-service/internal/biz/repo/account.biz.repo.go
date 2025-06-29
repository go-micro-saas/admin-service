package bizrepos

import (
	"context"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
)

type UserBizRepo interface {
	GetUserByUid(ctx context.Context, uid uint64) (*po.User, error)
	GetUsersByUidList(ctx context.Context, uidList []uint64) ([]*po.User, error)
	ListUser(ctx context.Context, param *bo.UserListParam) ([]*po.User, int64, error)
	CreateUser(ctx context.Context, param *bo.CreateUserParam) (*po.User, error)
	CreateUserByEmail(ctx context.Context, param *bo.CreateUserParam) (*po.User, error)
	CreateOrGetUserByEmail(ctx context.Context, param *bo.CreateUserParam) (dataModel *po.User, isCreate bool, err error)
	CreateUserByPhone(ctx context.Context, param *bo.CreateUserParam) (*po.User, error)
	CreateOrGetUserByPhone(ctx context.Context, param *bo.CreateUserParam) (dataModel *po.User, isCreate bool, err error)
}
