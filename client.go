package gohttpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	config HttpClientConfig
}

type HttpClientConfig struct {
	RequestTimeout time.Duration
}

func NewHttpClient(config HttpClientConfig) *HttpClient {
	return &HttpClient{
		config: config,
	}
}

func (c *HttpClient) Post(ctx context.Context, url string, data any, headers map[string]string) (Response, error) {
	body, err := encodeBody(data, headers)
	if err != nil {
		return Response{}, err
	}
	return c.Request(ctx, http.MethodPost, url, body, headers)
}

func (c *HttpClient) Get(ctx context.Context, url string, params map[string]any, headers map[string]string) (Response, error) {
	if params != nil {
		queryParamsString, err := urlEncode(params)
		if err != nil {
			return Response{}, err
		}
		url = fmt.Sprintf("%s?%s", url, queryParamsString)
	}

	return c.Request(ctx, http.MethodGet, url, nil, headers)
}

func (c *HttpClient) Request(context context.Context, method string, url string, body []byte, headers map[string]string) (Response, error) {
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		return Response{}, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: c.config.RequestTimeout,
	}

	response, err := client.Do(req)

	if err != nil {
		return Response{}, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, fmt.Errorf("client: could not read response body with error: %s", err)
	}

	return Response{
		Code: response.StatusCode,
		Body: responseBody,
	}, nil

}
