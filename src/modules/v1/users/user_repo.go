package users

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

var response helpers.Response

type user_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *user_repo {
	return &user_repo{grm}
}

func (repo *user_repo) FindAll() (*models.Users, error) {

	var users models.Users

	result := repo.db.Order("user_id desc").Find(&users)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &users, nil
}

func (repo *user_repo) FindByEmail(email string) (*models.User, error) {

	var users models.User

	result := repo.db.First(&users, "email = ?", email)
	if result.RowsAffected < 1 {
		err := json.Unmarshal([]byte("email"), &users)
		return nil, err
	}

	if result.Error != nil {
		err := json.Unmarshal([]byte("Tidak dapat mengambil data"), &users)
		return nil, err
	}

	return &users, nil
}

func (repo *user_repo) Add(data *models.User) (*models.User, error) {

	var users models.User

	getEmail := repo.db.Where("email = ?", &data.Email).First(&users)
	if getEmail.RowsAffected != 0 {
		return nil, errors.New("email sudah terdaftar")
	}

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambahkan data")
	}

	getData := repo.db.First(&users, &data.User_ID)
	if getData.RowsAffected < 1 {
		return nil, errors.New("email sudah terdaftar")
	}

	return &users, nil
}

func (repo *user_repo) Delete(id int) (*models.User, error) {

	var users models.User

	getData := repo.db.First(&users, id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	result := repo.db.Delete(&models.User{}, &id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	return &users, nil
}

func (repo *user_repo) Update(id int, data *models.User) (*models.User, error) {

	var users *models.User

	result := repo.db.Model(&models.User{}).Where("user_id = ?", &id).Updates(&models.User{Name: data.Name, Gender: data.Gender, Email: data.Email, Phone: data.Phone, Birth: data.Birth, Address: data.Address, Password: data.Password})

	if result.Error != nil {
		return nil, errors.New("gagal meng-update data")
	}

	getData := repo.db.First(&users, id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return users, nil
}
