package genshinapi

import "testing"

// Test GetElements
func TestGetElements(t *testing.T) {
	res, err := GetElements()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetElement
func TestGetElement(t *testing.T) {
	// Test with valid input
	res, err := GetElement("anemo")
	if err != nil {
		t.Error("Test with valid input failed:", err)
	} else {
		t.Logf("Test with valid input passed: %+v", res)
	}

	// Test with invalid input
	_, err = GetElement("dirt")
	if err == nil {
		t.Error("Test with invalid input failed.")
	} else {
		t.Log("Test with invalid input passed:", err)
	}
}
