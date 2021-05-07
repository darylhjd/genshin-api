package genshinapi

import "testing"

// Test GetDomains
func TestGetDomains(t *testing.T) {
	res, err := GetDomains()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetDomain
func TestGetDomain(t *testing.T) {
	// Test with valid input
	res, err := GetDomain("cecilia-garden")
	if err != nil {
		t.Error("Test with valid input failed:", err)
	} else {
		t.Logf("Test with valid input passed: %+v", res)
	}

	// Test with invalid input
	_, err = GetDomain("my-house")
	if err == nil {
		t.Error("Test with invalid input failed.")
	} else {
		t.Log("Test with invalid input passed:", err)
	}
}
