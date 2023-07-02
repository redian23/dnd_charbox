package backgr

import (
	"encoding/json"
	"fmt"
	"log"
)

func GenerateBackground(className string) Background {

	var bg Background
	var backData BackgroundJsonStruct
	//var jsonData = db.SelectJsonFromPgTable("SELECT * FROM backgrounds_json;")

	err := json.Unmarshal(backJsonData, &backData)
	if err != nil {
		log.Println(err)
	}

	bg.BackgroundName, bg.BackgroundNameRu = setRandomBackgroundName(backData)
	bg.Type = setBackgroundType(bg.BackgroundName, backData)
	bg.Description = setDescription(bg.BackgroundName, backData)
	bg.Personalization = setPersonalization(bg.BackgroundName, backData)
	bg.CharacterTrait = setCharacterTrait(bg.BackgroundName, backData)
	bg.Ideal = setIdeal(bg.BackgroundName, backData)
	bg.Affection = setAffection(bg.BackgroundName, backData)
	bg.Weakness = setWeakness(bg.BackgroundName, backData)

	fmt.Println(bg)
	return bg
}
