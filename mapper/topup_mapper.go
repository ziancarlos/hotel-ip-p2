package mapper

import (
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/model/web/response"
	"strconv"
)

func ToTopupDomain(req request.TopupWebhookRequest) (domain.Topup, error) {
	amount, err := strconv.ParseFloat(req.GrossAmount, 64)
	if err != nil {
		return domain.Topup{}, err
	}

	return domain.Topup{
		MidtransTransactionID: req.TransactionID,
		MidtransOrderID:       req.OrderID,
		Amount:                amount,
		Status:                req.TransactionStatus,
	}, nil
}

func ToTopupResponse(topup domain.Topup) response.TopupResponse {
	return response.TopupResponse{
		ID:                    topup.ID,
		UserID:                topup.UserID,
		MidtransTransactionID: topup.MidtransTransactionID,
		MidtransOrderID:       topup.MidtransOrderID,
		Amount:                topup.Amount,
		Status:                topup.Status,
		CreatedAt:             topup.CreatedAt,
		UpdatedAt:             topup.UpdatedAt,
	}
}
