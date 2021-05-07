package genshinapi

import "testing"

// Test GetCharacters
func TestGetCharacters(t *testing.T) {
	res, err := GetCharacters()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetCharacter
func TestGetCharacter(t *testing.T) {
	// Test with valid input
	res, err := GetCharacter("amber")
	if err != nil {
		t.Error("Test with valid input failed:", err)
	} else {
		t.Logf("Test with valid input passed: %+v", res)
	}

	// Test with invalid input
	_, err = GetCharacter("bruh")
	if err == nil {
		t.Error("Test with invalid input failed.")
	} else {
		t.Log("Test with invalid input passed:", err)
	}
}
