package spells

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"pregen/pkg/classes"
	"pregen/pkg/db"
	"pregen/pkg/races"
)

func GetSpellsFormDB() []SpellsBSON {
	var results []SpellsBSON
	client := db.ConnectToDB()

	ctx := context.Background()

	coll := client.Database(db.DBNAME).Collection("spells_all")
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)
	return results
}

func SetSpellsForCharacter(lvl int) []string {
	var raceSpells []string

	classSpells := GetHtmlFormattedClassSpells(lvl)
	raceSpells = GetHtmlFormattedRaceSpells(lvl)

	spellList := append(classSpells, raceSpells...)
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
