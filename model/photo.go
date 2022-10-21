package model

import "time"

type Photo struct {
	Id        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
