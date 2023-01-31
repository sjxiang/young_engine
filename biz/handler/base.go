package handler

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	
	"github.com/gin-gonic/gin"
)

const (
	SuccessMsg = "success"

	SuccessCode      = 0
	ServiceErrCode   = 10001
	ParamErrCode     = 10002
	CompileErrCode   = 20001
	RuleNotExistCode = 20002
	RuleExecErrCode  = 20003
)

type BaseResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BindResp(c *app.RequestContext, code int, msg string, data interface{}) {
	c.JSON(200, BaseResp{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}


// 序列化 serializer
// common

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	// Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		// Code:    Err.ErrCode,
		// Message: Err.ErrMsg,
		Data:    data,
	})
}
