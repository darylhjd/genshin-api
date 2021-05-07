package genshinapi

import "testing"

// Test GetConsumables
func TestGetConsumables(t *testing.T) {
	res, err := GetConsumables()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetFoodList
func TestGetFoodList(t *testing.T) {
	res, err := GetFoodList()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}

// Test GetPotionList
func TestGetPotionList(t *testing.T) {
	res, err := GetPotionList()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}
}
