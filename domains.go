package genshinapi

import "encoding/json"

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
	Level            string `json:"level"`
	AdventureRank    string `json:"adventureRank"`
	RecommendedLevel string `json:"recommendedLevel"`
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
	Level                   string `json:"level"`
	AdventureExperience     string `json:"adventureExperience"`
	CompanionshipExperience string `json:"companionshipExperience"`
	Mora                    string `json:"mora"`
	Drops                   []Drop `json:"drops"`
}

type Drop struct {
	Name string `json:"name"`
	Rate string `json:"rate"`
}

func (d Domain) EntryName() string {
	return d.Name
}

// --------- Helper functions
// GetDomains : Get list of domain names in Genshin Impact
func GetDomains() ([]string, error) {
	return GetDataTypeItemsList(DomainsDType)
}

// GetDomain : Return a pointer to a Domain struct containing info on a domain
func GetDomain(name string) (Domain, error) {
	reqBody := []string{
		DomainsDType,
		name,
	}

	var domain Domain
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return domain, err
	}

	err = json.Unmarshal(bytes, &domain)
	if err != nil {
		return domain, err
	}
	return domain, nil
}