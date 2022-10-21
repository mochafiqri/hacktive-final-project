package model

import "time"

type SocialMedia struct {
	Id             int
	Name           string
	SocialMediaUrl string
	UserId         int
	User           User `gorm:"foreignKey:UserId"`
	CreatedAt      time.Time
}
