package httpparser

// HTTPPayload - payload of http request
type HTTPPayload struct {
	Path    string
	Method  string
	Headers map[string]string
	Payload string
}
