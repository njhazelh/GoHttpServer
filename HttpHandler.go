package httpserver

/**
 * This interface defines a class that handles an HTTP Request
 * for one or more URLs.  A class that implements this interface
 * would be added to an instance of HttpServer.  When a client
 * requests the url that this class is matched to, the server
 * dispatches the request to the handler, by calling the
 * Handle Method.
 */
type HttpHandler interface {
	Handle(req *HttpRequest, res *HttpResponse)
	String() string
}
