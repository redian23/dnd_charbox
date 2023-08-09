package races

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mazen160/go-random"
	"io/ioutil"
	"log"
	"pregen/db"
)

var (
	raceData      = getRacesFormDB()
	raceName      string
	raceGender    string
	RacePhotoPath string
)

func readDirectory(path string) ([]string, []string) {
	var folderList []string
	var fileListInFolder []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() == true {
			folderList = append(folderList, file.Name())
		} else {
			fileListInFolder = append(fileListInFolder, file.Name())
		}
	}
	return folderList, fileListInFolder
}

func setRacePhoto(raceName, gender string) racePhoto {
	var photo racePhoto
	var allRacesFolderList, _ = readDirectory(RacePhotoPath)

	for _, folderName := range allRacesFolderList {
		if folderName == raceName {
			if gender == "Мужской" {
				var manPhotoPath = folderName + "/m/"

				_, racePhotoList := readDirectory(RacePhotoPath + manPhotoPath)
				randNum, _ := random.IntRange(0, len(racePhotoList))

				photo.Path = manPhotoPath
				photo.FileName = racePhotoList[randNum]
				return photo
			} else {
				var womanPhotoPath = folderName + "/w/"

				_, racePhotoList := readDirectory(RacePhotoPath + womanPhotoPath)
				randNum, _ := random.IntRange(0, len(racePhotoList))

				photo.Path = womanPhotoPath
				photo.FileName = racePhotoList[randNum]
				return photo
			}
		}
	}
	return racePhoto{}
}

func InsertRacesToDB() {
	client := db.ConnectToDB()
	coll := client.Database("data").Collection("races")

	var race RacesJsonStruct
	json.Unmarshal([]byte{}, &race) // пока что ничего не записывает

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
			// Вряд ли кто-то будет играть за глубокого старика
			// поэтому минус 25% от макс возраста
			maxAge := (race.MaxAge * 75) / 100
			age, _ := random.IntRange(race.MinAge, maxAge)
			return age
		}
	}
	return 0
}

func setHeight() (int, string) {
	for _, race := range raceData {
		if race.RaceName == raceName {
			height, _ := random.IntRange(race.MinHeight, race.MaxHeight)
			heightFt := fmt.Sprintf("%.1f", float32(height)*0.0328)
			return height, heightFt
		}
	}
	return 0, ""
}

func setWeight() (int, int) {
	for _, race := range raceData {
		if race.RaceName == raceName {
			weight, _ := random.IntRange(race.MinWeight, race.MaxWeight)
			weightLb := (weight * 220) / 100
			return weight, weightLb
		}
	}
	return 0, 0
}

func setBodySize() string {
	for _, race := range raceData {
		if race.RaceName == raceName {
			return race.BodySize
		}
	}
	return ""
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
	count, _ := random.IntRange(1, 10)
	if count%2 == 0 {
		gender = "Мужской"
	} else {
		gender = "Женский"
	}
	raceGender = gender // присвоение для других методов в пакете
	return gender
}

func setFirstName() string {
	for _, race := range raceData {
		if race.RaceName == raceName && raceGender == "Мужской" && len(race.Names.Man) != 0 {
			rollNum, _ := random.IntRange(0, len(race.Names.Man))
			return race.Names.Man[rollNum]
		}
		if race.RaceName == raceName && raceGender == "Женский" && len(race.Names.Man) != 0 {
			rollNum, _ := random.IntRange(0, len(race.Names.Woman))
			return race.Names.Woman[rollNum]
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
	var colors = []string{"Желтые", "Белые", "Черные", "Синие", "Красные",
		"Зеленые", "Каштановые", "Бирюзовые", "Рыжие"}

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

func setResist() string {
	for _, race := range raceData {
		if race.RaceName == raceName {
			return race.Resist
		}
	}
	return ""
}

func setDarkvision() bool {
	for _, race := range raceData {
		if race.RaceName == raceName {
			return race.Darkvision
		}
	}
	return false
}

func setRaceAbilities(raceTypename string) []raceAbility {
	for _, race := range raceData {
		if race.RaceName == raceName {
			for _, rType := range race.Type {
				if raceTypename == rType.TypeRaceNameRu {
					return rType.RaceAbilities
				}
			}
		}
	}
	return nil
}
