package service

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user domain.User) (domain.User, error)
	Login(email, password string) (domain.User, error)
	GetById(id int) (domain.User, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}
func (service *userServiceImpl) Register(user domain.User) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, exception.NewCustomError(http.StatusInternalServerError, "failed to hash password")
	}
	user.Password = string(hashedPassword)

	result, err := service.UserRepository.Register(service.DB, user)
	if err != nil {
		return domain.User{}, exception.NewCustomError(http.StatusBadRequest, "email already exists")
	}
	return result, nil
}

func (service *userServiceImpl) Login(email, password string) (domain.User, error) {
	user, err := service.UserRepository.FindByEmail(service.DB, email)
	if err != nil {
		return domain.User{}, exception.NewCustomError(http.StatusNotFound, "user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, exception.NewCustomError(http.StatusUnauthorized, "invalid credentials")
	}

	return user, nil
}

func (service *userServiceImpl) GetById(id int) (domain.User, error) {
	user, err := service.UserRepository.FindById(service.DB, id)
	if err != nil {
		return domain.User{}, exception.NewCustomError(http.StatusNotFound, "user not found")
	}
	return user, nil
}
