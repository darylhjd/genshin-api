package genshinapi

import "errors"

const (
	MaterialsDType = "materials"

	// Assets only available for cooking ingredients.
	// For asset names, get key from CookingIngredientsMaterials map.

	BossMaterialType                = "boss-material"
	CharacterAscensionMaterialType  = "character-ascension"
	CharacterExperienceMaterialType = "character-experience"
	CommonAscensionMaterialType     = "common-ascension"
	CookingIngredientsMaterialType  = "cooking-ingredients"
	LocalSpecialtiesMaterialType    = "local-specialties"
	TalentBookMaterialType          = "talent-book"
	TalentBossMaterialType          = "talent-boss"
	WeaponAscensionMaterialType     = "weapon-ascension"
	WeaponExperienceMaterialType    = "weapon-experience"
)

type Material interface {
	MaterialType() string
}

type MaterialItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

// BossMaterials : Struct containing material from each elemental boss.
type BossMaterials map[string]BossMaterial

func (bm BossMaterials) MaterialType() string {
	return BossMaterialType
}

type BossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

func (bm BossMaterial) EntryName() string {
	return bm.Name
}

// CharacterAscensionMaterials : Struct containing each elemental character ascension material.
type CharacterAscensionMaterials map[string]CharacterAscensionMaterial

func (cam CharacterAscensionMaterials) MaterialType() string {
	return CharacterAscensionMaterialType
}

type CharacterAscensionMaterial struct {
	Sliver   GemPiece `json:"sliver"`
	Fragment GemPiece `json:"fragment"`
	Chunk    GemPiece `json:"chunk"`
	Gemstone GemPiece `json:"gemstone"`
}

type GemPiece struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	Rarity  int      `json:"rarity"`
}

func (gp CharacterAscensionMaterial) EntryName() string {
	return CharacterAscensionMaterialType
}

// CharacterExperienceMaterials : Struct containing list of each CharacterExperienceMaterial.
type CharacterExperienceMaterials struct {
	Items []CharacterExperienceMaterial `json:"items"`
}

func (cem CharacterExperienceMaterials) MaterialType() string {
	return CharacterExperienceMaterialType
}

type CharacterExperienceMaterial struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Experience int    `json:"experience"`
	Rarity     int    `json:"rarity"`
}

func (cem CharacterExperienceMaterial) EntryName() string {
	return cem.Name
}

// CommonAscensionMaterials : Struct containing information on each common ascension material
type CommonAscensionMaterials map[string]CommonAscensionMaterial

func (coam CommonAscensionMaterials) MaterialType() string {
	return CommonAscensionMaterialType
}

type CommonAscensionMaterial struct {
	Characters []string       `json:"characters"`
	Items      []MaterialItem `json:"items"`
	Sources    []string       `json:"sources"`
}

func (coam CommonAscensionMaterial) EntryName() string {
	return CommonAscensionMaterialType
}

// CookingIngredientsMaterials : Struct containing list of CookingIngredientMaterial.
type CookingIngredientsMaterials map[string]CookingIngredientMaterial

func (cim CookingIngredientsMaterials) MaterialType() string {
	return CookingIngredientsMaterialType
}

type CookingIngredientMaterial struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Sources     []string `json:"sources"`
}

func (cim CookingIngredientMaterial) EntryName() string {
	return cim.Name
}

// LocalSpecialtiesMaterials : Struct containing list of LocalSpecialtyMaterial for each nation.
type LocalSpecialtiesMaterials map[string]LocalSpecialtyMaterial

func (lsm LocalSpecialtiesMaterials) MaterialType() string {
	return LocalSpecialtiesMaterialType
}

type LocalSpecialtyMaterial struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func (lsm LocalSpecialtyMaterial) EntryName() string {
	return lsm.Name
}

// TalentBookMaterials : Struct containing each TalentBookMaterials
type TalentBookMaterials map[string]TalentBookMaterial

func (tbookm TalentBookMaterials) MaterialType() string {
	return TalentBookMaterialType
}

type TalentBookMaterial struct {
	Characters   []string       `json:"characters"`
	Availability []string       `json:"availability"`
	Source       string         `json:"source"`
	Items        []MaterialItem `json:"items"`
}

func (tbookm TalentBookMaterial) EntryName() string {
	return TalentBookMaterialType
}

// TalentBossMaterials : Sturct containing each TalentBossMaterial
type TalentBossMaterials map[string]TalentBossMaterial

func (tbossm TalentBossMaterials) MaterialType() string {
	return TalentBossMaterialType
}

type TalentBossMaterial struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func (tbossm TalentBossMaterial) EntryName() string {
	return tbossm.Name
}

// WeaponAscensionMaterials : Struct containing each WeaponAscensionMaterial
type WeaponAscensionMaterials map[string]WeaponAscensionMaterial

func (wam WeaponAscensionMaterials) MaterialType() string {
	return WeaponAscensionMaterialType
}

type WeaponAscensionMaterial struct {
	Weapons      []string       `json:"weapons"`
	Availability []string       `json:"availability"`
	Source       string         `json:"source"`
	Items        []MaterialItem `json:"items"`
}

func (wam WeaponAscensionMaterial) EntryName() string {
	return WeaponAscensionMaterialType
}

// WeaponExperienceMaterials : Struct containing list of WeaponExperienceMaterial
type WeaponExperienceMaterials struct {
	Items []WeaponExperienceMaterial `json:"items"`
}

func (wem WeaponExperienceMaterials) MaterialType() string {
	return WeaponExperienceMaterialType
}

type WeaponExperienceMaterial struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Experience int      `json:"experience"`
	Rarity     int      `json:"rarity"`
	Source     []string `json:"source"`
}

func (wem WeaponExperienceMaterial) EntryName() string {
	return wem.Name
}

// GetMaterialTypes : Get list of material names.
func GetMaterialTypes() ([]string, error) {
	return GetDataTypeItemsList(MaterialsDType)
}

// GetMaterialType : Return a Material struct corresponding to requested MaterialType.
func GetMaterialType(mname string) (Material, error) {
	reqBody := []string{
		MaterialsDType,
		mname,
	}

	var (
		err error
	)
	switch mname {
	case BossMaterialType:
		t := BossMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case CharacterAscensionMaterialType:
		t := CharacterAscensionMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case CharacterExperienceMaterialType:
		t := CharacterExperienceMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case CommonAscensionMaterialType:
		t := CommonAscensionMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case CookingIngredientsMaterialType:
		t := CookingIngredientsMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case LocalSpecialtiesMaterialType:
		t := LocalSpecialtiesMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case TalentBookMaterialType:
		t := TalentBookMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case TalentBossMaterialType:
		t := TalentBossMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case WeaponAscensionMaterialType:
		t := WeaponAscensionMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	case WeaponExperienceMaterialType:
		t := WeaponExperienceMaterials{}
		err = getDataAndUnmarshal(reqBody, &t)
		return t, err
	default:
		return nil, errors.New("unknown MaterialType")
	}
}
