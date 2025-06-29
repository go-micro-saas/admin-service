package datarepos

import (
	"context"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
)

type VerifyCodeCacheRepo interface {
	SaveCode(ctx context.Context, param *po.VerifyCodeParam) error
	VerifyCode(ctx context.Context, param *po.VerifyCodeParam) (bool, error)
	DeleteCode(ctx context.Context, param *po.VerifyCodeParam) error
	GetCode(ctx context.Context, param *po.VerifyCodeParam) (string, error)
}
