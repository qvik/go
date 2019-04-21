package rest

import (
	"context"
	"encoding/json"
	"net/http"
)

// WriteJSONResponse is an utility method for returning a JSON response
func WriteJSONResponse(ctx context.Context, w http.ResponseWriter,
	response interface{}, apiErr error) {
	var responseObject interface{}
	statusCode := http.StatusOK

	if apiErr != nil {
		statusCode = http.StatusInternalServerError
		apiResp := &Response{Status: "error"}

		restError, ok := apiErr.(*Error)
		if ok {
			statusCode = restError.StatusCode
			apiResp.Message = restError.Message
		}
		responseObject = apiResp
	} else {
		responseObject = response
	}

	jsondata, err := json.Marshal(responseObject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	w.Write(jsondata)
}

// Rest allows for a handler function to simply return a response payload
// that is JSON-marshallable and/or an error and it will take care of the
// output.
func Rest(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allocate a context
		ctx := r.Context()

		// Call the endpoint handler
		response, err := handler(ctx, r)
		WriteJSONResponse(ctx, w, response, err)
	}
}
