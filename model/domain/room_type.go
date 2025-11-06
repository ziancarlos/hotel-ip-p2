package domain

type RoomType struct {
	ID    int     `gorm:"primaryKey;autoIncrement"`
	Name  string  `gorm:"type:varchar(100);not null"`
	Price float64 `gorm:"type:decimal(19,2);not null"`
}

func (RoomType) TableName() string {
	return "room_types"
}
