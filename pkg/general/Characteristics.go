package general

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/classes"
	"pregen/pkg/races"
	"slices"
	"sort"
)

var abs classes.AbilityScore
var GlobalCharLevel = 8

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

func getSortedAbilitiesNamesArray(characteristicsMap map[string]int) []string {
	// Создаем слайс для хранения пар ключ-значение
	var keyValuePairs []struct {
		Key   string
		Value int
	}

	// Заполняем слайс парами ключ-значение из карты
	for k, v := range characteristicsMap {
		keyValuePairs = append(keyValuePairs, struct {
			Key   string
			Value int
		}{k, v})
	}

	// Сортируем слайс по значению
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value < keyValuePairs[j].Value
	})

	// Создаем слайс только с ключами в отсортированном порядке
	var sortedKeys []string
	for _, pair := range keyValuePairs {
		sortedKeys = append(sortedKeys, pair.Key)
	}

	slices.Reverse(sortedKeys)
	return sortedKeys
}

func randomRollPoints() int {
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
		abilityScoreArray = append(abilityScoreArray, randomRollPoints())
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

func setClassAbilities(charMap map[string]int) {
	abs = classes.AbilityScore{} //очистка при создании нового обьекта

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
}

func addRaceAbilitiesToClassAbilities(raceInfo *races.Race) {
	var raceAbilities = raceInfo.Type.AbilityScorePlus

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
}

func abilityUp(ability string, point int) {
	switch ability {
	case "Strength":
		abs.Strength += point
	case "Dexterity":
		abs.Dexterity += point
	case "BodyDifficulty":
		abs.BodyDifficulty += point
	case "Intelligence":
		abs.Intelligence += point
	case "Wisdom":
		abs.Wisdom += point
	case "Charisma":
		abs.Charisma += point
	}
}

func abilitiesLevelUp(charMap map[string]int) {

	var loopCount int
	if GlobalCharLevel >= 4 && GlobalCharLevel <= 7 {
		loopCount = 1
	}
	if GlobalCharLevel == 8 {
		loopCount = 2
	}
	abilities := getSortedAbilitiesNamesArray(charMap)

	for i := 0; i < loopCount; i++ {
		randNum, _ := random.IntRange(0, 2)
		if randNum == 0 {
			ability := abilities[i]
			switch ability {
			case "Strength":
				if abs.Strength <= 18 {
					abilityUp(abilities[i], 2)
					break
				}
				if abs.Strength > 18 {
					abilityUp(abilities[i+1], 2)
				}
			case "Dexterity":
				if abs.Dexterity <= 18 {
					abilityUp(abilities[i], 2)
					break
				}
				if abs.Dexterity > 18 {
					abilityUp(abilities[i+1], 2)
				}
			case "BodyDifficulty":
				if abs.BodyDifficulty <= 18 {
					abilityUp(abilities[i], 2)
					break
				}
				if abs.BodyDifficulty > 18 {
					abilityUp(abilities[i+1], 2)
				}
			case "Intelligence":
				if abs.Intelligence <= 18 {
					abilityUp(abilities[i], 2)
					break
				}
				if abs.Intelligence > 18 {
					abilityUp(abilities[i+1], 2)
				}
			case "Wisdom":
				if abs.Wisdom <= 18 {
					abilityUp(abilities[i], 2)
					break
				}
				if abs.Wisdom > 18 {
					abilityUp(abilities[i+1], 2)
				}
			case "Charisma":
				if abs.Charisma <= 18 {
					abilityUp(abilities[i], 2)
					break
				}
				if abs.Charisma > 18 {
					abilityUp(abilities[i+1], 2)
				}
			}
		} else {
			var iter int
			for abn, ability := range abilities {
				iter++
				switch ability {
				case "Strength":
					if abs.Strength <= 19 {
						abilityUp(ability, 1)
						break
					}
					if abs.Strength > 19 {
						abilityUp(abilities[abn+1], 1)
					}
				case "Dexterity":
					if abs.Dexterity <= 19 {
						abilityUp(ability, 1)
						break
					}
					if abs.Dexterity > 19 {
						abilityUp(abilities[abn+1], 1)
					}
				case "BodyDifficulty":
					if abs.BodyDifficulty <= 19 {
						abilityUp(ability, 1)
						break
					}
					if abs.BodyDifficulty > 19 {
						abilityUp(abilities[abn+1], 1)
					}
				case "Intelligence":
					if abs.Intelligence <= 19 {
						abilityUp(ability, 1)
						break
					}
					if abs.Intelligence > 19 {
						abilityUp(abilities[abn+1], 1)
					}
				case "Wisdom":
					if abs.Wisdom <= 19 {
						abilityUp(ability, 1)
						break
					}
					if abs.Wisdom > 19 {
						abilityUp(abilities[abn+1], 1)
					}
				case "Charisma":
					if abs.Charisma <= 19 {
						abilityUp(ability, 1)
						break
					}
					if abs.Charisma > 19 {
						abilityUp(abilities[abn+1], 1)
					}
				}
				if iter > 1 {
					break
				}
			}
		}

	}
}

func GetClassAbilitiesScore(abilityRequest []string, raceInfo *races.Race) classes.AbilityScore {
	var charMap = characteristicsDistribution(abilityRequest)

	setClassAbilities(charMap)
	addRaceAbilitiesToClassAbilities(raceInfo)
	abilitiesLevelUp(charMap)

	return abs
}
