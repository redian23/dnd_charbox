package core

import (
	"fmt"
	"pregen/pkg/classes"
	"pregen/pkg/races"
	"slices"
	"sort"
)

var abs classes.AbilityScore

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func minimum(a []int) (int, int) {
	minimal := a[0]
	var index int
	for i, v := range a {
		if v < minimal {
			minimal = v
			index = i
		}
	}
	return minimal, index
}

func summa(a []int) int {
	var sum int
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func RandomRollPoints() int {
	firstTry := D6.RollDice()
	secondTry := D6.RollDice()
	thirdTry := D6.RollDice()
	fourthTry := D6.RollDice()
	arrayTry := []int{firstTry, secondTry, thirdTry, fourthTry}
	_, index := minimum(arrayTry)
	resultArray := remove(arrayTry, index)

	return summa(resultArray)
}

func characteristicsDistribution(abilityRequest []string) map[string]int {
	var abilityScoreArray []int
	for i := 0; i < 6; i++ {
		abilityScoreArray = append(abilityScoreArray, RandomRollPoints())
	}
	sort.Ints(abilityScoreArray)
	slices.Reverse(abilityScoreArray)

	var characteristicsMap = map[string]int{
		"Strength":       0,
		"Dexterity":      0,
		"BodyDifficulty": 0,
		"Intelligence":   0,
		"Wisdom":         0,
		"Charisma":       0,
	}

	for char, _ := range characteristicsMap {
		found := false
		for i, charReq := range abilityRequest {
			if char == charReq {
				fmt.Println(char, charReq)
				characteristicsMap[char] = abilityScoreArray[i]
				found = true
				break
			}
		}
		if !found {
			// Если характеристика не в списке требуемых, берем значение из массива
			characteristicsMap[char] = abilityScoreArray[len(abilityScoreArray)-1]
			abilityScoreArray = abilityScoreArray[:len(abilityScoreArray)-1]
		}
	}

	return characteristicsMap
}

func addRaceAbilitiesToClassAbilities(charMap map[string]int) classes.AbilityScore {

	var raceAbilities = races.RaceInfo.Type.AbilityScorePlus

	for key, value := range charMap {
		switch key {
		case "Strength":
			abs.Strength = value
		case "Dexterity":
			abs.Dexterity = value
		case "BodyDifficulty":
			abs.BodyDifficulty = value
		case "Intelligence":
			abs.Intelligence = value
		case "Wisdom":
			abs.Wisdom = value
		case "Charisma":
			abs.Charisma = value
		}
	}

	for key, value := range raceAbilities {
		switch key {
		case "Strength":
			abs.Strength += value
		case "Dexterity":
			abs.Dexterity += value
		case "BodyDifficulty":
			abs.BodyDifficulty += value
		case "Intelligence":
			abs.Intelligence += value
		case "Wisdom":
			abs.Wisdom += value
		case "Charisma":
			abs.Charisma += value
		}
	}

	return abs
}

func abilitiesLevelUp(abilityRequest []string) {
	var lvl = classes.Level

	if lvl >= 4 && lvl <= 7 {
		switch abilityRequest[0] {
		case "Strength":
			abs.Strength += 2
		case "Dexterity":
			abs.Dexterity += 2
		case "BodyDifficulty":
			abs.BodyDifficulty += 2
		case "Intelligence":
			abs.Intelligence += 2
		case "Wisdom":
			abs.Wisdom += 2
		case "Charisma":
			abs.Charisma += 2
		}
	} else {
		for _, value := range abilityRequest {
			switch value {
			case "Strength":
				abs.Strength += 1
			case "Dexterity":
				abs.Dexterity += 1
			case "BodyDifficulty":
				abs.BodyDifficulty += 1
			case "Intelligence":
				abs.Intelligence += 1
			case "Wisdom":
				abs.Wisdom += 1
			case "Charisma":
				abs.Charisma += 1
			}
		}
	}
}

func GetClassAbilitiesScore(abilityRequest []string) classes.AbilityScore {
	var charMap = characteristicsDistribution(abilityRequest)
	fmt.Println(charMap)
	return addRaceAbilitiesToClassAbilities(charMap)
}
