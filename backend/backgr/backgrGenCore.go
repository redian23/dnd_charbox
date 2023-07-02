package backgr

import (
	"github.com/mazen160/go-random"
)

func setRandomBackgroundName(backData BackgroundJsonStruct) (string, string) {
	randElem, _ := random.IntRange(0, len(backData.Backgrounds))
	return backData.Backgrounds[randElem].BackgroundName, backData.Backgrounds[randElem].BackgroundNameRu
}

func setBackgroundType(bgName string, backData BackgroundJsonStruct) string {
	var bgText string

	for _, background := range backData.Backgrounds {
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
			}
		}
	}
	if bgText == "" { // В свиче дефаутл не работает, если к верхнему ифу кинуть елсе то он затирает переменную. Да так надо.
		bgText = "Default"
	}
	return bgText
}
func setPersonalization(bgName string, backData BackgroundJsonStruct) string {
	for _, background := range backData.Backgrounds {
		if background.BackgroundName == bgName {
			return background.Personalization
		}
	}
	return ""
}

func setDescription(bgName string, backData BackgroundJsonStruct) string {
	for _, background := range backData.Backgrounds {
		if background.BackgroundName == bgName {
			return background.Description
		}
	}
	return ""
}

func setCharacterTrait(bgName string, backData BackgroundJsonStruct) string {
	for _, background := range backData.Backgrounds {
		if background.BackgroundName == bgName {
			var rollNumOfType int
			rollNumOfType, _ = random.IntRange(0, 7)
			return background.CharacterTrait.Value[rollNumOfType]
		}
	}
	return ""
}

func setIdeal(bgName string, backData BackgroundJsonStruct) string {
	for _, background := range backData.Backgrounds {
		var wvArray = []string{background.Ideal.Any, background.Ideal.Good, background.Ideal.Lawful,
			background.Ideal.Evil, background.Ideal.Neutral, background.Ideal.Chaotic}

		if background.BackgroundName == bgName {
			var rollNumOfType int
			rollNumOfType, _ = random.IntRange(0, 5)
			return wvArray[rollNumOfType]
		}
	}
	return ""
}

func setAffection(bgName string, backData BackgroundJsonStruct) string {
	for _, background := range backData.Backgrounds {
		if background.BackgroundName == bgName {
			var rollNumOfType int
			rollNumOfType, _ = random.IntRange(0, 5)
			return background.Affection.Value[rollNumOfType]
		}
	}
	return ""
}
func setWeakness(bgName string, backData BackgroundJsonStruct) string {
	for _, background := range backData.Backgrounds {
		if background.BackgroundName == bgName {
			var rollNumOfType int
			rollNumOfType, _ = random.IntRange(0, 5)
			return background.Weakness.Value[rollNumOfType]
		}
	}
	return ""
}

func setSkillMastery(bgName string, backData BackgroundJsonStruct) []string {
	for _, background := range backData.Backgrounds {
		if background.BackgroundName == bgName {
			return background.SkillMastery
		}
	}
	return nil
}
