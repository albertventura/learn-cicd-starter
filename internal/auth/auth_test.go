package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "Valid key",
			headers: http.Header{"Authorization": []string{"ApiKey validKey"}},
			wantKey: "validKey",
			wantErr: false,
		},
		{
			name:    "Invalid key",
			headers: http.Header{"Authorization": []string{"ApiKey invalidkey"}},
			wantKey: "invalidkey",
			wantErr: true,
		},
		{
			name:    "malformed header",
			headers: http.Header{"Authorization": []string{"potato apikey"}},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "No auth header",
			headers: http.Header{},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetApiKey got error %v, and wantErr = %v", err, tt.wantErr)
				}
			}

			if key != tt.wantKey {
				t.Fatalf("GetApiKey() => expected %v, got %v", tt.wantKey, key)
			}
		})
	}
}
