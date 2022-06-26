package interfaces

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
)

type VehicleRepo interface{
	FindAll() (*models.Vehicles, error)
	FindByID(id int) (*models.Vehicle, error)
	Search(category ...interface{}) (*models.Vehicles, error)
	SortByPrice(price int) (*models.Vehicles, error)
	Popular(rating int) (*models.Vehicles, error)
	Add(data *models.Vehicle) (*models.Vehicle, error)
	Delete(id int) (*models.Vehicle, error)
	Update(id int, data *models.Vehicle) (*models.Vehicle, error)
}

type VehicleService interface {
	FindAll() (*helpers.Response, error)
	FindByID(id int) (*helpers.Response, error)
	Search(category ...interface{}) (*helpers.Response, error)
	SortByPrice(price int) (*helpers.Response, error)
	Popular(rating int) (*helpers.Response, error)
	Save(data *models.Vehicle) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
	Update(id int, data *models.Vehicle) (*helpers.Response, error)
}