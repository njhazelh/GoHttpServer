package httpserver

import (
	"net"
)

type HttpHandler interface {
	Handle(r *HttpRequest, c net.Conn) error
	String() string
}
