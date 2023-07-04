package backgr

import (
	"encoding/json"
	"log"
)

func GenerateBackground(className string) Background {

	var bg Background
	var backData BackgroundJsonStruct

	err := json.Unmarshal(BackJsonData, &backData)
	if err != nil {
		log.Println(err)
	}

	bg.BackgroundName = backgroundAnalyze(className)
	bg.BackgroundNameRu = setBackgroundNameRU(bg.BackgroundName, backData)
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
