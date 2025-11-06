package service

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRoomTypeService_Create_Success(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	roomType := domain.RoomType{
		Name:  "Deluxe",
		Price: 500000,
	}

	expectedRoomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe",
		Price: 500000,
	}

	mockRoomTypeRepo.On("FindByName", &gorm.DB{}, "Deluxe").Return(domain.RoomType{}, gorm.ErrRecordNotFound)
	mockRoomTypeRepo.On("Create", &gorm.DB{}, roomType).Return(expectedRoomType, nil)

	result, err := service.Create(roomType)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoomType.ID, result.ID)
	assert.Equal(t, expectedRoomType.Name, result.Name)
	assert.Equal(t, expectedRoomType.Price, result.Price)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomTypeService_Create_DuplicateName(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	roomType := domain.RoomType{
		Name:  "Deluxe",
		Price: 500000,
	}

	existingRoomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe",
		Price: 450000,
	}

	mockRoomTypeRepo.On("FindByName", &gorm.DB{}, "Deluxe").Return(existingRoomType, nil)

	_, err := service.Create(roomType)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room type name already exists", customErr.Message)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomTypeService_FindAll_Success(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	expectedRoomTypes := []domain.RoomType{
		{ID: 1, Name: "Standard", Price: 300000},
		{ID: 2, Name: "Deluxe", Price: 500000},
	}

	mockRoomTypeRepo.On("FindAll", &gorm.DB{}).Return(expectedRoomTypes, nil)

	result, err := service.FindAll()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedRoomTypes[0].Name, result[0].Name)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomTypeService_FindById_Success(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	expectedRoomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe",
		Price: 500000,
	}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 1).Return(expectedRoomType, nil)

	result, err := service.FindById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoomType.ID, result.ID)
	assert.Equal(t, expectedRoomType.Name, result.Name)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomTypeService_FindById_NotFound(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 999).Return(domain.RoomType{}, gorm.ErrRecordNotFound)

	_, err := service.FindById(999)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room type not found", customErr.Message)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomTypeService_Update_Success(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	roomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe Updated",
		Price: 550000,
	}

	existingRoomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe",
		Price: 500000,
	}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 1).Return(existingRoomType, nil)
	mockRoomTypeRepo.On("FindByName", &gorm.DB{}, "Deluxe Updated").Return(domain.RoomType{}, gorm.ErrRecordNotFound)
	mockRoomTypeRepo.On("Update", &gorm.DB{}, roomType).Return(roomType, nil)

	result, err := service.Update(roomType)

	assert.NoError(t, err)
	assert.Equal(t, roomType.Name, result.Name)
	assert.Equal(t, roomType.Price, result.Price)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomTypeService_Update_NotFound(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	roomType := domain.RoomType{
		ID:    999,
		Name:  "Deluxe",
		Price: 500000,
	}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 999).Return(domain.RoomType{}, gorm.ErrRecordNotFound)

	_, err := service.Update(roomType)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room type not found", customErr.Message)
	mockRoomTypeRepo.AssertExpectations(t)
}

func TestRoomTypeService_Delete_Success(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	existingRoomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe",
		Price: 500000,
	}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 1).Return(existingRoomType, nil)
	mockRoomRepo.On("FindByRoomTypeId", &gorm.DB{}, 1).Return([]domain.Room{}, nil)
	mockRoomTypeRepo.On("Delete", &gorm.DB{}, 1).Return(nil)

	err := service.Delete(1)

	assert.NoError(t, err)
	mockRoomTypeRepo.AssertExpectations(t)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomTypeService_Delete_HasRooms(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	existingRoomType := domain.RoomType{
		ID:    1,
		Name:  "Deluxe",
		Price: 500000,
	}

	rooms := []domain.Room{
		{ID: 1, RoomTypeID: 1, RoomNumber: "101"},
	}

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 1).Return(existingRoomType, nil)
	mockRoomRepo.On("FindByRoomTypeId", &gorm.DB{}, 1).Return(rooms, nil)

	err := service.Delete(1)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Cannot delete room type that is being used by rooms", customErr.Message)
	mockRoomTypeRepo.AssertExpectations(t)
	mockRoomRepo.AssertExpectations(t)
}

func TestRoomTypeService_Delete_NotFound(t *testing.T) {
	mockRoomTypeRepo := new(mock.RoomTypeRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	service := NewRoomTypeService(mockRoomTypeRepo, mockRoomRepo, &gorm.DB{})

	mockRoomTypeRepo.On("FindById", &gorm.DB{}, 999).Return(domain.RoomType{}, gorm.ErrRecordNotFound)

	err := service.Delete(999)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room type not found", customErr.Message)
	mockRoomTypeRepo.AssertExpectations(t)
}
