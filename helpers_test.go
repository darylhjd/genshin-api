package genshinapi

import "testing"

// Test genshinAPIGetDataList
func TestGenshinAPIGetDataList(t *testing.T) {
	// Test with proper data type
	_, err := genshinAPIGetDataList(ArtifactsDType)
	if err != nil {
		t.Error("Testing valid data type failed. Expected nil, got:", err)
	} else {
		t.Log("Testing valid data type passed.")
	}

	// Test with improper data type
	_, err = genshinAPIGetDataList("wrong")
	if err == nil {
		t.Error("Testing invalid data type failed. Expected error, got nil.")
	} else {
		t.Log("Testing invalid data type passed.")
	}
	_, err = genshinAPIGetDataList("///////")
	if err == nil {
		t.Error("Testing invalid data type failed. Expected error, got nil.")
	} else {
		t.Log("Testing invalid data type passed.")
	}
}

func TestGetTypes(t *testing.T) {
	_, err := GetDataTypes()
	if err != nil {
		t.Error("Testing GetDataTypes failed")
	} else {
		t.Log("Testing GetDataTypes passed")
	}
}

func TestGetArtifacts(t *testing.T) {
	_, err := GetArtifacts()
	if err != nil {
		t.Error("Testing GetArtifacts failed")
	} else {
		t.Log("Testing GetArtifacts passed")
	}
}

func TestGetArtifact(t *testing.T) {
	// Test with valid input
	_, err := GetArtifact("adventurer")
	if err != nil {
		t.Error("Testing GetArtifact with valid input failed")
	} else {
		t.Log("Testing GetArtifact with valid input passed")
	}

	// Test with invalid input
	_, err = GetArtifact("non-existant-set")
	if err == nil {
		t.Error("Testing GetArtifact with invalid input failed")
	} else {
		t.Log("Testing GetArtifact with invalid input passed")
	}
}

func TestGetCharacters(t *testing.T) {
	_, err := GetCharacters()
	if err != nil {
		t.Error("Testing GetCharacters failed")
	} else {
		t.Log("Testing GetCharacters passed")
	}
}

func TestGetCharacter(t *testing.T) {
	// Test with valid input
	_, err := GetCharacter("amber")
	if err != nil {
		t.Error("Testing GetCharacter with valid input failed")
	} else {
		t.Log("Testing GetCharacter with valid input passed")
	}

	// Test with invalid input
	_, err = GetCharacter("bruh")
	if err == nil {
		t.Error("Testing GetCharacter with invalid input failed")
	} else {
		t.Log("Testing GetCharacter with invalid input passed")
	}
}

func TestGetDomains(t *testing.T) {
	_, err := GetDomains()
	if err != nil {
		t.Error("Testing GetDomains failed")
	} else {
		t.Log("Testing GetDomains passed")
	}
}

func TestGetDomain(t *testing.T) {
	// Test with valid input
	_, err := GetDomain("cecilia-garden")
	if err != nil {
		t.Error("Testing GetDomain with valid input failed")
	} else {
		t.Log("Testing GetDomain with valid input passed")
	}

	// Test with invalid input
	_, err = GetDomain("my-house")
	if err == nil {
		t.Error("Testing GetDomain with invalid input failed")
	} else {
		t.Log("Testing GetDomain with invalid input passed")
	}
}

func TestGetElements(t *testing.T) {
	_, err := GetElements()
	if err != nil {
		t.Error("Testing GetElements failed")
	} else {
		t.Log("Testing GetElements passed")
	}
}

func TestGetElement(t *testing.T) {
	// Test with valid input
	_, err := GetElement("anemo")
	if err != nil {
		t.Error("Testing GetElement with valid input failed")
	} else {
		t.Log("Testing GetElement with valid input passed")
	}

	// Test with invalid input
	_, err = GetElement("dirt")
	if err == nil {
		t.Error("Testing GetElement with invalid input failed")
	} else {
		t.Log("Testing GetElement with invalid input passed")
	}
}

func TestGetNations(t *testing.T) {
	_, err := GetNations()
	if err != nil {
		t.Error("Testing GetNations failed")
	} else {
		t.Log("Testing GetNations passed")
	}
}

func TestGetNation(t *testing.T) {
	// Test with valid input
	_, err := GetNation("liyue")
	if err != nil {
		t.Error("Testing GetNation with valid input failed")
	} else {
		t.Log("Testing GetNation with valid input passed")
	}

	// Test with invalid input
	_, err = GetNation("country")
	if err == nil {
		t.Error("Testing GetNation with invalid input failed")
	} else {
		t.Log("Testing GetNation with invalid input passed")
	}
}

func TestGetWeapons(t *testing.T) {
	_, err := GetWeapons()
	if err != nil {
		t.Error("Testing GetWeapons failed")
	} else {
		t.Log("Testing GetWeapons passed")
	}
}

func TestGetWeapon(t *testing.T) {
	// Test with valid input
	_, err := GetWeapon("rust")
	if err != nil {
		t.Error("Testing GetWeapon with valid input failed")
	} else {
		t.Log("Testing GetWeapon with valid input passed")
	}

	// Test with invalid input
	_, err = GetWeapon("sword")
	if err == nil {
		t.Error("Testing GetWeapon with invalid input failed")
	} else {
		t.Log("Testing GetWeapon with invalid input passed")
	}
}

func TestGetImage(t *testing.T) {
	// Test with valid input
	_, err := GetImage(ArtifactsDType, "adventurer", ArtifactPhoto)
	if err != nil {
		t.Error("Testing GetImage with valid input failed", err)
	} else {
		t.Log("Testing GetImage with valid input passed")
	}
	_, err = GetImage(CharactersDType, "amber", CharacterPortrait)
	if err != nil {
		t.Error("Testing GetImage with valid input failed", err)
	} else {
		t.Log("Testing GetImage with valid input passed")
	}
	_, err = GetImage(ElementsDType, "anemo", ElementIcon)
	if err != nil {
		t.Error("Testing GetImage with valid input failed", err)
	} else {
		t.Log("Testing GetImage with valid input passed")
	}
	_, err = GetImage(WeaponsDType, "rust", WeaponIcon)
	if err != nil {
		t.Error("Testing GetImage with valid input failed", err)
	} else {
		t.Log("Testing GetImage with valid input passed")
	}

	// Test GetImage with invalid data type name
	_, err = GetImage(CharactersDType, "no-such-char", CharacterIcon)
	if err == nil {
		t.Error("Testing GetImage with invalid data type name failed")
	} else {
		t.Log("Testing GetImage with invalid data type name passed")
	}

	// Test GetImage with invalid image type
	_, err = GetImage(ArtifactsDType, "adventurer", "wrongimage")
	if err == nil {
		t.Error("Testing GetImage with invalid image type failed")
	} else {
		t.Log("Testing GetImage with invalid image type passed")
	}

	// Test GetImage with invalid data type
	_, err = GetImage("wrongtype", "adventurer", ArtifactPhoto)
	if err == nil {
		t.Error("Testing GetImage with invalid data type failed")
	} else {
		t.Log("Testing GetImage with invalid data type passed")
	}
}
