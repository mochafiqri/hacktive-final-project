package handler

import (
	"gorm.io/gorm"
)

type Handler struct {
	Db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		Db: db,
	}
}
