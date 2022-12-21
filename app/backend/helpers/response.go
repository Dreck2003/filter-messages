package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dreck2003/api/backend/types"
)

type ResponseI interface {
	Send(http.ResponseWriter) // Send the HTTP
}

type Response struct {
	errorStatus  uint16
	errorMessage *string
	content      *interface{}
	ResponseI
}

func CreateResponse() *Response {
	return &Response{}
}

func (r *Response) AddContent(content interface{}) *Response {
	r.content = &content
	return r
}

func (r *Response) AddStatus(status uint16) *Response {
	r.errorStatus = status
	return r
}

func (r *Response) AddError(err string) *Response {
	r.errorMessage = &err
	return r
}

func (r *Response) Send(w *http.ResponseWriter) {
	parseResponse, err := json.Marshal(types.GenericResponse{
		Error:   r.errorMessage,
		Content: r.content,
	})

	if err != nil {
		fmt.Println("Error parsing httpResponse")
		log.Fatal(err)
	}

	(*w).WriteHeader(int(r.errorStatus))
	(*w).Write(parseResponse)
}
