package transport

import "github.com/IRFAN374/gojwtsvc/model"

type (
	LoginRequest struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}
	LoginResponse struct {
		LoginRes model.LoginResponse `json:"loginRes"`
	}
)
