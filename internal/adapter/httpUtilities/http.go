package httpUtilities

import (
	"context"
	"github.com/Sabyradinov/go-dating-matchmaking/common"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/httpClient"
	"github.com/go-resty/resty/v2"
	"time"
)

type appHttpClient struct {
	client            *resty.Client
	defaultTimeoutSec time.Duration
}

// New creates a new instance of the appHttpClient
func New(defaultTimeoutSec time.Duration) httpClient.IClient {
	return &appHttpClient{defaultTimeoutSec: defaultTimeoutSec, client: resty.New()}
}

// GetRequest sends a GET request to the specified URL
func (c *appHttpClient) GetRequest(request common.HttpArgs) (httpStatus int, responseBody []byte, err error) {
	if request.TimeoutSecond == 0 {
		request.TimeoutSecond = c.defaultTimeoutSec
	}
	ctx, cancel := context.WithTimeout(request.Context, request.TimeoutSecond*time.Second)
	defer cancel()

	resp, err := c.client.R().
		SetResult(request.ResponseStruct).
		SetHeaders(request.Headers).
		SetContext(ctx).
		Get(request.Url)

	if err != nil {
		return
	}

	httpStatus = resp.StatusCode()
	responseBody = resp.Body()

	return
}

// PostRequest sends a POST request to the specified URL
func (c *appHttpClient) PostRequest(request common.HttpArgs) (httpStatus int, responseBody []byte, err error) {
	if request.TimeoutSecond == 0 {
		request.TimeoutSecond = c.defaultTimeoutSec
	}
	ctx, cancel := context.WithTimeout(request.Context, request.TimeoutSecond*time.Second)
	defer cancel()

	resp, err := c.client.R().
		SetResult(request.ResponseStruct).
		ForceContentType("application/json").
		SetHeaders(request.Headers).
		SetContext(ctx).
		Post(request.Url)

	if err != nil {
		return
	}

	httpStatus = resp.StatusCode()
	responseBody = resp.Body()

	return
}

// PutRequest sends a PUT request to the specified URL
func (c *appHttpClient) PutRequest(request common.HttpArgs) (httpStatus int, responseBody []byte, err error) {
	if request.TimeoutSecond == 0 {
		request.TimeoutSecond = c.defaultTimeoutSec
	}
	ctx, cancel := context.WithTimeout(request.Context, request.TimeoutSecond*time.Second)
	defer cancel()

	resp, err := c.client.R().
		SetResult(request.ResponseStruct).
		ForceContentType("application/json").
		SetHeaders(request.Headers).
		SetContext(ctx).
		Put(request.Url)

	if err != nil {
		return
	}

	httpStatus = resp.StatusCode()
	responseBody = resp.Body()

	return
}
