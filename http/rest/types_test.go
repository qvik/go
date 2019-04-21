package rest

import (
	"testing"
)

func TestErrors(t *testing.T) {
	e1 := NewNotFoundError("error message")
	if e1.Error() != "error message" {
		t.Error("Error message mismatch")
	}
	if e1.StatusCode != 404 {
		t.Error("Error Status code mismatch")
	}

	//TODO
}
