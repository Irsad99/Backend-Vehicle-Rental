package interfaces

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
)

type HistoryRepo interface{
	FindAll() (*models.Histories, error)
	FindHistoryByID(id int) (*models.Results, error)
	SortByStart() (*models.Histories, error)
	Add(data *models.History) (*models.History, error)
	Delete(id int) (*models.History, error)
	Update(id int, status string) (*models.History, error)
}

type HistoryService interface {
	FindAll() (*helpers.Response, error)
	FindHistoryByID(id int) (*helpers.Response, error)
	SortByStart() (*helpers.Response, error)
	Save(data *models.History) (*helpers.Response, error)
	Delete(id int) (*helpers.Response, error)
	Update(id int, status string) (*helpers.Response, error)
}