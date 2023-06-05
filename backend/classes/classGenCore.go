package classes

import (
	"fmt"
	"pregen/backend/dice"
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
	sortStatMap(RerollClassStats())
	//classAnalyzer(RerollClassStats()) //defug
	return RerollClassStats()
}

func sortStatMap(class Class) {
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

	for i, k := range keys {
		if i == 0 || i == 1 || i == 2 { //первые 3 значения
			fmt.Println(k, statsMap[k])
		}
	}
}

// make for other classes metrics for generation
func classAnalyzer(class Class) Class {
	var sum int
	sum = class.BodyDifficulty + class.Dexterity + class.Intelligence +
		class.Charisma + class.Strength + class.Wisdom
	//balancer
	if sum < 67 || sum > 78 {
		class = RerollClassStats()
		sum = class.BodyDifficulty + class.Dexterity + class.Intelligence +
			class.Charisma + class.Strength + class.Wisdom
		//fmt.Println("Перерол на", sum)
	}
	return class
}
