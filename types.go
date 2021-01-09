package genshinapi

const (
	// Data types provided by the API
	ArtifactsDType  = "artifacts"
	CharactersDType = "characters"
	DomainsDType    = "domains"
	ElementsDType   = "elements"
	NationsDType    = "nations"
	WeaponsDType    = "weapons"
)

const (
	// Asset types
	ArtifactPhoto     = "flower-of-life"
	CharacterIcon     = "icon"
	CharacterPortrait = "portrait"
	ElementIcon       = "icon"
	WeaponIcon        = "icon"
)

type APIError struct {
	Error           *string  `json:"error"`
	AvailableTypes  []string `json:"availableTypes,omitempty"`  // For wrong data type
	AvailableImages []string `json:"availableImages,omitempty"` // For wrong image type
}

type DataTypes struct {
	Types []string `json:"types"`
}

type DataEntry interface {
	EntryName() string
}

type Artifact struct {
	Name           string `json:"name"`
	MaxRarity      int    `json:"max_rarity"`
	TwoPieceBonus  string `json:"2-piece_bonus"`
	FourPieceBonus string `json:"4-piece_bonus"`
}

func (a *Artifact) EntryName() string {
	return a.Name
}

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

func (c *Character) EntryName() string {
	return c.Name
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

func (d *Domain) EntryName() string {
	return d.Name
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

type Element struct {
	Name string `json:"name"`
}

func (e *Element) EntryName() string {
	return e.Name
}

type Nation struct {
	Name              string `json:"name"`
	Element           string `json:"element"`
	Archon            string `json:"archon"`
	ControllingEntity string `json:"controllingEntity"`
}

func (n *Nation) EntryName() string {
	return n.Name
}

type Weapon struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Rarity      int    `json:"rarity"`
	BaseAttack  string `json:"baseAttack"`
	SubStat     string `json:"subStat"`
	PassiveName string `json:"passiveName"`
	PassiveDesc string `json:"passiveDesc"`
	Location    string `json:"location"`
}

func (w *Weapon) EntryName() string {
	return w.Name
}
