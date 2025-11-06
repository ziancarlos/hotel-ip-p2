package repository

import (
	"hotel_ip-p2/model/domain"
	"time"

	"gorm.io/gorm"
)

type BookRoomRepository interface {
	Create(db *gorm.DB, bookRoom domain.BookRoom) (domain.BookRoom, error)
	FindByUserId(db *gorm.DB, userId int) ([]domain.BookRoom, error)
	FindByRoomIdAndDate(db *gorm.DB, roomId int, date time.Time) (domain.BookRoom, error)
}

type BookRoomRepositoryImpl struct{}

func NewBookRoomRepository() BookRoomRepository {
	return &BookRoomRepositoryImpl{}
}

func (r *BookRoomRepositoryImpl) Create(db *gorm.DB, bookRoom domain.BookRoom) (domain.BookRoom, error) {
	err := db.Create(&bookRoom).Error
	if err != nil {
		return bookRoom, err
	}
	err = db.Preload("Room.RoomType").Preload("User").First(&bookRoom, bookRoom.ID).Error
	return bookRoom, err
}

func (r *BookRoomRepositoryImpl) FindByUserId(db *gorm.DB, userId int) ([]domain.BookRoom, error) {
	var bookRooms []domain.BookRoom
	err := db.Preload("Room.RoomType").Preload("User").Where("user_id = ?", userId).Find(&bookRooms).Error
	return bookRooms, err
}

func (r *BookRoomRepositoryImpl) FindByRoomIdAndDate(db *gorm.DB, roomId int, date time.Time) (domain.BookRoom, error) {
	var bookRoom domain.BookRoom
	err := db.Where("room_id = ? AND date = ?", roomId, date).First(&bookRoom).Error
	return bookRoom, err
}
