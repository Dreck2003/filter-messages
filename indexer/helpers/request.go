package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
)

type CustomRequest struct {
	Request *http.Request
	method  string
	url     string
}

func CreateRequest(method string) *CustomRequest {
	return &CustomRequest{
		method: method,
	}
}

func (r *CustomRequest) AddUrl(url string) *CustomRequest {
	r.url = url
	return r
}

type Person struct {
	Name string
}

func (r *CustomRequest) Build(data interface{}) (*CustomRequest, error) {
	if reflect.TypeOf(data).Kind().String() == "string" {
		req, err := http.NewRequest(r.method, r.url, strings.NewReader(data.(string)))
		if err != nil {
			return nil, err
		}
		r.Request = req
		return r, nil
	}

	if reflect.TypeOf(data).Kind().String() == "slice" {
		req, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(data.([]byte)))
		if err != nil {
			return nil, err
		}
		r.Request = req
		return r, nil
	}

	parseJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(parseJson))
	if err != nil {
		return nil, err
	}
	r.Request = req
	return r, nil
}

func (r *CustomRequest) Send() (*http.Response, error) {
	res, err := NetClient.Do(r.Request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
