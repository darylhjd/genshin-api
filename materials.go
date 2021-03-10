package genshinapi

const (
	// API DataType name
	MaterialsDTpye = "materials"
	// Assets
	// -- There are no assets for Materials
)

// --------- Structs for Material
type Material struct {
}

// --------- Helper functions
// GetMaterials : Get list of material names in Genshin Impact
func GetMaterials() ([]string, error) {
	return GetDataTypeItemsList(MaterialsDTpye)
}
