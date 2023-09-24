package spells

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"pregen/pkg/classes"
	"pregen/pkg/db"
	"pregen/pkg/races"
	"time"
)

func GetSpellsFormDB() []SpellsBSON {
	var results []SpellsBSON
	client := db.ConnectToDB()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

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

func GetSpellsZeroLevelForCharacter() []string {
	classSpells := GetHtmlFormattedClassSpells(0)
	raceSpells := GetRaceSpellsZeroLevel()

	spellList := append(classSpells, raceSpells...)
	return spellList
}

func GetHtmlFormattedClassSpells(lvl int) []string {
	spells := GetSpellsFormDB()
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
	spellsListFromBD := GetSpellsFormDB()

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
	spellsListFromBD := GetSpellsFormDB()

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
