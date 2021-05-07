package genshinapi

import "testing"

// Test GetDataTypeItemsList
func TestGetDataTypeItemsList(t *testing.T) {
	// Test with proper data type
	_, err := GetDataTypeItemsList(ArtifactsDType)
	if err != nil {
		t.Error("Testing valid data type failed. Expected nil, got:", err)
	} else {
		t.Log("Testing valid data type passed.")
	}

	// Test with improper data type
	_, err = GetDataTypeItemsList("wrong")
	if err == nil {
		t.Error("Testing invalid data type failed. Expected error, got nil.")
	} else {
		t.Log("Testing invalid data type passed.")
	}
	_, err = GetDataTypeItemsList("///////")
	if err == nil {
		t.Error("Testing invalid data type failed. Expected error, got nil.")
	} else {
		t.Log("Testing invalid data type passed.")
	}
}

// Test GetDataTypes
func TestGetDataTypes(t *testing.T) {
	_, err := GetDataTypes()
	if err != nil {
		t.Error("Testing GetDataTypes failed")
	} else {
		t.Log("Testing GetDataTypes passed")
	}
}

// Test GetImage
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
