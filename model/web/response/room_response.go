package response

type RoomResponse struct {
	ID         int              `json:"id"`
	RoomTypeID int              `json:"room_type_id"`
	RoomNumber string           `json:"room_number"`
	RoomType   RoomTypeResponse `json:"room_type"`
}
