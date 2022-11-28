package chttp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/IRFAN374/gojwtsvc/common/cerror"
)

type Meta struct {
	TimeStamp int64 `json:"ts"`
}

type Wrapper struct {
	Code    string      `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta"`
}

type WrapperJson struct {
	Code    string          `json:"code"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
	Meta    Meta            `json:"meta"`
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(&Wrapper{
		Code:    "S001",
		Status:  "success",
		Message: "processed successfully",
		Data:    response,
		Meta: Meta{
			TimeStamp: time.Now().Unix(),
		},
	})
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var data interface{}

	var codeErr *cerror.Error
	if errors.As(err, &codeErr) {
		switch codeErr.Code() {
		case "E001":
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		codeErr = cerror.ErrUnKnownError(err).(*cerror.Error)
	}

	json.NewEncoder(w).Encode(&Wrapper{
		Code:    "E001",
		Status:  "failure",
		Message: "Error in failure",
		Data:    data,
		Meta: Meta{
			TimeStamp: time.Now().Unix(),
		},
	})
}

func DecodeResponse(ctx context.Context, r *http.Response, resp interface{}) error {
	switch r.StatusCode {
	case http.StatusOK, http.StatusConflict, http.StatusInternalServerError:
	default:
		_, err := io.Copy(io.Discard, r.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("http error, statuscode:%d status:%s", r.StatusCode, r.Status)
	}

	var wrapper WrapperJson

	err := json.NewDecoder(r.Body).Decode(&wrapper)
	if err != nil {
		return err
	}

	if wrapper.Code != "S001" {
		return DecodeError(wrapper)
	}

	return json.Unmarshal(wrapper.Data, &resp)

}

func DecodeError(w WrapperJson) error {
	err := cerror.New(w.Code, w.Message)

	return err
}
