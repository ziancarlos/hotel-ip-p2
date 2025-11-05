package service

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/repository"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type TopupService interface {
	ProcessWebhook(topup domain.Topup) (domain.Topup, error)
}

type topupServiceImpl struct {
	TopupRepository repository.TopupRepository
	UserRepository  repository.UserRepository
	DB              *gorm.DB
}

func NewTopupService(topupRepository repository.TopupRepository, userRepository repository.UserRepository, db *gorm.DB) TopupService {
	return &topupServiceImpl{
		TopupRepository: topupRepository,
		UserRepository:  userRepository,
		DB:              db,
	}
}

func (service *topupServiceImpl) ProcessWebhook(topup domain.Topup) (domain.Topup, error) {
	if topup.Status != "settlement" {
		return domain.Topup{}, nil
	}

	parts := strings.Split(topup.MidtransOrderID, "-")
	if len(parts) != 3 {
		return domain.Topup{}, exception.NewCustomError(http.StatusBadRequest, "invalid order id format")
	}

	userID, err := strconv.Atoi(parts[1])
	if err != nil {
		return domain.Topup{}, exception.NewCustomError(http.StatusBadRequest, "invalid user id in order id")
	}

	topup.UserID = userID

	var result domain.Topup

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		result, err = service.TopupRepository.Create(tx, topup)
		if err != nil {
			return exception.NewCustomError(http.StatusBadRequest, "failed to create topup record")
		}

		user, err := service.UserRepository.FindById(tx, userID)
		if err != nil {
			return exception.NewCustomError(http.StatusNotFound, "user not found")
		}

		user.Balance += topup.Amount
		_, err = service.UserRepository.Update(tx, user)
		if err != nil {
			return exception.NewCustomError(http.StatusInternalServerError, "failed to update balance")
		}

		return nil
	})

	if err != nil {
		return domain.Topup{}, err
	}

	return result, nil
}
