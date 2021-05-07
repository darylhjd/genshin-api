package genshinapi

import "testing"

// Test GetWeapons
func TestGetWeapons(t *testing.T) {
	res, err := GetWeapons()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetWeapon
func TestGetWeapon(t *testing.T) {
	// Test with valid input
	res, err := GetWeapon("rust")
	if err != nil {
		t.Error("Test with valid input failed:", err)
	} else {
		t.Logf("Test with valid input passed: %+v", res)
	}

	// Test with invalid input
	_, err = GetWeapon("sword")
	if err == nil {
		t.Error("Test with invalid input failed.")
	} else {
		t.Log("Test with invalid input passed:", err)
	}
}
