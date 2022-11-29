package user

import (
	"context"
	// "os"
	// "strconv"
	// "time"

	"github.com/IRFAN374/gojwtsvc/model"
	"github.com/IRFAN374/gojwtsvc/token"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/twinj/uuid"
)

type Service interface {
	Login(ctx context.Context, name string, password string) (loginRes model.LoginResponse, err error)
}

type service struct {
	tokenSvc token.Service
}

func NewService(tokenSvc token.Service) *service {
	return &service{
		tokenSvc: tokenSvc,
	}
}

var user = model.User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func (svc *service) Login(ctx context.Context, name string, password string) (loginRes model.LoginResponse, err error) {

	ts, err := svc.tokenSvc.CreateToken(ctx, user.ID)
	if err != nil {
		return model.LoginResponse{}, err
	}

	err = svc.tokenSvc.CreateAuth(ctx, user.ID, ts)
	if err != nil {
		return model.LoginResponse{}, err
	}

	loginRes = model.LoginResponse{
		UserId:       int(user.ID),
		UserName:     name,
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}

	return
}

// func CreateToken(userid uint64) (*model.TokenDetails, error) {
// 	td := &model.TokenDetails{}
// 	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
// 	td.AccessUuid = uuid.NewV4().String()

// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = uuid.NewV4().String()

// 	var err error
// 	//Creating Access Token
// 	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["access_uuid"] = td.AccessUuid
// 	atClaims["user_id"] = userid
// 	atClaims["exp"] = td.AtExpires
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	//Creating Refresh Token
// 	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
// 	rtClaims := jwt.MapClaims{}
// 	rtClaims["refresh_uuid"] = td.RefreshUuid
// 	rtClaims["user_id"] = userid
// 	rtClaims["exp"] = td.RtExpires
// 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
// 	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return td, nil
// }

// func (svc *service) CreateAuth(userid uint64, td *model.TokenDetails) error {
// 	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
// 	rt := time.Unix(td.RtExpires, 0)
// 	now := time.Now()

// 	errAccess := svc.client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
// 	if errAccess != nil {
// 		return errAccess
// 	}
// 	errRefresh := svc.client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
// 	if errRefresh != nil {
// 		return errRefresh
// 	}
// 	return nil
// }
