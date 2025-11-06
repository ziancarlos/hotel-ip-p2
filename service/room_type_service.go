package service

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository"
	"net/http"

	"gorm.io/gorm"
)

type RoomTypeService interface {
	Create(roomType domain.RoomType) (domain.RoomType, error)
	FindAll() ([]domain.RoomType, error)
	FindById(id int) (domain.RoomType, error)
	Update(roomType domain.RoomType) (domain.RoomType, error)
	Delete(id int) error
}

type RoomTypeServiceImpl struct {
	RoomTypeRepository repository.RoomTypeRepository
	RoomRepository     repository.RoomRepository
	DB                 *gorm.DB
}

func NewRoomTypeService(roomTypeRepository repository.RoomTypeRepository, roomRepository repository.RoomRepository, db *gorm.DB) RoomTypeService {
	return &RoomTypeServiceImpl{
		RoomTypeRepository: roomTypeRepository,
		RoomRepository:     roomRepository,
		DB:                 db,
	}
}

func (s *RoomTypeServiceImpl) Create(roomType domain.RoomType) (domain.RoomType, error) {
	existingRoomType, err := s.RoomTypeRepository.FindByName(s.DB, roomType.Name)
	if err == nil && existingRoomType.ID != 0 {
		return roomType, exception.NewCustomError(http.StatusBadRequest, "Room type name already exists")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return roomType, err
	}

	return s.RoomTypeRepository.Create(s.DB, roomType)
}

func (s *RoomTypeServiceImpl) FindAll() ([]domain.RoomType, error) {
	return s.RoomTypeRepository.FindAll(s.DB)
}

func (s *RoomTypeServiceImpl) FindById(id int) (domain.RoomType, error) {
	roomType, err := s.RoomTypeRepository.FindById(s.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return roomType, exception.NewCustomError(http.StatusNotFound, "Room type not found")
		}
		return roomType, err
	}
	return roomType, nil
}

func (s *RoomTypeServiceImpl) Update(roomType domain.RoomType) (domain.RoomType, error) {
	_, err := s.RoomTypeRepository.FindById(s.DB, roomType.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return roomType, exception.NewCustomError(http.StatusNotFound, "Room type not found")
		}
		return roomType, err
	}

	existingRoomType, err := s.RoomTypeRepository.FindByName(s.DB, roomType.Name)
	if err == nil && existingRoomType.ID != 0 && existingRoomType.ID != roomType.ID {
		return roomType, exception.NewCustomError(http.StatusBadRequest, "Room type name already exists")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return roomType, err
	}

	return s.RoomTypeRepository.Update(s.DB, roomType)
}

func (s *RoomTypeServiceImpl) Delete(id int) error {
	_, err := s.RoomTypeRepository.FindById(s.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewCustomError(http.StatusNotFound, "Room type not found")
		}
		return err
	}

	rooms, err := s.RoomRepository.FindByRoomTypeId(s.DB, id)
	if err != nil {
		return err
	}

	if len(rooms) > 0 {
		return exception.NewCustomError(http.StatusBadRequest, "Cannot delete room type that is being used by rooms")
	}

	return s.RoomTypeRepository.Delete(s.DB, id)
}
