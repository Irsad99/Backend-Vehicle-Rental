package histories

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/helpers"
	"errors"

	"gorm.io/gorm"
)

var response helpers.Response

type history_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *history_repo {
	return &history_repo{grm}
}

func (repo *history_repo) FindAll() (*models.Histories, error) {

	var histories models.Histories

	result := repo.db.Order("history_id desc").Find(&histories)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &histories, nil
}

func (repo *history_repo) FindHistoryByID(id int) (*models.Results, error) {

	var results models.Results

	result := repo.db.Raw(
		" select u.name as users, v.name as vehicle, h.start_date , h.end_date , h.prepayment , h.status, h.quantity"+
			" from histories h , users u , vehicles v"+
			" where u.user_id = ?"+
			" and u.user_id = h.id_user and v.vehicle_id = h.id_vehicle", id).Scan(&results)

	if result.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &results, nil
}

func (repo *history_repo) SortByStart() (*models.Histories, error) {

	var histories models.Histories

	result := repo.db.Order("start_date").Find(&histories)

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}

	return &histories, nil
}

func (repo *history_repo) Add(data *models.History) (*models.History, error) {

	var histories models.History

	result := repo.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menambahkan data")
	}

	getData := repo.db.First(&histories, &data.History_Id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &histories, nil
}

func (repo *history_repo) Delete(id int) (*models.History, error) {

	var histories models.History

	getData := repo.db.First(&histories, id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	result := repo.db.Delete(&models.History{}, id)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	return &histories, nil
}

func (repo *history_repo) Update(id int, status string) (*models.History, error) {

	var histories models.History

	result := repo.db.Model(&models.History{}).Where("history_id = ?", id).Update("status", status)

	if result.Error != nil {
		return nil, errors.New("gagal meng-update data")
	}

	getData := repo.db.First(&histories, &id)
	if getData.RowsAffected < 1 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &histories, nil
}
