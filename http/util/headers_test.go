package util

import (
	"net/http"
	"testing"
)

func createRequest(authHeader string) *http.Request {
	r, _ := http.NewRequest("GET", "http://example.com", nil)
	r.Header.Add("Authorization", authHeader)

	return r
}

func TestParseBearerToken(t *testing.T) {
	hdr1 := ""
	res1, err1 := ParseBearerToken(createRequest(hdr1))
	if res1 != "" {
		t.Error("Invalid parse result")
	}

	if err1 != ErrMissingAuthorization {
		t.Error("Invalid parse result")
	}

	hdr2 := "Bearer foo"
	res2, err2 := ParseBearerToken(createRequest(hdr2))
	if res2 != "foo" {
		t.Error("Invalid parse result")
	}

	if err2 != nil {
		t.Error("Invalid parse result")
	}
}
