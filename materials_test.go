package genshinapi

import "testing"

// Test GetMaterialTypes
func TestGetMaterialTypes(t *testing.T) {
	res, err := GetMaterialTypes()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetMaterialType
func TestGetMaterialType(t *testing.T) {
	// Test with valid input
	res, err := GetMaterialType(BossMaterialType)
	if err != nil {
		t.Log(err)
		t.Error("Test with valid input failed:", err)
	} else {
		t.Logf("Test with valid input passed: %+v", res)
	}

	// Test with invalid input
	_, err = GetMaterialType("no-such-material")
	if err == nil {
		t.Error("Test with invalid input failed.")
	} else {
		t.Log("Test with invalid input passed:", err)
	}
}
