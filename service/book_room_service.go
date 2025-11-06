package service

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository"
	"net/http"

	"gorm.io/gorm"
)

type BookRoomService interface {
	Create(bookRoom domain.BookRoom) (domain.BookRoom, error)
	FindByUserId(userId int) ([]domain.BookRoom, error)
}

type BookRoomServiceImpl struct {
	BookRoomRepository repository.BookRoomRepository
	RoomRepository     repository.RoomRepository
	UserRepository     repository.UserRepository
	DB                 *gorm.DB
}

func NewBookRoomService(bookRoomRepository repository.BookRoomRepository, roomRepository repository.RoomRepository, userRepository repository.UserRepository, db *gorm.DB) BookRoomService {
	return &BookRoomServiceImpl{
		BookRoomRepository: bookRoomRepository,
		RoomRepository:     roomRepository,
		UserRepository:     userRepository,
		DB:                 db,
	}
}

func (s *BookRoomServiceImpl) Create(bookRoom domain.BookRoom) (domain.BookRoom, error) {
	var result domain.BookRoom

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		room, err := s.RoomRepository.FindById(tx, bookRoom.RoomID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return exception.NewCustomError(http.StatusNotFound, "Room not found")
			}
			return err
		}

		user, err := s.UserRepository.FindById(tx, bookRoom.UserID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return exception.NewCustomError(http.StatusNotFound, "User not found")
			}
			return err
		}

		existingBooking, err := s.BookRoomRepository.FindByRoomIdAndDate(tx, bookRoom.RoomID, bookRoom.Date)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if existingBooking.ID != 0 {
			return exception.NewCustomError(http.StatusBadRequest, "Room is already booked for this date")
		}

		if user.Balance < room.RoomType.Price {
			return exception.NewCustomError(http.StatusBadRequest, "Insufficient balance")
		}

		bookRoom.Price = room.RoomType.Price

		user.Balance = user.Balance - room.RoomType.Price
		_, err = s.UserRepository.Update(tx, user)
		if err != nil {
			return err
		}

		result, err = s.BookRoomRepository.Create(tx, bookRoom)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func (s *BookRoomServiceImpl) FindByUserId(userId int) ([]domain.BookRoom, error) {
	return s.BookRoomRepository.FindByUserId(s.DB, userId)
}
