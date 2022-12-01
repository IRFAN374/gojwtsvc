package service

import (
	"context"
	"time"

	"github.com/IRFAN374/gojwtsvc/model"
	service "github.com/IRFAN374/gojwtsvc/token"
	log "github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (M loggingMiddleware) CreateToken(arg0 context.Context, arg1 uint64) (res0 *model.TokenDetails, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateToken",
			"request", logCreateToken{
				Userid: arg1,
			},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.CreateToken(arg0, arg1)
}

func (M loggingMiddleware) CreateAuth(arg0 context.Context, arg1 uint64, arg2 *model.TokenDetails) (res0 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateAuth",
			"request", logCreateAuth{
				Userid: arg1,
				Td:     arg2,
			},
			"err", res0,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.CreateAuth(arg0, arg1, arg2)
}

func (M loggingMiddleware) Refresh(arg0 context.Context) (res0 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Refresh",
			"request", logRefresh{},
			"err", res0,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.Refresh(arg0)
}

func (M loggingMiddleware) VerifyToken(arg0 context.Context) (res0 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "VerifyToken",
			"request", logVerifyToken{},
			"err", res0,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.VerifyToken(arg0)
}

type (
	logCreateToken struct {
		Userid uint64 `json:"userid"`
	}

	logCreateAuth struct {
		Userid uint64              `json:"userid"`
		Td     *model.TokenDetails `json:"td"`
	}

	logRefresh struct{}

	logVerifyToken struct{}
)
