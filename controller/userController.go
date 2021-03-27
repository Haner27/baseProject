package controller

import (
	"baseProject/entity"
	"baseProject/service"
	e "baseProject/util/error"
	resp "baseProject/util/response"
	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
		userService service.IUserService
	}
)

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var req entity.UserRegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.ErrorResp(ctx, e.ParamsInvalid, "")
		return
	}
	respData := uc.userService.RegisterUser(&req)
	resp.SuccessResp(ctx, respData)
}
