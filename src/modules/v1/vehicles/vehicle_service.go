package vehicles

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
	"BackendGo/src/interfaces"

	"github.com/asaskevich/govalidator"
)

type vehicle_service struct {
	repo interfaces.VehicleRepo
}

func NewService(svc interfaces.VehicleRepo) *vehicle_service {
	return &vehicle_service{svc}
}

func (svc *vehicle_service) FindAll() (*helpers.Response, error) {

	result, err := svc.repo.FindAll()
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) FindByID(id int) (*helpers.Response, error) {

	result, err := svc.repo.FindByID(id)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) Search(name string, location string) (*helpers.Response, error) {

	result, err := svc.repo.Search(name, location)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) SortByPrice(price int) (*helpers.Response, error) {

	result, err := svc.repo.SortByPrice(price)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) SortByType(category string) (*helpers.Response, error) {

	result, err := svc.repo.SortByType(category)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) SortByLocation(location string) (*helpers.Response, error) {

	result, err := svc.repo.SortByLocation(location)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) Popular(rating int) (*helpers.Response, error) {

	result, err := svc.repo.Popular(rating)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) Save(data *models.Vehicle) (*helpers.Response, error) {

	var vehicles models.Vehicles

	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		res := response.ResponseJSON(400, vehicles)
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Add(data)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) Delete(id int) (*helpers.Response, error) {

	_, err := govalidator.ToInt(id)
	if err != nil {
		res := response.ResponseJSON(400, "Id yang anda masukan salah")
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Delete(id)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *vehicle_service) Update(id int, data *models.Vehicle) (*helpers.Response, error) {

	_, err := govalidator.ToInt(id)
	if err != nil {
		res := response.ResponseJSON(400, "Id yang anda masukan salah")
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Update(id, data)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)

	return res, nil
}
