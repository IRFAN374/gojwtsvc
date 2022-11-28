package transporthttp

import (
	"net/url"

	transport "github.com/IRFAN374/gojwtsvc/user/transport"
	httpkit "github.com/go-kit/kit/transport/http"
)


func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {
	return transport.EndpointsSet{
		LoginEndpoint: httpkit.NewClient(
			"POST", u,
			Encode_Login_Request,
			Decode_Login_Response,
			opts...,
		).Endpoint(),
	}
}