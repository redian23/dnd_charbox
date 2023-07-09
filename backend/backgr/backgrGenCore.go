package backgr

import (
	"github.com/mazen160/go-random"
	"pregen/backend/classes"
)

func backgroundAnalyze(className string) string {
	var chars = classes.GetRaceCharactsFormDB()

	for _, char := range chars {
		if className == char.ClassName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(char.Background)-1)
			return char.Background[rollNum]
		}
	}
	return ""
}
func setBackgroundNameRU(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			return background.BackgroundNameRu
		}
	}
	return ""
}

func setBackgroundType(bgName string, backData []BackgroundBson) string {
	var bgText string

	for _, background := range backData {
		if background.BackgroundName == bgName && background.Type.Value != nil {
			var rollNumOfType int

			switch background.Type.Dice {
			case "D6":
				rollNumOfType, _ = random.IntRange(0, 5)
				bgText = background.Type.Value[rollNumOfType]
			case "D8":
				rollNumOfType, _ = random.IntRange(0, 7)
				bgText = background.Type.Value[rollNumOfType]
			case "D10":
				rollNumOfType, _ = random.IntRange(0, 9)
				bgText = background.Type.Value[rollNumOfType]
			case "D20":
				rollNumOfType, _ = random.IntRange(0, 19)
				bgText = background.Type.Value[rollNumOfType]
			}
		}
	}
	if bgText == "" { // В свиче дефаутл не работает, если к верхнему ифу кинуть елсе то он затирает переменную. Да так надо.
		bgText = "Стандартный"
	}
	return bgText
}

func setAdvice(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			return background.Advice
		} else {
			background.Advice = "Смотрите в книге игрока."
			return background.Advice
		}
	}
	return ""
}

func setPersonalization(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			return background.Personalization
		}
	}
	return ""
}

func setDescription(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			return background.Description
		}
	}
	return ""
}

func setCharacterTrait(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			var rollNum int
			rollNum, _ = random.IntRange(0, 7)
			return background.CharacterTrait.Value[rollNum]
		}
	}
	return ""
}

func setIdeal(bgName string, backData []BackgroundBson) (string, string) {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			var rollNum int
			rollNum, _ = random.IntRange(0, 5)
			return background.Ideal.Value[rollNum].Text, background.Ideal.Value[rollNum].WorldviewRu
		}
	}
	return "", ""
}

func setAffection(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			var rollNum int
			rollNum, _ = random.IntRange(0, 5)
			return background.Affection.Value[rollNum]
		}
	}
	return ""
}
func setWeakness(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			var rollNum int
			rollNum, _ = random.IntRange(0, 5)
			return background.Weakness.Value[rollNum]
		}
	}
	return ""
}

func setSkillMastery(bgName string, backData []BackgroundBson) []string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			return background.SkillMastery
		}
	}
	return nil
}
