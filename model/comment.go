package model

import "time"

type Comment struct {
	Id        int
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
	PhotoId   int
	Photo     Photo `gorm:"foreignKey:PhotoId"`
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
