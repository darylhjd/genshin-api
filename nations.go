package genshinapi

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
// GetNations : Get list of nation names.
func GetNations() ([]string, error) {
	return GetDataTypeItemsList(NationsDType)
}

// GetNation : Return a requested Nation struct.
func GetNation(name string) (Nation, error) {
	reqBody := []string{
		NationsDType,
		name,
	}

	var nation Nation
	err := getDataAndUnmarshal(reqBody, &nation)
	return nation, err
}
