package dto

import (
	resourcev1 "github.com/go-micro-saas/admin-service/api/testing-service/v1/resources"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/biz/bo"
)

var (
	TestingDto testingDto
)

type testingDto struct{}

func (s *testingDto) ToPbHelloWorldParam(req *resourcev1.TestReq) *bo.HelloWorldParam {
	res := &bo.HelloWorldParam{
		Message: req.Message,
	}
	return res
}

func (s *testingDto) ToPbTestRespData(msg string) *resourcev1.TestRespData {
	res := &resourcev1.TestRespData{
		Message: msg,
	}

	return res
}
