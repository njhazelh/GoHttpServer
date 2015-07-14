package httpserver

type HttpHandler interface {
	Handle(req *HttpRequest, res *HttpResponse)
	String() string
}
