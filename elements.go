package genshinapi

import "encoding/json"

const (
	// API DataType name
	ElementsDType = "elements"
	// Assets
	ElementIcon = "icon"
)

// --------- Structs for Element
type Element struct {
	Name string `json:"name"`
}

func (e Element) EntryName() string {
	return e.Name
}

// --------- Helper functions
// GetElements : Get list of element names in Genshin Impact
func GetElements() ([]string, error) {
	return GetDataTypeItemsList(ElementsDType)
}

// GetElement : Return a pointer to an Element struct containing info on an element
func GetElement(name string) (Element, error) {
	reqBody := []string{
		ElementsDType,
		name,
	}

	var element Element
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return element, err
	}

	err = json.Unmarshal(bytes, &element)
	if err != nil {
		return element, err
	}
	return element, nil
}
