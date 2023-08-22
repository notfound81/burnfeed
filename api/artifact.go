package api

import (
	"encoding/json"
	"io"
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

	// We can write everything to ipfs - and no content check if comform with
	// our data structure since FE puts them together (theoretically correct)
	bodyBytes, _ := io.ReadAll(r.Body)

	ipfsCID, err := h.ArtifactService.CreateArtifact(string(bodyBytes))
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
