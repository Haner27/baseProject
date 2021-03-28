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

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param nickname body string true "昵称"
// @Param gender body int true "性别"
// @Param birthday body string true "生日,如：1992-03-14"
// @Success 200 {object} resp.Response "{"code":0,"msg":"成功","data":{"id":5,"username":"hannengfang","nickname":"","gender":1,"genderTxt":"male","age":2020}}"
// @Failure 400 {object} resp.Response "{"code":3,"msg":"参数无效","data":null}"
// @Router /v1/user/register [post]
func (uc *UserController) Register(ctx *gin.Context) {
	var req entity.UserRegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.ErrorResp(ctx, e.ParamsInvalid, "")
		return
	}
	respData := uc.userService.RegisterUser(&req)
	resp.SuccessResp(ctx, respData)
}

// @Summary 获取用户
// @Description 更具用户名获取用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param userId body int true "用户ID"
// @Success 200 {object} resp.Response "{"code":0,"msg":"成功","data":{"id":4,"username":"haner27","nickname":"","gender":1,"genderTxt":"male","age":29}}"
// @Failure 400 {object} resp.Response "{"code":3,"msg":"参数无效","data":null}"
// @Router /v1/user/GetUserById [post]
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

// @Summary 获取用户列表
// @Description 翻页获取用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param maxUserId body int true "上一页最大的用户ID"
// @Success 200 {object} resp.Response "{"code":0,"msg":"成功","data":[{"id":4,"username":"haner27","nickname":"","gender":1,"genderTxt":"male","age":29},{"id":5,"username":"hannengfang","nickname":"","gender":1,"genderTxt":"male","age":2020}]}"
// @Failure 400 {object} resp.Response "{"code":3,"msg":"参数无效","data":null}"
// @Router /v1/user/GetUsers [post]
func (uc *UserController) GetUsers(ctx *gin.Context) {
	var req entity.UsersLoaderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.ErrorResp(ctx, e.ParamsInvalid, "")
		return
	}
	respData := uc.userService.LoadUsers(&req)
	resp.SuccessResp(ctx, respData)
}