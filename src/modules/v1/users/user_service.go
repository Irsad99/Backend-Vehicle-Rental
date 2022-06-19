package users

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
	"BackendGo/src/interfaces"

	"github.com/asaskevich/govalidator"
)

type user_service struct {
	repo interfaces.UserRepo
}

func NewService(svc interfaces.UserRepo) *user_service {
	return &user_service{svc}
}

func (svc *user_service) FindAll() (*helpers.Response, error) {

	result, err := svc.repo.FindAll()
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, len(*result))
	return res, nil
}

func (svc *user_service) FindByEmail(email string) (*helpers.Response, error) {

	result, err := svc.repo.FindByEmail(email)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)

	return res, nil
}

func (svc *user_service) Save(data *models.User) (*helpers.Response, error) {

	// var users models.Users

	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		res := response.ResponseJSON(400, data)
		res.Message = err.Error()
		return res, nil
	}

	hsPass, err := helpers.HassPassword(data.Password)
	if err != nil {
		res := response.ResponseJSON(400, hsPass)
		return res, nil
	}

	data.Password = hsPass
	result, err := svc.repo.Add(data)
	if err != nil {
		res := response.ResponseJSON(400, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)
	return res, nil
}

func (svc *user_service) Delete(id int) (*helpers.Response, error) {

	_, err := govalidator.ToInt(id)
	if err != nil {
		res := response.ResponseJSON(400, "Id yang anda masukan salah")
		res.Message = err.Error()
		return res, nil
	}

	result, err := svc.repo.Delete(id)
	if err != nil {
		res := response.ResponseJSON(404, result)
		res.Message = err.Error()
		return res, nil
	}

	res := response.ResponseJSON(200, result)

	return res, nil
}

func (svc *user_service) Update(id int, data *models.User) (*helpers.Response, error) {

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
