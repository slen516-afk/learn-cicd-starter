package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Define the table of test cases
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-super-secret-key"},
			},
			expectedKey:   "my-super-secret-key",
			expectedError: "",
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "Malformed Header - Wrong Prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer my-super-secret-key"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "Malformed Header - Missing Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
	}

	// Loop through each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			// Check if we expected an error
			if tc.expectedError != "" {
				if err == nil {
					t.Fatalf("expected error: %v, got nil", tc.expectedError)
				}
				if err.Error() != tc.expectedError {
					t.Fatalf("expected error message: %v, got: %v", tc.expectedError, err.Error())
				}
				return // Test passed this case, move to next
			}

			// If we didn't expect an error but got one, fail
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check if the returned key matches what we expected
			if key != tc.expectedKey {
				t.Fatalf("expected key: %v, got: %v", tc.expectedKey, key)
			}
		})
	}
}
