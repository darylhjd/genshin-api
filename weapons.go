package genshinapi

const (
	// API DataType name
	WeaponsDType = "weapons"
	// Assets
	WeaponIcon = "icon"
)

// --------- Structs for Weapon

type Weapon struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Rarity      int    `json:"rarity"`
	BaseAttack  string `json:"baseAttack"`
	SubStat     string `json:"subStat"`
	PassiveName string `json:"passiveName"`
	PassiveDesc string `json:"passiveDesc"`
	Location    string `json:"location"`
}

func (w Weapon) EntryName() string {
	return w.Name
}

// --------- Helper functions

// GetWeapons : Get list of weapon names.
func GetWeapons() ([]string, error) {
	return GetDataTypeItemsList(WeaponsDType)
}

// GetWeapon : Return a requested Weapon struct.
func GetWeapon(name string) (Weapon, error) {
	reqBody := []string{
		WeaponsDType,
		name,
	}

	var weapon Weapon
	err := getDataAndUnmarshal(reqBody, &weapon)
	return weapon, err
}
