package interfaces

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	FindById(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Add(data *models.User) (*models.User, error)
	Delete(id int) (*models.User, error)
	Update(id int, data *models.User) (*models.User, error)
}

type UserService interface {
	FindAll() (*helpers.Response, error)
	FindById(id int) (*helpers.Response, error)
	FindByEmail(email string) (*helpers.Response, error)
	Save(data *models.User) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
	Update(id int, data *models.User) (*helpers.Response, error)
}
