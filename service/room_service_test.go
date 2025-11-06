package service

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRoomService_Create_Success(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	room := domain.Room{
		RoomTypeID: 1,
		RoomNumber: "101",
	}

	roomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe",
		Price: 500000,
	}

	expectedRoom := domain.Room{
		ID:         1,
		RoomTypeID: 1,
		RoomNumber: "101",
	}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 1).Return(roomType, nil)
	mockRoomRepo.On("FindByRoomNumber", &gorm.DB{}, "101").Return(domain.Room{}, gorm.ErrRecordNotFound)
	mockRoomRepo.On("Create", &gorm.DB{}, room).Return(expectedRoom, nil)

	result, err := service.Create(room)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoom.ID, result.ID)
	assert.Equal(t, expectedRoom.RoomNumber, result.RoomNumber)
	mockRoomRepo.AssertExpectations(t)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomService_Create_RoomTypeNotFound(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	room := domain.Room{
		RoomTypeID: 999,
		RoomNumber: "101",
	}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 999).Return(domain.RoomType{}, gorm.ErrRecordNotFound)

	_, err := service.Create(room)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room type not found", customErr.Message)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomService_Create_DuplicateRoomNumber(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	room := domain.Room{
		RoomTypeID: 1,
		RoomNumber: "101",
	}

	roomType := domain.RoomType{ID: 1, Name: "Deluxe", Price: 500000}
	existingRoom := domain.Room{ID: 1, RoomTypeID: 1, RoomNumber: "101"}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 1).Return(roomType, nil)
	mockRoomRepo.On("FindByRoomNumber", &gorm.DB{}, "101").Return(existingRoom, nil)

	_, err := service.Create(room)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room number already exists", customErr.Message)
	mockRoomTypeRepo.AssertExpectations(t)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomService_FindAll_Success(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	expectedRooms := []domain.Room{
		{ID: 1, RoomTypeID: 1, RoomNumber: "101"},
		{ID: 2, RoomTypeID: 1, RoomNumber: "102"},
	}

	mockRoomRepo.On("FindAll", &gorm.DB{}).Return(expectedRooms, nil)

	result, err := service.FindAll()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomService_FindById_Success(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	expectedRoom := domain.Room{
		ID:         1,
		RoomTypeID: 1,
		RoomNumber: "101",
	}

	mockRoomRepo.On("FindById", &gorm.DB{}, 1).Return(expectedRoom, nil)

	result, err := service.FindById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoom.ID, result.ID)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomService_FindById_NotFound(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	mockRoomRepo.On("FindById", &gorm.DB{}, 999).Return(domain.Room{}, gorm.ErrRecordNotFound)

	_, err := service.FindById(999)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room not found", customErr.Message)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomService_Update_Success(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	room := domain.Room{
		ID:         1,
		RoomTypeID: 1,
		RoomNumber: "101A",
	}

	existingRoom := domain.Room{ID: 1, RoomTypeID: 1, RoomNumber: "101"}
	roomType := domain.RoomType{ID: 1, Name: "Deluxe", Price: 500000}

	mockRoomRepo.On("FindById", &gorm.DB{}, 1).Return(existingRoom, nil)
	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 1).Return(roomType, nil)
	mockRoomRepo.On("FindByRoomNumber", &gorm.DB{}, "101A").Return(domain.Room{}, gorm.ErrRecordNotFound)
	mockRoomRepo.On("Update", &gorm.DB{}, room).Return(room, nil)

	result, err := service.Update(room)

	assert.NoError(t, err)
	assert.Equal(t, room.RoomNumber, result.RoomNumber)
	mockRoomRepo.AssertExpectations(t)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomService_Update_NotFound(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	room := domain.Room{
		ID:         999,
		RoomTypeID: 1,
		RoomNumber: "101",
	}

	mockRoomRepo.On("FindById", &gorm.DB{}, 999).Return(domain.Room{}, gorm.ErrRecordNotFound)

	_, err := service.Update(room)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room not found", customErr.Message)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomService_Delete_Success(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	existingRoom := domain.Room{ID: 1, RoomTypeID: 1, RoomNumber: "101"}

	mockRoomRepo.On("FindById", &gorm.DB{}, 1).Return(existingRoom, nil)
	mockRoomRepo.On("Delete", &gorm.DB{}, 1).Return(nil)

	err := service.Delete(1)

	assert.NoError(t, err)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomService_Delete_NotFound(t *testing.T) {
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	service := NewRoomService(mockRoomRepo, mockRoomTypeRepo, &gorm.DB{})

	mockRoomRepo.On("FindById", &gorm.DB{}, 999).Return(domain.Room{}, gorm.ErrRecordNotFound)

	err := service.Delete(999)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room not found", customErr.Message)
	mockRoomRepo.AssertExpectations(t)
}
