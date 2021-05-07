package genshinapi

const (
	ElementsDType = "elements"

	ElementIcon = "icon"
)

type Element struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func (e Element) EntryName() string {
	return e.Name
}

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
