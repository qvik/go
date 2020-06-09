package http

import "net/http"

// CustomHeaderTransport implements net/http/RoundTripper and
// provides a mechanism for constructing a http.Client
// that will include a fixed set of headers into every request.
type CustomHeaderTransport struct {
	Headers map[string]string
}

// RoundTrip adds the defined set of custom headers into every request.
func (t CustomHeaderTransport) RoundTrip(
	r *http.Request) (*http.Response, error) {

	for k, v := range t.Headers {
		r.Header.Set(k, v)
	}

	return http.DefaultTransport.RoundTrip(r)
}
