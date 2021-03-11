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

type BossMaterial struct{}
type CharacterAscensionMaterial struct{}
type CharacterExperienceMaterial struct{}
type CommonAscensionMaterial struct{}
type LocalSpecialitiesMaterial struct{}
type TalentBookMaterial struct{}
type TalentBossMaterial struct{}
type WeaponAscensionMaterial struct{}
type WeaponExperienceMaterial struct{}

// --------- Helper functions
// GetMaterials : Get list of material names in Genshin Impact
func GetMaterials() ([]string, error) {
	return GetDataTypeItemsList(MaterialsDTpye)
}
