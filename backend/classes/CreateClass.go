package classes

import (
	"context"
	"pregen/db"
)

var ProficiencyBonus = 2

func GenerateClass() Class {
	//Start To Generate
	//Default Values
	var class Class

	class.ClassName = "HomeBrew"
	class.Inspiration = true
	class.ProficiencyBonus = ProficiencyBonus

	class.ClassName, class.ClassNameRU, class.Ability = statAnalyze(class)
	class.Modifier = setModifiersForClass(class.Ability)
	class.SavingThrows = setSaveThrowsForClass(class.Modifier)
	class.PassiveWisdom = setPassiveWisdom(class.Modifier.Wisdom)

	return class
}

func GetRaceCharactsFormDB() RaceCharacteristicsBSON {
	var results RaceCharacteristicsBSON
	var cursor = db.ReadFromDB("race_characteristics")

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}
