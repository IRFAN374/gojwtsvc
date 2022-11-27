package transport

import (
	"context"

	"github.com/IRFAN374/gojwtsvc/model"
)

func (set EndpointsSet) Login(arg0 context.Context, arg1 string, arg2 string) (loginRes model.LoginResponse, res1 error) {
	request := LoginRequest {
		Name: arg1,
		Password: arg2,
	}

	res0, res1 := set.LoginEndpoint(arg0, &request)

	if res1 != nil {
		return
	}

	return res0.(*LoginResponse).LoginRes, res1
}
