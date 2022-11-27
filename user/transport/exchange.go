package transport

import "github.com/IRFAN374/gojwtsvc/model"

type (
	LoginRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	LoginResponse struct {
		LoginRes model.LoginResponse `json:"loginRes"`
	}
)
