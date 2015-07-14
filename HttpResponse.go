package httpserver

import (
	"errors"
	"fmt"
)

// A structure that provides an interface for replying to HTTP messages
// Version: The version of HTTP being used. eg. HTTP/1.1
// statusCode: The number of the status. eg. 404
// statusMsg: The msg that comes to the right of the code. eg. Not Found
// headers: A collection of HTTP Headers
// body: The body of the response.
type HttpResponse struct {
	Version    string
	statusCode int
	statusMsg  string
	Headers    map[string]string
	body       []byte
}

// Create a default HttpResponse
func NewResponse() *HttpResponse {
	r := new(HttpResponse)
	r.Version = "HTTP/1.1"
	err := r.SetStatus("Not Found")
	if err != nil {
		fmt.Println(err)
	}
	r.body = []byte(nil)
	return r
}

// Set the statusMsg and statusCode of the HttpResponse
// name: The statusMsg to set to. Case matters.
func (r *HttpResponse) SetStatus(name string) error {
	code, ok := HTTP_STATUS[name]
	if !ok {
		return errors.New(name + " is not a status")
	}
	r.statusCode = code
	r.statusMsg = name
	return nil
}

func (r *HttpResponse) Write(bytes []byte) (int, error) {
	r.body = append(r.body, bytes...)
	return len(bytes), nil
}

func (r *HttpResponse) String() string {
	str := fmt.Sprintf("%s %d %s\r\n", r.Version, r.statusCode, r.statusMsg)
	for key, val := range r.Headers {
		str += fmt.Sprintf("%s: %v\r\n", key, val)
	}
	if len(r.body) > 0 {
		str += fmt.Sprintf("Content-length: %d\r\n", len(r.body))
	}
	str += "\r\n"
	str += string(r.body)
	return str
}
