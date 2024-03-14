package spells

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"pregen/pkg/db"
)

var SpellData = GetSpellsFormDB()

type SpellsJSON struct {
	ID          primitive.ObjectID `json:"_id"`
	Source      string             `json:"source"`
	Classes     []string           `json:"classes"`
	SpellLevel  int                `json:"spellLevel"`
	SpellName   string             `json:"spellNameEng"`
	SpellNameRu string             `json:"spellNameRu"`
	URL         string             `json:"url"`
}

func GetSpellsFormDB() []SpellsJSON {
	var results []SpellsJSON
	fileContent, err := os.Open(db.DataBasePath + "spells_all.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile(db.DataBasePath + "spells_all.json")

	json.Unmarshal(byteResult, &results)
	return results
}

func FindSpellInDB(spellName string) SpellsJSON {
	for _, spell := range SpellData {
		if spell.SpellNameRu == spellName {
			return spell
		}
	}
	return SpellsJSON{}
}

func GetListSpellsFromDB(spellNames ...string) []SpellsJSON {
	var spellList []SpellsJSON

	for _, spell := range SpellData {
		for _, spellName := range spellNames {
			if spell.SpellNameRu == spellName {
				spellList = append(spellList, spell)
			}
		}
	}
	return spellList
}
