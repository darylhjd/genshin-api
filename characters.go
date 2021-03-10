package genshinapi

import "encoding/json"

const (
	// API DataType name
	CharactersDType = "characters"
	// Assets
	CharacterIcon     = "icon"
	CharacterPortrait = "portrait"
)

// --------- Structs for Character
type Character struct {
	Name           string   `json:"name"`
	Title          string   `json:"title"`
	Vision         string   `json:"vision"`
	Weapon         string   `json:"weapon"`
	Gender         string   `json:"gender"`
	Nation         string   `json:"nation"`
	Affiliation    string   `json:"affiliation"`
	SpecialDish    string   `json:"specialDish"`
	Rarity         int      `json:"rarity"`
	Constellation  string   `json:"constellation"`
	Birthday       string   `json:"birthday"`
	Description    string   `json:"description"`
	SkillTalents   []Talent `json:"skillTalents"`
	PassiveTalents []Talent `json:"passiveTalents"`
	Constellations []Talent `json:"constellations"`
}

type Talent struct {
	Name        string         `json:"name"`
	Unlock      string         `json:"unlock"`
	Description string         `json:"description"`
	Upgrades    []UpgradeLevel `json:"upgrades,omitempty"`
}

type UpgradeLevel struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (c Character) EntryName() string {
	return c.Name
}

// --------- Helper functions
// GetCharacters : Get list of character names in Genshin Impact
func GetCharacters() ([]string, error) {
	return GetDataTypeItemsList(CharactersDType)
}

// GetCharacter : Return a pointer to a Character struct containing info on a character
func GetCharacter(name string) (Character, error) {
	reqBody := []string{
		CharactersDType,
		name,
	}

	var character Character
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return character, err
	}

	err = json.Unmarshal(bytes, &character)
	if err != nil {
		return character, err
	}
	return character, nil
}
