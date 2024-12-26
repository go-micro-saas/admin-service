package bizrepos

import (
	"context"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
)

type SendEmailHandler func(ctx context.Context, param *bo.SendVerifyCodeEventParam) (bo.SendVerifyCodeReply, error)

type SendEmailCodeEventRepo interface {
	Publish(ctx context.Context, param *bo.SendVerifyCodeEventParam) error
	Consume(ctx context.Context, handler SendEmailHandler) error
	Close(ctx context.Context) error
}
