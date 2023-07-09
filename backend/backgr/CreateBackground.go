package backgr

import (
	"context"
	"pregen/db"
)

func GenerateBackground(className string) BackgroundAnswer {

	var bg BackgroundAnswer
	var backData = GetBackgroundsFormDB()

	bg.BackgroundName = backgroundAnalyze(className)
	bg.BackgroundNameRu = setBackgroundNameRU(bg.BackgroundName, backData)
	bg.Type = setBackgroundType(bg.BackgroundName, backData)
	bg.Description = setDescription(bg.BackgroundName, backData)
	bg.Advice = setAdvice(bg.BackgroundName, backData)
	bg.SkillMastery = setSkillMastery(bg.BackgroundName, backData)
	bg.Personalization = setPersonalization(bg.BackgroundName, backData)
	bg.CharacterTrait = setCharacterTrait(bg.BackgroundName, backData)
	bg.Ideal, bg.Worldview = setIdeal(bg.BackgroundName, backData)
	bg.Affection = setAffection(bg.BackgroundName, backData)
	bg.Weakness = setWeakness(bg.BackgroundName, backData)
	return bg
}

func GetBackgroundsFormDB() []BackgroundBson {
	var results []BackgroundBson
	var cursor = db.ReadFromDB("backgrounds")

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}
