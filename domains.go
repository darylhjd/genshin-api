package genshinapi

const (
	// API DataType name
	DomainsDType = "domains"
	// Assets
	// -- There are no assets for Domains
)

// --------- Structs for Domain

type Domain struct {
	Name                string               `json:"name"`
	Type                string               `json:"type"`
	Description         string               `json:"description"`
	Location            string               `json:"location"`
	Nation              string               `json:"nation"`
	Requirements        []RequirementLevel   `json:"requirements"`
	RecommendedElements []RecommendedElement `json:"recommendedElements"`
	Rewards             []Reward             `json:"rewards"`
}

type RequirementLevel struct {
	Level            int    `json:"level"`
	AdventureRank    int    `json:"adventureRank"`
	RecommendedLevel int    `json:"recommendedLevel"`
	LeyLineDisorder  string `json:"leyLineDisorder"`
}

type RecommendedElement struct {
	Element string `json:"element"`
}

type Reward struct {
	Day     string         `json:"day"`
	Details []RewardDetail `json:"details"`
}

type RewardDetail struct {
	Level                   int    `json:"level"`
	AdventureExperience     int    `json:"adventureExperience"`
	CompanionshipExperience int    `json:"companionshipExperience"`
	Mora                    int    `json:"mora"`
	Drops                   []Drop `json:"drops"`
}

type Drop struct {
	Name    string `json:"name"`
	DropMin int    `json:"drop_min"`
	DropMax int    `json:"drop_max"`
}

func (d Domain) EntryName() string {
	return d.Name
}

// --------- Helper functions

// GetDomains : Get list of domain names.
func GetDomains() ([]string, error) {
	return GetDataTypeItemsList(DomainsDType)
}

// GetDomain : Return a requested Domain struct.
func GetDomain(name string) (Domain, error) {
	reqBody := []string{
		DomainsDType,
		name,
	}

	var domain Domain
	err := getDataAndUnmarshal(reqBody, &domain)
	return domain, err
}
