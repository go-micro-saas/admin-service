package bizrepos

import (
	"context"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
)

type SendEmailCodeBizRepo interface {
	SendEmailCode(ctx context.Context, param *bo.SendEmailCodeParam) (*bo.SendEmailCodeReply, error)
}
