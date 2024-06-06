package skills

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

var sk Skills

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

type Passive struct {
	PassiveWisdomInsight          int `json:"passive_wisdom_insight"`
	PassiveWisdomPerception       int `json:"passive_wisdom_perception"`
	PassiveIntellectInvestigation int `json:"passive_intellect_investigation"`
}

func SetSkillsForCharacter(raceInfo *races.Race, backgInfo *backgrounds.Background, classInfo *classes.Class, lvl int) *Skills {
	sk = Skills{}
	var skillsArray []string

	// принимает нужный экземпляр обьекта и в себя же записывает новые данные
	setSkillsNames()
	setSkillModifierValue(classInfo)

	skillsArray = append(skillsArray, raceInfo.RaceSkill...)
	skillsArray = append(skillsArray, backgInfo.BackgroundSkills...)
	skillsArray = append(skillsArray, classInfo.Proficiencies.SkillsOfClass...)

	setSkillProficiency(skillsArray, classInfo.ProficiencyBonus)

	var doubleSkillProf []string
	var iter int
	if classInfo.ClassNameRU == "Плут" ||
		classInfo.ClassNameRU == "Бард" && lvl >= 3 {
		for _, skl := range skillsArray {
			iter++
			doubleSkillProf = append(doubleSkillProf, skl)
			if iter >= 2 {
				break
			}
		}
	}

	//if FighterArchetypeName == "Самурай" {
	//	switch {
	//	case sk.Persuasion.Proficiency == false:
	//		sk.Persuasion.Proficiency = true
	//		sk.Persuasion.ProfLSS = 1
	//		sk.Persuasion.ModifierValue = sk.Persuasion.ModifierValue + profBonus
	//		break
	//	case sk.Perception.Proficiency == false:
	//		sk.Perception.Proficiency = true
	//		sk.Perception.ProfLSS = 1
	//		sk.Perception.ModifierValue = sk.Perception.ModifierValue + profBonus
	//		break
	//	case sk.History.Proficiency == false:
	//		sk.History.Proficiency = true
	//		sk.History.ProfLSS = 1
	//		sk.History.ModifierValue = sk.History.ModifierValue + profBonus
	//		break
	//	case sk.Performance.Proficiency == false:
	//		sk.Performance.Proficiency = true
	//		sk.Performance.ProfLSS = 1
	//		sk.Performance.ModifierValue = sk.Performance.ModifierValue + profBonus
	//		break
	//	}
	//}

	sk = setDoubleSkillProficiency(doubleSkillProf, sk, classInfo.ProficiencyBonus)

	return &sk
}

func GetPassive() *Passive {
	return &Passive{
		10 + sk.Insight.ModifierValue,
		10 + sk.Perception.ModifierValue,
		10 + sk.Investigation.ModifierValue,
	}
}

func setSkillsNames() {
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

}

func setSkillModifierValue(classInfo *classes.Class) {
	var modifier = classInfo.AbilityModifier

	sk.Athletics.ModifierValue = modifier.Strength
	sk.Acrobatics.ModifierValue = modifier.Dexterity
	sk.SleightOfHand.ModifierValue = modifier.Dexterity
	sk.Stealth.ModifierValue = modifier.Dexterity
	sk.Arcana.ModifierValue = modifier.Intelligence
	sk.History.ModifierValue = modifier.Intelligence
	sk.Investigation.ModifierValue = modifier.Intelligence
	sk.Nature.ModifierValue = modifier.Intelligence
	sk.Religion.ModifierValue = modifier.Intelligence
	sk.AnimalHandling.ModifierValue = modifier.Wisdom
	sk.Insight.ModifierValue = modifier.Wisdom
	sk.Medicine.ModifierValue = modifier.Wisdom
	sk.Perception.ModifierValue = modifier.Wisdom
	sk.Survival.ModifierValue = modifier.Wisdom
	sk.Deception.ModifierValue = modifier.Charisma
	sk.Intimidation.ModifierValue = modifier.Charisma
	sk.Performance.ModifierValue = modifier.Charisma
	sk.Persuasion.ModifierValue = modifier.Charisma
}

func setSkillProficiency(skillsArray []string, profBonus int) {
	for _, profSkill := range skillsArray {
		switch profSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.Proficiency = true
			sk.Athletics.ProfLSS = 1
			sk.Athletics.ModifierValue = sk.Athletics.ModifierValue + profBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.Proficiency = true
			sk.Acrobatics.ProfLSS = 1
			sk.Acrobatics.ModifierValue = sk.Acrobatics.ModifierValue + profBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.Proficiency = true
			sk.SleightOfHand.ProfLSS = 1
			sk.SleightOfHand.ModifierValue = sk.SleightOfHand.ModifierValue + profBonus
		case sk.Stealth.SkillName:
			sk.Stealth.Proficiency = true
			sk.Stealth.ProfLSS = 1
			sk.Stealth.ModifierValue = sk.Stealth.ModifierValue + profBonus
		case sk.Arcana.SkillName:
			sk.Arcana.Proficiency = true
			sk.Arcana.ProfLSS = 1
			sk.Arcana.ModifierValue = sk.Arcana.ModifierValue + profBonus
		case sk.History.SkillName:
			sk.History.Proficiency = true
			sk.History.ProfLSS = 1
			sk.History.ModifierValue = sk.History.ModifierValue + profBonus
		case sk.Investigation.SkillName:
			sk.Investigation.Proficiency = true
			sk.Investigation.ProfLSS = 1
			sk.Investigation.ModifierValue = sk.Investigation.ModifierValue + profBonus
		case sk.Nature.SkillName:
			sk.Nature.Proficiency = true
			sk.Nature.ProfLSS = 1
			sk.Nature.ModifierValue = sk.Nature.ModifierValue + profBonus
		case sk.Religion.SkillName:
			sk.Religion.Proficiency = true
			sk.Religion.ProfLSS = 1
			sk.Religion.ModifierValue = sk.Religion.ModifierValue + profBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.Proficiency = true
			sk.AnimalHandling.ProfLSS = 1
			sk.AnimalHandling.ModifierValue = sk.AnimalHandling.ModifierValue + profBonus
		case sk.Insight.SkillName:
			sk.Insight.Proficiency = true
			sk.Insight.ProfLSS = 1
			sk.Insight.ModifierValue = sk.Insight.ModifierValue + profBonus
		case sk.Medicine.SkillName:
			sk.Medicine.Proficiency = true
			sk.Medicine.ProfLSS = 1
			sk.Medicine.ModifierValue = sk.Medicine.ModifierValue + profBonus
		case sk.Perception.SkillName:
			sk.Perception.Proficiency = true
			sk.Perception.ProfLSS = 1
			sk.Perception.ModifierValue = sk.Perception.ModifierValue + profBonus
		case sk.Survival.SkillName:
			sk.Survival.Proficiency = true
			sk.Survival.ProfLSS = 1
			sk.Survival.ModifierValue = sk.Survival.ModifierValue + profBonus
		case sk.Deception.SkillName:
			sk.Deception.Proficiency = true
			sk.Deception.ProfLSS = 1
			sk.Deception.ModifierValue = sk.Deception.ModifierValue + profBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.Proficiency = true
			sk.Intimidation.ProfLSS = 1
			sk.Intimidation.ModifierValue = sk.Intimidation.ModifierValue + profBonus
		case sk.Performance.SkillName:
			sk.Performance.Proficiency = true
			sk.Performance.ProfLSS = 1
			sk.Performance.ModifierValue = sk.Performance.ModifierValue + profBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.Proficiency = true
			sk.Persuasion.ProfLSS = 1
			sk.Persuasion.ModifierValue = sk.Persuasion.ModifierValue + profBonus
		}
	}
}

func setDoubleSkillProficiency(dblMap []string, sk Skills, profBonus int) Skills {
	for _, dblProfSkill := range dblMap {
		switch dblProfSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.DoubleProficiency = true
			sk.Athletics.ProfLSS = 2
			sk.Athletics.ModifierValue += profBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.DoubleProficiency = true
			sk.Acrobatics.ProfLSS = 2
			sk.Acrobatics.ModifierValue += profBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.DoubleProficiency = true
			sk.SleightOfHand.ProfLSS = 2
			sk.SleightOfHand.ModifierValue += profBonus
		case sk.Stealth.SkillName:
			sk.Stealth.DoubleProficiency = true
			sk.Stealth.ProfLSS = 2
			sk.Stealth.ModifierValue += profBonus
		case sk.Arcana.SkillName:
			sk.Arcana.DoubleProficiency = true
			sk.Arcana.ProfLSS = 2
			sk.Arcana.ModifierValue += profBonus
		case sk.History.SkillName:
			sk.History.DoubleProficiency = true
			sk.History.ProfLSS = 2
			sk.History.ModifierValue += profBonus
		case sk.Investigation.SkillName:
			sk.Investigation.DoubleProficiency = true
			sk.Investigation.ProfLSS = 2
			sk.Investigation.ModifierValue += profBonus
		case sk.Nature.SkillName:
			sk.Nature.DoubleProficiency = true
			sk.Nature.ProfLSS = 2
			sk.Nature.ModifierValue += profBonus
		case sk.Religion.SkillName:
			sk.Religion.DoubleProficiency = true
			sk.Religion.ProfLSS = 2
			sk.Religion.ModifierValue += profBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.DoubleProficiency = true
			sk.AnimalHandling.ProfLSS = 2
			sk.AnimalHandling.ModifierValue += profBonus
		case sk.Insight.SkillName:
			sk.Insight.DoubleProficiency = true
			sk.Insight.ProfLSS = 2
			sk.Insight.ModifierValue += profBonus
		case sk.Medicine.SkillName:
			sk.Medicine.DoubleProficiency = true
			sk.Medicine.ProfLSS = 2
			sk.Medicine.ModifierValue += profBonus
		case sk.Perception.SkillName:
			sk.Perception.DoubleProficiency = true
			sk.Perception.ProfLSS = 2
			sk.Perception.ModifierValue += profBonus
		case sk.Survival.SkillName:
			sk.Survival.DoubleProficiency = true
			sk.Survival.ProfLSS = 2
			sk.Survival.ModifierValue += profBonus
		case sk.Deception.SkillName:
			sk.Deception.DoubleProficiency = true
			sk.Deception.ProfLSS = 2
			sk.Deception.ModifierValue += profBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.DoubleProficiency = true
			sk.Intimidation.ProfLSS = 2
			sk.Intimidation.ModifierValue += profBonus
		case sk.Performance.SkillName:
			sk.Performance.DoubleProficiency = true
			sk.Performance.ProfLSS = 2
			sk.Performance.ModifierValue += profBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.DoubleProficiency = true
			sk.Persuasion.ProfLSS = 2
			sk.Persuasion.ModifierValue += profBonus
		}
	}
	return sk
}
