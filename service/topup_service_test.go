package service

import (
	"database/sql"
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository/mock"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTopupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	var (
		db  *sql.DB
		err error
	)
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	return gormDB, mockSQL, err
}

func TestTopupService_ProcessWebhook_Success(t *testing.T) {
	mockTopupRepo := new(mock.TopupRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	db, sqlMock, _ := setupTopupMockDB()
	service := NewTopupService(mockTopupRepo, mockUserRepo, db)

	topup := domain.Topup{
		MidtransTransactionID: "TRX-123",
		MidtransOrderID:       "TOPUP-1-123456",
		Amount:                100000,
		Status:                "settlement",
	}

	user := domain.User{
		ID:      1,
		Name:    "John Doe",
		Email:   "john@example.com",
		Balance: 50000,
	}

	updatedUser := domain.User{
		ID:      1,
		Name:    "John Doe",
		Email:   "john@example.com",
		Balance: 150000,
	}

	expectedTopup := domain.Topup{
		ID:                    1,
		UserID:                1,
		MidtransTransactionID: "TRX-123",
		MidtransOrderID:       "TOPUP-1-123456",
		Amount:                100000,
		Status:                "settlement",
	}

	// Mock the transaction behavior
	sqlMock.ExpectBegin()
	mockTopupRepo.On("Create", testifymock.Anything, testifymock.MatchedBy(func(t domain.Topup) bool {
		return t.UserID == 1 && t.Amount == 100000
	})).Return(expectedTopup, nil)
	mockUserRepo.On("FindById", testifymock.Anything, 1).Return(user, nil)
	mockUserRepo.On("Update", testifymock.Anything, testifymock.MatchedBy(func(u domain.User) bool {
		return u.ID == 1 && u.Balance == 150000
	})).Return(updatedUser, nil)
	sqlMock.ExpectCommit()

	result, err := service.ProcessWebhook(topup)

	assert.NoError(t, err)
	assert.Equal(t, expectedTopup.ID, result.ID)
	assert.Equal(t, 1, result.UserID)
}

func TestTopupService_ProcessWebhook_NotSettlement(t *testing.T) {
	mockTopupRepo := new(mock.TopupRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	service := NewTopupService(mockTopupRepo, mockUserRepo, &gorm.DB{})

	topup := domain.Topup{
		MidtransTransactionID: "TRX-123",
		MidtransOrderID:       "TOPUP-1-123456",
		Amount:                100000,
		Status:                "pending",
	}

	result, err := service.ProcessWebhook(topup)

	assert.NoError(t, err)
	assert.Equal(t, domain.Topup{}, result)
	mockTopupRepo.AssertNotCalled(t, "Create")
	mockUserRepo.AssertNotCalled(t, "FindById")
}

func TestTopupService_ProcessWebhook_InvalidOrderIDFormat(t *testing.T) {
	mockTopupRepo := new(mock.TopupRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	service := NewTopupService(mockTopupRepo, mockUserRepo, &gorm.DB{})

	topup := domain.Topup{
		MidtransTransactionID: "TRX-123",
		MidtransOrderID:       "INVALID",
		Amount:                100000,
		Status:                "settlement",
	}

	_, err := service.ProcessWebhook(topup)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "invalid order id format", customErr.Message)
}

func TestTopupService_ProcessWebhook_InvalidUserIDInOrderID(t *testing.T) {
	mockTopupRepo := new(mock.TopupRepositoryMock)
	mockUserRepo := new(mock.UserRepositoryMock)
	service := NewTopupService(mockTopupRepo, mockUserRepo, &gorm.DB{})

	topup := domain.Topup{
		MidtransTransactionID: "TRX-123",
		MidtransOrderID:       "TOPUP-ABC-123456",
		Amount:                100000,
		Status:                "settlement",
	}

	_, err := service.ProcessWebhook(topup)

	assert.Error(t, err)
	customErr, ok := err.(*exception.CustomError)
	assert.True(t, ok)
	assert.Equal(t, "invalid user id in order id", customErr.Message)
}
