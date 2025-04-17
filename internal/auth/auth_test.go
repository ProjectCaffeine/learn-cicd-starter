package auth

import (
	"net/http"
	"testing"
)

func TestNoHeader(t *testing.T) {
	header := http.Header{}

	_, err := GetAPIKey(header)

	if err.Error() != "no authorization header included" {
		t.Fatal("No error received!")
	}
}
