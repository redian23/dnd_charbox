package classes

import (
	"encoding/json"
	"fmt"
	"pregen/backend/dice"
	"reflect"
	"sort"
)

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func minimum(a []int) (int, int) {
	min := a[0]
	var index int
	for i, v := range a {
		if v < min {
			min = v
			index = i
		}
	}
	return min, index
}

func summa(a []int) int {
	var sum int
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func RandomRollPoints() int {
	dice := dice.D6
	firstTry := dice.RollDice()
	secondTry := dice.RollDice()
	thirdTry := dice.RollDice()
	fourthTry := dice.RollDice()
	arrayTry := []int{firstTry, secondTry, thirdTry, fourthTry}
	_, index := minimum(arrayTry)
	resultArray := remove(arrayTry, index)

	return summa(resultArray)
}

func rerollClassAbilitiesStats() Ability {
	TargetClassAbilities := Ability{
		Strength:       RandomRollPoints(),
		Dexterity:      RandomRollPoints(),
		BodyDifficulty: RandomRollPoints(),
		Intelligence:   RandomRollPoints(),
		Wisdom:         RandomRollPoints(),
		Charisma:       RandomRollPoints(),
	}
	TargetClassAbilities.Total = TargetClassAbilities.Strength + TargetClassAbilities.Dexterity +
		TargetClassAbilities.BodyDifficulty + TargetClassAbilities.Wisdom + TargetClassAbilities.Charisma +
		TargetClassAbilities.Intelligence

	return TargetClassAbilities
}

func sortMapCustom(statMap map[string]int) ([]string, map[string]int) {
	keys := make([]string, 0, len(statMap))

	for key := range statMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return statMap[keys[i]] > statMap[keys[j]]
	})
	return keys, statMap
}
func sortStats(abil Ability) []string {
	var statsMap = map[string]int{
		"Strength":       abil.Strength,
		"Dexterity":      abil.Dexterity,
		"BodyDifficulty": abil.BodyDifficulty,
		"Intelligence":   abil.Intelligence,
		"Wisdom":         abil.Wisdom,
		"Charisma":       abil.Charisma,
	}
	keys, _ := sortMapCustom(statsMap)
	statsForFindClassSpec := []string{} //для анализа класса
	for i, k := range keys {
		if i == 0 || i == 1 || i == 2 { //первые 3 значения
			statsForFindClassSpec = append(statsForFindClassSpec, k)
			//fmt.Println(k, statsMap[k])
		}
	}
	return statsForFindClassSpec
}
func statAnalyze(cl Class) (string, string, Ability) {
START:
	cl.Ability = rerollClassAbilitiesStats()
	var stats = sortStats(cl.Ability)
	//var jsonData = db.SelectJsonFromPgTable("SELECT * FROM race_characteristic_json;")
	var chars CharacteristicsForClass
	json.Unmarshal(RaceCharactsJsonData, &chars)

	for _, char := range chars {
		for _, cla := range char.CharReq {
			if reflect.DeepEqual(stats, cla) {
				cl.ClassName = char.ClassName
				cl.ClassNameRU = char.ClassNameRU
			}
		}
	}

	if cl.ClassName == "HomeBrew" {
		cl.Ability = rerollClassAbilitiesStats()
		goto START
	}
	//fmt.Println(cl.ClassName)
	//fmt.Println(cl.Ability)
	return cl.ClassName, cl.ClassNameRU, cl.Ability
}

func setModifiersForClass(ab Ability) Modifier {
	abilitiesArray := []int{ab.Strength, ab.Dexterity, ab.BodyDifficulty,
		ab.Intelligence, ab.Wisdom, ab.Charisma}
	var modifier Modifier
	modifierArray := []int{}
	for _, ability := range abilitiesArray {
		mod := (ability - 10) / 2
		modifierArray = append(modifierArray, mod)
	}

	for range modifierArray {
		modifier.Strength = modifierArray[0]
		modifier.Dexterity = modifierArray[1]
		modifier.BodyDifficulty = modifierArray[2]
		modifier.Intelligence = modifierArray[3]
		modifier.Wisdom = modifierArray[4]
		modifier.Charisma = modifierArray[5]
	}
	//fmt.Println(modifier)
	return modifier
}

func setSaveThrowsForClass(modifier Modifier) SavingThrows {
	var modifierMap = map[string]int{
		"Strength":       modifier.Strength,
		"Dexterity":      modifier.Dexterity,
		"Intelligence":   modifier.Intelligence,
		"BodyDifficulty": modifier.BodyDifficulty,
		"Wisdom":         modifier.Wisdom,
		"Charisma":       modifier.Charisma,
	}

	var saveTh SavingThrows
	keys, _ := sortMapCustom(modifierMap)
	statsForSV := []string{} //для анализа класса

	for k, value := range modifierMap { //первичное заполнение
		switch k {
		case "Strength":
			saveTh.Strength.Name = "Strength"
			saveTh.Strength.Point = value
		case "Dexterity":
			saveTh.Dexterity.Name = "Dexterity"
			saveTh.Dexterity.Point = value
		case "BodyDifficulty":
			saveTh.BodyDifficulty.Name = "BodyDifficulty"
			saveTh.BodyDifficulty.Point = value
		case "Intelligence":
			saveTh.Intelligence.Name = "Intelligence"
			saveTh.Intelligence.Point = value
		case "Wisdom":
			saveTh.Wisdom.Name = "Wisdom"
			saveTh.Wisdom.Point = value
		case "Charisma":
			saveTh.Charisma.Name = "Charisma"
			saveTh.Charisma.Point = value
		}
	}

	for i, k := range keys { //поиск 2х наивысших стат
		if i == 0 || i == 1 { //первые 2 значения
			statsForSV = append(statsForSV, k)
			//fmt.Println(k, statMap[k])
		}
	}

	for _, stat := range statsForSV {
		switch stat {
		case saveTh.Strength.Name:
			saveTh.Strength.Mastery = true
		case saveTh.Dexterity.Name:
			saveTh.Dexterity.Mastery = true
		case saveTh.BodyDifficulty.Name:
			saveTh.BodyDifficulty.Mastery = true
		case saveTh.Intelligence.Name:
			saveTh.Intelligence.Mastery = true
		case saveTh.Wisdom.Name:
			saveTh.Wisdom.Mastery = true
		case saveTh.Charisma.Name:
			saveTh.Charisma.Mastery = true
		}
	}
	//fmt.Println("2 - ", saveTh)
	return saveTh
}

func SetSkillsForClass(profSkills []string, modifier Modifier) Skills {

	fmt.Println("KURWA", profSkills)
	var sk Skills
	mobifierArray := []int{modifier.Strength, modifier.Dexterity,
		modifier.Intelligence, modifier.Wisdom, modifier.Charisma}

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

	for i, _ := range mobifierArray {
		switch {
		case modifier.Strength == mobifierArray[i]:
			sk.Athletics.ModifierValue = modifier.Strength
		case modifier.Dexterity == mobifierArray[i]:
			sk.Acrobatics.ModifierValue = modifier.Dexterity
			sk.SleightOfHand.ModifierValue = modifier.Dexterity
			sk.Stealth.ModifierValue = modifier.Dexterity
		case modifier.Intelligence == mobifierArray[i]:
			sk.Arcana.ModifierValue = modifier.Intelligence
			sk.History.ModifierValue = modifier.Intelligence
			sk.Investigation.ModifierValue = modifier.Intelligence
			sk.Nature.ModifierValue = modifier.Intelligence
			sk.Religion.ModifierValue = modifier.Intelligence
		case modifier.Wisdom == mobifierArray[i]:
			sk.AnimalHandling.ModifierValue = modifier.Wisdom
			sk.Insight.ModifierValue = modifier.Wisdom
			sk.Medicine.ModifierValue = modifier.Wisdom
			sk.Perception.ModifierValue = modifier.Wisdom
			sk.Survival.ModifierValue = modifier.Wisdom
		case modifier.Charisma == mobifierArray[i]:
			sk.Deception.ModifierValue = modifier.Charisma
			sk.Intimidation.ModifierValue = modifier.Charisma
			sk.Performance.ModifierValue = modifier.Charisma
			sk.Persuasion.ModifierValue = modifier.Charisma
		}
	}

	for _, profSkill := range profSkills {
		switch profSkill {
		case sk.Athletics.SkillName:
			sk.Athletics.Proficiency = true
			sk.Athletics.ModifierValue = sk.Athletics.ModifierValue + ProficiencyBonus
		case sk.Acrobatics.SkillName:
			sk.Acrobatics.Proficiency = true
			sk.Acrobatics.ModifierValue = sk.Acrobatics.ModifierValue + ProficiencyBonus
		case sk.SleightOfHand.SkillName:
			sk.SleightOfHand.Proficiency = true
			sk.SleightOfHand.ModifierValue = sk.SleightOfHand.ModifierValue + ProficiencyBonus
		case sk.Stealth.SkillName:
			sk.Stealth.Proficiency = true
			sk.Stealth.ModifierValue = sk.Stealth.ModifierValue + ProficiencyBonus
		case sk.Arcana.SkillName:
			sk.Arcana.Proficiency = true
			sk.Arcana.ModifierValue = sk.Arcana.ModifierValue + ProficiencyBonus
		case sk.History.SkillName:
			sk.History.Proficiency = true
			sk.History.ModifierValue = sk.History.ModifierValue + ProficiencyBonus
		case sk.Investigation.SkillName:
			sk.Investigation.Proficiency = true
			sk.Investigation.ModifierValue = sk.Investigation.ModifierValue + ProficiencyBonus
		case sk.Nature.SkillName:
			sk.Nature.Proficiency = true
			sk.Nature.ModifierValue = sk.Nature.ModifierValue + ProficiencyBonus
		case sk.Religion.SkillName:
			sk.Religion.Proficiency = true
			sk.Religion.ModifierValue = sk.Religion.ModifierValue + ProficiencyBonus
		case sk.AnimalHandling.SkillName:
			sk.AnimalHandling.Proficiency = true
			sk.AnimalHandling.ModifierValue = sk.AnimalHandling.ModifierValue + ProficiencyBonus
		case sk.Insight.SkillName:
			sk.Insight.Proficiency = true
			sk.Insight.ModifierValue = sk.Insight.ModifierValue + ProficiencyBonus
		case sk.Medicine.SkillName:
			sk.Medicine.Proficiency = true
			sk.Medicine.ModifierValue = sk.Medicine.ModifierValue + ProficiencyBonus
		case sk.Perception.SkillName:
			sk.Perception.Proficiency = true
			sk.Perception.ModifierValue = sk.Perception.ModifierValue + ProficiencyBonus
		case sk.Survival.SkillName:
			sk.Survival.Proficiency = true
			sk.Survival.ModifierValue = sk.Survival.ModifierValue + ProficiencyBonus
		case sk.Deception.SkillName:
			sk.Deception.Proficiency = true
			sk.Deception.ModifierValue = sk.Deception.ModifierValue + ProficiencyBonus
		case sk.Intimidation.SkillName:
			sk.Intimidation.Proficiency = true
			sk.Intimidation.ModifierValue = sk.Intimidation.ModifierValue + ProficiencyBonus
		case sk.Performance.SkillName:
			sk.Performance.Proficiency = true
			sk.Performance.ModifierValue = sk.Performance.ModifierValue + ProficiencyBonus
		case sk.Persuasion.SkillName:
			sk.Persuasion.Proficiency = true
			sk.Persuasion.ModifierValue = sk.Persuasion.ModifierValue + ProficiencyBonus
		}
	}
	return sk
}

func setPassiveWisdom(modWisdom int) int {
	var passWisdom int
	passWisdom = 10 + modWisdom
	return passWisdom
}
