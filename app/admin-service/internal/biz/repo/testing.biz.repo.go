package bizrepos

import (
	"context"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
)

type TestingBizRepo interface {
	HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error)
}
