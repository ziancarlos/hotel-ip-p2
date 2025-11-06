package service

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository"
	"net/http"

	"gorm.io/gorm"
)

type RoomService interface {
	Create(room domain.Room) (domain.Room, error)
	FindAll() ([]domain.Room, error)
	FindById(id int) (domain.Room, error)
	Update(room domain.Room) (domain.Room, error)
	Delete(id int) error
}

type RoomServiceImpl struct {
	RoomRepository     repository.RoomRepository
	RoomTypeRepository repository.RoomTypeRepository
	DB                 *gorm.DB
}

func NewRoomService(roomRepository repository.RoomRepository, roomTypeRepository repository.RoomTypeRepository, db *gorm.DB) RoomService {
	return &RoomServiceImpl{
		RoomRepository:     roomRepository,
		RoomTypeRepository: roomTypeRepository,
		DB:                 db,
	}
}

func (s *RoomServiceImpl) Create(room domain.Room) (domain.Room, error) {
	_, err := s.RoomTypeRepository.FindById(s.DB, room.RoomTypeID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return room, exception.NewCustomError(http.StatusNotFound, "Room type not found")
		}
		return room, err
	}

	// Check if room number already exists
	existingRoom, err := s.RoomRepository.FindByRoomNumber(s.DB, room.RoomNumber)
	if err == nil && existingRoom.ID != 0 {
		return room, exception.NewCustomError(http.StatusBadRequest, "Room number already exists")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return room, err
	}

	return s.RoomRepository.Create(s.DB, room)
}

func (s *RoomServiceImpl) FindAll() ([]domain.Room, error) {
	return s.RoomRepository.FindAll(s.DB)
}

func (s *RoomServiceImpl) FindById(id int) (domain.Room, error) {
	room, err := s.RoomRepository.FindById(s.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return room, exception.NewCustomError(http.StatusNotFound, "Room not found")
		}
		return room, err
	}
	return room, nil
}

func (s *RoomServiceImpl) Update(room domain.Room) (domain.Room, error) {
	_, err := s.RoomRepository.FindById(s.DB, room.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return room, exception.NewCustomError(http.StatusNotFound, "Room not found")
		}
		return room, err
	}

	_, err = s.RoomTypeRepository.FindById(s.DB, room.RoomTypeID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return room, exception.NewCustomError(http.StatusNotFound, "Room type not found")
		}
		return room, err
	}

	existingRoom, err := s.RoomRepository.FindByRoomNumber(s.DB, room.RoomNumber)
	if err == nil && existingRoom.ID != 0 && existingRoom.ID != room.ID {
		return room, exception.NewCustomError(http.StatusBadRequest, "Room number already exists")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return room, err
	}

	return s.RoomRepository.Update(s.DB, room)
}

func (s *RoomServiceImpl) Delete(id int) error {
	_, err := s.RoomRepository.FindById(s.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewCustomError(http.StatusNotFound, "Room not found")
		}
		return err
	}

	return s.RoomRepository.Delete(s.DB, id)
}
