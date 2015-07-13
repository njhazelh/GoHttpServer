package httpserver

type HttpHeaders map[string]string

func (hdrs HttpHeaders) String() string {
	str := ""
	for key, val := range hdrs {
		str += key + ": " + val + "\n"
	}
	return str
}
