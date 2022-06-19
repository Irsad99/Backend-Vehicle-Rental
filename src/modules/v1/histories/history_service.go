package histories

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
	"BackendGo/src/interfaces"

	"github.com/asaskevich/govalidator"
)

type histories_service struct {
	repo interfaces.HistoryRepo
}

func NewService(svc interfaces.HistoryRepo) *histories_service {
	return &histories_service{svc}
}

func (svc *histories_service) FindAll() (*helpers.Response, error) {

	result, err := svc.repo.FindAll()
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *histories_service) FindHistoryByID(id int) (*helpers.Response, error) {

	_, err := govalidator.ToInt(id)
	if err != nil {
		res := response.ResponseJSON(400, "Id yang anda masukan salah")
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.FindHistoryByID(id)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *histories_service) SortByStart() (*helpers.Response, error) {

	result, err := svc.repo.SortByStart()
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}
	
	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *histories_service) Save(data *models.History) (*helpers.Response, error) {

	var histories models.History

	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		res := response.ResponseJSON(400, histories)
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Add(data)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(201, result)
	return res, nil
}

func (svc *histories_service) Delete(id int) (*helpers.Response, error) {

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

	res := response.ResponseJSON(204, result)
	return res, nil
}

func (svc *histories_service) Update(id int, status string) (*helpers.Response, error) {

	_, err := govalidator.ToInt(id)
	if err != nil {
		res := response.ResponseJSON(400, "Id yang anda masukan salah")
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Update(id, status)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(201, result)
	return res, nil
}
