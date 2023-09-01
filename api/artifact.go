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

	// Retrieve the value of X-Action-Type header
	// This will indicate what action is done
	artifactType := r.Header.Get("X-Action-Type")
	if artifactType == "" {
		http.Error(w, "No X-Action-Type set", http.StatusBadRequest)
		return
	}

	// We can write everything to ipfs - and no content check if comform with
	// our data structure since FE puts them together (theoretically correct)
	bodyBytes, _ := io.ReadAll(r.Body)

	ipfsCID, err := h.ArtifactService.CreateArtifact(artifactType, string(bodyBytes))
	if err != nil {
		http.Error(w, "Error creating artifact", http.StatusInternalServerError)
		return
	}

	// If creating the artifact file, we just need to retrun the CID
	// Otherwise we need the subtype so that broser/client side we can
	// put together the artifact file
	if artifactType == "artifact_file" {
		response := struct {
			IPFS_CID string `json:"ipfsCID"`
		}{
			IPFS_CID: ipfsCID,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response := struct {
			IPFS_CID string `json:"ipfsCID"`
			SUBTYPE  string `json:"subtype"`
		}{
			IPFS_CID: ipfsCID,
			SUBTYPE:  artifactType,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}
