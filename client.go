package gohttpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type httpClient struct {
	config HttpClientConfig
}

type HttpClientConfig struct {
	RequestTimeout time.Duration
}

func NewHttpClient(config HttpClientConfig) *httpClient {
	return &httpClient{
		config: config,
	}
}

func (c *httpClient) Post(ctx context.Context, url string, data interface{}, headers map[string]string) ([]byte, error) {
	return c.Request(ctx, http.MethodPost, url, data, headers)
}

func (c *httpClient) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return c.Request(ctx, http.MethodGet, url, nil, headers)
}

func (c *httpClient) Request(context context.Context, method string, url string, data interface{}, headers map[string]string) ([]byte, error) {
	jsonBody, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: c.config.RequestTimeout,
	}

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 201 {
		return nil, fmt.Errorf("client: bad response code (%d) from dest (%s)", response.StatusCode, url)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read response body with error: %s", err)
	}

	return responseBody, nil

}
