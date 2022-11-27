package service

import (
	"context"
	"time"

	model "github.com/IRFAN374/gojwtsvc/model"
	service "github.com/IRFAN374/gojwtsvc/user"
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

func (M loggingMiddleware) Login(arg0 context.Context, arg1 string, arg2 string) (res0 model.LoginResponse, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Login",
			"request", logLoginRequest{
				Id:       1,
				UserName: arg1,
				Password: arg2,
			},
			"err", res1,
			"took", time.Since(begin),
		)
	}(time.Now())

	return M.next.Login(arg0, arg1, arg2)
}

type (
	logLoginRequest struct {
		Id       int    `json:"id"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
)
