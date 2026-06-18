package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    any `json:"data,omitempty"`
}

const (
	SuccessCode    = 0
	ErrorCode      = 1
	AuthErrorCode  = 401
	ForbiddenCode  = 403
	NotFoundCode   = 404
	ValidationCode = 400
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    SuccessCode,
		Message: "success",
		Data:    data,
	})
}

func SuccessWithMsg(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    SuccessCode,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    ErrorCode,
		Message: message,
	})
}

func ErrorWithCode(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    ValidationCode,
		Message: message,
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    AuthErrorCode,
		Message: "unauthorized",
	})
	c.Abort()
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		Code:    ForbiddenCode,
		Message: "forbidden",
	})
	c.Abort()
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code:    NotFoundCode,
		Message: "not found",
	})
	c.Abort()
}

func ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    ErrorCode,
		Message: message,
	})
}