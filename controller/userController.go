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
		userService *service.UserService
	}
)

func NewUserController(userService *service.UserService) *UserController {
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

func (uc *UserController) GetUserById(ctx *gin.Context) {
	var req entity.UserLoaderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.ErrorResp(ctx, e.ParamsInvalid, "")
		return
	}
	respData := uc.userService.LoadUserById(&req)
	if respData == nil {
		resp.ErrorResp(ctx, e.UserNotFound, "")
		return
	}
	resp.SuccessResp(ctx, respData)
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	var req entity.UsersLoaderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.ErrorResp(ctx, e.ParamsInvalid, "")
		return
	}
	respData := uc.userService.LoadUsers(&req)
	resp.SuccessResp(ctx, respData)
}