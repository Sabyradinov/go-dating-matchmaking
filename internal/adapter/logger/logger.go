package logger

import (
	"fmt"
	"github.com/Sabyradinov/go-dating-matchmaking/common"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/logger"
	goLogger "github.com/apsdehal/go-logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type app struct {
	log *goLogger.Logger
}

// New constructor for application logger instance
func New() logger.AppLogger {

	log, err := goLogger.New("appLogger", 1, os.Stdout)
	if err != nil {
		panic(err)
	}

	return &app{
		log: log,
	}
}

// HttpPanicHandler for handling panic function
func (l app) HttpPanicHandler(ctx *gin.Context, recovered interface{}) {
	code, mess := common.PanicErrorCode, fmt.Sprint(recovered)

	l.log.Error(mess)

	ctx.JSON(http.StatusUnprocessableEntity, errorToBase(code, mess))
	return
}

// ErrorWithCode log error and return error code
func (l app) ErrorWithCode(ctx *gin.Context, action string, code int, message string, opt *common.LogOptions) {
	if opt == nil {
		opt = &common.LogOptions{
			CustomMessage: message,
		}
	}

	l.log.Errorf("%v: %v", action, message)

	ctx.JSON(http.StatusUnprocessableEntity, errorToBase(code, message))
}

// Info log info
func (l app) Info(action string, message string, opt *common.LogOptions) {
	if opt == nil {
		opt = &common.LogOptions{}
	}

	l.log.InfoF("%v %v %v", action, message, opt)

}

// Debug log debug
func (l app) Debug(action string, message string, opt *common.LogOptions) {
	if opt == nil {
		opt = &common.LogOptions{}
	}

	l.log.DebugF("%v: %v: %v", action, message, opt)
}

// Warn log warning
func (l app) Warn(action string, message string, opt *common.LogOptions) {
	if opt == nil {
		opt = &common.LogOptions{}
	}

	l.log.WarningF("%v: %v: %v", action, message, opt)

}

// Error log error
func (l app) Error(action string, code int, message string, opt *common.LogOptions) {
	if opt == nil {
		opt = &common.LogOptions{}
	}

	l.log.Errorf("%v: %v: %v : %v", action, code, message, opt)
}

// errorToBase function to return error struct
func errorToBase(code int, message string) *common.BaseResponse {
	return &common.BaseResponse{Code: code, Message: message}
}
