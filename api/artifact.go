package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api-server/internal/artifact"
)

// ArtifactHandler handles requests related to artifacts.
type ArtifactHandler struct {
	ArtifactService artifact.ArtifactService
}

func NewArtifactHandler(service artifact.ArtifactService) *ArtifactHandler {
	return &ArtifactHandler{ArtifactService: service}
}

func (h *ArtifactHandler) CreateArtifact(w http.ResponseWriter, r *http.Request) {
	var requestBody artifact.Artifact
	// Authorizing if header token is still valid
	authHeader := r.Header.Get("Authorization")
	fmt.Println(authHeader)
	// if authHeader == "" {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// if tokenString == "" {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	// authChecker := authentication.NewAuthChecker()
	// err := authChecker.ValidateToken(tokenString, "user_nonce") // Replace with the actual user nonce
	// if err != nil {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ipfsCID, err := h.ArtifactService.CreateArtifact(requestBody)
	if err != nil {
		http.Error(w, "Error creating artifact", http.StatusInternalServerError)
		return
	}

	response := struct {
		IPFS_CID string `json:"ipfsCID"`
	}{IPFS_CID: ipfsCID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
