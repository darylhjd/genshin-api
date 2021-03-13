package genshinapi

import (
	"errors"
)

const (
	// API DataType name
	MaterialsDType = "materials"
	// Assets
	// -- There are no assets for Materials
)

const (
	// Material types
	BossMaterialType                = "boss-material"
	CharacterAscensionMaterialType  = "character-ascension"
	CharacterExperienceMaterialType = "character-experience"
	CommonAscensionMaterialType     = "common-ascension"
	LocalSpecialtiesMaterialType    = "local-specialties"
	TalentBookMaterialType          = "talent-book"
	TalentBossMaterialType          = "talent-boss"
	WeaponAscensionMaterialType     = "weapon-ascension"
	WeaponExperienceMaterialType    = "weapon-experience"
)

// Material : This interface is unique to a Material DataEntry.
// Any Material interface will also implement the normal DataEntry interface.
type Material interface {
	MaterialType() string
}

// --------- Structs for Material

type MaterialItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

// BossMaterials : Struct containing material from each elemental boss.
type BossMaterials struct {
	Anemo   BossMaterial `json:"anemo"`
	Cryo    BossMaterial `json:"cryo"`
	Electro BossMaterial `json:"electro"`
	Geo     BossMaterial `json:"geo"`
	Hydro   BossMaterial `json:"hydro"`
	Pyro    BossMaterial `json:"pyro"`
}

type BossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

func (bm BossMaterials) MaterialType() string {
	return BossMaterialType
}

func (bm BossMaterials) EntryName() string {
	return BossMaterialType
}

// CharacterAscensionMaterials : Struct containing each elemental character ascension material.
type CharacterAscensionMaterials struct {
	Anemo    CharacterAscensionMaterial `json:"anemo"`
	Cryo     CharacterAscensionMaterial `json:"cryo"`
	Electro  CharacterAscensionMaterial `json:"electro"`
	Geo      CharacterAscensionMaterial `json:"geo"`
	Hydro    CharacterAscensionMaterial `json:"hydro"`
	Pyro     CharacterAscensionMaterial `json:"pyro"`
	Traveler CharacterAscensionMaterial `json:"traveler"`
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

func (cam CharacterAscensionMaterials) MaterialType() string {
	return CharacterAscensionMaterialType
}

func (cam CharacterAscensionMaterials) EntryName() string {
	return CharacterAscensionMaterialType
}

// CharacterExperienceMaterials : Struct containing list of each CharacterExperienceMaterial.
type CharacterExperienceMaterials struct {
	Items []CharacterExperienceMaterial `json:"items"`
}

type CharacterExperienceMaterial struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Experience int    `json:"experience"`
	Rarity     int    `json:"rarity"`
}

func (cem CharacterExperienceMaterials) MaterialType() string {
	return CharacterExperienceMaterialType
}

func (cem CharacterExperienceMaterials) EntryName() string {
	return CharacterExperienceMaterialType
}

// CommonAscensionMaterials : Struct containing information on each common ascension material
type CommonAscensionMaterials struct {
	Slime                    CommonAscensionMaterial `json:"slime"`
	HilichurlMasks           CommonAscensionMaterial `json:"hilichurl-masks"`
	HilichurlArrowheads      CommonAscensionMaterial `json:"hilichurl-arrowheads"`
	SamachurlScrolls         CommonAscensionMaterial `json:"samachurl-scrolls"`
	TreasureHoarderInsignias CommonAscensionMaterial `json:"treasure-hoarder-insignias"`
	FatuiInsignias           CommonAscensionMaterial `json:"fatui-insignias"`
	WhopperflowerNectar      CommonAscensionMaterial `json:"whopperflower-nectar"`
	HilichurlHorns           CommonAscensionMaterial `json:"hilichurl-horns"`
	LeyLine                  CommonAscensionMaterial `json:"ley-line"`
	BoneShards               CommonAscensionMaterial `json:"bone-shards"`
	MistGrass                CommonAscensionMaterial `json:"mist-grass"`
	FatuiKnives              CommonAscensionMaterial `json:"fatui-knives"`
	ChaosParts               CommonAscensionMaterial `json:"chaos-parts"`
}

type CommonAscensionMaterial struct {
	Characters []string       `json:"characters"`
	Items      []MaterialItem `json:"items"`
	Sources    []string       `json:"sources"`
}

func (coam CommonAscensionMaterials) MaterialType() string {
	return CommonAscensionMaterialType
}

func (coam CommonAscensionMaterials) EntryName() string {
	return CommonAscensionMaterialType
}

// LocalSpecialtiesMaterials : Struct containing list of LocalSpecialtyMaterial for each nation.
type LocalSpecialtiesMaterials struct {
	Mondstadt []LocalSpecialtyMaterial `json:"mondstadt"`
	Liyue     []LocalSpecialtyMaterial `json:"liyue"`
}

type LocalSpecialtyMaterial struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func (lsm LocalSpecialtiesMaterials) MaterialType() string {
	return LocalSpecialtiesMaterialType
}

func (lsm LocalSpecialtiesMaterials) EntryName() string {
	return LocalSpecialtiesMaterialType
}

// TalentBookMaterials : Struct containing each TalentBookMaterials
type TalentBookMaterials struct {
	Freedom    TalentBookMaterial `json:"freedom"`
	Resistance TalentBookMaterial `json:"resistance"`
	Ballad     TalentBookMaterial `json:"ballad"`
	Prosperity TalentBookMaterial `json:"prosperity"`
	Diligence  TalentBookMaterial `json:"diligence"`
	Gold       TalentBookMaterial `json:"gold"`
}

type TalentBookMaterial struct {
	Characters   []string       `json:"characters"`
	Availability []string       `json:"availability"`
	Source       string         `json:"source"`
	Items        []MaterialItem `json:"items"`
}

func (tbookm TalentBookMaterials) MaterialType() string {
	return TalentBookMaterialType
}

func (tbookm TalentBookMaterials) EntryName() string {
	return TalentBookMaterialType
}

// TalentBossMaterials : Sturct containing each TalentBossMaterial
type TalentBossMaterials struct {
	TailOfBoreas         TalentBossMaterial `json:"tail-of-boreas"`
	RingOfBoreas         TalentBossMaterial `json:"ring-of-boreas"`
	SpiritLocketOfBoreas TalentBossMaterial `json:"spirit-locket-of-boreas"`
	DvalinsPlume         TalentBossMaterial `json:"dvalin's-plume"`
	DvalinsClaw          TalentBossMaterial `json:"dvalin's-claw"`
	DvalinsSigh          TalentBossMaterial `json:"dvalin's-sigh"`
	TuskOfMonocerosCaeli TalentBossMaterial `json:"tusk-of-monoceros-caeli"`
	ShardOfAFoulLegacy   TalentBossMaterial `json:"shard-of-a-foul-legacy"`
	ShadowOfTheWarrior   TalentBossMaterial `json:"shadow-of-the-warrior"`
}

type TalentBossMaterial struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func (tbossm TalentBossMaterials) MaterialType() string {
	return TalentBossMaterialType
}

func (tbossm TalentBossMaterials) EntryName() string {
	return TalentBossMaterialType
}

// WeaponAscensionMaterials : Struct containing each WeaponAscensionMaterial
type WeaponAscensionMaterials struct {
	Decarabian   WeaponAscensionMaterial `json:"decarabian"`
	Boreal       WeaponAscensionMaterial `json:"boreal"`
	Dandelion    WeaponAscensionMaterial `json:"dandelion"`
	Guyun        WeaponAscensionMaterial `json:"guyun"`
	Elixir       WeaponAscensionMaterial `json:"elixir"`
	Aerosiderite WeaponAscensionMaterial `json:"aerosiderite"`
}

type WeaponAscensionMaterial struct {
	Weapons      []string       `json:"weapons"`
	Availability []string       `json:"availability"`
	Source       string         `json:"source"`
	Items        []MaterialItem `json:"items"`
}

func (wam WeaponAscensionMaterials) MaterialType() string {
	return WeaponAscensionMaterialType
}

func (wam WeaponAscensionMaterials) EntryName() string {
	return WeaponAscensionMaterialType
}

// WeaponExperienceMaterials : Struct containing list of WeaponExperienceMaterial
type WeaponExperienceMaterials struct {
	Items []WeaponExperienceMaterial `json:"items"`
}

type WeaponExperienceMaterial struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Experience int      `json:"experience"`
	Rarity     int      `json:"rarity"`
	Source     []string `json:"source"`
}

func (wem WeaponExperienceMaterials) MaterialType() string {
	return WeaponExperienceMaterialType
}

func (wem WeaponExperienceMaterials) EntryName() string {
	return WeaponExperienceMaterialType
}

// --------- Helper functions

// GetMaterials : Get list of material names.
func GetMaterials() ([]string, error) {
	return GetDataTypeItemsList(MaterialsDType)
}

var materialMap = map[string]func() (Material, error){
	BossMaterialType:                GetBossMaterial,
	CharacterAscensionMaterialType:  GetCharacterAscensionMaterial,
	CharacterExperienceMaterialType: GetCharacterExperienceMaterial,
	CommonAscensionMaterialType:     GetCommonAscensionMaterial,
	LocalSpecialtiesMaterialType:    GetLocalSpecialtiesMaterial,
	TalentBookMaterialType:          GetTalentBookMaterial,
	TalentBossMaterialType:          GetTalentBossMaterial,
	WeaponAscensionMaterialType:     GetWeaponAscensionMaterial,
	WeaponExperienceMaterialType:    GetWeaponExperienceMaterial,
}

// GetMaterial : Return a Material struct corresponding to requested MaterialType.
func GetMaterial(mname string) (Material, error) {
	// Get required function from the function mapper

	// I originally wanted it to be implemented such that the user would have to pass
	// an struct of Material interface into the function and it would unmarshal accordingly,
	// but I didn't want to bring the problem further down the line, so this is what I went with.
	fn, ok := materialMap[mname]
	if !ok {
		return nil, errors.New("material type does not exist")
	}
	return fn()
}

// GetBossMaterial : Return a BossMaterials struct.
func GetBossMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		BossMaterialType,
	}

	var bm BossMaterials
	err := getDataAndUnmarshal(reqBody, &bm)
	return bm, err
}

// GetCharacterAscensionMaterial : Return a CharacterAscensionMaterials struct.
func GetCharacterAscensionMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		CharacterAscensionMaterialType,
	}

	var cam CharacterAscensionMaterials
	err := getDataAndUnmarshal(reqBody, &cam)
	return cam, err
}

// GetCharacterExperienceMaterial : Return a CharacterExperienceMaterials struct.
func GetCharacterExperienceMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		CharacterExperienceMaterialType,
	}

	var cem CharacterExperienceMaterials
	err := getDataAndUnmarshal(reqBody, &cem)
	return cem, err
}

// GetCommonAscensionMaterial : Return a CommonAscensionMaterials struct
func GetCommonAscensionMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		CommonAscensionMaterialType,
	}

	var coam CommonAscensionMaterials
	err := getDataAndUnmarshal(reqBody, &coam)
	return coam, err
}

// GetLocalSpecialtiesMaterial : Return a LocalSpecialtiesMaterials struct
func GetLocalSpecialtiesMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		LocalSpecialtiesMaterialType,
	}

	var lsm LocalSpecialtiesMaterials
	err := getDataAndUnmarshal(reqBody, &lsm)
	return lsm, err
}

// GetTalentBookMaterial : Return a TalentBookMaterials struct
func GetTalentBookMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		TalentBookMaterialType,
	}

	var tbookm TalentBookMaterials
	err := getDataAndUnmarshal(reqBody, &tbookm)
	return tbookm, err
}

// GetTalentBossMaterial : Return a TalentBossMaterials struct
func GetTalentBossMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		TalentBossMaterialType,
	}

	var tbossm TalentBossMaterials
	err := getDataAndUnmarshal(reqBody, &tbossm)
	return tbossm, err
}

// GetWeaponAscensionMaterial : Return a WeaponAscensionMaterials struct
func GetWeaponAscensionMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		WeaponAscensionMaterialType,
	}

	var wam WeaponAscensionMaterials
	err := getDataAndUnmarshal(reqBody, &wam)
	return wam, err
}

// GetWeaponExperienceMaterial : Return a WeaponExperienceMaterials struct
func GetWeaponExperienceMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		WeaponExperienceMaterialType,
	}

	var wem WeaponExperienceMaterials
	err := getDataAndUnmarshal(reqBody, &wem)
	return wem, err
}
