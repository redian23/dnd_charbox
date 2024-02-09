package spells

import (
	"encoding/json"
	"log"
	"os"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

func GetSpellsFormDB() []SpellsBSON {
	var results []SpellsBSON
	fileContent, err := os.Open("./pkg/db/spells_all.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile("./pkg/db/spells_all.json")

	json.Unmarshal(byteResult, &results)
	return results
}

func removeDuplicates(slice []string) []string {
	// Create a map to store unique elements
	seen := make(map[string]bool)
	var result []string

	// Loop through the slice, adding elements to the map if they haven't been seen before
	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func SetSpellsForCharacter(lvl int) []string {
	var raceSpells []string

	classSpells := GetHtmlFormattedClassSpells(lvl)
	raceSpells = GetHtmlFormattedRaceSpells(lvl)

	spellList := append(classSpells, raceSpells...)
	spellList = removeDuplicates(spellList)

	return spellList
}

func GetClassSpells(level int, classNameReq string) []string {
	var castCount int

	var casting = classes.GetClassSpellBasicCharacteristic()
	var spellData = GetSpellsFormDB()

	switch level {
	case 0:
		castCount = casting.ZeroLevelSpellsKnownCount
	case 1:
		castCount = casting.OneLevelSpellsKnownCount
	case 2:
		castCount = casting.TwoLevelSpellsKnownCount
	case 3:
		castCount = casting.TreeLevelSpellsKnownCount
	case 4:
		castCount = casting.FourLevelSpellsKnownCount
	}

	var classSpellsMap = make(map[string]string)
	var classSpellsAnswer []string
	for _, classSpellValue := range spellData {
		for _, className := range classSpellValue.Classes {
			if classNameReq == className && classSpellValue.SpellLevel == level {
				classSpellsMap[classSpellValue.SpellName] = classSpellValue.SpellNameRu
			}
		}
	}
	var iter int
	for _, clSpell := range classSpellsMap {
		if iter < castCount {
			iter++
			classSpellsAnswer = append(classSpellsAnswer, clSpell)
		} else {
			break
		}
	}

	return classSpellsAnswer
}

func GetHtmlFormattedClassSpells(lvl int) []string {
	spells := GetSpellsFormDB()
	classSpells := GetClassSpells(lvl, classes.ClassNameGlobalRu)

	var spellAnswer []string

	for _, spell := range spells {
		for _, classSpell := range classSpells {
			if classSpell == spell.SpellNameRu && spell.SpellLevel == lvl {
				spellAnswer = append(spellAnswer, "<a href="+spell.URL+">"+spell.SpellNameRu+" ["+spell.SpellName+"]"+"</a>")
			}
		}
	}
	return spellAnswer
}

func GetHtmlFormattedRaceSpells(lvl int) []string {
	spells := GetSpellsFormDB()

	var spellAnswer []string
	var raceSpells []races.RaceLvLSpells

	for _, race := range races.RaceData {
		for _, rType := range race.Type {
			if races.RaceTypeGlobalRu == rType.TypeRaceNameRu {
				if lvl == 0 {
					raceSpells = rType.RaceZeroLvLSpells
				}
				if lvl == 1 {
					raceSpells = rType.RaceOneLvLSpells
				}
			}
		}
	}

	for _, spell := range spells {
		for _, rSpell := range raceSpells {
			if rSpell.SpellName == spell.SpellNameRu && spell.SpellLevel == lvl {
				spellAnswer = append(spellAnswer, "<a href="+spell.URL+">"+spell.SpellNameRu+" ["+spell.SpellName+"]"+"</a>")
			}
		}
	}
	return spellAnswer
}

func GetHtmlFormattedAddictionSpells(lvl, count int) []string {
	spells := GetSpellsFormDB()
	classSpells := GetClassSpells(lvl, classes.ClassNameGlobalRu)

	var spellAnswer []string

	var iter int
	for _, spell := range spells {
		for _, classSpell := range classSpells {
			if classSpell == spell.SpellNameRu && spell.SpellLevel == lvl {
				spellAnswer = append(spellAnswer, "<a href="+spell.URL+">"+spell.SpellNameRu+" ["+spell.SpellName+"]"+"</a>")
			}
			if iter == count {
				break
			}
		}
	}
	return spellAnswer
}
