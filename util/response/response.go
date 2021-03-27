package response

import (
	e "baseProject/util/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		code,
		msg,
		data,
	}
}

func SuccessResp(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponse(int(e.Success), e.GetErrorStr(e.Success), data))
}

func ErrorResp(ctx *gin.Context, code e.ErrCode, msg string) {
	if msg == "" {
		msg = e.GetErrorStr(code)
	}
	ctx.JSON(http.StatusOK, NewResponse(int(e.Success), msg, nil))
}