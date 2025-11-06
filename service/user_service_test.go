package service

import (
	"errors"
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestUserService_Register_Success(t *testing.T) {
	mockRepo := new(mock.UserRepositoryMock)
	service := NewUserService(mockRepo, &gorm.DB{})

	user := domain.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	expectedUser := domain.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("Register", &gorm.DB{}, testifymock.MatchedBy(func(u domain.User) bool {
		return u.Name == "John Doe" && u.Email == "john@example.com"
	})).Return(expectedUser, nil)

	result, err := service.Register(user)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Name, result.Name)
	assert.Equal(t, expectedUser.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Register_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(mock.UserRepositoryMock)
	service := NewUserService(mockRepo, &gorm.DB{})

	user := domain.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	mockRepo.On("Register", &gorm.DB{}, testifymock.MatchedBy(func(u domain.User) bool {
		return u.Name == "John Doe" && u.Email == "john@example.com"
	})).Return(domain.User{}, errors.New("duplicate email"))

	result, err := service.Register(user)

	assert.Error(t, err)
	assert.Equal(t, domain.User{}, result)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "email already exists", customErr.Message)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Login_Success(t *testing.T) {
	mockRepo := new(mock.UserRepositoryMock)
	service := NewUserService(mockRepo, &gorm.DB{})

	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	expectedUser := domain.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: string(hashedPassword),
		Balance:  0,
	}

	mockRepo.On("FindByEmail", &gorm.DB{}, "john@example.com").Return(expectedUser, nil)

	result, err := service.Login("john@example.com", password)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Login_UserNotFound(t *testing.T) {
	mockRepo := new(mock.UserRepositoryMock)
	service := NewUserService(mockRepo, &gorm.DB{})

	mockRepo.On("FindByEmail", &gorm.DB{}, "notfound@example.com").Return(domain.User{}, gorm.ErrRecordNotFound)

	result, err := service.Login("notfound@example.com", "password123")

	assert.Error(t, err)
	assert.Equal(t, domain.User{}, result)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "user not found", customErr.Message)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Login_InvalidPassword(t *testing.T) {
	mockRepo := new(mock.UserRepositoryMock)
	service := NewUserService(mockRepo, &gorm.DB{})

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	existingUser := domain.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: string(hashedPassword),
	}

	mockRepo.On("FindByEmail", &gorm.DB{}, "john@example.com").Return(existingUser, nil)

	result, err := service.Login("john@example.com", "wrongpassword")

	assert.Error(t, err)
	assert.Equal(t, domain.User{}, result)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "invalid credentials", customErr.Message)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetById_Success(t *testing.T) {
	mockRepo := new(mock.UserRepositoryMock)
	service := NewUserService(mockRepo, &gorm.DB{})

	expectedUser := domain.User{
		ID:      1,
		Name:    "John Doe",
		Email:   "john@example.com",
		Balance: 100.0,
	}

	mockRepo.On("FindById", &gorm.DB{}, 1).Return(expectedUser, nil)

	result, err := service.GetById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetById_UserNotFound(t *testing.T) {
	mockRepo := new(mock.UserRepositoryMock)
	service := NewUserService(mockRepo, &gorm.DB{})

	mockRepo.On("FindById", &gorm.DB{}, 999).Return(domain.User{}, gorm.ErrRecordNotFound)

	result, err := service.GetById(999)

	assert.Error(t, err)
	assert.Equal(t, domain.User{}, result)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "user not found", customErr.Message)
	mockRepo.AssertExpectations(t)
}
