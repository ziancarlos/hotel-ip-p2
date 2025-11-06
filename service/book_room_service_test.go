package service

import (
	"database/sql"
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository/mock"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	var (
		db  *sql.DB
		err error
	)
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	return gormDB, mock, err
}

func TestBookRoomService_Create_Success(t *testing.T) {
	mockBookRoomRepo := new(mock.BookRoomRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)

	db, sqlMock, _ := setupMockDB()
	service := NewBookRoomService(mockBookRoomRepo, mockRoomRepo, mockUserRepo, db)

	bookingDate := time.Now().AddDate(0, 0, 1)
	bookRoom := domain.BookRoom{
		RoomID: 1,
		UserID: 1,
		Date:   bookingDate,
	}

	room := domain.Room{
		ID:         1,
		RoomTypeID: 1,
		RoomNumber: "101",
		RoomType: domain.RoomType{
			ID:    1,
			Name:  "Deluxe",
			Price: 500000,
		},
	}

	user := domain.User{
		ID:      1,
		Name:    "John Doe",
		Email:   "john@example.com",
		Balance: 600000,
	}

	updatedUser := domain.User{
		ID:      1,
		Name:    "John Doe",
		Email:   "john@example.com",
		Balance: 100000,
	}

	expectedBooking := domain.BookRoom{
		ID:     1,
		RoomID: 1,
		UserID: 1,
		Date:   bookingDate,
		Price:  500000,
	}

	// Mock transaction
	sqlMock.ExpectBegin()
	mockRoomRepo.On("FindById", testifymock.Anything, 1).Return(room, nil)
	mockUserRepo.On("FindById", testifymock.Anything, 1).Return(user, nil)
	mockBookRoomRepo.On("FindByRoomIdAndDate", testifymock.Anything, 1, bookingDate).Return(domain.BookRoom{}, gorm.ErrRecordNotFound)
	mockUserRepo.On("Update", testifymock.Anything, testifymock.MatchedBy(func(u domain.User) bool {
		return u.ID == 1 && u.Balance == 100000
	})).Return(updatedUser, nil)
	mockBookRoomRepo.On("Create", testifymock.Anything, testifymock.MatchedBy(func(b domain.BookRoom) bool {
		return b.RoomID == 1 && b.UserID == 1 && b.Price == 500000
	})).Return(expectedBooking, nil)
	sqlMock.ExpectCommit()

	result, err := service.Create(bookRoom)

	assert.NoError(t, err)
	assert.Equal(t, expectedBooking.ID, result.ID)
	assert.Equal(t, float64(500000), result.Price)
}

func TestBookRoomService_Create_RoomNotFound(t *testing.T) {
	mockBookRoomRepo := new(mock.BookRoomRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	db, sqlMock, _ := setupMockDB()
	service := NewBookRoomService(mockBookRoomRepo, mockRoomRepo, mockUserRepo, db)

	bookRoom := domain.BookRoom{
		RoomID: 999,
		UserID: 1,
		Date:   time.Now(),
	}

	sqlMock.ExpectBegin()
	mockRoomRepo.On("FindById", testifymock.Anything, 999).Return(domain.Room{}, gorm.ErrRecordNotFound)
	sqlMock.ExpectRollback()

	_, err := service.Create(bookRoom)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room not found", customErr.Message)
}

func TestBookRoomService_Create_UserNotFound(t *testing.T) {
	mockBookRoomRepo := new(mock.BookRoomRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	db, sqlMock, _ := setupMockDB()
	service := NewBookRoomService(mockBookRoomRepo, mockRoomRepo, mockUserRepo, db)

	bookRoom := domain.BookRoom{
		RoomID: 1,
		UserID: 999,
		Date:   time.Now(),
	}

	room := domain.Room{
		ID:         1,
		RoomTypeID: 1,
		RoomNumber: "101",
		RoomType: domain.RoomType{
			ID:    1,
			Name:  "Deluxe",
			Price: 500000,
		},
	}

	sqlMock.ExpectBegin()
	mockRoomRepo.On("FindById", testifymock.Anything, 1).Return(room, nil)
	mockUserRepo.On("FindById", testifymock.Anything, 999).Return(domain.User{}, gorm.ErrRecordNotFound)
	sqlMock.ExpectRollback()

	_, err := service.Create(bookRoom)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "User not found", customErr.Message)
}

func TestBookRoomService_Create_RoomAlreadyBooked(t *testing.T) {
	mockBookRoomRepo := new(mock.BookRoomRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	db, sqlMock, _ := setupMockDB()
	service := NewBookRoomService(mockBookRoomRepo, mockRoomRepo, mockUserRepo, db)

	bookingDate := time.Now().AddDate(0, 0, 1)
	bookRoom := domain.BookRoom{
		RoomID: 1,
		UserID: 1,
		Date:   bookingDate,
	}

	room := domain.Room{
		ID:         1,
		RoomTypeID: 1,
		RoomNumber: "101",
		RoomType: domain.RoomType{
			ID:    1,
			Name:  "Deluxe",
			Price: 500000,
		},
	}

	user := domain.User{
		ID:      1,
		Name:    "John Doe",
		Email:   "john@example.com",
		Balance: 600000,
	}

	existingBooking := domain.BookRoom{
		ID:     1,
		RoomID: 1,
		UserID: 2,
		Date:   bookingDate,
		Price:  500000,
	}

	sqlMock.ExpectBegin()
	mockRoomRepo.On("FindById", testifymock.Anything, 1).Return(room, nil)
	mockUserRepo.On("FindById", testifymock.Anything, 1).Return(user, nil)
	mockBookRoomRepo.On("FindByRoomIdAndDate", testifymock.Anything, 1, bookingDate).Return(existingBooking, nil)
	sqlMock.ExpectRollback()

	_, err := service.Create(bookRoom)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Room is already booked for this date", customErr.Message)
}

func TestBookRoomService_Create_InsufficientBalance(t *testing.T) {
	mockBookRoomRepo := new(mock.BookRoomRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	db, sqlMock, _ := setupMockDB()
	service := NewBookRoomService(mockBookRoomRepo, mockRoomRepo, mockUserRepo, db)

	bookingDate := time.Now().AddDate(0, 0, 1)
	bookRoom := domain.BookRoom{
		RoomID: 1,
		UserID: 1,
		Date:   bookingDate,
	}

	room := domain.Room{
		ID:         1,
		RoomTypeID: 1,
		RoomNumber: "101",
		RoomType: domain.RoomType{
			ID:    1,
			Name:  "Deluxe",
			Price: 500000,
		},
	}

	user := domain.User{
		ID:      1,
		Name:    "John Doe",
		Email:   "john@example.com",
		Balance: 100000, // Insufficient balance
	}

	sqlMock.ExpectBegin()
	mockRoomRepo.On("FindById", testifymock.Anything, 1).Return(room, nil)
	mockUserRepo.On("FindById", testifymock.Anything, 1).Return(user, nil)
	mockBookRoomRepo.On("FindByRoomIdAndDate", testifymock.Anything, 1, bookingDate).Return(domain.BookRoom{}, gorm.ErrRecordNotFound)
	sqlMock.ExpectRollback()

	_, err := service.Create(bookRoom)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "Insufficient balance", customErr.Message)
}

func TestBookRoomService_FindByUserId_Success(t *testing.T) {
	mockBookRoomRepo := new(mock.BookRoomRepositoryMock)
	mockRoomRepo := new(mock.RoomRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	service := NewBookRoomService(mockBookRoomRepo, mockRoomRepo, mockUserRepo, &gorm.DB{})

	expectedBookings := []domain.BookRoom{
		{ID: 1, RoomID: 1, UserID: 1, Date: time.Now(), Price: 500000},
		{ID: 2, RoomID: 2, UserID: 1, Date: time.Now().AddDate(0, 0, 1), Price: 500000},
	}

	mockBookRoomRepo.On("FindByUserId", &gorm.DB{}, 1).Return(expectedBookings, nil)

	result, err := service.FindByUserId(1)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedBookings[0].ID, result[0].ID)
	mockBookRoomRepo.AssertExpectations(t)
}
