package users

import (
	"BackendGo/src/helpers"

	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

// var users Users
var response helpers.Response

type user_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *user_repo {
	return &user_repo{grm}
}

func (r *user_repo) FindAll() (*helpers.Response, error) {

	var users Users

	result := r.db.Order("user_id desc").Find(&users)

	if result.Error != nil {
		res := response.ResponseJSON(500, users)
		return res, nil
	}

	res := response.ResponseJSON(200, users)

	return res, nil
}

func (r *user_repo) Add(data *User) (*helpers.Response, error) {

	var users Users

	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		res := response.ResponseJSON(400, users)
		res.Message = err.Error()
		return res, nil
	}

	getEmail := r.db.Where("email = ?", &data.Email).First(&users)
	if getEmail.RowsAffected != 0 {
		res := response.ResponseJSON(300, users)
		return res, nil
	}

	result := r.db.Create(data)

	if result.Error != nil {
		res := response.ResponseJSON(400, users)
		return res, nil
	}

	getData := r.db.First(&users, &data.User_ID)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, users)
		return res, nil
	}

	res := response.ResponseJSON(201, users)

	return res, nil
}

func (r *user_repo) Delete(data *int) (*helpers.Response, error) {

	var users Users

	getData := r.db.First(&users, &data)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, users)
		return res, nil
	}

	result := r.db.Delete(&User{}, &data)

	if result.Error != nil {
		res := response.ResponseJSON(400, users)
		return res, nil
	}

	res := response.ResponseJSON(204, users)
	return res, nil
}

func (r *user_repo) Update(id *int, data *User) (*helpers.Response, error) {

	var users Users

	result := r.db.Model(&User{}).Where("user_id = ?", &id).Updates(&User{Name : data.Name, Gender: data.Gender, Email: data.Email, Phone: data.Phone, Birth: data.Birth, Address: data.Address})

	if result.Error != nil {
		res := response.ResponseJSON(400, users)
		return res, nil
	}

	getData := r.db.First(&users, &id)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, users)
		return res, nil
	}

	res := response.ResponseJSON(201, users)
	return res, nil
}
