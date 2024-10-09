package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	input1 := httptest.NewRequest("GET", "/", nil).Header
	input1.Set("Authorization", "ApiKey 1234567890")
	expected1 := "1234567890"
	var input2 = httptest.NewRequest("GET", "/", nil).Header
	input2.Set("Authorization", "ApiKey HelloBanana")
	expected2 := "HelloBanana"
	var input3 = httptest.NewRequest("GET", "/", nil).Header
	input3.Set("Authorization", "Bearer 1234567890")
	expected3 := ""
	var input4 = httptest.NewRequest("GET", "/", nil).Header
	input4.Set("Authorization", "ApiKey1234567890")
	expected4 := ""

	tests := []struct {
		name     string
		input    http.Header
		expected string
	}{
		{"Happy path", input1, expected1},
		{"Happy path 2", input2, expected2},
		{"Malformed API Key", input3, expected3},
		{"Malformed API Key 2", input4, expected4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetAPIKey(tt.input)
			if result != tt.expected {
				t.Errorf("GetAPIKey(%v) = %v; want %v\n", tt.input, result, tt.expected)
				t.Errorf("Error value: %v\n", err)
			}
		})
	}
}
