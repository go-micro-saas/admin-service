package events

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/account-service/app/account-service/internal/biz/repo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var SendEmailCodeEventTopic = "send_email_code"

var (
	_ bizrepos.SendEmailCodeEventRepo = (*sendEmailCodeEvent)(nil)
)

type sendEmailCodeEvent struct {
	logger log.Logger
	log    *log.Helper
	mqConn *rabbitmqpkg.ConnectionWrapper

	topic          string             // topic
	initPubSubOnce sync.Once          // init
	pub            message.Publisher  // 使用 getPublisherSubscriber
	sub            message.Subscriber // 使用 getPublisherSubscriber

	closing        chan struct{}
	receiveCounter uint64
}

func NewSendEmailCodeEventRepo(
	logger log.Logger,
	mqConn *rabbitmqpkg.ConnectionWrapper,
) bizrepos.SendEmailCodeEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "account-service/biz/event/send_email_code"))
	return &sendEmailCodeEvent{
		logger:  logger,
		log:     logHelper,
		mqConn:  mqConn,
		topic:   po.Key(SendEmailCodeEventTopic),
		closing: make(chan struct{}),
	}
}

func (s *sendEmailCodeEvent) getPublisherSubscriber() (message.Publisher, message.Subscriber, error) {
	var (
		err error
	)
	s.initPubSubOnce.Do(func() {
		s.pub, s.sub, err = rabbitmqpkg.NewPublisherAndSubscriberWithConnection(s.mqConn, rabbitmqpkg.WithKratosLogger(s.logger))
	})
	if err != nil {
		s.initPubSubOnce = sync.Once{}
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, nil, errorpkg.WithStack(e)
	}
	return s.pub, s.sub, nil
}

func (s *sendEmailCodeEvent) Publish(ctx context.Context, param *bo.SendVerifyCodeEventParam) error {
	publisher, _, err := s.getPublisherSubscriber()
	if err != nil {
		return err
	}
	payload, err := param.MarshalToJSON()
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuidpkg.New(), payload)
	msg.SetContext(ctx)
	err = publisher.Publish(s.topic, msg)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
	}
	return nil
}

func (s *sendEmailCodeEvent) Consume(ctx context.Context, handler bizrepos.SendEmailHandler) error {
	_, subscriber, err := s.getPublisherSubscriber()
	if err != nil {
		return err
	}
	m, err := subscriber.Subscribe(context.Background(), s.topic)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
	}
	for {
		select {
		case msg := <-m:
			s.receiveCounter++
			{
				s.log.WithContext(ctx).Infow("msg", "SendEmailCodeEvent.Consume Received message", "receiveCounter", s.receiveCounter, "msg.payload", string(msg.Payload))
				param := &bo.SendVerifyCodeEventParam{}
				err := param.UnmarshalFromJSON(msg.Payload)
				if err != nil {
					s.log.WithContext(ctx).Errorw("msg", "SendEmailCodeEvent.Consume SendVerifyCodeEventParam UnmarshalFromJSON failed",
						"err", err, "payload", string(msg.Payload))
					msg.Ack()
					continue
				}
				result, err := handler(context.Background(), param)
				if err != nil {
					s.log.WithContext(ctx).Errorw("msg", "SendEmailCodeEvent.Consume Process failed",
						"err", err, "payload", string(msg.Payload), "result", result)
					msg.Ack()
					continue
				}
				msg.Ack()
			}
		case <-s.closing:
			s.log.Debugw("msg", "SendEmailCodeEvent.Consume Stopping Consume")
			return nil
		}
	}
}

func (s *sendEmailCodeEvent) Close(ctx context.Context) error {
	s.closing <- struct{}{}
	return nil
}
