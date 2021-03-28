package entity

type (
	UserRegisterReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
		Gender   int    `json:"gender" binding:"required"`
		Birthday string `json:"birthday" binding:"required"`
	}
	UserResp struct {
		Id         int    `json:"id"`
		Username   string `json:"username"`
		Nickname   string `json:"nickname"`
		Gender     int    `json:"gender"`
		GenderText string `json:"genderTxt"`
		Age        int    `json:"age"`
	}
	UserLoaderReq struct {
		UserId int `json:"userId" binding:"required"`
	}
	UsersLoaderReq struct {
		MaxUserId int `json:"maxUserId"`
	}
)
