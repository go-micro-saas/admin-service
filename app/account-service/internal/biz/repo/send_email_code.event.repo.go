package bizrepos

import (
	"context"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
)

type SendEmailHandler func(ctx context.Context, param *bo.SendEmailCodeParam) (*bo.SendEmailCodeReply, error)

type SendEmailCodeEventRepo interface {
	Publish(ctx context.Context, param *bo.SendEmailCodeParam) error
	Consume(ctx context.Context, handler SendEmailHandler) error
	Close(ctx context.Context) error
}
