package repository

import (
	"hotel_ip-p2/model/domain"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(db *gorm.DB, room domain.Room) (domain.Room, error)
	FindAll(db *gorm.DB) ([]domain.Room, error)
	FindById(db *gorm.DB, id int) (domain.Room, error)
	FindByRoomNumber(db *gorm.DB, roomNumber string) (domain.Room, error)
	Update(db *gorm.DB, room domain.Room) (domain.Room, error)
	Delete(db *gorm.DB, id int) error
	FindByRoomTypeId(db *gorm.DB, roomTypeId int) ([]domain.Room, error)
}

type RoomRepositoryImpl struct{}

func NewRoomRepository() RoomRepository {
	return &RoomRepositoryImpl{}
}

func (r *RoomRepositoryImpl) Create(db *gorm.DB, room domain.Room) (domain.Room, error) {
	err := db.Create(&room).Error
	if err != nil {
		return room, err
	}
	err = db.Preload("RoomType").First(&room, room.ID).Error
	return room, err
}

func (r *RoomRepositoryImpl) FindAll(db *gorm.DB) ([]domain.Room, error) {
	var rooms []domain.Room
	err := db.Preload("RoomType").Find(&rooms).Error
	return rooms, err
}
func (r *RoomRepositoryImpl) FindById(db *gorm.DB, id int) (domain.Room, error) {
	var room domain.Room
	err := db.Preload("RoomType").First(&room, id).Error
	return room, err
}

func (r *RoomRepositoryImpl) FindByRoomNumber(db *gorm.DB, roomNumber string) (domain.Room, error) {
	var room domain.Room
	err := db.Where("room_number = ?", roomNumber).First(&room).Error
	return room, err
}

func (r *RoomRepositoryImpl) Update(db *gorm.DB, room domain.Room) (domain.Room, error) {
	err := db.Save(&room).Error
	if err != nil {
		return room, err
	}
	err = db.Preload("RoomType").First(&room, room.ID).Error
	return room, err
}

func (r *RoomRepositoryImpl) Delete(db *gorm.DB, id int) error {
	return db.Delete(&domain.Room{}, id).Error
}

func (r *RoomRepositoryImpl) FindByRoomTypeId(db *gorm.DB, roomTypeId int) ([]domain.Room, error) {
	var rooms []domain.Room
	err := db.Where("room_type_id = ?", roomTypeId).Find(&rooms).Error
	return rooms, err
}
