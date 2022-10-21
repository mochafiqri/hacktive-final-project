package parameter

import "time"

type (
	PhotoReq struct {
		Id       int
		UserId   int
		Title    string `json:"title" binding:"required"`
		PhotoUrl string `json:"photo_url" binding:"required"`
		Caption  string `json:"caption"`
	}

	PhotoResp struct {
		Id       int       `json:"id"`
		Title    string    `json:"title"`
		Caption  string    `json:"caption"`
		PhotoUrl string    `json:"photo_url"`
		UserId   int       `json:"user_id"`
		Created  time.Time `json:"created,omitempty"`
	}

	PhotoListResp struct {
		Id        int       `json:"id"`
		Title     string    `json:"title"`
		Caption   string    `json:"caption"`
		PhotoUrl  string    `json:"photo_url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      struct {
			Email    string `json:"email"`
			Username string `json:"username"`
		}
	}
)
