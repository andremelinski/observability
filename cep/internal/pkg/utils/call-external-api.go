package utils

import (
	"context"
	"io"
	"net/http"
	"time"
)

type HandlerExternalApi struct{}

func NewHandlerExternalApi() *HandlerExternalApi{
	return &HandlerExternalApi{}
}

func (hea *HandlerExternalApi)CallExternalApi(ctx context.Context, timeoutMs int, method string, url string) ([]byte, error){
	timeout := time.Duration(timeoutMs)*time.Millisecond

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return resp, err
}