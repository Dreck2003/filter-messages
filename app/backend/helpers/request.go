package helpers

import (
	"io"
	"net/http"
)

type RequestI interface {
	SetUrl(url string)
	Build(body io.Reader)
	Send()
}

type CustomRequest struct {
	method        TypeMethod
	url           string
	HttpRequest   *http.Request
	errrorMessage *error
}

type TypeMethod string

func CreateRequest(method TypeMethod) *CustomRequest {
	return &CustomRequest{
		method:        method,
		HttpRequest:   nil,
		errrorMessage: nil,
	}
}

func (r *CustomRequest) SetUrl(url string) *CustomRequest {
	r.url = url
	return r
}

func (r *CustomRequest) Build(body io.Reader) *CustomRequest {
	request, err := http.NewRequest(string(r.method), r.url, body)
	if err != nil {
		r.errrorMessage = &err
	} else {
		r.HttpRequest = request
	}
	return r
}

func (r *CustomRequest) Send() (*http.Response, error) {
	if r.errrorMessage != nil {
		return nil, *r.errrorMessage
	}
	return NetClient.Do(r.HttpRequest)
}
