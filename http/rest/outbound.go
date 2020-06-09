package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/qvik/gokit/http/util"
)

// Request performs a HTTP request. Returns the response payload as bytes -
// or error.
func Request(ctx context.Context, httpClient *http.Client,
	req *http.Request) ([]byte, error) {

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform HTTP request")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read HTTP response body")
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, errors.Errorf("HTTP request failed with code %v; "+
			"message: %v, body: %v",
			res.StatusCode, res.Status, string(body))
	}

	return body, nil
}

// PostJSON Performs a POST request with the given JSON request payload;
// returns the response payload as bytes, or error
func PostJSON(ctx context.Context, httpClient *http.Client, url string,
	payload interface{}) ([]byte, error) {

	payloadJSONBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to serialize JSON payload")
	}

	bodyIn := bytes.NewBuffer(payloadJSONBytes)

	req, err := http.NewRequest("POST", url, bodyIn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HTTP POST request")
	}

	req.Header.Set("Content-Type", "application/json")

	return Request(ctx, httpClient, req)
}

// GetJSON Performs a GET request using Accept: application/json header.
// Returns the response payload as bytes, or error.
// You may pass nil as queryParams (or empty map) to indicate no query params.
func GetJSON(ctx context.Context, httpClient *http.Client,
	url string, queryParams map[string]string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HTTP GET request")
	}

	if queryParams != nil {
		for k, v := range queryParams {
			util.AddQueryParam(req, k, v)
		}
	}

	req.Header.Set("Accept", "application/json")

	return Request(ctx, httpClient, req)
}
