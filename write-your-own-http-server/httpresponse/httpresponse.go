package httpresponse

import (
	"fmt"
	"strings"
)

var StatusCodeDescription = map[int32]string{
	200: "OK",
	400: "Bad Request",
	404: "Not Found",
}

type HTTPResponse struct {
	StatusCode int32
	Headers    map[string]string
	Body       string
}

func NewHTTPResponse(statusCode int32, headers map[string]string, body string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}

func (r HTTPResponse) ToHTTPFormat() string {
	headerLines := []string{
		fmt.Sprintf("HTTP/1.1 %d %s", r.StatusCode, StatusCodeDescription[r.StatusCode]),
	}
	for k, v := range r.Headers {
		headerLines = append(headerLines, fmt.Sprintf("%s: %s", k, v))
	}

	return strings.Join(headerLines, "\r\n") + "\r\n\r\n" + r.Body
}
