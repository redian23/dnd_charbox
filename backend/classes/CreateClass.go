package classes

import (
	"context"
	"pregen/db"
)

var ProficiencyBonus = 2

func GenerateClass() ClassAnswer {
	//Start To Generate
	//Default Values
	var class ClassAnswer

	class.ClassName = "HomeBrew"
	class.Inspiration = true
	class.ProficiencyBonus = ProficiencyBonus

	class.ClassName, class.ClassNameRU, class.Ability = statAnalyze(class)
	class.Modifier = setModifiersForClass(class.Ability)
	class.SavingThrows = setSaveThrowsForClass(class.ClassName, class.Modifier)
	class.PassiveWisdom = setPassiveWisdom(class.Modifier.Wisdom)
	class.Hits.HitDice = setHitDice(class.ClassName)
	class.Hits.HitCount = setHitCount(class.ClassName, class.Modifier.BodyDifficulty)
	class.SkillsOfClass = setClassSkills(class.ClassName)

	return class
}

func GetClassCharactsFormDB() ClassesBSON {
	var results ClassesBSON
	var cursor = db.ReadFromDB("classes")
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}
