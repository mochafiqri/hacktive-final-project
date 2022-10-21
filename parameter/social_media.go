package parameter

import "time"

type (
	SocialMediaReq struct {
		Id             int
		Name           string `json:"name" binding:"required"`
		SocialMediaUrl string `json:"social_media_url" binding:"required"`
		UserId         int
	}

	SocialMediaResp struct {
		Id             int       `json:"id"`
		Name           string    `json:"name"`
		SocialMediaUrl string    `json:"social_media_url"`
		UserId         int       `json:"user_id"`
		CreatedAt      time.Time `json:"created_at"`
	}

	SocialMediaList struct {
		Id             int       `json:"id"`
		Name           string    `json:"name"`
		SocialMediaUrl string    `json:"social_media_url"`
		UserId         int       `json:"user_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		User           UserResp  `json:"user"`
	}
)
