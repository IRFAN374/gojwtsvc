package transporthttp

import (
	http1 "net/http"

	transport "github.com/IRFAN374/gojwtsvc/user/transport"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()

	mux.Methods("POST").Path("/user/login").Handler(
		http.NewServer(
			endpoints.LoginEndpoint,
			Decode_Login_Request,
			Encode_Login_Response,
			opts...))

	return mux
}
