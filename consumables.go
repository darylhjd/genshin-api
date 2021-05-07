package genshinapi

import (
	"encoding/json"
)

const (
	ConsumablesDType = "consumables"

	// Assets only available for potions.
	// For asset names, get key from PotionList.

	FoodConsumable   = "food"
	PotionConsumable = "potions"
)

type FoodList map[string]Food

type Food struct {
	Name        string       `json:"name"`
	Rarity      int          `json:"rarity"`
	Type        string       `json:"type"`
	Effect      string       `json:"effect"`
	HasRecipe   bool         `json:"hasRecipe"`
	Description string       `json:"description"`
	Proficiency int          `json:"proficiency"`
	Recipe      []Ingredient `json:"recipe"`
}

func (f *Food) EntryName() string {
	return f.Name
}

type PotionList map[string]Potion

type Potion struct {
	Name     string       `json:"name"`
	Effect   string       `json:"effect"`
	Rarity   int          `json:"rarity"`
	Crafting []Ingredient `json:"crafting"`
}

func (p *Potion) EntryName() string {
	return p.Name
}

type Ingredient struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

// GetConsumables : Get a list of consumables
func GetConsumables() ([]string, error) {
	return GetDataTypeItemsList(ConsumablesDType)
}

// GetFoodList : Return a map of Food items.
func GetFoodList() (FoodList, error) {
	reqBody := []string{
		ConsumablesDType,
		FoodConsumable,
	}

	bytearray, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var list FoodList
	err = json.Unmarshal(bytearray, &list)
	return list, err
}

// GetPotionList : Get a list of potions
func GetPotionList() (PotionList, error) {
	reqBody := []string{
		ConsumablesDType,
		PotionConsumable,
	}

	bytearray, err := GetCustomBody(reqBody...)
	if err != nil {
		return nil, err
	}

	var list PotionList
	err = json.Unmarshal(bytearray, &list)
	return list, err
}
