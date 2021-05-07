package genshinapi

import "testing"

// Test GetNations
func TestGetNations(t *testing.T) {
	res, err := GetNations()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetNation
func TestGetNation(t *testing.T) {
	// Test with valid input
	res, err := GetNation("liyue")
	if err != nil {
		t.Error("Test with valid input failed:", err)
	} else {
		t.Logf("Test with valid input passed: %+v", res)
	}

	// Test with invalid input
	_, err = GetNation("country")
	if err == nil {
		t.Error("Test with invalid input failed.")
	} else {
		t.Log("Test with invalid input passed:", err)
	}
}
