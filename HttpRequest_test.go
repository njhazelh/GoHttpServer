package httpserver

import (
	"bufio"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	httpMsg := "POST /hello/world.xyz HTTP/1.1\r\n"
	httpMsg += "Connection: keep-alive\r\n"
	httpMsg += "Accept-Language: en-US,en;q=0.8\r\n"
	httpMsg += "Content-length: 11\r\n"
	httpMsg += "\r\n"
	httpMsg += "hello world"
	reader := bufio.NewReader(strings.NewReader(httpMsg))

	status, err := readStatus(reader)
	if err != nil {
		t.Error("Encountered an error while parsing a valid status:", err)
	}
	if status.method != "POST" {
		t.Error("Got the wrong method:", status.method)
	}
	if status.uri != "/hello/world.xyz" {
		t.Error("Got the wrong URI:", status.uri)
	}
	if status.version != "HTTP/1.1" {
		t.Error("Got the wrong HTTP version:", status.version)
	}

	headers, err := readHeaders(reader)
	if err != nil {
		t.Error("Got error while parsing valid headers:", err)
	}
	if val, ok := headers["Connection"]; !ok || val != "keep-alive" {
		t.Error("Didn't parse Connection header Correctly", val)
	}
	if val, ok := headers["Accept-Language"]; !ok || val != "en-US,en;q=0.8" {
		t.Error("Didn't parse Accept-Language header Correctly:", val)
	}
	if val, ok := headers["Content-length"]; !ok || val != "11" {
		t.Error("Didn't parse Accept-Language header Correctly:", val)
	}

	body, err := readBody(headers, reader)
	if err != nil || string(body) != "hello world" {
		t.Error("Got the wrong body:", string(body))
	}
}
