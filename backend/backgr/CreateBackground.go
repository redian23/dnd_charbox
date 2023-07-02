package backgr

import (
	"encoding/json"
	"github.com/mazen160/go-random"
	"log"
	"pregen/backend/classes"
)

func GenerateBackground(className string) Background {

	var bg Background
	var backData BackgroundJsonStruct
	//var jsonData = db.SelectJsonFromPgTable("SELECT * FROM backgrounds_json;")

	err := json.Unmarshal(backJsonData, &backData)
	if err != nil {
		log.Println(err)
	}
	backgroundAnalyze(className)

	bg.BackgroundName, bg.BackgroundNameRu = setRandomBackgroundName(backData)
	bg.Type = setBackgroundType(bg.BackgroundName, backData)
	bg.Description = setDescription(bg.BackgroundName, backData)
	bg.SkillMastery = setSkillMastery(bg.BackgroundName, backData)
	bg.Personalization = setPersonalization(bg.BackgroundName, backData)
	bg.CharacterTrait = setCharacterTrait(bg.BackgroundName, backData)
	bg.Ideal = setIdeal(bg.BackgroundName, backData)
	bg.Affection = setAffection(bg.BackgroundName, backData)
	bg.Weakness = setWeakness(bg.BackgroundName, backData)

	return bg
}

func backgroundAnalyze(className string) string {
	var chars classes.CharacteristicsForClass
	json.Unmarshal(classes.RaceCharactsJsonData, &chars)

	for _, char := range chars.Data {
		if className == char.ClassName {
			var rollNumOfType int
			rollNumOfType, _ = random.IntRange(0, len(char.Background)-1)
			return char.Background[rollNumOfType]
		}
	}
	return ""
}
