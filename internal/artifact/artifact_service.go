package artifact

type artifactService struct {
	repo ArtifactRepository
}

func NewArtifactService(repo ArtifactRepository) ArtifactService {
	return &artifactService{repo: repo}
}

func (s *artifactService) CreateArtifact(artifact Artifact) (string, error) {
	// Add business logic to create an artifact.
	// Right now just call repo's corresponding function.
	return s.repo.CreateArtifact(artifact)
}

func (s *artifactService) GetArtifactByID(ID string) (*Artifact, error) {
	// Add business logic to retrieve an artifact by ID.
	// Right now just call repo's corresponding function.
	return s.repo.GetArtifactByID(ID)
}
