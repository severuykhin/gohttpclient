package gohttpclient

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int
	Body []byte
}

func (r Response) IsOk() bool {
	return r.Code == http.StatusOK
}

func (r Response) IsAccepted() bool {
	return r.Code == http.StatusAccepted
}

func (r Response) IsCreated() bool {
	return r.Code == http.StatusCreated
}

func (r Response) IsConflict() bool {
	return r.Code == http.StatusConflict
}

func (r Response) IsNotFound() bool {
	return r.Code == http.StatusNotFound
}

func (r Response) IsForbidden() bool {
	return r.Code == http.StatusForbidden
}

func (r Response) IsInternalServerError() bool {
	return r.Code == http.StatusInternalServerError
}

func (r Response) Unmarshal(dest interface{}) error {
	return json.Unmarshal(r.Body, dest)
}
