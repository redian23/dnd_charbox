package classes

import (
	"context"
	"github.com/mazen160/go-random"
	"pregen/backend/dice"
	"pregen/backend/races"
	"pregen/db"
	"reflect"
	"sort"
)

var (
	proficiencyBonus = 2
	chars            = GetClassCharactsFormDB()
	mod              Modifier
	specialSkills    []string
)

func GetClassCharactsFormDB() ClassesBSON {
	var results ClassesBSON
	var cursor = db.ReadFromDB("classes")
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func RandomRollPoints() int {
	d6 := dice.D6
	firstTry := d6.RollDice()
	secondTry := d6.RollDice()
	thirdTry := d6.RollDice()
	fourthTry := d6.RollDice()
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

	raceAbilities := races.GetRaceAbilities()

	TargetClassAbilities.Total = (TargetClassAbilities.Strength + raceAbilities.Strength) +
		(TargetClassAbilities.Dexterity + raceAbilities.Dexterity) +
		(TargetClassAbilities.BodyDifficulty + raceAbilities.BodyDifficulty) +
		(TargetClassAbilities.Intelligence + raceAbilities.Intelligence) +
		(TargetClassAbilities.Wisdom + raceAbilities.Wisdom) +
		(TargetClassAbilities.Charisma + raceAbilities.Charisma)
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
func extractStats(abil Ability) []string {
	var statsMap = map[string]int{
		"Strength":       abil.Strength,
		"Dexterity":      abil.Dexterity,
		"BodyDifficulty": abil.BodyDifficulty,
		"Intelligence":   abil.Intelligence,
		"Wisdom":         abil.Wisdom,
		"Charisma":       abil.Charisma,
	}
	keys, _ := sortMapCustom(statsMap)
	var statsForFindClassSpec []string //для анализа класса
	for i, k := range keys {
		if i == 0 || i == 1 || i == 2 { //первые 3 значения
			statsForFindClassSpec = append(statsForFindClassSpec, k)
			//fmt.Println(k, statsMap[k])
		}
	}
	return statsForFindClassSpec
}

func statAnalyze() (string, string, Ability) {
START:
	var abilities = rerollClassAbilitiesStats()
	var stats = extractStats(abilities)
	var name, nameRu string

	for _, char := range chars {
		for _, cla := range char.CharReq {
			if reflect.DeepEqual(stats, cla) {
				name = char.ClassName
				nameRu = char.ClassNameRU
			}
		}
	}

	if name == "" {
		goto START
	}

	return name, nameRu, abilities
}

func setModifiersForClass(ab Ability) Modifier {
	abilitiesArray := []int{ab.Strength, ab.Dexterity, ab.BodyDifficulty,
		ab.Intelligence, ab.Wisdom, ab.Charisma}
	var modifier Modifier
	var modifierArray []int

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
	mod = modifier //присвоение в переменную для всего пакета
	return modifier
}

func setSaveThrowsForClass(mod Modifier) SavingThrows {

	var modifierMap = map[string]int{
		"Strength":       mod.Strength,
		"Dexterity":      mod.Dexterity,
		"Intelligence":   mod.Intelligence,
		"BodyDifficulty": mod.BodyDifficulty,
		"Wisdom":         mod.Wisdom,
		"Charisma":       mod.Charisma,
	}

	var saveTh SavingThrows
	var saveThrArray []string

	for _, char := range chars {
		if char.ClassName == ClassNameGlobal {
			saveThrArray = char.SavingThrows
		}
	}

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

	for _, stat := range saveThrArray {
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
	specialSkills = skills
	return skills
}

func SetSkillsForCharacter(profSkills []string) Skills {
	var sk Skills
	modifierArray := []int{mod.Strength, mod.Dexterity,
		mod.Intelligence, mod.Wisdom, mod.Charisma}

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

	var raceSkill = races.GetRaceSkill()
	var skillsSliceTmp = append(profSkills, specialSkills...)
	var skillsSlice = append(skillsSliceTmp, raceSkill)

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

func setPassiveWisdom(modWisdom int) int {
	passWisdom := 10 + modWisdom
	return passWisdom
}

func setHitDice() string {
	var hitDice string
	for _, char := range chars {
		if char.ClassName == ClassNameGlobal {
			hitDice = char.Hits.HitDice
		}
	}
	return hitDice
}
func setHitCount(modBody int) int {
	var hitCount int
	for _, char := range chars {
		if char.ClassName == ClassNameGlobal {
			hitCount = char.Hits.HitCount + modBody
		}
	}
	return hitCount
}

func setProficiencies() Proficiencies {
	for _, char := range chars {
		if char.ClassName == ClassNameGlobal {
			return char.Proficiencies
		}
	}
	return Proficiencies{}
}

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
