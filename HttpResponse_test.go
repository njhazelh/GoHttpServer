package httpserver

import (
	"fmt"
	"testing"
)

func TestDefaultResponse(t *testing.T) {
	r := NewResponse()

	if r.Version != "HTTP/1.1" {
		t.Error("Version should be HTTP/1.1, but was:", r.Version)
	}
	if r.statusCode != 404 {
		t.Error("Default status should be 404, but was:", r.statusCode)
	}
	if r.statusMsg != "Not Found" {
		t.Error("Default status should be 'Not Found', was:", r.statusMsg)
	}

	fmt.Fprint(r, "Hello world")
	if string(r.body) != "Hello world" {
		t.Errorf(
			"Body writing did not work correctly.  Wrote %s. Saw %s",
			"Hello world",
			r.body)
	}

	desired := "HTTP/1.1 404 Not Found\r\n"
	desired += "Content-length: 11\r\n"
	desired += "\r\n"
	desired += "Hello world"
	if msg := r.String(); msg != desired {
		t.Errorf(
			"Message not converted to string correctly\nExpected: %s\nSaw: %s",
			desired,
			msg)
	}
}
