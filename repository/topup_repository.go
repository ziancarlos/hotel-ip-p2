package repository

import (
	"hotel_ip-p2/model/domain"

	"gorm.io/gorm"
)

type TopupRepository interface {
	Create(db *gorm.DB, topup domain.Topup) (domain.Topup, error)
	FindByOrderID(db *gorm.DB, orderID string) (domain.Topup, error)
}

type topupRepositoryImpl struct {
}

func NewTopupRepository() TopupRepository {
	return &topupRepositoryImpl{}
}

func (repository *topupRepositoryImpl) Create(db *gorm.DB, topup domain.Topup) (domain.Topup, error) {
	err := db.Create(&topup).Error
	if err != nil {
		return domain.Topup{}, err
	}
	return topup, nil
}

func (repository *topupRepositoryImpl) FindByOrderID(db *gorm.DB, orderID string) (domain.Topup, error) {
	var topup domain.Topup
	err := db.Where("midtrans_order_id = ?", orderID).First(&topup).Error
	if err != nil {
		return domain.Topup{}, err
	}
	return topup, nil
}
