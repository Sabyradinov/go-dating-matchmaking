package httpClient

import "github.com/Sabyradinov/go-dating-matchmaking/common"

type IClient interface {
	GetRequest(request common.HttpArgs) (httpStatus int, responseBody []byte, err error)
	PostRequest(request common.HttpArgs) (httpStatus int, responseBody []byte, err error)
	PutRequest(request common.HttpArgs) (httpStatus int, responseBody []byte, err error)
}
