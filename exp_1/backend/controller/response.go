package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SuccessWithMessage 响应成功
func SuccessWithMessage(msg string, data interface{}, c *gin.Context) {
	res := Response{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

// FailWithMessage 响应失败
func FailWithMessage(msg string, c *gin.Context) {
	res := Response{
		Code: 0,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, res)
}
