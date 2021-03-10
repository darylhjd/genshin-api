package genshinapi

import (
	"encoding/json"
)

// genshinAPIGetDataList : Get a list of items of a particular data type
func genshinAPIGetDataList(t string) ([]string, error) {
	resp, err := GetCustomBody(t)
	if err != nil {
		return nil, err
	}

	var list []string
	err = json.Unmarshal(resp, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

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
func GetArtifact(name string) (Artifact, error) {
	reqBody := []string{
		ArtifactsDType,
		name,
	}

	var artifact Artifact
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return artifact, err
	}

	err = json.Unmarshal(bytes, &artifact)
	if err != nil {
		return artifact, err
	}
	return artifact, nil
}

// GetCharacters : Get list of character names in Genshin Impact
func GetCharacters() ([]string, error) {
	return genshinAPIGetDataList(CharactersDType)
}

// GetCharacter : Return a pointer to a Character struct containing info on a character
func GetCharacter(name string) (Character, error) {
	reqBody := []string{
		CharactersDType,
		name,
	}

	var character Character
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return character, err
	}

	err = json.Unmarshal(bytes, &character)
	if err != nil {
		return character, err
	}
	return character, nil
}

// GetDomains : Get list of domain names in Genshin Impact
func GetDomains() ([]string, error) {
	return genshinAPIGetDataList(DomainsDType)
}

// GetDomain : Return a pointer to a Domain struct containing info on a domain
func GetDomain(name string) (Domain, error) {
	reqBody := []string{
		DomainsDType,
		name,
	}

	var domain Domain
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return domain, err
	}

	err = json.Unmarshal(bytes, &domain)
	if err != nil {
		return domain, err
	}
	return domain, nil
}

// GetElements : Get list of element names in Genshin Impact
func GetElements() ([]string, error) {
	return genshinAPIGetDataList(ElementsDType)
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

// GetNations : Get list of nation names in Genshin Impact
func GetNations() ([]string, error) {
	return genshinAPIGetDataList(NationsDType)
}

// GetNation : Return a pointer to a Nation struct containing info on a nation
func GetNation(name string) (Nation, error) {
	reqBody := []string{
		NationsDType,
		name,
	}

	var nation Nation
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return nation, err
	}

	err = json.Unmarshal(bytes, &nation)
	if err != nil {
		return nation, err
	}
	return nation, nil
}

// GetWeapons : Get list of weapon names in Genshin Impact
func GetWeapons() ([]string, error) {
	return genshinAPIGetDataList(WeaponsDType)
}

// GetWeapon : Return a pointer to a Weapon struct containing info on a weapon
func GetWeapon(name string) (Weapon, error) {
	reqBody := []string{
		WeaponsDType,
		name,
	}

	var weapon Weapon
	bytes, err := GetCustomBody(reqBody...)
	if err != nil {
		return weapon, err
	}

	err = json.Unmarshal(bytes, &weapon)
	if err != nil {
		return weapon, err
	}
	return weapon, nil
}

// GetImage : Return a byte array of image data for a data type entry
// Returns error if request is not valid (validity of request is guaranteed).
func GetImage(dtype, name, imagetype string) ([]byte, error) {
	return GetCustomBody(dtype, name, imagetype)
}
