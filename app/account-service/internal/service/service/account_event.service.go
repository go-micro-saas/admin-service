package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	servicev1 "github.com/go-micro-saas/account-service/api/account-service/v1/services"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	threadpkg "github.com/ikaiguang/go-srv-kit/kratos/thread"
	"sync/atomic"
)

// accountEventService ...
type accountEventService struct {
	servicev1.UnimplementedSrvAccountEventV1Server

	log                    *log.Helper
	sendEmailCodeBizRepo   bizrepos.SendEmailCodeBizRepo
	sendEmailCodeEventRepo bizrepos.SendEmailCodeEventRepo

	isConsumingSendEmailCodeEvent atomic.Int64
}

func NewAccountEventService(
	logger log.Logger,
	sendEmailCodeBizRepo bizrepos.SendEmailCodeBizRepo,
	sendEmailCodeEventRepo bizrepos.SendEmailCodeEventRepo,
) servicev1.SrvAccountEventV1Server {
	return &accountEventService{
		log:                    log.NewHelper(log.With(logger, "module", "account-service/service/account_event")),
		sendEmailCodeBizRepo:   sendEmailCodeBizRepo,
		sendEmailCodeEventRepo: sendEmailCodeEventRepo,
	}
}

// Ping ping pong
func (s *accountEventService) Ping(ctx context.Context, in *resourcev1.PingReq) (out *resourcev1.PingResp, err error) {
	out = &resourcev1.PingResp{
		Message: in.Message,
	}
	return out, err
}

func (s *accountEventService) SubscribeSendEmailCodeEvent(ctx context.Context, req *resourcev1.SubscribeSendEmailCodeEventReq) (*resourcev1.SubscribeSendEmailCodeEventResp, error) {
	if s.isConsumingSendEmailCodeEvent.Load() < 1 {
		threadpkg.GoSafe(func() {
			s.isConsumingSendEmailCodeEvent.Add(1)
			defer func() { s.isConsumingSendEmailCodeEvent.Add(-1) }()
			err := s.sendEmailCodeEventRepo.Consume(ctx, s.sendEmailCodeBizRepo.SendEmailCode)
			if err != nil {
				s.log.WithContext(ctx).Errorw("msg", "SubscribeSendEmailCodeEvent failed!", "err", err)
			}
		})
	}
	return &resourcev1.SubscribeSendEmailCodeEventResp{
		Data: &resourcev1.SubscribeSendEmailCodeEventRespData{
			ConsumerCounter: s.isConsumingSendEmailCodeEvent.Load(),
		},
	}, nil
}

func (s *accountEventService) StopSendEmailCodedEvent(ctx context.Context, req *resourcev1.StopSendEmailCodeEventReq) (*resourcev1.StopSendEmailCodeEventResp, error) {
	err := s.sendEmailCodeEventRepo.Close(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorw("msg", "run StopSendEmailCodeEvent failed!", "err", err)
		return nil, err
	}
	return &resourcev1.StopSendEmailCodeEventResp{
		Data: &resourcev1.StopSendEmailCodeEventRespData{
			ConsumerCounter: s.isConsumingSendEmailCodeEvent.Load(),
		},
	}, nil
}
