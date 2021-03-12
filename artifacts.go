package genshinapi

const (
	// API DataType name
	ArtifactsDType = "artifacts"
	// Assets
	ArtifactPhoto = "flower-of-life"
)

// --------- Structs for Artifact
type Artifact struct {
	Name           string `json:"name"`
	MaxRarity      int    `json:"max_rarity"`
	TwoPieceBonus  string `json:"2-piece_bonus"`
	FourPieceBonus string `json:"4-piece_bonus"`
}

func (a Artifact) EntryName() string {
	return a.Name
}

// --------- Helper functions
// GetArtifacts : Get list of artifact names.
func GetArtifacts() ([]string, error) {
	return GetDataTypeItemsList(ArtifactsDType)
}

// GetArtifact : Return a requested Artifact struct.
func GetArtifact(name string) (Artifact, error) {
	reqBody := []string{
		ArtifactsDType,
		name,
	}

	var artifact Artifact
	err := getDataAndUnmarshal(reqBody, &artifact)
	return artifact, err
}
