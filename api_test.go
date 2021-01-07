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
