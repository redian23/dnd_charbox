package spells

import (
	"context"
	"pregen/pkg/db"
)

func GetSpellsFormDB(collName string) []SpellsBSON {
	var results []SpellsBSON
	var cursor = db.ReadFromDB(collName)

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func GetSpellsZeroLevelForClass(classNameRu string, count int) []string {
	spells := GetSpellsFormDB("spells_zero_lvl")
	var spellList = make(map[string]string)
	var spellAnswer []string
	for _, spell := range spells {
		for _, clas := range spell.Classes {
			if clas == classNameRu {
				spellList[spell.SpellNameRu] = spell.SpellNameRu + " [" + spell.SpellName + "]"
			}
		}
	}

	var iter int
	for _, s2 := range spellList {
		spellAnswer = append(spellAnswer, s2)
		iter++
		if iter >= count {
			break
		}
	}
	return spellAnswer
}
