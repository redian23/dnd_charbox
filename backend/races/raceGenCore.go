package races

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mazen160/go-random"
	"pregen/db"
)

var raceData = getRacesFormDB()
var raceName string
var raceGender string

func InsertRacesToDB() {
	client := db.ConnectToDB()
	coll := client.Database("data").Collection("races")

	var race RacesJsonStruct
	json.Unmarshal([]byte{}, &race) //пока что ничего не записывает

	docs := []interface{}{}

	for _, cl := range race {
		fmt.Println(cl)
		docs = append(docs, cl)
	}

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
}

func getRacesFormDB() []RacesBSON {
	var results []RacesBSON
	var cursor = db.ReadFromDB("races")

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func setRaceName() (string, string) {
	rollNum, _ := random.IntRange(0, len(raceData))
	raceName = raceData[rollNum].RaceName

	return raceData[rollNum].RaceName, raceData[rollNum].RaceNameRu
}

func setRaceType() (string, string) {
	for _, race := range raceData {
		if race.RaceName == raceName {
			rollNum, _ := random.IntRange(0, len(race.Type))
			return race.Type[rollNum].TypeRaceName, race.Type[rollNum].TypeRaceNameRu
		}
	}
	return "", ""
}

func GetRaceAbilities() StatsUp {
	for _, race := range raceData {
		if race.RaceName == raceName {
			for _, typppe := range race.Type {
				return typppe.StatsUp
			}
		}
	}
	return StatsUp{}
}

func setRaceSkills() []string {
	for _, race := range raceData {
		if race.RaceName == raceName {
			return race.RaceSkill
		}
	}
	return nil
}

func GetRaceSkill() string {
	for _, race := range raceData {
		if race.RaceName == raceName && len(race.RaceSkill) > 0 {
			rollNum, _ := random.IntRange(0, len(race.RaceSkill))
			return race.RaceSkill[rollNum]
		}
	}
	return ""
}

func setAge() int {
	for _, race := range raceData {
		if race.RaceName == raceName {
			age, _ := random.IntRange(race.MinAge, race.MaxAge)
			return age
		}
	}
	return 0
}

func setHeight() int {

	for _, race := range raceData {
		if race.RaceName == raceName {
			height, _ := random.IntRange(race.MinHeight, race.MaxHeight)
			return height
		}
	}
	return 0
}

func setWeight() int {
	for _, race := range raceData {
		if race.RaceName == raceName {
			weight, _ := random.IntRange(race.MinWeight, race.MaxWeight)
			return weight
		}
	}
	return 0
}

func setBodySize() {
	//брать из данных из базы
}

func setSpeed() int {
	for _, race := range raceData {
		if race.RaceName == raceName {
			return race.Speed
		}
	}
	return 0
}

func setGender() string {
	var gender string
	count, _ := random.IntRange(0, 10)
	if count%2 == 0 {
		gender = "Мужской"
	} else {
		gender = "Женский"
	}
	raceGender = gender //присвоение для других методов в пакете
	return gender
}

func setFirstName() string {
	for _, race := range raceData {
		if race.RaceName == raceName && raceGender == "Мужской" && len(race.Names.Man) != 0 {
			rollNum, _ := random.IntRange(0, len(race.Names.Man))
			return race.Names.Man[rollNum].(string)
		}
		if race.RaceName == raceName && raceGender == "Женский" && len(race.Names.Man) != 0 {
			rollNum, _ := random.IntRange(0, len(race.Names.Woman))
			return race.Names.Woman[rollNum].(string)
		}
	}
	return ""
}

func setLastName() string {
	for _, race := range raceData {
		if race.RaceName == raceName && len(race.LastNames) != 0 {
			rollNum, _ := random.IntRange(0, len(race.LastNames))
			return race.LastNames[rollNum].(string)
		}
	}
	return ""
}

func setEyesColor() string {
	var colors = []string{"Желтый", "Синий", "Красный", "Зеленый", "Карий", "Морская волна"}

	color, err := random.Choice(colors)
	if err != nil {
		fmt.Println(err)
	}
	return color
}

func setHairColor() string {
	var colors = []string{"Желтые", "Белые", "Черные", "Синие", "Красные", "Зеленые", "Каштановые", "Бирюзовые", "Рыжие"}

	color, err := random.Choice(colors)
	if err != nil {
		fmt.Println(err)
	}
	return color
}

func setKnowLanguages() []string {
	for _, race := range raceData {
		if race.RaceName == raceName {
			return race.Langs
		}
	}
	return nil
}

func setResist() []string {
	for _, race := range raceData {
		if race.RaceName == raceName {
			return race.Resist
		}
	}
	return nil
}
