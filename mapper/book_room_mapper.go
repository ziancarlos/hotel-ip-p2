package mapper

import (
	"hotel_ip-p2/model/domain"
	"hotel_ip-p2/model/web/request"
	"hotel_ip-p2/model/web/response"
	"time"
)

func ToBookRoomDomain(req request.BookRoomRequest, userID int) (domain.BookRoom, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return domain.BookRoom{}, err
	}

	return domain.BookRoom{
		RoomID: req.RoomID,
		UserID: userID,
		Date:   date,
	}, nil
}

func ToBookRoomResponse(bookRoom domain.BookRoom) response.BookRoomResponse {
	return response.BookRoomResponse{
		ID:     bookRoom.ID,
		RoomID: bookRoom.RoomID,
		UserID: bookRoom.UserID,
		Date:   bookRoom.Date.Format("2006-01-02"),
		Price:  bookRoom.Price,
		Room: response.RoomResponse{
			ID:         bookRoom.Room.ID,
			RoomTypeID: bookRoom.Room.RoomTypeID,
			RoomNumber: bookRoom.Room.RoomNumber,
			RoomType: response.RoomTypeResponse{
				ID:    bookRoom.Room.RoomType.ID,
				Name:  bookRoom.Room.RoomType.Name,
				Price: bookRoom.Room.RoomType.Price,
			},
		},
		User: response.UserResponse{
			ID:      bookRoom.User.ID,
			Name:    bookRoom.User.Name,
			Email:   bookRoom.User.Email,
			Balance: bookRoom.User.Balance,
		},
	}
}

func ToBookRoomResponses(bookRooms []domain.BookRoom) []response.BookRoomResponse {
	var responses []response.BookRoomResponse
	for _, bookRoom := range bookRooms {
		responses = append(responses, ToBookRoomResponse(bookRoom))
	}
	return responses
}
