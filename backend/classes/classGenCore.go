package classes

import (
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

func GetNewClass() Class {
	var class Class
	class.ClassName = "HomeBrew" //default name

	class.Ability = RerollClassAbilitiesStats()
	class.ClassName = statAnalyze(sortStats(class.Ability), class)
	class.Modifier = getModifiersForClass(class.Ability)
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

func RerollClassAbilitiesStats() Ability {
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
	fmt.Println(modifier)

	return modifier
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

	keys := make([]string, 0, len(statsMap))

	for key := range statsMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return statsMap[keys[i]] > statsMap[keys[j]]
	})
	top3stats := []string{}
	for i, k := range keys {
		if i == 0 || i == 1 || i == 2 { //первые 3 значения
			top3stats = append(top3stats, k)
			//fmt.Println(k, statsMap[k])
		}
	}
	fmt.Println(top3stats)
	return top3stats
}
func statAnalyze(stats []string, cl Class) string {
	var (
		statMap = map[string][]string{
			"Classic Wizard": {"Intelligence", "BodyDifficulty", "Charisma"},
			"Snake Wizard":   {"Intelligence", "Dexterity", "Charisma"},
			"Charm Wizard":   {"Intelligence", "Wisdom", "Charisma"},

			"Classic Fighter": {"Strength", "Dexterity", "Intelligence"},
			"Bag Fighter":     {"Strength", "Dexterity", "BodyDifficulty"},

			"Wisdom Barbarian": {"Strength", "BodyDifficulty", "Wisdom"},
			"Charm Barbarian":  {"Strength", "BodyDifficulty", "Charisma"},

			"Big Paladin":     {"Strength", "Charisma", "BodyDifficulty"},
			"Wisdom Paladin2": {"Strength", "Charisma", "Wisdom"},

			"statForMonk":  {"Dexterity", "Wisdom", "Strength"},
			"statForMonk2": {"Dexterity", "Wisdom", "Charisma"},

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
	fmt.Println(cl.ClassName)
	return cl.ClassName
}

// make for other classes metrics for generation
//func classAnalyzer(class Class) Class {
//	var sum int
//	sum = class.BodyDifficulty + class.Dexterity + class.Intelligence +
//		class.Charisma + class.Strength + class.Wisdom
//	//balancer
//	if sum < 67 || sum > 78 {
//		class = RerollClassStats()
//		sum = class.BodyDifficulty + class.Dexterity + class.Intelligence +
//			class.Charisma + class.Strength + class.Wisdom
//		//fmt.Println("Перерол на", sum)
//	}
//	return class
//}
