package types

import "net/http"

type RouterTypeMethod string

const (
	Get    RouterTypeMethod = "GET"
	Post   RouterTypeMethod = "POST"
	Delete RouterTypeMethod = "DELETE"
	Put    RouterTypeMethod = "PUT"
)

type RouterType struct {
	Pattern    string
	Router     http.HandlerFunc
	TypeMethod RouterTypeMethod
}
