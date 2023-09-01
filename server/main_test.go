package main

import (
	"api-server/api"
	"api-server/internal/artifact"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateArtifact(t *testing.T) {
	repo := artifact.NewDatabaseArtifactRepository()
	service := artifact.NewArtifactService(repo)
	handler := api.NewArtifactHandler(service)

	// Define your request payload
	requestPayload := struct {
		Type    string `json:"type"`
		Version string `json:"version"`
		Subtype string `json:"subtype"`
		User    string `json:"user"`
		Content string `json:"content"`
	}{
		Type:    "artifact",
		Version: "0.1",
		Subtype: "tweet",
		User:    "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
		Content: "This serves as an example of a tweet.",
	}

	// Encode the request payload to JSON
	payloadBytes, _ := json.Marshal(requestPayload)

	// Create a mock HTTP request with the JSON payload
	req := httptest.NewRequest("POST", "/artifact", bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Call the handler function
	handler.CreateArtifact(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Parse the response body if needed
	var responseBody map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Failed to parse response JSON: %v", err)
	}

	// Add additional checks based on your expectations

	// For example, you can check the IPFS CID
	ipfsCID, found := responseBody["ipfsCID"].(string)
	if !found || ipfsCID == "" {
		t.Errorf("Expected non-empty IPFS CID in response")
	} else {
		t.Log("ipfsCID is:", ipfsCID)
	}
}
