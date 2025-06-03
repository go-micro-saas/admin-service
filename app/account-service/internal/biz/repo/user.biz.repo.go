package bizrepos

import (
	"context"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
)

type UserBizRepo interface {
	GetUserByUid(ctx context.Context, uid uint64) (*po.User, error)
	GetUsersByUidList(ctx context.Context, uidList []uint64) ([]*po.User, error)
}
