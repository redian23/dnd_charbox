package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
)

var PassWisdom int

type Skills struct {
	Acrobatics     Skill `json:"acrobatics"`
	AnimalHandling Skill `json:"animal_handling"`
	Arcana         Skill `json:"arcana"`
	Athletics      Skill `json:"athletics"`
	Deception      Skill `json:"deception"`
	History        Skill `json:"history"`
	Insight        Skill `json:"insight"`
	Intimidation   Skill `json:"intimidation"`
	Investigation  Skill `json:"investigation"`
	Medicine       Skill `json:"medicine"`
	Nature         Skill `json:"nature"`
	Perception     Skill `json:"perception"`
	Performance    Skill `json:"performance"`
	Persuasion     Skill `json:"persuasion"`
	Religion       Skill `json:"religion"`
	SleightOfHand  Skill `json:"sleight_of_hand"`
	Stealth        Skill `json:"stealth"`
	Survival       Skill `json:"survival"`
}

type Skill struct {
	SkillName         string `json:"skill_name"`
	SkillNameRu       string `json:"skill_name_ru"`
	ModifierValue     int    `json:"modifier_value"`
	Proficiency       bool   `json:"proficiency"`
	DoubleProficiency bool   `json:"double_proficiency"`
}

func SetSkillsForCharacter() Skills {
	sk := Skills{}

	// принимает нужный экземпляр обьекта и в себя же записывает новые данные
	sk = setSkillsNames(sk)
	sk = setSkillModifierValue(sk)

	skillMap, dblProfMap := classes.GetAnalyzedSkillSlice(backgrounds.BackgroundSkillMastery)
	sk = setSkillProficiency(skillMap, sk)

	sk = setDoubleSkillProficiency(dblProfMap, sk)
	return sk
}

func setSkillsNames(sk Skills) Skills {
	sk.Athletics.SkillName = "Athletics"
	sk.Acrobatics.SkillName = "Acrobatics"
	sk.SleightOfHand.SkillName = "Sleight Of Hand"
	sk.Stealth.SkillName = "Stealth"
	sk.Arcana.SkillName = "Arcana"
	sk.History.SkillName = "History"
	sk.Investigation.SkillName = "Investigation"
	sk.Nature.SkillName = "Nature"
	sk.Religion.SkillName = "Religion"
	sk.AnimalHandling.SkillName = "Animal Handling"
	sk.Insight.SkillName = "Insight"
	sk.Medicine.SkillName = "Medicine"
	sk.Perception.SkillName = "Perception"
	sk.Survival.SkillName = "Survival"
	sk.Deception.SkillName = "Deception"
	sk.Intimidation.SkillName = "Intimidation"
	sk.Performance.SkillName = "Performance"
	sk.Persuasion.SkillName = "Persuasion"

	sk.Athletics.SkillNameRu = "Атлетика"
	sk.Acrobatics.SkillNameRu = "Акробатика"
	sk.SleightOfHand.SkillNameRu = "Ловкость рук"
	sk.Stealth.SkillNameRu = "Скрытность"
	sk.Arcana.SkillNameRu = "Магия"
	sk.History.SkillNameRu = "История"
	sk.Investigation.SkillNameRu = "Анализ"
	sk.Nature.SkillNameRu = "Природа"
	sk.Religion.SkillNameRu = "Религия"
	sk.AnimalHandling.SkillNameRu = "Уход за животными"
	sk.Insight.SkillNameRu = "Восприятие"
	sk.Medicine.SkillNameRu = "Медицина"
	sk.Perception.SkillNameRu = "Проницательность"
	sk.Survival.SkillNameRu = "Выживание"
	sk.Deception.SkillNameRu = "Обман"
	sk.Intimidation.SkillNameRu = "Запугивание"
	sk.Performance.SkillNameRu = "Выступление"
	sk.Persuasion.SkillNameRu = "Убеждение"

	return sk
}
func setSkillModifierValue(sk Skills) Skills {

	sk.Athletics.ModifierValue = classes.ModifierGlobal.Strength
	sk.Acrobatics.ModifierValue = classes.ModifierGlobal.Dexterity
	sk.SleightOfHand.ModifierValue = classes.ModifierGlobal.Dexterity
	sk.Stealth.ModifierValue = classes.ModifierGlobal.Dexterity
	sk.Arcana.ModifierValue = classes.ModifierGlobal.Intelligence
	sk.History.ModifierValue = classes.ModifierGlobal.Intelligence
	sk.Investigation.ModifierValue = classes.ModifierGlobal.Intelligence
	sk.Nature.ModifierValue = classes.ModifierGlobal.Intelligence
	sk.Religion.ModifierValue = classes.ModifierGlobal.Intelligence
	sk.AnimalHandling.ModifierValue = classes.ModifierGlobal.Wisdom
	sk.Insight.ModifierValue = classes.ModifierGlobal.Wisdom
	sk.Medicine.ModifierValue = classes.ModifierGlobal.Wisdom
	sk.Perception.ModifierValue = classes.ModifierGlobal.Wisdom
	sk.Survival.ModifierValue = classes.ModifierGlobal.Wisdom
	sk.Deception.ModifierValue = classes.ModifierGlobal.Charisma
	sk.Intimidation.ModifierValue = classes.ModifierGlobal.Charisma
	sk.Performance.ModifierValue = classes.ModifierGlobal.Charisma
	sk.Persuasion.ModifierValue = classes.ModifierGlobal.Charisma

	return sk
}

func setSkillProficiency(skillMap map[int]string, sk Skills) Skills {
	for _, profSkill := range skillMap {
		switch profSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.Proficiency = true
			sk.Athletics.ModifierValue = sk.Athletics.ModifierValue + 2
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.Proficiency = true
			sk.Acrobatics.ModifierValue = sk.Acrobatics.ModifierValue + 2
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.Proficiency = true
			sk.SleightOfHand.ModifierValue = sk.SleightOfHand.ModifierValue + 2
		case sk.Stealth.SkillName:
			sk.Stealth.Proficiency = true
			sk.Stealth.ModifierValue = sk.Stealth.ModifierValue + 2
		case sk.Arcana.SkillName:
			sk.Arcana.Proficiency = true
			sk.Arcana.ModifierValue = sk.Arcana.ModifierValue + 2
		case sk.History.SkillName:
			sk.History.Proficiency = true
			sk.History.ModifierValue = sk.History.ModifierValue + 2
		case sk.Investigation.SkillName:
			sk.Investigation.Proficiency = true
			sk.Investigation.ModifierValue = sk.Investigation.ModifierValue + 2
		case sk.Nature.SkillName:
			sk.Nature.Proficiency = true
			sk.Nature.ModifierValue = sk.Nature.ModifierValue + 2
		case sk.Religion.SkillName:
			sk.Religion.Proficiency = true
			sk.Religion.ModifierValue = sk.Religion.ModifierValue + 2
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.Proficiency = true
			sk.AnimalHandling.ModifierValue = sk.AnimalHandling.ModifierValue + 2
		case sk.Insight.SkillName:
			sk.Insight.Proficiency = true
			sk.Insight.ModifierValue = sk.Insight.ModifierValue + 2
		case sk.Medicine.SkillName:
			sk.Medicine.Proficiency = true
			sk.Medicine.ModifierValue = sk.Medicine.ModifierValue + 2
		case sk.Perception.SkillName:
			sk.Perception.Proficiency = true
			sk.Perception.ModifierValue = sk.Perception.ModifierValue + 2
		case sk.Survival.SkillName:
			sk.Survival.Proficiency = true
			sk.Survival.ModifierValue = sk.Survival.ModifierValue + 2
		case sk.Deception.SkillName:
			sk.Deception.Proficiency = true
			sk.Deception.ModifierValue = sk.Deception.ModifierValue + 2
		case sk.Intimidation.SkillName:
			sk.Intimidation.Proficiency = true
			sk.Intimidation.ModifierValue = sk.Intimidation.ModifierValue + 2
		case sk.Performance.SkillName:
			sk.Performance.Proficiency = true
			sk.Performance.ModifierValue = sk.Performance.ModifierValue + 2
		case sk.Persuasion.SkillName:
			sk.Persuasion.Proficiency = true
			sk.Persuasion.ModifierValue = sk.Persuasion.ModifierValue + 2
		}
	}
	setPassiveWisdom(sk)
	return sk
}

func setDoubleSkillProficiency(dblMap map[string]string, sk Skills) Skills {
	for _, dblProfSkill := range dblMap {
		switch dblProfSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.DoubleProficiency = true
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.DoubleProficiency = true
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.DoubleProficiency = true
		case sk.Stealth.SkillName:
			sk.Stealth.DoubleProficiency = true
		case sk.Arcana.SkillName:
			sk.Arcana.DoubleProficiency = true
		case sk.History.SkillName:
			sk.History.DoubleProficiency = true
		case sk.Investigation.SkillName:
			sk.Investigation.DoubleProficiency = true
		case sk.Nature.SkillName:
			sk.Nature.DoubleProficiency = true
		case sk.Religion.SkillName:
			sk.Religion.DoubleProficiency = true
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.DoubleProficiency = true
		case sk.Insight.SkillName:
			sk.Insight.DoubleProficiency = true
		case sk.Medicine.SkillName:
			sk.Medicine.DoubleProficiency = true
		case sk.Perception.SkillName:
			sk.Perception.DoubleProficiency = true
		case sk.Survival.SkillName:
			sk.Survival.DoubleProficiency = true
		case sk.Deception.SkillName:
			sk.Deception.DoubleProficiency = true
		case sk.Intimidation.SkillName:
			sk.Intimidation.DoubleProficiency = true
		case sk.Performance.SkillName:
			sk.Performance.DoubleProficiency = true
		case sk.Persuasion.SkillName:
			sk.Persuasion.DoubleProficiency = true
		}
	}
	return sk
}

func setPassiveWisdom(sk Skills) {
	PassWisdom = 10 + sk.Perception.ModifierValue
}
