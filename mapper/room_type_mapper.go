package mapper

import (
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/model/web/response"
)

func ToRoomTypeDomain(req request.RoomTypeRequest) domain.RoomType {
	return domain.RoomType{
		Name:  req.Name,
		Price: req.Price,
	}
}

func ToRoomTypeResponse(roomType domain.RoomType) response.RoomTypeResponse {
	return response.RoomTypeResponse{
		ID:    roomType.ID,
		Name:  roomType.Name,
		Price: roomType.Price,
	}
}

func ToRoomTypeResponses(roomTypes []domain.RoomType) []response.RoomTypeResponse {
	var responses []response.RoomTypeResponse
	for _, roomType := range roomTypes {
		responses = append(responses, ToRoomTypeResponse(roomType))
	}
	return responses
}
