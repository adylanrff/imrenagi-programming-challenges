package httpparser

import (
	"bufio"
	"net"
	"strings"
)

// ParseHTTP - parse http request
func ParseHTTP(conn net.Conn) (*HTTPPayload, error) {
	var (
		method  string
		path    string
		headers map[string]string
		payload string
	)

	sc := bufio.NewScanner(conn)
	for !sc.Scan() {
	}

	if sc.Err() != nil {
		return nil, sc.Err()
	}

	splittedBuf := strings.Split(sc.Text(), "\r\n\r\n")

	meta := splittedBuf[0]
	if len(splittedBuf) > 1 {
		payload = splittedBuf[1]
	}

	method, path, headers = parseMeta(meta)

	return &HTTPPayload{
		Path:    path,
		Method:  method,
		Headers: headers,
		Payload: payload,
	}, nil
}

func parseMeta(meta string) (method string, path string, headers map[string]string) {
	headers = make(map[string]string)

	// Split meta by CRLF
	splittedMeta := strings.Split(meta, "\r\n")

	// Parse method line
	methodLine := strings.Split(splittedMeta[0], " ")
	method, path = methodLine[0], methodLine[1]

	// Parse header lines
	for _, headerLine := range splittedMeta[1:] {
		splittedHeaderLine := strings.Split(headerLine, ":")
		headers[splittedHeaderLine[0]] = splittedHeaderLine[1]
	}

	return
}
