package spells

import (
	"encoding/json"
	"log"
	"os"
	"pregen/pkg/db"
)

var SpellData = GetSpellsFormDB()

type SpellsJSON struct {
	Source      string   `json:"source"`
	Classes     []string `json:"classes"`
	SpellLevel  int      `json:"spellLevel"`
	SpellName   string   `json:"spellNameEng"`
	SpellNameRu string   `json:"spellNameRu"`
	URL         string   `json:"url"`
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
		if spell.SpellNameRu == spellName || spell.SpellName == spellName {
			return spell
		}
	}
	return SpellsJSON{}
}

func FindSpellListInDB(spellsArray []string) []SpellsJSON {
	var spellList []SpellsJSON
	for _, spell := range SpellData {
		for _, spellName := range spellsArray {
			if spell.SpellNameRu == spellName {
				spellList = append(spellList, spell)
			}
		}

	}
	return spellList
}

func GetAllSpellForClass(className string, spellLevel int) []SpellsJSON {
	var spellList []SpellsJSON

	for _, spell := range SpellData {
		for _, class := range spell.Classes {
			if class == className && spell.SpellLevel <= spellLevel {
				spellList = append(spellList, spell)
			}
		}
	}
	return spellList
}

func GetRandomSpellForClass(className string, spellLevel, count int) []SpellsJSON {
	var spellMap = make(map[int]SpellsJSON)

	for i, spell := range SpellData {
		for _, class := range spell.Classes {
			if class == className && spell.SpellLevel == spellLevel {
				spellMap[i] = spell
			}
		}
	}

	var iter = 0
	var spellList []SpellsJSON

	for _, spellJSON := range spellMap {
		if iter >= count {
			break
		}
		iter++
		spellList = append(spellList, spellJSON)
	}
	return spellList
}

func GetRandomSpellsForClassFromCustomSpellList(className string, spellList []SpellsJSON, spellLevel, count int) []SpellsJSON {
	var spellMap = make(map[int]SpellsJSON)

	for i, spell := range spellList {
		for _, class := range spell.Classes {
			if class == className && spell.SpellLevel == spellLevel {
				spellMap[i] = spell
			}
		}
	}

	var iter = 0
	var answerSpellList []SpellsJSON

	for _, spellJSON := range spellMap {
		if iter >= count {
			break
		}
		iter++
		answerSpellList = append(answerSpellList, spellJSON)
	}
	return answerSpellList
}
