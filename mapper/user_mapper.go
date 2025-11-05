package mapper

import (
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/model/web/response"
)

func ToUserDomain(req request.UserRequest) domain.User {
	return domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToUserResponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
