package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		headers        http.Header
		expectedApiKey string
		expectedError  error
	}{
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234"},
			},
			expectedApiKey: "1234",
			expectedError:  nil,
		},
		{
			headers:        http.Header{},
			expectedApiKey: "",
			expectedError:  ErrNoAuthHeaderIncluded,
		},
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedApiKey: "",
			expectedError:  ErrMalformedAuthHeader,
		},
	}

	for _, c := range cases {
		apiKey, err := GetAPIKey(c.headers)
		if !errors.Is(err, c.expectedError) {
			t.Errorf("Expected no error, got %v", err)
		}
		if apiKey != c.expectedApiKey {
			t.Errorf("Expected %s, got %s", c.expectedApiKey, apiKey)
		}
	}
}
