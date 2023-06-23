package classes

import (
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

func GetNewClass() Class {
	//Start To Generate
	//Default Values
	var class Class
	class.ClassName = "HomeBrew"
	class.Inspiration = false
	class.ProficiencyBonus = 2
	class.Ability = rerollClassAbilitiesStats()

	var statsForClass = sortStats(class.Ability)

	class.ClassName = statAnalyze(statsForClass, class)
	class.Modifier = getModifiersForClass(class.Ability)
	class.SavingThrows = getSaveThrowsForClass(class.Modifier)
	return class
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
	//fmt.Println("1 - ", statsForFindClassSpec)
	return statsForFindClassSpec
}
func statAnalyze(stats []string, cl Class) string {
	var (
		statMap = map[string][]string{ //залить в базу
			"Classic Wizard": {"Intelligence", "BodyDifficulty", "Charisma"},
			"Snake Wizard":   {"Intelligence", "Dexterity", "Charisma"},
			"Charm Wizard":   {"Intelligence", "Wisdom", "Charisma"},

			"Classic Fighter": {"Strength", "Dexterity", "Intelligence"},
			"Bag Fighter":     {"Strength", "Dexterity", "BodyDifficulty"},

			"Wisdom Barbarian": {"Strength", "BodyDifficulty", "Wisdom"},
			"Charm Barbarian":  {"Strength", "BodyDifficulty", "Charisma"},

			"Big Paladin":     {"Strength", "Charisma", "BodyDifficulty"},
			"Wisdom Paladin2": {"Strength", "Charisma", "Wisdom"},

			"Classic Monk": {"Dexterity", "Wisdom", "Strength"},
			"Charm Monk":   {"Dexterity", "Wisdom", "Charisma"},

			"Classic Rogue": {"Dexterity", "Intelligence", "Charisma"},
			"Wisdom Rogue":  {"Dexterity", "Intelligence", "Wisdom"},

			"Classic Ranger": {"Dexterity", "Wisdom", "Charisma"}, //stock
			"Strong Ranger":  {"Strength", "Wisdom", "Dexterity"}, // 2th weapon

			"Classic Druid": {"Wisdom", "BodyDifficulty", "Dexterity"},

			"Big Cleric":    {"Wisdom", "BodyDifficulty", "Strength"},
			"Strong Cleric": {"Wisdom", "Strength", "BodyDifficulty"},

			"Classic Warlock": {"Charisma", "BodyDifficulty", "Wisdom"},

			"Classic Sorcerer": {"Charisma", "BodyDifficulty", "Wisdom"},
			"Smart Sorcerer":   {"Charisma", "BodyDifficulty", "Intelligence"},

			"Classic Bard": {"Charisma", "Dexterity", "Wisdom"},
			"Big Bard":     {"Charisma", "Dexterity", "BodyDifficulty"},
			"Smart Bard":   {"Charisma", "Dexterity", "Intelligence"},
		}
	)

	for name, statsArray := range statMap {
		switch {
		case reflect.DeepEqual(stats, statsArray):
			cl.ClassName = name
		}
	}
	//fmt.Println(cl.ClassName)
	return cl.ClassName
}

func getModifiersForClass(ab Ability) Modifier {
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

func getSaveThrowsForClass(modifier Modifier) SavingThrows {
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

//func getSkillsForClass(modifier Modifier) Skills {
//
//	var sk Skills
//	var prof = false
//	mobifierArray := []int{modifier.Strength, modifier.Dexterity,
//		modifier.Intelligence, modifier.Wisdom, modifier.Charisma}
//
//	for i, _ := range mobifierArray {
//		switch {
//		case modifier.Strength == mobifierArray[i]:
//			sk.Athletics = skill{modifier.Strength, prof}
//		case modifier.Dexterity == mobifierArray[i]:
//			sk.Acrobatics = skill{modifier.Dexterity, prof}
//			sk.SleightOfHand = skill{modifier.Dexterity, prof}
//			sk.Stealth = skill{modifier.Dexterity, prof}
//		case modifier.Intelligence == mobifierArray[i]:
//			sk.Arcana = skill{modifier.Intelligence, prof}
//			sk.History = skill{modifier.Intelligence, prof}
//			sk.Investigation = skill{modifier.Intelligence, prof}
//			sk.Nature = skill{modifier.Intelligence, prof}
//			sk.Religion = skill{modifier.Intelligence, prof}
//		case modifier.Wisdom == mobifierArray[i]:
//			sk.AnimalHandling = skill{modifier.Wisdom, prof}
//			sk.Insight = skill{modifier.Wisdom, prof}
//			sk.Medicine = skill{modifier.Wisdom, prof}
//			sk.Perception = skill{modifier.Wisdom, prof}
//			sk.Survival = skill{modifier.Wisdom, prof}
//		case modifier.Charisma == mobifierArray[i]:
//			sk.Deception = skill{modifier.Charisma, prof}
//			sk.Intimidation = skill{modifier.Charisma, prof}
//			sk.Performance = skill{modifier.Charisma, prof}
//			sk.Persuasion = skill{modifier.Charisma, prof}
//		}
//	}
//	return sk
//}
