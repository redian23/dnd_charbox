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

func RerollClassStats() Class {
	TargetClass := Class{
		ClassName:      "NAME_CLASS_XXX",
		Description:    "My AI to generated this.",
		Strength:       RandomRollPoints(),
		Dexterity:      RandomRollPoints(),
		BodyDifficulty: RandomRollPoints(),
		Intelligence:   RandomRollPoints(),
		Wisdom:         RandomRollPoints(),
		Charisma:       RandomRollPoints(),
	}
	return TargetClass
}

func GetNewClass() Class {
	randClass := RerollClassStats()
	randClass.ClassName = statAnalyze(sortStats(randClass), randClass)
	//classAnalyzer(RerollClassStats()) //defug
	return randClass
}

func sortStats(class Class) []string {
	var statsMap = map[string]int{
		"Strength":       class.Strength,
		"Dexterity":      class.Dexterity,
		"BodyDifficulty": class.BodyDifficulty,
		"Intelligence":   class.Intelligence,
		"Wisdom":         class.Wisdom,
		"Charisma":       class.Charisma,
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
			fmt.Println(k, statsMap[k])
		}
	}
	fmt.Println(top3stats)
	return top3stats
}
func statAnalyze(stats []string, cl Class) string {

	var statForMage = []string{"Intelligence", "BodyDifficulty", "Charisma"}
	var statForMage2 = []string{"Intelligence", "Dexterity", "Charisma"}
	var statForMage3 = []string{"Intelligence", "Wisdom", "Charisma"}

	var statForFighter = []string{"Strength", "Dexterity", "Intelligence"}
	var statForFighter2 = []string{"Strength", "Dexterity", "BodyDifficulty"}

	var statForRogue = []string{"Dexterity", "Intelligence", "Charisma"}
	var statForRogue2 = []string{"Dexterity", "Intelligence", "Wisdom"}

	var statForWarrior = []string{"Strength", "BodyDifficulty", "Wisdom"}
	var statForWarrior2 = []string{"Strength", "BodyDifficulty", "Charisma"}

	var statForBard = []string{"Charisma", "Dexterity", "Wisdom"}
	var statForBard2 = []string{"Charisma", "Dexterity", "BodyDifficulty"}
	var statForBard3 = []string{"Charisma", "Dexterity", "Intelligence"}

	var statForDruid = []string{"Wisdom", "BodyDifficulty", "Dexterity"}
	var statForDruid2 = []string{"Wisdom", "BodyDifficulty", "Strength"}

	var statForRanger = []string{"Dexterity", "Wisdom", "Charisma"}
	var statForRanger2 = []string{"Dexterity", "Wisdom", "Strength"}

	tempSortArray := [][]string{statForMage, statForMage2, statForMage3,
		statForFighter, statForFighter2, statForRogue, statForRogue2, statForWarrior, statForWarrior2,
		statForBard, statForBard2, statForBard3, statForDruid, statForDruid2, statForRanger, statForRanger2}

	for _, arrstrings := range tempSortArray {
		sort.Strings(arrstrings)
	}
	switch {
	case reflect.DeepEqual(stats, statForMage):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForMage2):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForMage3):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForFighter):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForFighter2):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForRogue):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForRogue2):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForWarrior):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForWarrior2):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForBard):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForBard2):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForBard3):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForDruid):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForDruid2):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForRanger):
		cl.ClassName = "Def MAg"
	case reflect.DeepEqual(stats, statForRanger2):
		cl.ClassName = "Def MAg"
	default:
		cl.ClassName = "????"

	}

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
