package vehicles

import (
	"BackendGo/src/helpers"

	"gorm.io/gorm"
)

var vehicles Vehicles
var response helpers.Response

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *vehicle_repo {
	return &vehicle_repo{grm}
}

func (r *vehicle_repo) FindAll() (*helpers.Response, error) {

	result := r.db.Order("vehicle_id desc").Find(&vehicles)

	if result.Error != nil {
		res := response.ResponseJSON(500, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}

func (r *vehicle_repo) Search(category ...interface{}) (*helpers.Response, error) {

	result := r.db.Where("category = ? AND location = ? ", category[0], category[1]).Find(&vehicles)

	if result.Error != nil {
		res := response.ResponseJSON(500, vehicles)
		return res, nil
	}
	
	if result.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}

func (r *vehicle_repo) SortByPrice(price *int) (*helpers.Response, error) {

	result := r.db.Order("price desc").Where("CAST(price AS int) > ?", price).Find(&vehicles)

	if result.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	if result.Error != nil {
		res := response.ResponseJSON(500, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}

func (r *vehicle_repo) Popular(rating *int) (*helpers.Response, error) {

	result := r.db.Where("rating >= ?", rating).Order("rating desc").Find(&vehicles)

	if result.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	if result.Error != nil {
		res := response.ResponseJSON(500, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}

func (r *vehicle_repo) Add(data *Vehicle) (*helpers.Response, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	getdata := r.db.First(&vehicles, &data.Vehicle_ID)
	if getdata.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(201, vehicles)
	return res, nil
}

func (r *vehicle_repo) Delete(data *int) (*helpers.Response, error) {

	getdata := r.db.First(&vehicles, &data)
	if getdata.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	result := r.db.Delete(&Vehicle{}, &data)

	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(200, vehicles)
	return res, nil
}

func (r *vehicle_repo) Update(id *int, data *Vehicle) (*helpers.Response, error) {

	result := r.db.Model(&Vehicle{}).Where("vehicle_id = ?", &id).Updates(&Vehicle{Price: data.Price, Status: data.Status})

	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	getdata := r.db.First(&vehicles, &id)
	if getdata.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(201, vehicles)
	return res, nil
}
