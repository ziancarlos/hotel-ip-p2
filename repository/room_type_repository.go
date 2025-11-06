package repository

import (
	"hotel_ip-p2/model/domain"

	"gorm.io/gorm"
)

type RoomTypeRepository interface {
	Create(db *gorm.DB, roomType domain.RoomType) (domain.RoomType, error)
	FindAll(db *gorm.DB) ([]domain.RoomType, error)
	FindById(db *gorm.DB, id int) (domain.RoomType, error)
	FindByName(db *gorm.DB, name string) (domain.RoomType, error)
	Update(db *gorm.DB, roomType domain.RoomType) (domain.RoomType, error)
	Delete(db *gorm.DB, id int) error
}

type RoomTypeRepositoryImpl struct{}

func NewRoomTypeRepository() RoomTypeRepository {
	return &RoomTypeRepositoryImpl{}
}

func (r *RoomTypeRepositoryImpl) Create(db *gorm.DB, roomType domain.RoomType) (domain.RoomType, error) {
	err := db.Create(&roomType).Error
	return roomType, err
}

func (r *RoomTypeRepositoryImpl) FindAll(db *gorm.DB) ([]domain.RoomType, error) {
	var roomTypes []domain.RoomType
	err := db.Find(&roomTypes).Error
	return roomTypes, err
}
func (r *RoomTypeRepositoryImpl) FindById(db *gorm.DB, id int) (domain.RoomType, error) {
	var roomType domain.RoomType
	err := db.First(&roomType, id).Error
	return roomType, err
}

func (r *RoomTypeRepositoryImpl) FindByName(db *gorm.DB, name string) (domain.RoomType, error) {
	var roomType domain.RoomType
	err := db.Where("name = ?", name).First(&roomType).Error
	return roomType, err
}

func (r *RoomTypeRepositoryImpl) Update(db *gorm.DB, roomType domain.RoomType) (domain.RoomType, error) {
	err := db.Save(&roomType).Error
	return roomType, err
}

func (r *RoomTypeRepositoryImpl) Delete(db *gorm.DB, id int) error {
	return db.Delete(&domain.RoomType{}, id).Error
}
