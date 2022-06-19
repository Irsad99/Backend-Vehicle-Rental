package interfaces

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
)

type AuthService interface {
	Login(body models.User) (*helpers.Response, error)
}