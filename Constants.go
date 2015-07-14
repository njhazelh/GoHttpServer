package httpserver

var HTTP_STATUS = map[string]int{
	// Informational
	"Continue":            100,
	"Switching Protocols": 101,

	// Success
	"OK":                           200,
	"Created":                      201,
	"Accepted":                     202,
	"Non-Authoriative Information": 203,
	"No Content":                   204,
	"Reset Content":                205,
	"Partial Content":              206,

	// Redirection
	"Multiple Choices":   300,
	"Moved Permanently":  301,
	"Found":              302,
	"See Other":          303,
	"Not Modified":       304,
	"Use Proxy":          305,
	"Temporary Redirect": 307,

	// Client Errors
	"Bad Request":                     400,
	"Unauthorized":                    401,
	"Payment Required":                402,
	"Forbidden":                       403,
	"Not Found":                       404,
	"Method Not Allowed":              405,
	"Not Acceptable":                  406,
	"Proxy Authentication Required":   407,
	"Request Time-out":                408,
	"Conflict":                        409,
	"Gone":                            410,
	"Length Required":                 411,
	"Precondition Failed":             412,
	"Request Entity Too Large":        413,
	"Request-URI Too Large":           414,
	"Unsupported Media Type":          415,
	"Requested range not satisfiable": 416,
	"Expectation Failed":              417,

	// Server Errors
	"Internal Server Error":      500,
	"Not Implemented":            501,
	"Bad Gateway":                502,
	"Service Unavailable":        503,
	"Gateway Time-out":           504,
	"HTTP Version not supported": 505,
}
