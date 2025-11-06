package mock

import (
	"hotel_ip-p2/model/domain"
	"time"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type TopupRepositoryMock struct {
	mock.Mock
}

func (m *TopupRepositoryMock) Create(db *gorm.DB, topup domain.Topup) (domain.Topup, error) {
	args := m.Called(db, topup)
	return args.Get(0).(domain.Topup), args.Error(1)
}

func (m *TopupRepositoryMock) FindByOrderID(db *gorm.DB, orderID string) (domain.Topup, error) {
	args := m.Called(db, orderID)
	return args.Get(0).(domain.Topup), args.Error(1)
}

func (m *TopupRepositoryMock) FindByMidtransOrderID(db *gorm.DB, orderID string) (domain.Topup, error) {
	args := m.Called(db, orderID)
	return args.Get(0).(domain.Topup), args.Error(1)
}

func (m *TopupRepositoryMock) FindByUserId(db *gorm.DB, userId int) ([]domain.Topup, error) {
	args := m.Called(db, userId)
	return args.Get(0).([]domain.Topup), args.Error(1)
}

type RoomTypeRepositoryMock struct {
	mock.Mock
}

func (m *RoomTypeRepositoryMock) Create(db *gorm.DB, roomType domain.RoomType) (domain.RoomType, error) {
	args := m.Called(db, roomType)
	return args.Get(0).(domain.RoomType), args.Error(1)
}

func (m *RoomTypeRepositoryMock) FindAll(db *gorm.DB) ([]domain.RoomType, error) {
	args := m.Called(db)
	return args.Get(0).([]domain.RoomType), args.Error(1)
}

func (m *RoomTypeRepositoryMock) FindById(db *gorm.DB, id int) (domain.RoomType, error) {
	args := m.Called(db, id)
	return args.Get(0).(domain.RoomType), args.Error(1)
}

func (m *RoomTypeRepositoryMock) FindByName(db *gorm.DB, name string) (domain.RoomType, error) {
	args := m.Called(db, name)
	return args.Get(0).(domain.RoomType), args.Error(1)
}

func (m *RoomTypeRepositoryMock) Update(db *gorm.DB, roomType domain.RoomType) (domain.RoomType, error) {
	args := m.Called(db, roomType)
	return args.Get(0).(domain.RoomType), args.Error(1)
}

func (m *RoomTypeRepositoryMock) Delete(db *gorm.DB, id int) error {
	args := m.Called(db, id)
	return args.Error(0)
}

type RoomRepositoryMock struct {
	mock.Mock
}

func (m *RoomRepositoryMock) Create(db *gorm.DB, room domain.Room) (domain.Room, error) {
	args := m.Called(db, room)
	return args.Get(0).(domain.Room), args.Error(1)
}

func (m *RoomRepositoryMock) FindAll(db *gorm.DB) ([]domain.Room, error) {
	args := m.Called(db)
	return args.Get(0).([]domain.Room), args.Error(1)
}

func (m *RoomRepositoryMock) FindById(db *gorm.DB, id int) (domain.Room, error) {
	args := m.Called(db, id)
	return args.Get(0).(domain.Room), args.Error(1)
}

func (m *RoomRepositoryMock) FindByRoomNumber(db *gorm.DB, roomNumber string) (domain.Room, error) {
	args := m.Called(db, roomNumber)
	return args.Get(0).(domain.Room), args.Error(1)
}

func (m *RoomRepositoryMock) FindByRoomTypeId(db *gorm.DB, roomTypeId int) ([]domain.Room, error) {
	args := m.Called(db, roomTypeId)
	return args.Get(0).([]domain.Room), args.Error(1)
}

func (m *RoomRepositoryMock) Update(db *gorm.DB, room domain.Room) (domain.Room, error) {
	args := m.Called(db, room)
	return args.Get(0).(domain.Room), args.Error(1)
}

func (m *RoomRepositoryMock) Delete(db *gorm.DB, id int) error {
	args := m.Called(db, id)
	return args.Error(0)
}

type BookRoomRepositoryMock struct {
	mock.Mock
}

func (m *BookRoomRepositoryMock) Create(db *gorm.DB, bookRoom domain.BookRoom) (domain.BookRoom, error) {
	args := m.Called(db, bookRoom)
	return args.Get(0).(domain.BookRoom), args.Error(1)
}

func (m *BookRoomRepositoryMock) FindByUserId(db *gorm.DB, userId int) ([]domain.BookRoom, error) {
	args := m.Called(db, userId)
	return args.Get(0).([]domain.BookRoom), args.Error(1)
}

func (m *BookRoomRepositoryMock) FindByRoomIdAndDate(db *gorm.DB, roomId int, date time.Time) (domain.BookRoom, error) {
	args := m.Called(db, roomId, date)
	return args.Get(0).(domain.BookRoom), args.Error(1)
}
