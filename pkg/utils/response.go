package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(200, "success", data))
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, NewResponse(code, message, nil))
}
