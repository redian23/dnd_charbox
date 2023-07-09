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
	class.SavingThrows = setSaveThrowsForClass(class.Modifier)
	class.PassiveWisdom = setPassiveWisdom(class.Modifier.Wisdom)

	return class
}

func GetRaceCharactsFormDB() ClassesBSON {
	var results ClassesBSON
	var cursor = db.ReadFromDB("classes")
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}
