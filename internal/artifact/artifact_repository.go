package artifact

import (
	"fmt"
	"net/http"
	"strings"

	ipfsApi "github.com/ipfs/go-ipfs-api"
)

// Or we can say
type StorageArtifactRepository struct {
	// Any database / ipfs connection or configurations needed
	ipfsClient *ipfsApi.Shell
}

func NewStorageArtifactRepository(endpoint, projectId, projectSecret string) *StorageArtifactRepository {
	// Initialize and return a new instance of StorageArtifactRepository
	var IPFSClient *ipfsApi.Shell
	if projectId != "" && projectSecret != "" {
		IPFSClient = ipfsApi.NewShellWithClient(endpoint, NewClient(projectId, projectSecret))
	} else {
		IPFSClient = ipfsApi.NewShell(endpoint)
	}
	artifactory := &StorageArtifactRepository{
		ipfsClient: IPFSClient,
	}

	return artifactory
}

func (r *StorageArtifactRepository) CreateArtifact(artifact Artifact) (string, error) {
	// Implementation for creating an artifact on ipfs
	// Mock data !
	fmt.Println("CreateArtifact")
	// Define a JSON string as a variable - THis is just a mock !
	jsonStr := `{
        "type": "action",
        "timestamp": "2023-08-20 13:02:09.846505 +0800 CST m=+0.285431793",
        "actions": [
            {
                "subtype": "tweet",
                "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            },
            {
                "subtype": "tweet",
                "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9",
                "retweetOf": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            },
            {
                "subtype": "follow",
                "user": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
                "followee": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
            },
            {
                "subtype": "like",
                "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            },
            {
                "subtype": "send_message",
                "to": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
                "message": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            }
        ]
    }`

	cid, err := r.ipfsClient.Add(strings.NewReader(jsonStr))

	return cid, err
}

func (r *StorageArtifactRepository) GetArtifactByID(ID string) (*Artifact, error) {
	// Implementation for getting an artifact from the database by ID
	// Mock implementation for testing - later connect it with DB or IPFS (or both).
	fmt.Println("GetArtifactByID")
	return &Artifact{"0x123456"}, nil
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
