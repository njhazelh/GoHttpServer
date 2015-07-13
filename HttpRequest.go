package httpserver

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type HttpRequest struct {
	status  HttpRequestStatus
	headers HttpHeaders
	body    []byte
}

func (msg *HttpRequest) String() string {
	status := msg.status.String()
	headers := msg.headers.String()
	return status + "\n" + headers
}

/******************************************************************************
 *
 * Network->Struct parsing functions
 *
 *****************************************************************************/

func readStatus(r *bufio.Reader) (HttpRequestStatus, error) {
	line, _, err := r.ReadLine()
	if err != nil {
		return HttpRequestStatus{}, err
	}
	parts := strings.SplitN(string(line), " ", 3)
	return HttpRequestStatus{
		method:  parts[0],
		uri:     parts[1],
		version: parts[2],
	}, nil
}

func readHeaders(r *bufio.Reader) (HttpHeaders, error) {
	headers := make(HttpHeaders)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		if string(line) == "" {
			break
		}
		parts := strings.SplitN(string(line), ": ", 2)
		headers[parts[0]] = parts[1]
	}
	return headers, nil
}

func readBody(headers map[string]string, r *bufio.Reader) ([]byte, error) {
	length_str, ok := headers["Content-length"]
	if !ok {
		length_str = "0"
	}
	length, err := strconv.Atoi(length_str)
	if err != nil {
		return nil, err
	}
	body := make([]byte, length)
	_, err = r.Read(body)
	return body, err
}

func ParseRequest(c net.Conn) (*HttpRequest, error) {
	reader := bufio.NewReader(c)
	status, err := readStatus(reader)
	if err != nil {
		return nil, err
	}
	headers, err := readHeaders(reader)
	if err != nil {
		return nil, err
	}
	body, err := readBody(headers, reader)
	if err != nil {
		fmt.Println("Failed to read body")
		return nil, err
	}
	request := &HttpRequest{
		status:  status,
		headers: headers,
		body:    body,
	}
	return request, nil
}
