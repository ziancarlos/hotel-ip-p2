package domain

import "time"

type BookRoom struct {
	ID     int       `gorm:"primaryKey;autoIncrement"`
	RoomID int       `gorm:"not null"`
	UserID int       `gorm:"not null"`
	Date   time.Time `gorm:"type:date;not null"`
	Price  float64   `gorm:"type:decimal(19,2);not null"`
	Room   Room      `gorm:"foreignKey:RoomID;references:ID"`
	User   User      `gorm:"foreignKey:UserID;references:ID"`
}

func (BookRoom) TableName() string {
	return "book_rooms"
}
