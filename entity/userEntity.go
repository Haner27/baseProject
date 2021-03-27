package entity

type (
	UserRegisterReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
		Gender   int    `json:"gender"`
		Birthday string `json:"birthday"`
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
		UserId int `json:"userId"`
	}
	UsersLoaderReq struct {
		MaxUserId int `json:"maxUserId"`
	}
)
