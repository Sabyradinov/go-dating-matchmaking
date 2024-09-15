package common

import (
	"context"
	"errors"
	"time"
)

// BaseResponse base response struct
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HttpArgs http request args
type HttpArgs struct {
	Context        context.Context
	Url            string
	Data           []byte
	Headers        map[string]string
	ResponseStruct interface{}
	Proxy          string
	TimeoutSecond  time.Duration
}

const (
	PanicErrorCode = -911
	DefaultLang    = "ENG"
)

var PanicErrorMsg = errors.New("panic error message")

// LogOptions log options
type LogOptions struct {
	CustomMessage    string
	Data             interface{}
	AdditionalFields *map[string]interface{}
}
