package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("missing Authorization header", func(t *testing.T) {
		headers := http.Header{} 
		key, err := GetAPIKey(headers)

		if key != "" {
    		t.Fatalf("expected empty key, got %q", key)
		}

		if err == nil {
    		t.Fatalf("expected an error, got nil")
		}
		if err != ErrNoAuthHeaderIncluded {
    		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})

	t.Run("malformed Authorization header", func(t *testing.T) {
    	headers := http.Header{}
    	headers.Set("Authorization", "ApiKey") // no token
    	key, err := GetAPIKey(headers)
    	// expect key == "" and err.Error() == "malformed authorization header"
	})

	t.Run("valid ApiKey", func(t *testing.T) {
    	headers := http.Header{}
    	headers.Set("Authorization", "ApiKey abc123")
    	key, err := GetAPIKey(headers)
    	// expect key == "abc123" and err == nil
	})
}