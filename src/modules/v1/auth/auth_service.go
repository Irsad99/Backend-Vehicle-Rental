package auth

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
	"BackendGo/src/interfaces"
)

var response helpers.Response

type token_response struct {
	Tokens string `json:"token"`
}

type Auth_service struct {
	rep interfaces.UserRepo
}

func NewService(svc interfaces.UserRepo) *Auth_service {
	return &Auth_service{svc}
}

func (svc *Auth_service) Login(body models.User) (*helpers.Response, error) {

	user, err := svc.rep.FindByEmail(body.Email)
	if err != nil {
		return response.ResponseJSON(400, "Email Salah"), nil
	}

	if !helpers.CheckPassword(user.Password, body.Password) {
		return response.ResponseJSON(400, "Password Salah"), nil
	}

	token := helpers.NewToken(user.User_ID, body.Email, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return nil, err
	}

	res := response.ResponseJSON(200, token_response{Tokens: theToken})
	return res, nil

}
