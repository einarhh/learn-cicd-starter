package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	expectedKey := "1234567890"
	header.Set("Authorization", "ApiKey "+expectedKey)

	key, err := GetAPIKey(header)
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if key != expectedKey {
		t.Error("Expected "+expectedKey+" string, got", key)
	}
}

func TestGetAPIKeyNoHeader(t *testing.T) {
	header := http.Header{}

	_, err := GetAPIKey(header)
	if err != ErrNoAuthHeaderIncluded {
		t.Error("Expected ErrNoAuthHeaderIncluded error, got", err)
	}
}
