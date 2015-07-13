package httpserver

import (
	"errors"
	"net"
	"reflect"
)

type BaseHandler struct{}

func (h *BaseHandler) Handle(r HttpRequest, c net.Conn) error {
	object := reflect.ValueOf(*h)
	methodName := "Do" + r.status.method
	method := object.MethodByName(methodName)
	if method.Int() == 0 {
		// Method not implemented
		// Reply with a 404 or some other error
		return errors.New("Method not found")
	}
	input := []reflect.Value{reflect.ValueOf(r), reflect.ValueOf(c)}
	method.Call(input)
	return nil
}
