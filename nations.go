package genshinapi

import "encoding/json"

const (
	// API DataType name
	NationsDType = "nations"
	// Assets
	// -- There are no assets for Nations
)

// --------- Structs for Nation
type Nation struct {
	Name              string `json:"name"`
	Element           string `json:"element"`
	Archon            string `json:"archon"`
	ControllingEntity string `json:"controllingEntity"`
}

func (n Nation) EntryName() string {
	return n.Name
}

// --------- Helper functions
// GetNations : Get list of nation names in Genshin Impact
func GetNations() ([]string, error) {
	return GetDataTypeItemsList(NationsDType)
}

// GetNation : Return a pointer to a Nation struct containing info on a nation
func GetNation(name string) (Nation, error) {
	reqBody := []string{
		NationsDType,
		name,
	}

	var nation Nation
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nation, err
	}

	err = json.Unmarshal(bytes, &nation)
	if err != nil {
		return nation, err
	}
	return nation, nil
}
