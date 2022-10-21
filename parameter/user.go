package parameter

type (
	RegisterReq struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Age      int    `json:"age" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	UserResp struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Age      int    `json:"age,omitempty"`
	}

	LoginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	LoginResp struct {
		Token string `json:"token"`
	}

	UpdateReq struct {
		Id       int
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
	}
)
