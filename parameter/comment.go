package parameter

import "time"

type (
	CommentReq struct {
		Id      int
		Message string `json:"message" binding:"required"`
		PhotoId int    `json:"photo_id" binding:"required"`
		UserId  int
	}

	CommentUpdate struct {
		Message string `json:"message" binding: "required"`
	}

	CommentResp struct {
		Id        int       `json:"id"`
		Message   string    `json:"message"`
		PhotoId   int       `json:"photo_id"`
		UserId    int       `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
	}

	CommentList struct {
		Id        int       `json:"id"`
		Message   string    `json:"message"`
		PhotoId   int       `json:"photo_id"`
		UserId    int       `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      UserResp  `json:"user"`
		Photo     PhotoResp `json:"photo"`
	}
)
