package request

type BookRoomRequest struct {
	RoomID int    `json:"room_id" validate:"required,gt=0"`
	Date   string `json:"date" validate:"required"`
}
