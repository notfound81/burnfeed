package main

import (
	"log"
	"net/http"

	"api-server/api"

	"api-server/internal/artifact"
)

func main() {
	// Initialize your artifact repository, service, and handler
	repo := artifact.NewDatabaseArtifactRepository() // Initialize your artifact repository (e.g., database repository)
	artifactService := artifact.NewArtifactService(repo)
	artifactHandler := api.NewArtifactHandler(artifactService)

	http.HandleFunc("/get-nonce", api.GetNonceHandler)
	http.HandleFunc("/sign-in", api.SignInHandler)
	http.HandleFunc("/create-artifact", logRequest(artifactHandler.CreateArtifact))

	port := ":8080"
	log.Printf("Server started on port %s", port)
	err := http.ListenAndServe(port, nil)
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
