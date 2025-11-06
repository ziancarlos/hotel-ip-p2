package request

type RoomRequest struct {
	RoomTypeID int    `json:"room_type_id" validate:"required,gt=0"`
	RoomNumber string `json:"room_number" validate:"required"`
}
