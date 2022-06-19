package mocks

import (
	"BackendGo/src/database/gorm/models"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (m *RepoMock) FindAll() (*models.Users, error) {
	args := m.Mock.Called()
	return args.Get(0).(*models.Users), args.Error(1)
}

func (m *RepoMock) FindByEmail(email string) (*models.User, error) {
	args := m.Mock.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *RepoMock) Add(data *models.User) (*models.User, error) {
	args := m.Mock.Called(data)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *RepoMock) Delete(id int) (*models.User, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *RepoMock) Update(id int, data *models.User) (*models.User, error) {
	args := m.Mock.Called(id, data)
	return args.Get(0).(*models.User), args.Error(1)
}
