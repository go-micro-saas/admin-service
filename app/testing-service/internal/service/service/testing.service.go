package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/admin-service/api/testing-service/v1/resources"
	servicev1 "github.com/go-micro-saas/admin-service/api/testing-service/v1/services"
	bizrepos "github.com/go-micro-saas/admin-service/app/testing-service/internal/biz/repo"
	"github.com/go-micro-saas/admin-service/app/testing-service/internal/service/dto"
)

type testingV1Service struct {
	servicev1.UnimplementedSrvTestdataServer

	log        *log.Helper
	testingBiz bizrepos.TestingBizRepo
}

func NewTestingV1Service(logger log.Logger, testingBiz bizrepos.TestingBizRepo) servicev1.SrvTestdataServer {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/service/service"))
	return &testingV1Service{
		log:        logHelper,
		testingBiz: testingBiz,
	}
}

func (s *testingV1Service) Get(ctx context.Context, req *resourcev1.TestReq) (*resourcev1.TestResp, error) {
	param := dto.TestingDto.ToBHelloWorldParam(req)
	reply, err := s.testingBiz.HelloWorld(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.TestResp{
		Data: dto.TestingDto.ToPbTestRespData(reply.Message),
	}, nil
}
