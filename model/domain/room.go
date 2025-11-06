package domain

type Room struct {
	ID         int      `gorm:"primaryKey;autoIncrement"`
	RoomTypeID int      `gorm:"not null"`
	RoomNumber string   `gorm:"type:varchar(50);not null;unique"`
	RoomType   RoomType `gorm:"foreignKey:RoomTypeID;references:ID"`
}

func (Room) TableName() string {
	return "rooms"
}
