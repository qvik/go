package util

import "net/http"

// AddQueryParam adds a query parameter with given name, value to
// HTTP request.
func AddQueryParam(r *http.Request, name, value string) {
	q := r.URL.Query()
	q.Add(name, value)
	r.URL.RawQuery = q.Encode()
}
