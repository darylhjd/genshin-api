package genshinapi

import (
	"encoding/json"
)

type GetDataNameListFnType func() ([]string, error)

// GetDataTypes : Get data types provided by the API
func GetDataTypes() ([]string, error) {
	bytes, err := GetCustomBody()
	if err != nil {
		return nil, err
	}

	var dts DataTypes
	err = json.Unmarshal(bytes, &dts)
	if err != nil {
		return nil, err
	}
	return dts.Types, nil
}

// GetArtifacts : Get list of artifact names in Genshin Impact
func GetArtifacts() ([]string, error) {
	return genshinAPIGetDataList(ArtifactsDType)
}

// GetArtifact : Return a pointer to an Artifact struct containing info on an artifact
func GetArtifact(name string) (*Artifact, error) {
	reqBody := []string{
		ArtifactsDType,
		name,
	}

	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var artifact Artifact
	err = json.Unmarshal(bytes, &artifact)
	if err != nil {
		return nil, err
	}
	return &artifact, nil
}

// GetCharacters : Get list of character names in Genshin Impact
func GetCharacters() ([]string, error) {
	return genshinAPIGetDataList(CharactersDType)
}

// GetCharacter : Return a pointer to a Character struct containing info on a character
func GetCharacter(name string) (*Character, error) {
	reqBody := []string{
		CharactersDType,
		name,
	}

	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var character Character
	err = json.Unmarshal(bytes, &character)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

// GetDomains : Get list of domain names in Genshin Impact
func GetDomains() ([]string, error) {
	return genshinAPIGetDataList(DomainsDType)
}

// GetDomain : Return a pointer to a Domain struct containing info on a domain
func GetDomain(name string) (*Domain, error) {
	reqBody := []string{
		DomainsDType,
		name,
	}

	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var domain Domain
	err = json.Unmarshal(bytes, &domain)
	if err != nil {
		return nil, err
	}
	return &domain, nil
}

// GetElements : Get list of element names in Genshin Impact
func GetElements() ([]string, error) {
	return genshinAPIGetDataList(ElementsDType)
}

// GetElement : Return a pointer to an Element struct containing info on an element
func GetElement(name string) (*Element, error) {
	reqBody := []string{
		ElementsDType,
		name,
	}

	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var element Element
	err = json.Unmarshal(bytes, &element)
	if err != nil {
		return nil, err
	}
	return &element, nil
}

// GetNations : Get list of nation names in Genshin Impact
func GetNations() ([]string, error) {
	return genshinAPIGetDataList(NationsDType)
}

// GetNation : Return a pointer to a Nation struct containing info on a nation
func GetNation(name string) (*Nation, error) {
	reqBody := []string{
		NationsDType,
		name,
	}

	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var nation Nation
	err = json.Unmarshal(bytes, &nation)
	if err != nil {
		return nil, err
	}
	return &nation, nil
}

// GetWeapons : Get list of weapon names in Genshin Impact
func GetWeapons() ([]string, error) {
	return genshinAPIGetDataList(WeaponsDType)
}

// GetWeapon : Return a pointer to a Weapon struct containing info on a weapon
func GetWeapon(name string) (*Weapon, error) {
	reqBody := []string{
		WeaponsDType,
		name,
	}

	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var weapon Weapon
	err = json.Unmarshal(bytes, &weapon)
	if err != nil {
		return nil, err
	}
	return &weapon, nil
}

// GetImage : Return a byte array of image data for a data type entry
// Validity of types is not guaranteed.
func GetImage(dtype, name, imagetype string) ([]byte, error) {
	return GetCustomBody(dtype, name, imagetype)
}
