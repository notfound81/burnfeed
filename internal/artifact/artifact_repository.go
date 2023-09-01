package artifact

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	ipfsApi "github.com/ipfs/go-ipfs-api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Or we can say
type StorageArtifactRepository struct {
	// Any database / ipfs connection or configurations needed
	ipfsClient *ipfsApi.Shell
	db         *gorm.DB
}

func NewStorageArtifactRepository(endpoint, projectId, projectSecret, dbDsn string) *StorageArtifactRepository {
	// Initialize and return a new instance of StorageArtifactRepository
	var ipfsClient *ipfsApi.Shell
	if projectId != "" && projectSecret != "" {
		ipfsClient = ipfsApi.NewShellWithClient(endpoint, NewClient(projectId, projectSecret))
	} else {
		ipfsClient = ipfsApi.NewShell(endpoint)
	}

	db, err := gorm.Open(mysql.Open(dbDsn), nil)
	if err != nil {
		return nil
	}

	artifactory := &StorageArtifactRepository{
		ipfsClient: ipfsClient,
		db:         db,
	}

	return artifactory
}

func (r *StorageArtifactRepository) CreateArtifact(artifactType, artifactJsonString string) (string, error) {

	// Check if the action type is supported
	if artifactType == "send_message" || artifactType == "tweet" || artifactType == "retweetOf" || artifactType == "artifact_file" {
		cid, err := r.ipfsClient.Add(strings.NewReader(artifactJsonString))
		return cid, err
	} else {
		return "", errors.New("Unsupported action type")
	}
}

func (r *StorageArtifactRepository) GetArtifactByID(ID string) (*Artifact, error) {
	// Implementation for getting an artifact from the database by ID
	// Mock implementation for testing - later connect it with DB or IPFS (or both).
	fmt.Println("GetArtifactByID")
	return &Artifact{"0x123456"}, nil
}

func (s *StorageArtifactRepository) GetUserArtifacts(wallet string) (string, error) {
	// Add business logic to retrieve all entries per DB per user
	// Right now just call repo's corresponding function.
	return "TO_BE_IMPLEMENTED", nil
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: &authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}
