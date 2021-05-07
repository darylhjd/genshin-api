package genshinapi

import "testing"

// Test GetArtifacts
func TestGetArtifacts(t *testing.T) {
	res, err := GetArtifacts()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetArtifact
func TestGetArtifact(t *testing.T) {
	// Test with valid input
	res, err := GetArtifact("adventurer")
	if err != nil {
		t.Error("Test with valid input failed:", res)
	} else {
		t.Logf("Test with valid input passed: %+v", res)
	}

	// Test with invalid input
	res, err = GetArtifact("non-existant-set")
	if err == nil {
		t.Error("Test with invalid input failed.")
	} else {
		t.Log("Test with invalid input passed:", err)
	}
}
