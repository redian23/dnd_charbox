package spells

import (
	"context"
	"pregen/pkg/classes"
	"pregen/pkg/db"
	"pregen/pkg/races"
)

func GetSpellsFormDB(collName string) []SpellsBSON {
	var results []SpellsBSON
	var cursor = db.ReadFromDB(collName)

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func GetSpellsZeroLevelForCharacter() []string {
	classSpells := GetHtmlFormattedClassSpells(0)
	raceSpells := GetRaceSpellsZeroLevel()

	spellList := append(classSpells, raceSpells...)
	return spellList
}

func GetHtmlFormattedClassSpells(lvl int) []string {
	spells := GetSpellsFormDB("spells_all")
	classSpells := classes.GetClassSpells(lvl)

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

func GetRaceSpellsZeroLevel() []string {
	spellsListFromBD := GetSpellsFormDB("spells_all")

	var raceZeroSpellsList []races.RaceZeroLvLSpells
	var raceSpells []string
	for _, race := range races.RaceData {
		for _, rType := range race.Type {
			if races.RaceTypeGlobalRu == rType.TypeRaceNameRu {
				raceZeroSpellsList = rType.RaceZeroLvLSpells
			}
		}
	}

	for _, ZeroDBSpl := range spellsListFromBD {
		for _, raceZeroSpl := range raceZeroSpellsList {
			if ZeroDBSpl.SpellNameRu == raceZeroSpl.SpellName {
				if raceZeroSpl.SpellName == "" {
					break
				}
				spl := "<a href=" + ZeroDBSpl.URL + ">" + ZeroDBSpl.SpellNameRu + " [" + ZeroDBSpl.SpellName + "]" + "</a>"
				raceSpells = append(raceSpells, spl)
			}
		}
	}

	return raceSpells
}

func GetSpellsOneLevelForCharacter() []string {
	classSpells := GetHtmlFormattedClassSpells(1)
	raceSpells := GetRaceSpellsOneLevel()

	spellList := append(raceSpells, classSpells...)
	spellList = append(spellList, classes.ClassSpecialSpells...)
	return spellList
}

func GetRaceSpellsOneLevel() []string {
	spellsListFromBD := GetSpellsFormDB("spells_all")

	var raceOneSpellsList []races.RaceOneLvLSpells
	var raceSpells []string

	for _, race := range races.RaceData {
		for _, rType := range race.Type {
			if races.RaceTypeGlobalRu == rType.TypeRaceNameRu {
				raceOneSpellsList = rType.RaceOneLvLSpells
			}
		}
	}

	for _, ZeroDBSpl := range spellsListFromBD {
		for _, raceZeroSpl := range raceOneSpellsList {
			if ZeroDBSpl.SpellNameRu == raceZeroSpl.SpellName {
				if raceZeroSpl.SpellName == "" {
					break
				}
				spl := "<a href=" + ZeroDBSpl.URL + ">" + ZeroDBSpl.SpellNameRu + " [" + ZeroDBSpl.SpellName + "]" + "</a>"
				raceSpells = append(raceSpells, spl)
			}
		}
	}

	return raceSpells
}
