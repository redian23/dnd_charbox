package backgr

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/classes"
)

var BackgroundSkillMastery []string

func GenerateBackground() *BackgroundAnswer {

	backgroundName = backgroundAnalyze(classes.ClassNameGlobal)
	var backgroundNameRu string
	var advice string
	var personalization string
	var description string
	var backgroundType string
	var characterTrait string
	var worldviewRu string
	var ideal string

	var affection string
	var weakness string

	var masteryOfTools []string

	var bgAbility backgroundAbility
	var equipmentList []string

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
		}

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
	}
}
