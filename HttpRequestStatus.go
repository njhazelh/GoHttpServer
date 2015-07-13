package httpserver

import "fmt"

type HttpRequestStatus struct {
	method  string
	uri     string
	version string
}

func (status HttpRequestStatus) String() string {
	return fmt.Sprintf("%s %s %s", status.method, status.uri, status.version)
}
