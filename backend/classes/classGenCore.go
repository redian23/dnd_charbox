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
	fmt.Println("1 - ", statsForFindClassSpec)
	return statsForFindClassSpec
}
func statAnalyze(cl Class) (string, Ability) {
START:
	cl.Ability = rerollClassAbilitiesStats()
	var stats = sortStats(cl.Ability)
	//var jsonData = db.SelectJsonFromPgTable("SELECT * FROM race_characteristic_json;")
	var chars CharacteristicsForClass
	json.Unmarshal(raceCharactsJsonData, &chars)

	for _, char := range chars.Data {
		for _, cla := range char.CharReq {
			if reflect.DeepEqual(stats, cla) {
				cl.ClassName = char.ClassName
			}
		}
	}

	if cl.ClassName == "HomeBrew" {
		cl.Ability = rerollClassAbilitiesStats()
		goto START
	}
	//fmt.Println(cl.ClassName)
	//fmt.Println(cl.Ability)
	return cl.ClassName, cl.Ability
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

func setSkillsForClass(profSkills []string, modifier Modifier) Skills {

	var sk Skills
	var prof = false
	mobifierArray := []int{modifier.Strength, modifier.Dexterity,
		modifier.Intelligence, modifier.Wisdom, modifier.Charisma}

	for i, _ := range mobifierArray {
		switch {
		case modifier.Strength == mobifierArray[i]:
			sk.Athletics = skill{"Athletics", modifier.Strength, prof}
		case modifier.Dexterity == mobifierArray[i]:
			sk.Acrobatics = skill{"Acrobatics", modifier.Dexterity, prof}
			sk.SleightOfHand = skill{"SleightOfHand", modifier.Dexterity, prof}
			sk.Stealth = skill{"Stealth", modifier.Dexterity, prof}
		case modifier.Intelligence == mobifierArray[i]:
			sk.Arcana = skill{"Arcana", modifier.Intelligence, prof}
			sk.History = skill{"History", modifier.Intelligence, prof}
			sk.Investigation = skill{"Investigation", modifier.Intelligence, prof}
			sk.Magic = skill{"Magic", modifier.Intelligence, prof}
			sk.Nature = skill{"Nature", modifier.Intelligence, prof}
			sk.Religion = skill{"Religion", modifier.Intelligence, prof}
		case modifier.Wisdom == mobifierArray[i]:
			sk.AnimalHandling = skill{"AnimalHandling", modifier.Wisdom, prof}
			sk.Insight = skill{"Insight", modifier.Wisdom, prof}
			sk.Medicine = skill{"Medicine", modifier.Wisdom, prof}
			sk.Perception = skill{"Perception", modifier.Wisdom, prof}
			sk.Survival = skill{"Survival", modifier.Wisdom, prof}
		case modifier.Charisma == mobifierArray[i]:
			sk.Deception = skill{"Deception", modifier.Charisma, prof}
			sk.Intimidation = skill{"Intimidation", modifier.Charisma, prof}
			sk.Performance = skill{"Performance", modifier.Charisma, prof}
			sk.Persuasion = skill{"Persuasion", modifier.Charisma, prof}
		}
	}
	return sk
}

func setPassiveWisdom(modWisdom int) int {
	var passWisdom int
	passWisdom = 10 + modWisdom
	return passWisdom
}
