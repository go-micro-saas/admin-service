package caches

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/global"
	"github.com/go-micro-saas/admin-service/app/admin-service/internal/data/po"
	datarepos "github.com/go-micro-saas/admin-service/app/admin-service/internal/data/repo"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type verifyCodeCache struct {
	log     *log.Helper
	redisCC redis.UniversalClient
}

func NewVerifyCodeCache(logger log.Logger, redisCC redis.UniversalClient) datarepos.VerifyCodeCacheRepo {
	logHandler := log.NewHelper(log.With(logger, "module", "admin-service/data/cache/verify_code"))
	return &verifyCodeCache{
		log:     logHandler,
		redisCC: redisCC,
	}
}

func (s *verifyCodeCache) key(param *po.VerifyCodeParam) string {
	key := "verifyCode:" + param.VerifyAccount +
		":" + strconv.FormatInt(int64(param.VerifyType), 10) +
		":" + param.VerifyCode

	return global.Key(key)
}

func (s *verifyCodeCache) SaveCode(ctx context.Context, param *po.VerifyCodeParam) error {
	key := s.key(param)
	_, err := s.redisCC.Set(ctx, key, param.VerifyCode, param.TTL).Result()
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return errorpkg.WithStack(e)
	}
	return nil
}

func (s *verifyCodeCache) GetCode(ctx context.Context, param *po.VerifyCodeParam) (string, error) {
	key := s.key(param)
	code, err := s.redisCC.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		e := errorpkg.ErrorInternalError(err.Error())
		return "", errorpkg.WithStack(e)
	}
	return code, nil
}

func (s *verifyCodeCache) DeleteCode(ctx context.Context, param *po.VerifyCodeParam) error {
	key := s.key(param)
	_, err := s.redisCC.Del(ctx, key).Result()
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return errorpkg.WithStack(e)
	}
	return nil
}

func (s *verifyCodeCache) VerifyCode(ctx context.Context, param *po.VerifyCodeParam) (bool, error) {
	code, err := s.GetCode(ctx, param)
	if err != nil {
		return false, err
	}
	defer func() {
		deleteErr := s.DeleteCode(ctx, param)
		if deleteErr != nil {
			s.log.WithContext(ctx).Warnw("msg", "delete verify code error", "error", deleteErr)
		}
	}()
	if code == "" {
		return false, nil
	}
	return code == param.VerifyCode, nil
}
