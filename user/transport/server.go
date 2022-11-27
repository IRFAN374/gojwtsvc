package transport

import (
	"context"

	"github.com/IRFAN374/gojwtsvc/user"
	endpoint "github.com/go-kit/kit/endpoint"
)

func Endpoints(svc user.Service) EndpointsSet {
	return EndpointsSet{
		LoginEndpoint: LoginEndpoint(svc),
	}
}

func LoginEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*LoginRequest)
		res0, res1 := svc.Login(arg0, req.Name, req.Password)
		return &LoginResponse{
			LoginRes: res0,
		}, res1
	}
}
