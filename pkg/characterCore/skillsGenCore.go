package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
)

var (
	PassWisdom int
)

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
	ProfLSS           int    `json:"prof_lss"`
}

func SetSkillsForCharacter() Skills {
	sk := Skills{}

	// принимает нужный экземпляр обьекта и в себя же записывает новые данные
	sk = setSkillsNames(sk)
	sk = setSkillModifierValue(sk)

	skillList := classes.GetAnalyzedSkillSlice(backgrounds.BackgroundSkills)
	sk = setSkillProficiency(skillList, sk)

	var doubleSkillProf []string
	var iter int
	if classes.ClassNameGlobalRu == "Плут" ||
		classes.ClassNameGlobalRu == "Бард" && classes.CharacterLevelGlobal >= 3 {
		for _, skl := range skillList {
			iter++
			doubleSkillProf = append(doubleSkillProf, skl)
			if iter >= 2 {
				break
			}
		}
	}
	sk = setDoubleSkillProficiency(doubleSkillProf, sk)
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

func setSkillProficiency(skillMap []string, sk Skills) Skills {
	for _, profSkill := range skillMap {
		switch profSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.Proficiency = true
			sk.Athletics.ProfLSS = 1
			sk.Athletics.ModifierValue = sk.Athletics.ModifierValue + classes.ProficiencyBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.Proficiency = true
			sk.Acrobatics.ProfLSS = 1
			sk.Acrobatics.ModifierValue = sk.Acrobatics.ModifierValue + classes.ProficiencyBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.Proficiency = true
			sk.SleightOfHand.ProfLSS = 1
			sk.SleightOfHand.ModifierValue = sk.SleightOfHand.ModifierValue + classes.ProficiencyBonus
		case sk.Stealth.SkillName:
			sk.Stealth.Proficiency = true
			sk.Stealth.ProfLSS = 1
			sk.Stealth.ModifierValue = sk.Stealth.ModifierValue + classes.ProficiencyBonus
		case sk.Arcana.SkillName:
			sk.Arcana.Proficiency = true
			sk.Arcana.ProfLSS = 1
			sk.Arcana.ModifierValue = sk.Arcana.ModifierValue + classes.ProficiencyBonus
		case sk.History.SkillName:
			sk.History.Proficiency = true
			sk.History.ProfLSS = 1
			sk.History.ModifierValue = sk.History.ModifierValue + classes.ProficiencyBonus
		case sk.Investigation.SkillName:
			sk.Investigation.Proficiency = true
			sk.Investigation.ProfLSS = 1
			sk.Investigation.ModifierValue = sk.Investigation.ModifierValue + classes.ProficiencyBonus
		case sk.Nature.SkillName:
			sk.Nature.Proficiency = true
			sk.Nature.ProfLSS = 1
			sk.Nature.ModifierValue = sk.Nature.ModifierValue + classes.ProficiencyBonus
		case sk.Religion.SkillName:
			sk.Religion.Proficiency = true
			sk.Religion.ProfLSS = 1
			sk.Religion.ModifierValue = sk.Religion.ModifierValue + classes.ProficiencyBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.Proficiency = true
			sk.AnimalHandling.ProfLSS = 1
			sk.AnimalHandling.ModifierValue = sk.AnimalHandling.ModifierValue + classes.ProficiencyBonus
		case sk.Insight.SkillName:
			sk.Insight.Proficiency = true
			sk.Insight.ProfLSS = 1
			sk.Insight.ModifierValue = sk.Insight.ModifierValue + classes.ProficiencyBonus
		case sk.Medicine.SkillName:
			sk.Medicine.Proficiency = true
			sk.Medicine.ProfLSS = 1
			sk.Medicine.ModifierValue = sk.Medicine.ModifierValue + classes.ProficiencyBonus
		case sk.Perception.SkillName:
			sk.Perception.Proficiency = true
			sk.Perception.ProfLSS = 1
			sk.Perception.ModifierValue = sk.Perception.ModifierValue + classes.ProficiencyBonus
		case sk.Survival.SkillName:
			sk.Survival.Proficiency = true
			sk.Survival.ProfLSS = 1
			sk.Survival.ModifierValue = sk.Survival.ModifierValue + classes.ProficiencyBonus
		case sk.Deception.SkillName:
			sk.Deception.Proficiency = true
			sk.Deception.ProfLSS = 1
			sk.Deception.ModifierValue = sk.Deception.ModifierValue + classes.ProficiencyBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.Proficiency = true
			sk.Intimidation.ProfLSS = 1
			sk.Intimidation.ModifierValue = sk.Intimidation.ModifierValue + classes.ProficiencyBonus
		case sk.Performance.SkillName:
			sk.Performance.Proficiency = true
			sk.Performance.ProfLSS = 1
			sk.Performance.ModifierValue = sk.Performance.ModifierValue + classes.ProficiencyBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.Proficiency = true
			sk.Persuasion.ProfLSS = 1
			sk.Persuasion.ModifierValue = sk.Persuasion.ModifierValue + classes.ProficiencyBonus
		}
	}

	setPassiveWisdom(sk)
	return sk
}

// Useless & Defected
func setDoubleSkillProficiency(dblMap []string, sk Skills) Skills {
	for _, dblProfSkill := range dblMap {
		switch dblProfSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.DoubleProficiency = true
			sk.Athletics.ProfLSS = 2
			sk.Athletics.ModifierValue += classes.ProficiencyBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.DoubleProficiency = true
			sk.Acrobatics.ProfLSS = 2
			sk.Acrobatics.ModifierValue += classes.ProficiencyBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.DoubleProficiency = true
			sk.SleightOfHand.ProfLSS = 2
			sk.SleightOfHand.ModifierValue += classes.ProficiencyBonus
		case sk.Stealth.SkillName:
			sk.Stealth.DoubleProficiency = true
			sk.Stealth.ProfLSS = 2
			sk.Stealth.ModifierValue += classes.ProficiencyBonus
		case sk.Arcana.SkillName:
			sk.Arcana.DoubleProficiency = true
			sk.Arcana.ProfLSS = 2
			sk.Arcana.ModifierValue += classes.ProficiencyBonus
		case sk.History.SkillName:
			sk.History.DoubleProficiency = true
			sk.History.ProfLSS = 2
			sk.History.ModifierValue += classes.ProficiencyBonus
		case sk.Investigation.SkillName:
			sk.Investigation.DoubleProficiency = true
			sk.Investigation.ProfLSS = 2
			sk.Investigation.ModifierValue += classes.ProficiencyBonus
		case sk.Nature.SkillName:
			sk.Nature.DoubleProficiency = true
			sk.Nature.ProfLSS = 2
			sk.Nature.ModifierValue += classes.ProficiencyBonus
		case sk.Religion.SkillName:
			sk.Religion.DoubleProficiency = true
			sk.Religion.ProfLSS = 2
			sk.Religion.ModifierValue += classes.ProficiencyBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.DoubleProficiency = true
			sk.AnimalHandling.ProfLSS = 2
			sk.AnimalHandling.ModifierValue += classes.ProficiencyBonus
		case sk.Insight.SkillName:
			sk.Insight.DoubleProficiency = true
			sk.Insight.ProfLSS = 2
			sk.Insight.ModifierValue += classes.ProficiencyBonus
		case sk.Medicine.SkillName:
			sk.Medicine.DoubleProficiency = true
			sk.Medicine.ProfLSS = 2
			sk.Medicine.ModifierValue += classes.ProficiencyBonus
		case sk.Perception.SkillName:
			sk.Perception.DoubleProficiency = true
			sk.Perception.ProfLSS = 2
			sk.Perception.ModifierValue += classes.ProficiencyBonus
		case sk.Survival.SkillName:
			sk.Survival.DoubleProficiency = true
			sk.Survival.ProfLSS = 2
			sk.Survival.ModifierValue += classes.ProficiencyBonus
		case sk.Deception.SkillName:
			sk.Deception.DoubleProficiency = true
			sk.Deception.ProfLSS = 2
			sk.Deception.ModifierValue += classes.ProficiencyBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.DoubleProficiency = true
			sk.Intimidation.ProfLSS = 2
			sk.Intimidation.ModifierValue += classes.ProficiencyBonus
		case sk.Performance.SkillName:
			sk.Performance.DoubleProficiency = true
			sk.Performance.ProfLSS = 2
			sk.Performance.ModifierValue += classes.ProficiencyBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.DoubleProficiency = true
			sk.Persuasion.ProfLSS = 2
			sk.Persuasion.ModifierValue += classes.ProficiencyBonus
		}
	}
	return sk
}

func setPassiveWisdom(sk Skills) {
	PassWisdom = 10 + sk.Perception.ModifierValue
}
