package http

import (
	"net/http"

	"github.com/qvik/gokit/http/util"
)

// QueryParamTransport implements net/http/RoundTripper and
// provides a mechanism for constructing a http.Client
// that will add a set of query params to every GET request.
type QueryParamTransport struct {
	Params map[string]string
}

// RoundTrip adds the defined set of custom headers into every request.
func (t QueryParamTransport) RoundTrip(
	r *http.Request) (*http.Response, error) {

	if r.Method == "GET" {
		for k, v := range t.Params {
			util.AddQueryParam(r, k, v)
		}
	}

	return http.DefaultTransport.RoundTrip(r)
}
