package genshinapi

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

// GetElements : Get list of element names.
func GetElements() ([]string, error) {
	return GetDataTypeItemsList(ElementsDType)
}

// GetElement : Return a requested Element struct.
func GetElement(name string) (Element, error) {
	reqBody := []string{
		ElementsDType,
		name,
	}

	var element Element
	err := getDataAndUnmarshal(reqBody, &element)
	return element, err
}
