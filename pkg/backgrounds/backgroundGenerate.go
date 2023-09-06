package backgrounds

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/classes"
)

var (
	backgroundNameRu       string
	advice                 string
	personalization        string
	description            string
	backgroundType         string
	characterTrait         string
	worldviewRu            string
	ideal                  string
	affection              string
	weakness               string
	gold                   int
	masteryOfTools         []string
	bgAbility              backgroundAbility
	equipmentList          []string
	BackgroundSkillMastery []string
)

func GenerateBackground() *BackgroundAnswer {
	backData := getBackgroundsFormDB()
	backgroundName = backgroundAnalyze(classes.ClassNameGlobal)

	for _, background := range backData {
		if background.BackgroundName == backgroundName {

			backgroundNameRu = background.BackgroundNameRu
			personalization = background.Personalization
			advice = background.Advice
			if background.Advice == "" {
				background.Advice = "Смотрите в книге игрока или спросите у мастера."
				advice = background.Advice
			}
			description = background.Description

			if background.Type.Value != nil {
				var rollNumOfType int
				rollNumOfType, _ = random.IntRange(0, len(background.Type.Value))
				backgroundType = background.Type.Value[rollNumOfType]
			} else {
				backgroundType = "Стандартный"
			}

			rollNum, _ := random.IntRange(0, len(background.CharacterTrait.Value))
			characterTrait = background.CharacterTrait.Value[rollNum]

			rollNum, _ = random.IntRange(0, len(background.Ideal.Value))
			ideal = background.Ideal.Value[rollNum].Text
			worldviewRu = background.Ideal.Value[rollNum].WorldviewRu

			rollNum, _ = random.IntRange(0, len(background.Affection.Value))
			affection = background.Affection.Value[rollNum]

			rollNum, _ = random.IntRange(0, len(background.Weakness.Value))
			weakness = background.Weakness.Value[rollNum]

			BackgroundSkillMastery = background.SkillMastery
			bgAbility = background.BackgroundAbility
			masteryOfTools = background.MasteryOfTools
			equipmentList = background.Equipment
			gold = background.Gold
		}
	}

	return &BackgroundAnswer{
		BackgroundName:    backgroundName,
		BackgroundNameRu:  backgroundNameRu,
		Type:              backgroundType,
		Description:       description,
		Advice:            advice,
		Personalization:   personalization,
		SkillMastery:      BackgroundSkillMastery,
		CharacterTrait:    characterTrait,
		Ideal:             ideal,
		Worldview:         worldviewRu,
		Affection:         affection,
		Weakness:          weakness,
		BackgroundAbility: bgAbility,
		MasteryOfTools:    masteryOfTools,
		Equipment:         equipmentList,
		Gold:              gold,
	}
}
