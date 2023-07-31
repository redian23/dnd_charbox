package backgr

import (
	"context"
	"github.com/mazen160/go-random"
	"pregen/backend/classes"
	"pregen/db"
)

var backgroundName string
var backData = getBackgroundsFormDB()

func getBackgroundsFormDB() []BackgroundBson {
	var results []BackgroundBson
	var cursor = db.ReadFromDB("backgrounds")

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func backgroundAnalyze(className string) string {
	var chars = classes.GetClassCharactsFormDB()

	for _, char := range chars {
		if className == char.ClassName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(char.Background))

			backgroundName = char.Background[rollNum] //get Background name for all methods in package
			return char.Background[rollNum]
		}
	}
	return ""
}
func setBackgroundNameRU() string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.BackgroundNameRu
		}
	}
	return ""
}

func setBackgroundType() string {
	var bgText string

	for _, background := range backData {
		if background.BackgroundName == backgroundName && background.Type.Value != nil {
			var rollNumOfType int
			rollNumOfType, _ = random.IntRange(0, len(background.Type.Value))
			bgText = background.Type.Value[rollNumOfType]
		}
	}
	if bgText == "" { // В свиче дефаутл не работает, если к верхнему ифу кинуть елсе то он затирает переменную. Да так надо.
		bgText = "Стандартный"
	}
	return bgText
}

func setAdvice() string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.Advice
		} else {
			background.Advice = "Смотрите в книге игрока или спросите у мастера."
			return background.Advice
		}
	}
	return ""
}

func setPersonalization() string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.Personalization
		}
	}
	return ""
}

func setDescription() string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.Description
		}
	}
	return ""
}

func setCharacterTrait() string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(background.CharacterTrait.Value))
			return background.CharacterTrait.Value[rollNum]
		}
	}
	return ""
}

func setIdeal() (string, string) {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(background.Ideal.Value))
			return background.Ideal.Value[rollNum].Text, background.Ideal.Value[rollNum].WorldviewRu
		}
	}
	return "", ""
}

func setAffection() string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(background.Affection.Value))
			return background.Affection.Value[rollNum]
		}
	}
	return ""
}
func setWeakness() string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(background.Weakness.Value))
			return background.Weakness.Value[rollNum]
		}
	}
	return ""
}

func setSkillMastery() []string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.SkillMastery
		}
	}
	return nil
}

func setBackgroundAbility() backgroundAbility {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.BackgroundAbility
		}
	}
	return backgroundAbility{}
}

func setMasteryOfTools() []string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.MasteryOfTools
		}
	}
	return nil
}

func setEquipment() []string {
	for _, background := range backData {
		if background.BackgroundName == backgroundName {
			return background.Equipment
		}
	}
	return nil
}
