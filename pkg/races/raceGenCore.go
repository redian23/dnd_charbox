package races

import (
	"encoding/json"
	"fmt"
	"github.com/mazen160/go-random"
	"log"
	"os"
)

var (
	RaceData         []RacesBSON
	RaceNameGlobalRu string
	RaceSpeedGlobal  int

	RaceTypeGlobalRu string
	RacePhotoPath    string
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
	fmt.Println(allRacesFolderList)
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

func getRacesFormDB() []RacesBSON {
	var results []RacesBSON
	fileContent, err := os.Open("/etc/charbox.d/db/races.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile("/etc/charbox.d/db/races.json")

	json.Unmarshal(byteResult, &results)
	return results
}

func GetRaceAbilities() StatsUp {
	for _, race := range RaceData {
		if race.RaceNameRu == RaceTypeGlobalRu {
			for _, typppe := range race.Type {
				return typppe.StatsUp
			}
		}
	}
	return StatsUp{}
}

func GetRaceSkill() ([]string, int) {
	var skillsArray []string
	var randSkillCount int

	for _, race := range RaceData {
		if race.RaceName == RaceTypeGlobalRu {
			skillsArray = race.RaceSkill.SkillsList
			randSkillCount = race.RaceSkill.RandomCount
		}
	}
	return skillsArray, randSkillCount
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
