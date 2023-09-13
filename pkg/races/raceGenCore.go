package races

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mazen160/go-random"
	"log"
	"os"
	"pregen/pkg/db"
)

var (
	raceData       []RacesBSON
	raceNameGlobal string
	RacePhotoPath  string
)

func readDirectory(path string) ([]string, []string) {
	var folderList []string
	var fileListInFolder []string
	files, err := os.ReadDir(path)
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

				photo.Path = "race_imgs/" + manPhotoPath
				photo.FileName = racePhotoList[randNum]
				return photo
			} else {
				var womanPhotoPath = folderName + "/w/"

				_, racePhotoList := readDirectory(RacePhotoPath + womanPhotoPath)
				randNum, _ := random.IntRange(0, len(racePhotoList))

				photo.Path = "race_imgs/" + womanPhotoPath
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

func GetRaceAbilities() StatsUp {
	for _, race := range raceData {
		if race.RaceName == raceNameGlobal {
			for _, typppe := range race.Type {
				return typppe.StatsUp
			}
		}
	}
	return StatsUp{}
}

func GetRaceSkill() []string {
	var skills []string
	var skillsArray []string
	var randSkillCount int
	skillMap := make(map[int]string)

	for _, race := range raceData {
		if race.RaceName == raceNameGlobal {
			skillsArray = race.RaceSkill.SkillsList
			randSkillCount = race.RaceSkill.RandomCount
		}
	}

	for i, skill := range skillsArray {
		skillMap[i] = skill
	}

	var iter int
	for _, s2 := range skillMap {
		skills = append(skills, s2)
		iter++
		if iter >= randSkillCount {
			break
		}
	}
	return skills
}

func setEyesColor() string {
	var colors = []string{"Желтый", "Синий", "Красный", "Зеленый", "Карий", "Морская волна"}

	color, err := random.Choice(colors)
	if err != nil {
		log.Println(err)
	}
	return color
}

func setHairColor() string {
	var colors = []string{"Желтые", "Белые", "Черные", "Синие", "Красные",
		"Зеленые", "Каштановые", "Бирюзовые", "Рыжие"}

	color, err := random.Choice(colors)
	if err != nil {
		log.Println(err)
	}
	return color
}
