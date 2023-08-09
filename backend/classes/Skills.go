package classes

import (
	"github.com/mazen160/go-random"
	"pregen/backend/races"
	"reflect"
)

var (
	proficiencyBonus = 2
)

func SetSkillsForCharacter(backgroundSkills []string) Skills {
	var sk Skills
	// принимает нужный экземпляр обьекта и в себя же записывает новые данные
	sk = setSkillsNames(sk)
	sk = setSkillModifierValue(sk)

	skillsSlice := getAnalyzedSkillSlice(backgroundSkills)
	sk = setSkillProficiency(skillsSlice, sk)

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
	modifierArray := []int{mod.Strength, mod.Dexterity,
		mod.Intelligence, mod.Wisdom, mod.Charisma}

	for i := range modifierArray {
		switch {
		case mod.Strength == modifierArray[i]:
			sk.Athletics.ModifierValue = mod.Strength
		case mod.Dexterity == modifierArray[i]:
			sk.Acrobatics.ModifierValue = mod.Dexterity
			sk.SleightOfHand.ModifierValue = mod.Dexterity
			sk.Stealth.ModifierValue = mod.Dexterity
		case mod.Intelligence == modifierArray[i]:
			sk.Arcana.ModifierValue = mod.Intelligence
			sk.History.ModifierValue = mod.Intelligence
			sk.Investigation.ModifierValue = mod.Intelligence
			sk.Nature.ModifierValue = mod.Intelligence
			sk.Religion.ModifierValue = mod.Intelligence
		case mod.Wisdom == modifierArray[i]:
			sk.AnimalHandling.ModifierValue = mod.Wisdom
			sk.Insight.ModifierValue = mod.Wisdom
			sk.Medicine.ModifierValue = mod.Wisdom
			sk.Perception.ModifierValue = mod.Wisdom
			sk.Survival.ModifierValue = mod.Wisdom
		case mod.Charisma == modifierArray[i]:
			sk.Deception.ModifierValue = mod.Charisma
			sk.Intimidation.ModifierValue = mod.Charisma
			sk.Performance.ModifierValue = mod.Charisma
			sk.Persuasion.ModifierValue = mod.Charisma
		}
	}
	return sk
}

func setClassSkills() []string {
	var skillsArray []string
	var randSkillCount int
	var skills []string

	for _, char := range chars {
		if char.ClassName == ClassNameGlobal {
			skillsArray = char.SkillsInDB.SkillsList
			randSkillCount = char.SkillsInDB.RandomCount
		}
	}

	for i := range skillsArray {
		if i < randSkillCount {
			rollNum, _ := random.IntRange(0, len(skillsArray))
			skills = append(skills, skillsArray[rollNum])
		}
	}
	classSkills = skills
	return skills
}

func getAnalyzedSkillSlice(backgroundSkills []string) []string {
	var raceSkill = races.GetRaceSkill()

ReRollClassSkills:

	if reflect.DeepEqual(classSkills, backgroundSkills) == true {
		classSkills = setClassSkills()
		goto ReRollClassSkills
	}
	for _, classSkl := range classSkills {
		for _, backgroundSkl := range backgroundSkills {
			if classSkl == raceSkill {
				classSkills = setClassSkills()
				goto ReRollClassSkills
			}
			if classSkl == backgroundSkl {
				classSkills = setClassSkills()
				goto ReRollClassSkills
			}
		}
	}
	for i, classSkl1 := range classSkills {
		for j, classSkl2 := range classSkills {
			if classSkl1 == classSkl2 && i != j {
				classSkills = setClassSkills()
				goto ReRollClassSkills
			}
		}
	}

	var skillsSliceTmp = append(backgroundSkills, classSkills...)
	var skillsSlice = append(skillsSliceTmp, raceSkill)

	return skillsSlice
}

func setSkillProficiency(skillsSlice []string, sk Skills) Skills {
	for _, profSkill := range skillsSlice {
		switch profSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.Proficiency = true
			sk.Athletics.ModifierValue = sk.Athletics.ModifierValue + proficiencyBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.Proficiency = true
			sk.Acrobatics.ModifierValue = sk.Acrobatics.ModifierValue + proficiencyBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.Proficiency = true
			sk.SleightOfHand.ModifierValue = sk.SleightOfHand.ModifierValue + proficiencyBonus
		case sk.Stealth.SkillName:
			sk.Stealth.Proficiency = true
			sk.Stealth.ModifierValue = sk.Stealth.ModifierValue + proficiencyBonus
		case sk.Arcana.SkillName:
			sk.Arcana.Proficiency = true
			sk.Arcana.ModifierValue = sk.Arcana.ModifierValue + proficiencyBonus
		case sk.History.SkillName:
			sk.History.Proficiency = true
			sk.History.ModifierValue = sk.History.ModifierValue + proficiencyBonus
		case sk.Investigation.SkillName:
			sk.Investigation.Proficiency = true
			sk.Investigation.ModifierValue = sk.Investigation.ModifierValue + proficiencyBonus
		case sk.Nature.SkillName:
			sk.Nature.Proficiency = true
			sk.Nature.ModifierValue = sk.Nature.ModifierValue + proficiencyBonus
		case sk.Religion.SkillName:
			sk.Religion.Proficiency = true
			sk.Religion.ModifierValue = sk.Religion.ModifierValue + proficiencyBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.Proficiency = true
			sk.AnimalHandling.ModifierValue = sk.AnimalHandling.ModifierValue + proficiencyBonus
		case sk.Insight.SkillName:
			sk.Insight.Proficiency = true
			sk.Insight.ModifierValue = sk.Insight.ModifierValue + proficiencyBonus
		case sk.Medicine.SkillName:
			sk.Medicine.Proficiency = true
			sk.Medicine.ModifierValue = sk.Medicine.ModifierValue + proficiencyBonus
		case sk.Perception.SkillName:
			sk.Perception.Proficiency = true
			sk.Perception.ModifierValue = sk.Perception.ModifierValue + proficiencyBonus
		case sk.Survival.SkillName:
			sk.Survival.Proficiency = true
			sk.Survival.ModifierValue = sk.Survival.ModifierValue + proficiencyBonus
		case sk.Deception.SkillName:
			sk.Deception.Proficiency = true
			sk.Deception.ModifierValue = sk.Deception.ModifierValue + proficiencyBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.Proficiency = true
			sk.Intimidation.ModifierValue = sk.Intimidation.ModifierValue + proficiencyBonus
		case sk.Performance.SkillName:
			sk.Performance.Proficiency = true
			sk.Performance.ModifierValue = sk.Performance.ModifierValue + proficiencyBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.Proficiency = true
			sk.Persuasion.ModifierValue = sk.Persuasion.ModifierValue + proficiencyBonus
		}
	}
	return sk
}
