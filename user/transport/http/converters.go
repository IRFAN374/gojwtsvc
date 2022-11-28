package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/IRFAN374/gojwtsvc/common/chttp"
	transport "github.com/IRFAN374/gojwtsvc/user/transport"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(&buf)

	return nil
}

func CommonHTTPResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return chttp.EncodeResponse(ctx, w, response)
}

func Decode_Login_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("Decoding request body", req)

	return &req, err
}

func Decode_Login_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.LoginResponse
	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func Encode_Login_Request(ctx context.Context, r *http.Request, request interface{}) error {
	_ = request.(*transport.LoginRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/login"))
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func Encode_Login_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
