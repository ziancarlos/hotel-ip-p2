package mock

import (
	"hotel_ip-p2/model/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Register(db *gorm.DB, user domain.User) (domain.User, error) {
	args := m.Called(db, user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByEmail(db *gorm.DB, email string) (domain.User, error) {
	args := m.Called(db, email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) FindById(db *gorm.DB, id int) (domain.User, error) {
	args := m.Called(db, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(db *gorm.DB, user domain.User) (domain.User, error) {
	args := m.Called(db, user)
	return args.Get(0).(domain.User), args.Error(1)
}
