package main

import (
	"fmt"
)
// ============= BÀI 8: BUILDER PATTERN =============

// TODO: Builder Pattern
type HTTPRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
}

type RequestBuilder struct {
	request HTTPRequest
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: HTTPRequest{
			Headers: make(map[string]string),
		},
	}
}

func (b *RequestBuilder) URL(url string) *RequestBuilder {
	// Implement here
	b.request.URL = url
	return b
}

func (b *RequestBuilder) Method(method string) *RequestBuilder {
	// Implement here
	b.request.Method = method
	return b
}

func (b *RequestBuilder) Header(key, value string) *RequestBuilder {
	// Implement here
	b.request.Headers[key] = value
	return b
}

func (b *RequestBuilder) Body(body []byte) *RequestBuilder {
	// Implement here
	b.request.Body = body
	return b
}

func (b *RequestBuilder) Build() *HTTPRequest {
	return &b.request
}

func exercise8() {
	fmt.Println("\n=== BÀI 8: BUILDER PATTERN ===")

	// TODO: Test builder
	req := NewRequestBuilder().
		URL("https://api.example.com").
		Method("POST").
		Header("Content-Type", "application/json").
		Body([]byte(`{"name":"John"}`)).
		Build()

	fmt.Printf("Request: %+v\n", req)
}
func main() {
	exercise8()
}
