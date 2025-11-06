package response

type BookRoomResponse struct {
	ID     int          `json:"id"`
	RoomID int          `json:"room_id"`
	UserID int          `json:"user_id"`
	Date   string       `json:"date"`
	Price  float64      `json:"price"`
	Room   RoomResponse `json:"room"`
	User   UserResponse `json:"user"`
}
