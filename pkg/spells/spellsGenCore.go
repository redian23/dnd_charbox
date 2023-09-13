package spells

import (
	"context"
	"fmt"
	"pregen/pkg/db"
)

func GetSpellsFormDB() []SpellsBSON {
	var results []SpellsBSON
	var cursor = db.ReadFromDB("spells")

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func GetSpellsForClass(classNameRu string, count int) map[string]string {
	spells := GetSpellsFormDB()
	var spellList = make(map[string]string)
	var spellAnswer []string
	for _, spell := range spells {
		for _, clas := range spell.Classes {
			if clas == classNameRu {
				spellList[spell.SpellNameRu] = spell.SpellNameRu
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
	fmt.Println(spellAnswer)
	return spellList
}
