package genshinapi

import (
	"testing"
)

// Test GetCustomBody
func TestGetCustomBody(t *testing.T) {
	// Test valid URL arguments
	ext := []string{
		CharactersDType,
		"amber",
	}
	_, err := GetCustomBody(ext...)
	if err != nil {
		t.Error("Testing valid URL arguments failed. Expected nil, got:", err)
	} else {
		t.Log("Testing valid URL arguments passed.")
	}

	// Test invalid extension arguments
	_, err = GetCustomBody("//////////")
	if err == nil {
		t.Error("Testing invalid URL failed. Expected error, got nil.")
	} else {
		t.Log("Testing invalid URL arguments passed.")
	}
}
