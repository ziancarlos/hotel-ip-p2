package repository

import (
	"hotel_ip-p2/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(db *gorm.DB, user domain.User) (domain.User, error)
	FindByEmail(db *gorm.DB, email string) (domain.User, error)
	FindById(db *gorm.DB, id int) (domain.User, error)
	Update(db *gorm.DB, user domain.User) (domain.User, error)
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) Register(db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindByEmail(db *gorm.DB, email string) (domain.User, error) {
	var user domain.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindById(db *gorm.DB, id int) (domain.User, error) {
	var user domain.User
	err := db.First(&user, id).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) Update(db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":    user.Name,
		"email":   user.Email,
		"balance": user.Balance,
	}).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
