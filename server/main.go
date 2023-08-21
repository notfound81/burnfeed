package main

import (
	"api-server/api"
	"api-server/internal/artifact"
	"log"
	"net/http"
)

func main() {
	// Config parameters hard-coded here !!
	// IPFS related
	IPFS_ENDPOINT := "http://localhost:5001"
	IPFS_PROJECTID := ""
	IPFS_SECRET := ""

	// API SERVER REALTED
	PORT := ":8181"
	// Initialize your artifact repository, service, and handler
	//Here comes the DB entries as well!!
	repo := artifact.NewStorageArtifactRepository(IPFS_ENDPOINT, IPFS_PROJECTID, IPFS_SECRET) // Initialize your artifact repository (e.g., database repository)
	artifactService := artifact.NewArtifactService(repo)
	artifactHandler := api.NewArtifactHandler(artifactService)

	http.HandleFunc("/get-nonce", api.GetNonceHandler)
	http.HandleFunc("/sign-in", api.SignInHandler)
	http.HandleFunc("/create-artifact", logRequest(artifactHandler.CreateArtifact))

	log.Printf("Server started on port %s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("Server encountered an error:", err)
	}
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Preprocessing logic, if needed
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		// Call the original handler
		handler(w, r)

		// Postprocessing logic, if needed
		log.Println("Request processed")
	}
}
