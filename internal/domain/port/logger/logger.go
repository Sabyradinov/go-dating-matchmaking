package logger

import (
	"github.com/Sabyradinov/go-dating-matchmaking/common"
	"github.com/gin-gonic/gin"
)

type AppLogger interface {
	HttpPanicHandler(c *gin.Context, recovered interface{})
	ErrorWithCode(ctx *gin.Context, action string, code int, message string, opt *common.LogOptions)
	Error(action string, code int, message string, opt *common.LogOptions)
	Warn(action, message string, opt *common.LogOptions)
	Debug(action, message string, opt *common.LogOptions)
	Info(action, message string, opt *common.LogOptions)
}
