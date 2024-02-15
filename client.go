package gohttpclient

import (
	"bytes"
	"context"
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

func (c *httpClient) Post(ctx context.Context, url string, data any, headers map[string]string) ([]byte, error) {
	body, err := encodeBody(data, headers)
	if err != nil {
		return []byte{}, err
	}
	return c.Request(ctx, http.MethodPost, url, body, headers)
}

func (c *httpClient) Get(ctx context.Context, url string, params map[string]any, headers map[string]string) ([]byte, error) {
	if params != nil {
		queryParamsString, err := urlEncode(params)
		if err != nil {
			return nil, err
		}
		url = fmt.Sprintf("%s?%s", url, queryParamsString)
	}

	return c.Request(ctx, http.MethodGet, url, nil, headers)
}

func (c *httpClient) Request(context context.Context, method string, url string, body []byte, headers map[string]string) ([]byte, error) {
	bodyReader := bytes.NewReader(body)

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
