package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "Get Api Key Success",
			headers: func() http.Header {
				h := make(http.Header)
				h.Set("Authorization", "ApiKey abc123")
				return h
			}(),
			want:    "abc123",
			wantErr: nil,
		},
		{
			name: "No Authorization Header",
			headers: func() http.Header {
				h := make(http.Header)
				return h
			}(),
			want:    "",
			wantErr: errors.New("no authorization header included"),
		},
		{
			name: "Malformed Authorization Header",
			headers: func() http.Header {
				h := make(http.Header)
				h.Set("Authorization", "abc123")
				return h
			}(),
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && key != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", key, tt.want)
			}
		})
	}
}
