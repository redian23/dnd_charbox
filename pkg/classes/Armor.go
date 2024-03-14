package classes

import (
	"encoding/json"
	"log"
	"os"
	"pregen/pkg/db"
)

type Armor struct {
	Name            string `json:"armorName"`
	ArmorType       string `json:"armorType"`
	ArmorClassCount int    `json:"armorClassCount"`
	Stealth         bool   `json:"armorStealth"`
}

func GetArmorFormDB() []Armor {
	var results []Armor
	fileContent, err := os.Open(db.DataBasePath + "armor.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile(db.DataBasePath + "armor.json")

	json.Unmarshal(byteResult, &results)
	return results
}
