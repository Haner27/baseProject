package controller

import (
	"baseProject/entity"
	"baseProject/service"
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
	if err := ctx.ShouldBindJSON(&req);err != nil {

	}
}
