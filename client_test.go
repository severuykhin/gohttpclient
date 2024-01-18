package gohttpclient

import (
	"context"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	// Your test logic here
	c := NewHttpClient(HttpClientConfig{})

	params := map[string]any{
		"features": []string{
			"utm_source", "utm_medium",
		},
	}

	data, err := c.Get(context.Background(), "https://google.com", params, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(data)
}
