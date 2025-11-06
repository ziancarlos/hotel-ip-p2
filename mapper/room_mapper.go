package mapper

import (
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/model/web/response"
)

func ToRoomDomain(req request.RoomRequest) domain.Room {
	return domain.Room{
		RoomTypeID: req.RoomTypeID,
		RoomNumber: req.RoomNumber,
	}
}

func ToRoomResponse(room domain.Room) response.RoomResponse {
	return response.RoomResponse{
		ID:         room.ID,
		RoomTypeID: room.RoomTypeID,
		RoomNumber: room.RoomNumber,
		RoomType: response.RoomTypeResponse{
			ID:    room.RoomType.ID,
			Name:  room.RoomType.Name,
			Price: room.RoomType.Price,
		},
	}
}

func ToRoomResponses(rooms []domain.Room) []response.RoomResponse {
	var responses []response.RoomResponse
	for _, room := range rooms {
		responses = append(responses, ToRoomResponse(room))
	}
	return responses
}
