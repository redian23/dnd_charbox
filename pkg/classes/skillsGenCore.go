package classes

import "fmt"

var (
	PassWisdom       int
	sk               Skills
	SkillsListGlobal Skills
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

func SetSkillsForCharacter(bk []string) Skills {
	sk = Skills{}

	// принимает нужный экземпляр обьекта и в себя же записывает новые данные
	sk = setSkillsNames(sk)
	sk = setSkillModifierValue(sk)

	skillList := GetAnalyzedSkillSlice(bk)
	sk = setSkillProficiency(skillList, sk)

	var doubleSkillProf []string
	var iter int
	if ClassNameGlobalRu == "Плут" ||
		ClassNameGlobalRu == "Бард" && CharacterLevelGlobal >= 3 {
		for _, skl := range skillList {
			iter++
			doubleSkillProf = append(doubleSkillProf, skl)
			if iter >= 2 {
				break
			}
		}
	}

	sk = setDoubleSkillProficiency(doubleSkillProf, sk)
	SkillsListGlobal = sk
	return SkillsListGlobal
}

func getSkillsOfCharacter() Skills {
	return SkillsListGlobal
}

func getArraySkillsOfCharacter() map[string]bool {
	var skillArray = map[string]bool{
		SkillsListGlobal.Athletics.SkillName:      SkillsListGlobal.Athletics.Proficiency,
		SkillsListGlobal.Acrobatics.SkillName:     SkillsListGlobal.Acrobatics.Proficiency,
		SkillsListGlobal.SleightOfHand.SkillName:  SkillsListGlobal.SleightOfHand.Proficiency,
		SkillsListGlobal.Stealth.SkillName:        SkillsListGlobal.Stealth.Proficiency,
		SkillsListGlobal.Arcana.SkillName:         SkillsListGlobal.Arcana.Proficiency,
		SkillsListGlobal.History.SkillName:        SkillsListGlobal.History.Proficiency,
		SkillsListGlobal.Investigation.SkillName:  SkillsListGlobal.Investigation.Proficiency,
		SkillsListGlobal.Nature.SkillName:         SkillsListGlobal.Nature.Proficiency,
		SkillsListGlobal.Religion.SkillName:       SkillsListGlobal.Religion.Proficiency,
		SkillsListGlobal.AnimalHandling.SkillName: SkillsListGlobal.AnimalHandling.Proficiency,
		SkillsListGlobal.Insight.SkillName:        SkillsListGlobal.Insight.Proficiency,
		SkillsListGlobal.Medicine.SkillName:       SkillsListGlobal.Medicine.Proficiency,
		SkillsListGlobal.Perception.SkillName:     SkillsListGlobal.Perception.Proficiency,
		SkillsListGlobal.Survival.SkillName:       SkillsListGlobal.Survival.Proficiency,
		SkillsListGlobal.Deception.SkillName:      SkillsListGlobal.Deception.Proficiency,
		SkillsListGlobal.Intimidation.SkillName:   SkillsListGlobal.Intimidation.Proficiency,
		SkillsListGlobal.Performance.SkillName:    SkillsListGlobal.Performance.Proficiency,
		SkillsListGlobal.Persuasion.SkillName:     SkillsListGlobal.Persuasion.Proficiency,
	}
	return skillArray
}

func addMoreSkillsToCharacter(addSkillCount int) {
	skillMap := getArraySkillsOfCharacter()

	var iter int
	var freeSkills []string

	for skName, profStatus := range skillMap {
		if profStatus == false {
			freeSkills = append(freeSkills, skName)
			iter++
		}
		if iter == addSkillCount {
			break
		}
	}
	fmt.Println(freeSkills)
	for _, skName := range freeSkills {
		switch skName {
		case SkillsListGlobal.Athletics.SkillName:
			SkillsListGlobal.Athletics.Proficiency = true
			SkillsListGlobal.Athletics.ProfLSS = 1
			SkillsListGlobal.Athletics.ModifierValue = SkillsListGlobal.Athletics.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Acrobatics.SkillName:
			SkillsListGlobal.Acrobatics.Proficiency = true
			SkillsListGlobal.Acrobatics.ProfLSS = 1
			SkillsListGlobal.Acrobatics.ModifierValue = SkillsListGlobal.Acrobatics.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.SleightOfHand.SkillName:
			SkillsListGlobal.SleightOfHand.Proficiency = true
			SkillsListGlobal.SleightOfHand.ProfLSS = 1
			SkillsListGlobal.SleightOfHand.ModifierValue = SkillsListGlobal.SleightOfHand.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Stealth.SkillName:
			SkillsListGlobal.Stealth.Proficiency = true
			SkillsListGlobal.Stealth.ProfLSS = 1
			SkillsListGlobal.Stealth.ModifierValue = SkillsListGlobal.Stealth.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Arcana.SkillName:
			SkillsListGlobal.Arcana.Proficiency = true
			SkillsListGlobal.Arcana.ProfLSS = 1
			SkillsListGlobal.Arcana.ModifierValue = SkillsListGlobal.Arcana.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.History.SkillName:
			SkillsListGlobal.History.Proficiency = true
			SkillsListGlobal.History.ProfLSS = 1
			SkillsListGlobal.History.ModifierValue = SkillsListGlobal.History.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Investigation.SkillName:
			SkillsListGlobal.Investigation.Proficiency = true
			SkillsListGlobal.Investigation.ProfLSS = 1
			SkillsListGlobal.Investigation.ModifierValue = SkillsListGlobal.Investigation.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Nature.SkillName:
			SkillsListGlobal.Nature.Proficiency = true
			SkillsListGlobal.Nature.ProfLSS = 1
			SkillsListGlobal.Nature.ModifierValue = SkillsListGlobal.Nature.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Religion.SkillName:
			SkillsListGlobal.Religion.Proficiency = true
			SkillsListGlobal.Religion.ProfLSS = 1
			SkillsListGlobal.Religion.ModifierValue = SkillsListGlobal.Religion.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.AnimalHandling.SkillName:
			SkillsListGlobal.AnimalHandling.Proficiency = true
			SkillsListGlobal.AnimalHandling.ProfLSS = 1
			SkillsListGlobal.AnimalHandling.ModifierValue = SkillsListGlobal.AnimalHandling.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Insight.SkillName:
			SkillsListGlobal.Insight.Proficiency = true
			SkillsListGlobal.Insight.ProfLSS = 1
			SkillsListGlobal.Insight.ModifierValue = SkillsListGlobal.Insight.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Medicine.SkillName:
			SkillsListGlobal.Medicine.Proficiency = true
			SkillsListGlobal.Medicine.ProfLSS = 1
			SkillsListGlobal.Medicine.ModifierValue = SkillsListGlobal.Medicine.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Perception.SkillName:
			SkillsListGlobal.Perception.Proficiency = true
			SkillsListGlobal.Perception.ProfLSS = 1
			SkillsListGlobal.Perception.ModifierValue = SkillsListGlobal.Perception.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Survival.SkillName:
			SkillsListGlobal.Survival.Proficiency = true
			SkillsListGlobal.Survival.ProfLSS = 1
			SkillsListGlobal.Survival.ModifierValue = SkillsListGlobal.Survival.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Deception.SkillName:
			SkillsListGlobal.Deception.Proficiency = true
			SkillsListGlobal.Deception.ProfLSS = 1
			SkillsListGlobal.Deception.ModifierValue = SkillsListGlobal.Deception.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Intimidation.SkillName:
			SkillsListGlobal.Intimidation.Proficiency = true
			SkillsListGlobal.Intimidation.ProfLSS = 1
			SkillsListGlobal.Intimidation.ModifierValue = SkillsListGlobal.Intimidation.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Performance.SkillName:
			SkillsListGlobal.Performance.Proficiency = true
			SkillsListGlobal.Performance.ProfLSS = 1
			SkillsListGlobal.Performance.ModifierValue = SkillsListGlobal.Performance.ModifierValue + ProficiencyBonus
		case SkillsListGlobal.Persuasion.SkillName:
			SkillsListGlobal.Persuasion.Proficiency = true
			SkillsListGlobal.Persuasion.ProfLSS = 1
			SkillsListGlobal.Persuasion.ModifierValue = SkillsListGlobal.Persuasion.ModifierValue + ProficiencyBonus
		}
	}
	fmt.Println(SkillsListGlobal)
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

	sk.Athletics.ModifierValue = ModifierGlobal.Strength
	sk.Acrobatics.ModifierValue = ModifierGlobal.Dexterity
	sk.SleightOfHand.ModifierValue = ModifierGlobal.Dexterity
	sk.Stealth.ModifierValue = ModifierGlobal.Dexterity
	sk.Arcana.ModifierValue = ModifierGlobal.Intelligence
	sk.History.ModifierValue = ModifierGlobal.Intelligence
	sk.Investigation.ModifierValue = ModifierGlobal.Intelligence
	sk.Nature.ModifierValue = ModifierGlobal.Intelligence
	sk.Religion.ModifierValue = ModifierGlobal.Intelligence
	sk.AnimalHandling.ModifierValue = ModifierGlobal.Wisdom
	sk.Insight.ModifierValue = ModifierGlobal.Wisdom
	sk.Medicine.ModifierValue = ModifierGlobal.Wisdom
	sk.Perception.ModifierValue = ModifierGlobal.Wisdom
	sk.Survival.ModifierValue = ModifierGlobal.Wisdom
	sk.Deception.ModifierValue = ModifierGlobal.Charisma
	sk.Intimidation.ModifierValue = ModifierGlobal.Charisma
	sk.Performance.ModifierValue = ModifierGlobal.Charisma
	sk.Persuasion.ModifierValue = ModifierGlobal.Charisma

	return sk
}

func setSkillProficiency(skillMap []string, sk Skills) Skills {
	for _, profSkill := range skillMap {
		switch profSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.Proficiency = true
			sk.Athletics.ProfLSS = 1
			sk.Athletics.ModifierValue = sk.Athletics.ModifierValue + ProficiencyBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.Proficiency = true
			sk.Acrobatics.ProfLSS = 1
			sk.Acrobatics.ModifierValue = sk.Acrobatics.ModifierValue + ProficiencyBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.Proficiency = true
			sk.SleightOfHand.ProfLSS = 1
			sk.SleightOfHand.ModifierValue = sk.SleightOfHand.ModifierValue + ProficiencyBonus
		case sk.Stealth.SkillName:
			sk.Stealth.Proficiency = true
			sk.Stealth.ProfLSS = 1
			sk.Stealth.ModifierValue = sk.Stealth.ModifierValue + ProficiencyBonus
		case sk.Arcana.SkillName:
			sk.Arcana.Proficiency = true
			sk.Arcana.ProfLSS = 1
			sk.Arcana.ModifierValue = sk.Arcana.ModifierValue + ProficiencyBonus
		case sk.History.SkillName:
			sk.History.Proficiency = true
			sk.History.ProfLSS = 1
			sk.History.ModifierValue = sk.History.ModifierValue + ProficiencyBonus
		case sk.Investigation.SkillName:
			sk.Investigation.Proficiency = true
			sk.Investigation.ProfLSS = 1
			sk.Investigation.ModifierValue = sk.Investigation.ModifierValue + ProficiencyBonus
		case sk.Nature.SkillName:
			sk.Nature.Proficiency = true
			sk.Nature.ProfLSS = 1
			sk.Nature.ModifierValue = sk.Nature.ModifierValue + ProficiencyBonus
		case sk.Religion.SkillName:
			sk.Religion.Proficiency = true
			sk.Religion.ProfLSS = 1
			sk.Religion.ModifierValue = sk.Religion.ModifierValue + ProficiencyBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.Proficiency = true
			sk.AnimalHandling.ProfLSS = 1
			sk.AnimalHandling.ModifierValue = sk.AnimalHandling.ModifierValue + ProficiencyBonus
		case sk.Insight.SkillName:
			sk.Insight.Proficiency = true
			sk.Insight.ProfLSS = 1
			sk.Insight.ModifierValue = sk.Insight.ModifierValue + ProficiencyBonus
		case sk.Medicine.SkillName:
			sk.Medicine.Proficiency = true
			sk.Medicine.ProfLSS = 1
			sk.Medicine.ModifierValue = sk.Medicine.ModifierValue + ProficiencyBonus
		case sk.Perception.SkillName:
			sk.Perception.Proficiency = true
			sk.Perception.ProfLSS = 1
			sk.Perception.ModifierValue = sk.Perception.ModifierValue + ProficiencyBonus
		case sk.Survival.SkillName:
			sk.Survival.Proficiency = true
			sk.Survival.ProfLSS = 1
			sk.Survival.ModifierValue = sk.Survival.ModifierValue + ProficiencyBonus
		case sk.Deception.SkillName:
			sk.Deception.Proficiency = true
			sk.Deception.ProfLSS = 1
			sk.Deception.ModifierValue = sk.Deception.ModifierValue + ProficiencyBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.Proficiency = true
			sk.Intimidation.ProfLSS = 1
			sk.Intimidation.ModifierValue = sk.Intimidation.ModifierValue + ProficiencyBonus
		case sk.Performance.SkillName:
			sk.Performance.Proficiency = true
			sk.Performance.ProfLSS = 1
			sk.Performance.ModifierValue = sk.Performance.ModifierValue + ProficiencyBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.Proficiency = true
			sk.Persuasion.ProfLSS = 1
			sk.Persuasion.ModifierValue = sk.Persuasion.ModifierValue + ProficiencyBonus
		}
	}

	setPassiveWisdom(sk)
	return sk
}

func setDoubleSkillProficiency(dblMap []string, sk Skills) Skills {
	for _, dblProfSkill := range dblMap {
		switch dblProfSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.DoubleProficiency = true
			sk.Athletics.ProfLSS = 2
			sk.Athletics.ModifierValue += ProficiencyBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.DoubleProficiency = true
			sk.Acrobatics.ProfLSS = 2
			sk.Acrobatics.ModifierValue += ProficiencyBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.DoubleProficiency = true
			sk.SleightOfHand.ProfLSS = 2
			sk.SleightOfHand.ModifierValue += ProficiencyBonus
		case sk.Stealth.SkillName:
			sk.Stealth.DoubleProficiency = true
			sk.Stealth.ProfLSS = 2
			sk.Stealth.ModifierValue += ProficiencyBonus
		case sk.Arcana.SkillName:
			sk.Arcana.DoubleProficiency = true
			sk.Arcana.ProfLSS = 2
			sk.Arcana.ModifierValue += ProficiencyBonus
		case sk.History.SkillName:
			sk.History.DoubleProficiency = true
			sk.History.ProfLSS = 2
			sk.History.ModifierValue += ProficiencyBonus
		case sk.Investigation.SkillName:
			sk.Investigation.DoubleProficiency = true
			sk.Investigation.ProfLSS = 2
			sk.Investigation.ModifierValue += ProficiencyBonus
		case sk.Nature.SkillName:
			sk.Nature.DoubleProficiency = true
			sk.Nature.ProfLSS = 2
			sk.Nature.ModifierValue += ProficiencyBonus
		case sk.Religion.SkillName:
			sk.Religion.DoubleProficiency = true
			sk.Religion.ProfLSS = 2
			sk.Religion.ModifierValue += ProficiencyBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.DoubleProficiency = true
			sk.AnimalHandling.ProfLSS = 2
			sk.AnimalHandling.ModifierValue += ProficiencyBonus
		case sk.Insight.SkillName:
			sk.Insight.DoubleProficiency = true
			sk.Insight.ProfLSS = 2
			sk.Insight.ModifierValue += ProficiencyBonus
		case sk.Medicine.SkillName:
			sk.Medicine.DoubleProficiency = true
			sk.Medicine.ProfLSS = 2
			sk.Medicine.ModifierValue += ProficiencyBonus
		case sk.Perception.SkillName:
			sk.Perception.DoubleProficiency = true
			sk.Perception.ProfLSS = 2
			sk.Perception.ModifierValue += ProficiencyBonus
		case sk.Survival.SkillName:
			sk.Survival.DoubleProficiency = true
			sk.Survival.ProfLSS = 2
			sk.Survival.ModifierValue += ProficiencyBonus
		case sk.Deception.SkillName:
			sk.Deception.DoubleProficiency = true
			sk.Deception.ProfLSS = 2
			sk.Deception.ModifierValue += ProficiencyBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.DoubleProficiency = true
			sk.Intimidation.ProfLSS = 2
			sk.Intimidation.ModifierValue += ProficiencyBonus
		case sk.Performance.SkillName:
			sk.Performance.DoubleProficiency = true
			sk.Performance.ProfLSS = 2
			sk.Performance.ModifierValue += ProficiencyBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.DoubleProficiency = true
			sk.Persuasion.ProfLSS = 2
			sk.Persuasion.ModifierValue += ProficiencyBonus
		}
	}
	return sk
}

func setPassiveWisdom(sk Skills) {
	PassWisdom = 10 + sk.Perception.ModifierValue
}
