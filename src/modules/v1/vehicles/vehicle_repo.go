package vehicles

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
	"errors"

	"gorm.io/gorm"
)

var response helpers.Response
type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *vehicle_repo {
	return &vehicle_repo{grm}
}

func (repo *vehicle_repo) FindAll() (*models.Vehicles, error) {

	var vehicles models.Vehicles

	result := repo.db.Order("vehicle_id desc").Find(&vehicles)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) FindByID(id int) (*models.Vehicle, error) {

	var vehicles models.Vehicle

	result := repo.db.First(&vehicles, id)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) Search(name string, location string) (*models.Vehicles, error) {

	var vehicles models.Vehicles

	result := repo.db.Where(`vehicles."name" LIKE ? AND vehicles."location" LIKE ?`, "%"+name+"%", "%"+location+"%").Find(&vehicles)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) SortByPrice(price int) (*models.Vehicles, error) {

	var vehicles models.Vehicles

	result := repo.db.Order("CAST(price AS int) desc").Where("CAST(price AS int) > ?", price).Find(&vehicles)

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) SortByType(category string) (*models.Vehicles, error) {

	var vehicles models.Vehicles

	result := repo.db.Order("vehicle_id desc").Where("category = ?", category).Find(&vehicles)

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) SortByLocation(location string) (*models.Vehicles, error) {

	var vehicles models.Vehicles

	result := repo.db.Order("vehicle_id desc").Where("location = ?", location).Find(&vehicles)

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) Popular(rating int) (*models.Vehicles, error) {

	var vehicles models.Vehicles

	result := repo.db.Where("rating >= ?", rating).Order("rating desc").Find(&vehicles)

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) Add(data *models.Vehicle) (*models.Vehicle, error) {

	var vehicles models.Vehicle

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambahkan data")
	}

	getdata := repo.db.First(&vehicles, &data.Vehicle_ID)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) Delete(id int) (*models.Vehicle, error) {

	var vehicles models.Vehicle

	getdata := repo.db.First(&vehicles, id)
	if getdata.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	result := repo.db.Delete(&models.Vehicle{}, id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	return &vehicles, nil
}

func (repo *vehicle_repo) Update(id int, data *models.Vehicle) (*models.Vehicle, error) {

	var vehicles models.Vehicle

	result := repo.db.Model(&models.Vehicle{}).Where("vehicle_id = ?", id).Updates(&models.Vehicle{Name: data.Name, Location: data.Location, Description: data.Description, Price: data.Price, Status: data.Status, Stock: data.Stock, Category: data.Category, Image: data.Image, Rating: data.Rating})

	if result.Error != nil {
		return nil, errors.New("gagal meng-update data")
	}

	getData := repo.db.First(&vehicles, id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &vehicles, nil
}

