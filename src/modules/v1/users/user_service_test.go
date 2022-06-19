package users

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/modules/v1/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var dataUser models.Users

	repo.Mock.On("FindAll").Return(&dataUser, nil)
	data, err := service.FindAll()

	users := data.Message

	assert.Equal(t, "OK", users, "Expect len status = 200")
	assert.Nil(t, err)
}

func TestFindByEmail(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var modelMock = models.User{
		Email: "lutfi123@gmail.com",
	}

	repo.Mock.On("FindByEmail", "lutfi123@gmail.com").Return(&modelMock, nil)
	data, err := service.FindByEmail("lutfi123@gmail.com")

	result := data.Data.(*models.User)

	assert.Equal(t, "lutfi123@gmail.com", result.Email, "Expect len email = lutfi123@gmail.com")
	assert.Nil(t, err)
}

func TestAdd(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var AddMock = models.User{
		Name:     "Fatimah",
		Gender:   "Perempuan",
		Email:    "fatim@gmail.com",
		Phone:    "087750408756",
		Address:  "Magetan",
		Birth:    "2001-07-10",
		Password: "Fatim123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:     "user",
	}

	repo.Mock.On("Add", &AddMock).Return(&AddMock, nil)
	data, err := service.Save(&AddMock)

	result := data.Data.(*models.User)

	assert.Equal(t, "Fatimah", result.Name, "Expect len email = lutfi123@gmail.com")
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var deleteMock = models.User{
		User_ID : 25,
	}

	repo.Mock.On("Delete", 25).Return(&deleteMock, nil)
	data, err := service.Delete(25)

	result := data.Message

	assert.Equal(t, "OK", result, "Expect Message  = OK")
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var UpdateMock = models.User{
		Name:     "Fatimah",
		Gender:   "Perempuan",
		Email:    "fatim@gmail.com",
		Phone:    "087750408756",
		Address:  "Aceh",
		Birth:    "2001-07-10",
		Password: "Fatim123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:     "user",
	}

	repo.Mock.On("Update", 1, &UpdateMock).Return(&UpdateMock, nil)
	data, err := service.Update(1, &UpdateMock)

	result := data.Data.(*models.User)

	assert.Equal(t, "Aceh", result.Address, "Expect Address  = Aceh")
	assert.Nil(t, err)
}
