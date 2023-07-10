package backgr

import (
	"github.com/mazen160/go-random"
	"pregen/backend/classes"
)

func backgroundAnalyze(className string) string {
	var chars = classes.GetClassCharactsFormDB()

	for _, char := range chars {
		if className == char.ClassName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(char.Background))
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
			rollNumOfType, _ = random.IntRange(0, len(background.Type.Value))
			bgText = background.Type.Value[rollNumOfType]
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
			background.Advice = "Смотрите в книге игрока или спросите у мастера."
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
			rollNum, _ = random.IntRange(0, len(background.CharacterTrait.Value))
			return background.CharacterTrait.Value[rollNum]
		}
	}
	return ""
}

func setIdeal(bgName string, backData []BackgroundBson) (string, string) {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(background.Ideal.Value))
			return background.Ideal.Value[rollNum].Text, background.Ideal.Value[rollNum].WorldviewRu
		}
	}
	return "", ""
}

func setAffection(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(background.Affection.Value))
			return background.Affection.Value[rollNum]
		}
	}
	return ""
}
func setWeakness(bgName string, backData []BackgroundBson) string {
	for _, background := range backData {
		if background.BackgroundName == bgName {
			var rollNum int
			rollNum, _ = random.IntRange(0, len(background.Weakness.Value))
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
