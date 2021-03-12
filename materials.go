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
// ---- BossMaterial
type BossMaterial struct {
	Anemo   AnemoBossMaterial   `json:"anemo"`
	Cryo    CryoBossMaterial    `json:"cryo"`
	Electro ElectroBossMaterial `json:"electro"`
	Geo     GeoBossMaterial     `json:"geo"`
	Hydro   HydroBossMaterial   `json:"hydro"`
	Pyro    PyroBossMaterial    `json:"pyro"`
}

type AnemoBossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type CryoBossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type ElectroBossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type GeoBossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type HydroBossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type PyroBossMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

func (bm BossMaterial) MaterialType() string {
	return BossMaterialType
}

func (bm BossMaterial) EntryName() string {
	return BossMaterialType
}

// ---- CharacterAscensionMaterial
type CharacterAscensionMaterial struct {
	Anemo    AnemoCharacterAscensionMaterial    `json:"anemo"`
	Cryo     CryoCharacterAscensionMaterial     `json:"cryo"`
	Electro  ElectroCharacterAscensionMaterial  `json:"electro"`
	Geo      GeoCharacterAscensionMaterial      `json:"geo"`
	Hydro    HydroCharacterAscensionMaterial    `json:"hydro"`
	Pyro     PyroCharacterAscensionMaterial     `json:"pyro"`
	Traveler TravelerCharacterAscensionMaterial `json:"traveler"`
}

type Sliver struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	Rarity  int      `json:"rarity"`
}

type Fragment struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	Rarity  int      `json:"rarity"`
}

type Chunk struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	Rarity  int      `json:"rarity"`
}

type Gemstone struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	Rarity  int      `json:"rarity"`
}

type AnemoCharacterAscensionMaterial struct {
	Sliver   Sliver   `json:"sliver"`
	Fragment Fragment `json:"fragment"`
	Chunk    Chunk    `json:"chunk"`
	Gemstone Gemstone `json:"gemstone"`
}

type CryoCharacterAscensionMaterial struct {
	Sliver   Sliver   `json:"sliver"`
	Fragment Fragment `json:"fragment"`
	Chunk    Chunk    `json:"chunk"`
	Gemstone Gemstone `json:"gemstone"`
}

type ElectroCharacterAscensionMaterial struct {
	Sliver   Sliver   `json:"sliver"`
	Fragment Fragment `json:"fragment"`
	Chunk    Chunk    `json:"chunk"`
	Gemstone Gemstone `json:"gemstone"`
}

type GeoCharacterAscensionMaterial struct {
	Sliver   Sliver   `json:"sliver"`
	Fragment Fragment `json:"fragment"`
	Chunk    Chunk    `json:"chunk"`
	Gemstone Gemstone `json:"gemstone"`
}

type HydroCharacterAscensionMaterial struct {
	Sliver   Sliver   `json:"sliver"`
	Fragment Fragment `json:"fragment"`
	Chunk    Chunk    `json:"chunk"`
	Gemstone Gemstone `json:"gemstone"`
}

type PyroCharacterAscensionMaterial struct {
	Sliver   Sliver   `json:"sliver"`
	Fragment Fragment `json:"fragment"`
	Chunk    Chunk    `json:"chunk"`
	Gemstone Gemstone `json:"gemstone"`
}

type TravelerCharacterAscensionMaterial struct {
	Sliver   Sliver   `json:"sliver"`
	Fragment Fragment `json:"fragment"`
	Chunk    Chunk    `json:"chunk"`
	Gemstone Gemstone `json:"gemstone"`
}

func (cam CharacterAscensionMaterial) MaterialType() string {
	return CharacterAscensionMaterialType
}

func (cam CharacterAscensionMaterial) EntryName() string {
	return CharacterAscensionMaterialType
}

// ---- CharacterExperienceMaterial
type CharacterExperienceMaterial struct {
	Items []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Experience int    `json:"experience"`
		Rarity     int    `json:"rarity"`
	} `json:"items"`
}

func (cem CharacterExperienceMaterial) MaterialType() string {
	return CharacterExperienceMaterialType
}

func (cem CharacterExperienceMaterial) EntryName() string {
	return CharacterExperienceMaterialType
}

// ---- CommonAscensionMaterial
type CommonAscensionMaterial struct {
	Slime                    Slime                    `json:"slime"`
	HilichurlMasks           HilichurlMasks           `json:"hilichurl-masks"`
	HilichurlArrowheads      HilichurlArrowheads      `json:"hilichurl-arrowheads"`
	SamachurlScrolls         SamachurlScrolls         `json:"samachurl-scrolls"`
	TreasureHoarderInsignias TreasureHoarderInsignias `json:"treasure-hoarder-insignias"`
	FatuiInsignias           FatuiInsignias           `json:"fatui-insignias"`
	WhopperflowerNectar      WhopperflowerNectar      `json:"whopperflower-nectar"`
	HilichurlHorns           HilichurlHorns           `json:"hilichurl-horns"`
	LeyLine                  LeyLine                  `json:"ley-line"`
	BoneShards               BoneShards               `json:"bone-shards"`
	MistGrass                MistGrass                `json:"mist-grass"`
	FatuiKnives              FatuiKnives              `json:"fatui-knives"`
	ChaosParts               ChaosParts               `json:"chaos-parts"`
}

type CommonAscensionMaterialItems struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type Slime struct {
	Characters []string                       `json:"characters"`
	Items      []CommonAscensionMaterialItems `json:"items"`
	Sources    []string                       `json:"sources"`
}

type HilichurlMasks struct {
	Characters []string                       `json:"characters"`
	Items      []CommonAscensionMaterialItems `json:"items"`
	Sources    []string                       `json:"sources"`
}

type HilichurlArrowheads struct {
	Characters []string                       `json:"characters"`
	Items      []CommonAscensionMaterialItems `json:"items"`
	Sources    []string                       `json:"sources"`
}

type SamachurlScrolls struct {
	Characters []string                       `json:"characters"`
	Items      []CommonAscensionMaterialItems `json:"items"`
	Sources    []string                       `json:"sources"`
}

type TreasureHoarderInsignias struct {
	Characters []string                       `json:"characters"`
	Items      []CommonAscensionMaterialItems `json:"items"`
	Sources    []string                       `json:"sources"`
}

type FatuiInsignias struct {
	Characters []string                       `json:"characters"`
	Items      []CommonAscensionMaterialItems `json:"items"`
	Sources    []string                       `json:"sources"`
}

type WhopperflowerNectar struct {
	Characters []string                       `json:"characters"`
	Items      []CommonAscensionMaterialItems `json:"items"`
	Sources    []string                       `json:"sources"`
}

type HilichurlHorns struct {
	Weapons []string                       `json:"weapons"`
	Items   []CommonAscensionMaterialItems `json:"items"`
	Sources []string                       `json:"sources"`
}

type LeyLine struct {
	Weapons []string                       `json:"weapons"`
	Items   []CommonAscensionMaterialItems `json:"items"`
	Sources []string                       `json:"sources"`
}

type BoneShards struct {
	Weapons []string                       `json:"weapons"`
	Items   []CommonAscensionMaterialItems `json:"items"`
	Sources []string                       `json:"sources"`
}

type MistGrass struct {
	Weapons []string                       `json:"weapons"`
	Items   []CommonAscensionMaterialItems `json:"items"`
	Sources []string                       `json:"sources"`
}

type FatuiKnives struct {
	Weapons []string                       `json:"weapons"`
	Items   []CommonAscensionMaterialItems `json:"items"`
	Sources []string                       `json:"sources"`
}

type ChaosParts struct {
	Weapons []string                       `json:"weapons"`
	Items   []CommonAscensionMaterialItems `json:"items"`
	Sources []string                       `json:"sources"`
}

func (coam CommonAscensionMaterial) MaterialType() string {
	return CommonAscensionMaterialType
}

func (coam CommonAscensionMaterial) EntryName() string {
	return CommonAscensionMaterialType
}

// ---- LocalSpecialtiesMaterial
type LocalSpecialtiesMaterial struct {
	Mondstadt []MondstadtLocalSpecialty `json:"mondstadt"`
	Liyue     []LiyueLocalSpecialty     `json:"liyue"`
}

type MondstadtLocalSpecialty struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type LiyueLocalSpecialty struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func (lsm LocalSpecialtiesMaterial) MaterialType() string {
	return LocalSpecialtiesMaterialType
}

func (lsm LocalSpecialtiesMaterial) EntryName() string {
	return LocalSpecialtiesMaterialType
}

// ---- TalentBookMaterial
type TalentBookMaterial struct {
	Freedom    Freedom    `json:"freedom"`
	Resistance Resistance `json:"resistance"`
	Ballad     Ballad     `json:"ballad"`
	Prosperity Prosperity `json:"prosperity"`
	Diligence  Diligence  `json:"diligence"`
	Gold       Gold       `json:"gold"`
}

type TalentBookMaterialItems struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type Freedom struct {
	Characters   []string                  `json:"characters"`
	Availability []string                  `json:"availability"`
	Source       string                    `json:"source"`
	Items        []TalentBookMaterialItems `json:"items"`
}

type Resistance struct {
	Characters   []string                  `json:"characters"`
	Availability []string                  `json:"availability"`
	Source       string                    `json:"source"`
	Items        []TalentBookMaterialItems `json:"items"`
}

type Ballad struct {
	Characters   []string                  `json:"characters"`
	Availability []string                  `json:"availability"`
	Source       string                    `json:"source"`
	Items        []TalentBookMaterialItems `json:"items"`
}

type Prosperity struct {
	Characters   []string                  `json:"characters"`
	Availability []string                  `json:"availability"`
	Source       string                    `json:"source"`
	Items        []TalentBookMaterialItems `json:"items"`
}

type Diligence struct {
	Characters   []string                  `json:"characters"`
	Availability []string                  `json:"availability"`
	Source       string                    `json:"source"`
	Items        []TalentBookMaterialItems `json:"items"`
}

type Gold struct {
	Characters   []string                  `json:"characters"`
	Availability []string                  `json:"availability"`
	Source       string                    `json:"source"`
	Items        []TalentBookMaterialItems `json:"items"`
}

func (tbookm TalentBookMaterial) MaterialType() string {
	return TalentBookMaterialType
}

func (tbookm TalentBookMaterial) EntryName() string {
	return TalentBookMaterialType
}

// ---- TalentBossMaterial
type TalentBossMaterial struct {
	TailOfBoreas         TailOfBoreas         `json:"tail-of-boreas"`
	RingOfBoreas         RingOfBoreas         `json:"ring-of-boreas"`
	SpiritLocketOfBoreas SpiritLocketOfBoreas `json:"spirit-locket-of-boreas"`
	DvalinSPlume         DvalinsPlume         `json:"dvalin's-plume"`
	DvalinSClaw          DvalinsClaw          `json:"dvalin's-claw"`
	DvalinSSigh          DvalinsSigh          `json:"dvalin's-sigh"`
	TuskOfMonocerosCaeli TuskOfMonocerosCaeli `json:"tusk-of-monoceros-caeli"`
	ShardOfAFoulLegacy   ShardOfAFoulLegacy   `json:"shard-of-a-foul-legacy"`
	ShadowOfTheWarrior   ShadowOfTheWarrior   `json:"shadow-of-the-warrior"`
}

type TailOfBoreas struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type RingOfBoreas struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type SpiritLocketOfBoreas struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type DvalinsPlume struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type DvalinsClaw struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type DvalinsSigh struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type TuskOfMonocerosCaeli struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type ShardOfAFoulLegacy struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type ShadowOfTheWarrior struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

func (tbossm TalentBossMaterial) MaterialType() string {
	return TalentBossMaterialType
}

func (tbossm TalentBossMaterial) EntryName() string {
	return TalentBossMaterialType
}

// ---- WeaponAscensionMaterial
type WeaponAscensionMaterial struct {
	Decarabian   Decarabian   `json:"decarabian"`
	Boreal       Boreal       `json:"boreal"`
	Dandelion    Dandelion    `json:"dandelion"`
	Guyun        Guyun        `json:"guyun"`
	Elixir       Elixir       `json:"elixir"`
	Aerosiderite Aerosiderite `json:"aerosiderite"`
}

type WeaponAscensionMaterialItems struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type Decarabian struct {
	Weapons      []string                       `json:"weapons"`
	Availability []string                       `json:"availability"`
	Source       string                         `json:"source"`
	Items        []WeaponAscensionMaterialItems `json:"items"`
}

type Boreal struct {
	Weapons      []string                       `json:"weapons"`
	Availability []string                       `json:"availability"`
	Source       string                         `json:"source"`
	Items        []WeaponAscensionMaterialItems `json:"items"`
}

type Dandelion struct {
	Weapons      []string                       `json:"weapons"`
	Availability []string                       `json:"availability"`
	Source       string                         `json:"source"`
	Items        []WeaponAscensionMaterialItems `json:"items"`
}

type Guyun struct {
	Weapons      []string                       `json:"weapons"`
	Availability []string                       `json:"availability"`
	Source       string                         `json:"source"`
	Items        []WeaponAscensionMaterialItems `json:"items"`
}

type Elixir struct {
	Weapons      []string                       `json:"weapons"`
	Availability []string                       `json:"availability"`
	Source       string                         `json:"source"`
	Items        []WeaponAscensionMaterialItems `json:"items"`
}

type Aerosiderite struct {
	Weapons      []string                       `json:"weapons"`
	Availability []string                       `json:"availability"`
	Source       string                         `json:"source"`
	Items        []WeaponAscensionMaterialItems `json:"items"`
}

func (wam WeaponAscensionMaterial) MaterialType() string {
	return WeaponAscensionMaterialType
}

func (wam WeaponAscensionMaterial) EntryName() string {
	return WeaponAscensionMaterialType
}

// ---- WeaponExperienceMaterial
type WeaponExperienceMaterial struct {
	Items []struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		Experience int      `json:"experience"`
		Rarity     int      `json:"rarity"`
		Source     []string `json:"source"`
	} `json:"items"`
}

func (wem WeaponExperienceMaterial) MaterialType() string {
	return WeaponExperienceMaterialType
}

func (wem WeaponExperienceMaterial) EntryName() string {
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

// GetBossMaterial : Return a BossMaterial struct.
func GetBossMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		BossMaterialType,
	}

	var bm BossMaterial
	err := getDataAndUnmarshal(reqBody, &bm)
	return bm, err
}

// GetCharacterAscensionMaterial : Return a CharacterAscensionMaterial struct.
func GetCharacterAscensionMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		CharacterAscensionMaterialType,
	}

	var cam CharacterAscensionMaterial
	err := getDataAndUnmarshal(reqBody, &cam)
	return cam, err
}

// GetCharacterExperienceMaterial : Return a CharacterExperienceMaterial struct.
func GetCharacterExperienceMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		CharacterExperienceMaterialType,
	}

	var cem CharacterExperienceMaterial
	err := getDataAndUnmarshal(reqBody, &cem)
	return cem, err
}

// GetCommonAscensionMaterial : Return a CommonAscensionMaterial struct
func GetCommonAscensionMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		CommonAscensionMaterialType,
	}

	var coam CommonAscensionMaterial
	err := getDataAndUnmarshal(reqBody, &coam)
	return coam, err
}

// GetLocalSpecialtiesMaterial : Return a LocalSpecialtiesMaterial struct
func GetLocalSpecialtiesMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		LocalSpecialtiesMaterialType,
	}

	var lsm LocalSpecialtiesMaterial
	err := getDataAndUnmarshal(reqBody, &lsm)
	return lsm, err
}

// GetTalentBookMaterial : Return a TalentBookMaterial struct
func GetTalentBookMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		TalentBookMaterialType,
	}

	var tbookm TalentBookMaterial
	err := getDataAndUnmarshal(reqBody, &tbookm)
	return tbookm, err
}

// GetTalentBossMaterial : Return a TalentBossMaterial struct
func GetTalentBossMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		TalentBossMaterialType,
	}

	var tbossm TalentBossMaterial
	err := getDataAndUnmarshal(reqBody, &tbossm)
	return tbossm, err
}

// GetWeaponAscensionMaterial : Return a WeaponAscensionMaterial struct
func GetWeaponAscensionMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		WeaponAscensionMaterialType,
	}

	var wam WeaponAscensionMaterial
	err := getDataAndUnmarshal(reqBody, &wam)
	return wam, err
}

// GetWeaponExperienceMaterial : Return a WeaponExperienceMaterial struct
func GetWeaponExperienceMaterial() (Material, error) {
	reqBody := []string{
		MaterialsDType,
		WeaponExperienceMaterialType,
	}

	var wem WeaponExperienceMaterial
	err := getDataAndUnmarshal(reqBody, &wem)
	return wem, err
}
