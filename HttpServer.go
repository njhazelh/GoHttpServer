package httpserver

import (
	"fmt"
	"net"
)

type HttpServer struct {
	Handlers map[string]HttpHandler
}

func NewServer() *HttpServer {
	server := new(HttpServer)
	server.Handlers = make(map[string]HttpHandler)
	return server
}

func (s *HttpServer) AddHandle(uri string, handler HttpHandler) {
	s.Handlers[uri] = handler
}

func (s *HttpServer) handleConnection(c net.Conn) {
	defer c.Close()
	for {
		request, err := ParseRequest(c)
		if err != nil {
			// Send a closing message?
			fmt.Println("Failed to parse request: ", err)
			break
		}
		fmt.Println(request.String())
		handler, ok := s.Handlers[request.status.uri]
		response := NewResponse()
		if ok {
			handler.Handle(request, response)
		}
		fmt.Println(response)
		fmt.Fprint(c, response)
		if err != nil {
			fmt.Println("Failed to dispatch")
			break
		}
	}
	fmt.Println("Terminating Connection")
}

func (s *HttpServer) Run(source string) {
	ln, err := net.Listen("tcp", source)
	if err != nil {
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// Handle Error
		}
		go s.handleConnection(conn)
	}
}
