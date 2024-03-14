package classes

import (
	"encoding/json"
	"log"
	"os"
	"pregen/pkg/db"
)

var WeaponData = GetWeaponFormDB()

type Weapon struct {
	Name       string `json:"weaponName"`
	WeaponType string `json:"weaponType"`
	Damage     string `json:"damage"`
	Property   string `json:"property"`
	Count      int    `json:"count"`
}

func GetWeaponFormDB() []Weapon {
	var results []Weapon
	fileContent, err := os.Open(db.DataBasePath + "weapons.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile(db.DataBasePath + "weapons.json")

	json.Unmarshal(byteResult, &results)
	return results
}

func FindWeaponInDB(spellName string) Weapon {
	for _, weapon := range WeaponData {
		if weapon.Name == spellName {
			return weapon
		}
	}
	return Weapon{}
}
