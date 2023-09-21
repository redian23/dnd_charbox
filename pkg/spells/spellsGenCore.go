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

func GetClassSpellsZeroLevel() []string {
	spells := GetSpellsFormDB("spells_zero_lvl")
	spellCount := classes.GetClassZeroSpellCount()

	var spellList = make(map[string]string)
	var spellAnswer []string

	for _, spell := range spells {
		for _, clas := range spell.Classes {
			if clas == classes.ClassNameGlobalRu {
				spellList[spell.SpellNameRu] = spell.SpellNameRu + " [" + spell.SpellName + "]"
			}
		}
	}

	var iter int
	for _, s2 := range spellList {
		spellAnswer = append(spellAnswer, s2)
		iter++
		if iter >= spellCount {
			break
		}
	}
	return spellAnswer
}

func GetSpellsZeroLevelForCharacter() []string {
	classSpells := GetClassSpellsZeroLevel()
	raceSpells := races.GetRaceSpellsZeroLevel()

	spellList := append(classSpells, raceSpells...)
	return spellList
}

func GetSpellsOneLevelForCharacter() []string {
	classSpells := classes.GetClassOneSpellsList()
	raceSpells := races.GetRaceSpellsOneLevel()

	spellList := append(raceSpells, classSpells...)
	spellList = append(spellList, classes.ClassSpecialSpells...)
	return spellList
}
