package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPriceHandler(t *testing.T) {
	// Create a test server with the desired handler
	ts := httptest.NewServer(http.HandlerFunc(getPriceHandler))
	defer ts.Close()

	// Make a request to the test server
	resp, err := http.Get(ts.URL + "/price")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response code is OK (200)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode the response body into a map[string]interface{} to check the format
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check if the "rates" field is present and has the correct format
	rates, ok := responseBody["rates"].(map[string]interface{})
	if !ok {
		t.Error("Expected 'rates' field in the response body")
		return
	}

	// Check if each rate is a float64
	for _, rate := range rates {
		_, isFloat := rate.(float64)
		if !isFloat {
			t.Errorf("Expected 'float64' type for rate, but got %T", rate)
		}
	}
}
