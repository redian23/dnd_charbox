package classes

import (
	"context"
	"github.com/mazen160/go-random"
	"math"
	"pregen/pkg/db"
	"pregen/pkg/dice"
	"pregen/pkg/races"
	"reflect"
	"sort"
)

var (
	chars          ClassesBSON
	armorData      []ArmorAnswer
	weaponData     []WeaponAnswer
	equipmentList  []Variants
	classSkills    []string
	ModifierGlobal Modifier
)

func GetClassCharactsFormDB() ClassesBSON {
	var results ClassesBSON
	var cursor = db.ReadFromDB("classes")
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func GetArmorFormDB() []ArmorAnswer {
	var results []ArmorAnswer
	var cursor = db.ReadFromDB("armor")
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func GetWeaponFormDB() []WeaponAnswer {
	var results []WeaponAnswer
	var cursor = db.ReadFromDB("weapons")
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
	raceAbilities := races.GetRaceAbilities()

	TargetClassAbilities := Ability{
		Strength:       RandomRollPoints() + raceAbilities.Strength,
		Dexterity:      RandomRollPoints() + raceAbilities.Dexterity,
		BodyDifficulty: RandomRollPoints() + raceAbilities.BodyDifficulty,
		Intelligence:   RandomRollPoints() + raceAbilities.Intelligence,
		Wisdom:         RandomRollPoints() + raceAbilities.Wisdom,
		Charisma:       RandomRollPoints() + raceAbilities.Charisma,
	}

	TargetClassAbilities.Total = TargetClassAbilities.Strength + TargetClassAbilities.Dexterity +
		TargetClassAbilities.BodyDifficulty + TargetClassAbilities.Intelligence +
		TargetClassAbilities.Wisdom + TargetClassAbilities.Charisma

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
	var modifierArray []float64

	for _, ability := range abilitiesArray {
		modPoint := math.Floor(float64(ability-10) / 2)
		modifierArray = append(modifierArray, modPoint)
	}

	for range modifierArray {
		modifier.Strength = int(modifierArray[0])
		modifier.Dexterity = int(modifierArray[1])
		modifier.BodyDifficulty = int(modifierArray[2])
		modifier.Intelligence = int(modifierArray[3])
		modifier.Wisdom = int(modifierArray[4])
		modifier.Charisma = int(modifierArray[5])
	}
	//fmt.Println(modifier)
	ModifierGlobal = modifier //присвоение в переменную для всего пакета
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
			saveTh.Strength.Point += 2
			saveTh.Strength.Mastery = true
		case saveTh.Dexterity.Name:
			saveTh.Dexterity.Point += 2
			saveTh.Dexterity.Mastery = true
		case saveTh.BodyDifficulty.Name:
			saveTh.BodyDifficulty.Point += 2
			saveTh.BodyDifficulty.Mastery = true
		case saveTh.Intelligence.Name:
			saveTh.Intelligence.Point += 2
			saveTh.Intelligence.Mastery = true
		case saveTh.Wisdom.Name:
			saveTh.Wisdom.Point += 2
			saveTh.Wisdom.Mastery = true
		case saveTh.Charisma.Name:
			saveTh.Charisma.Point += 2
			saveTh.Charisma.Mastery = true
		}
	}
	//fmt.Println("2 - ", saveTh)
	return saveTh
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

func setClassEquipmentList() []Variants {
	var equipP equipPicks
	var equipB equipBasic
	var equipList []Variants
	for _, char := range chars {
		if char.ClassName == ClassNameGlobal {
			equipP = char.PicksEquipment
			equipB = char.BasicEquipment
		}
	}
	for _, item := range equipP {
		rollNum, _ := random.IntRange(0, len(item.Variants))
		equipList = append(equipList, item.Variants[rollNum])
	}
	for _, item := range equipB {
		equipList = append(equipList, item)
	}
	equipmentList = equipList
	return equipList
}

func getArmorList() []ArmorAnswer {
	var armList []ArmorAnswer
	for _, item := range equipmentList {
		if item.Type == "armor" {
			for _, armor := range armorData {
				if armor.ArmorName == item.ItemName {
					armList = append(armList, armor)
				}
			}
		}
	}
	return armList
}

func setArmor(className string) []ArmorAnswer {
	armorList := getArmorList()
	var armorAns ArmorAnswer
	var shield int
	var armorAnsList []ArmorAnswer

	for _, armor := range armorList {
		if armor.ArmorName == "Щит" {
			shield = 2
		}
	}
	for _, armor := range armorList {
		armorAns.ArmorName = armor.ArmorName
		armorAns.ArmorType = armor.ArmorType
		armorAns.Stealth = armor.Stealth
		switch armor.ArmorType {
		case "Лёгкий доспех":
			armorAns.ArmorClassCount = armor.ArmorClassCount + ModifierGlobal.Dexterity + shield
		case "Средний доспех":
			if ModifierGlobal.Dexterity > 2 {
				armorAns.ArmorClassCount = armor.ArmorClassCount + 2 + shield
			} else {
				armorAns.ArmorClassCount = armor.ArmorClassCount + ModifierGlobal.Dexterity + shield
			}
		case "Тяжёлый доспех":
			armorAns.ArmorClassCount = armor.ArmorClassCount + shield
		}
		armorAnsList = append(armorAnsList, armorAns)
	}

	if armorAnsList == nil {
		var noArmor = ArmorAnswer{ArmorName: "Без Доспеха", ArmorType: "Нет", Stealth: true}
		switch className {
		case "Монах":
			noArmor.ArmorClassCount = 10 + ModifierGlobal.Dexterity + ModifierGlobal.Wisdom
		case "Варвар":
			noArmor.ArmorClassCount = 10 + ModifierGlobal.Dexterity + ModifierGlobal.BodyDifficulty
		case "Волшебник":
			noArmor.ArmorClassCount = 10 + ModifierGlobal.Dexterity + ModifierGlobal.Intelligence
		}
		armorAnsList = append(armorAnsList, noArmor)
	}
	return armorAnsList
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

func GetAnalyzedSkillSlice(backgroundSkills []string) []string {
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

func setWeapons() []WeaponAnswer {
	var weaponList []WeaponAnswer

	for _, item := range equipmentList {
		if item.Type == "weapon" {
			switch item.ItemName {
			case "Простое оружие":
				rollNum, _ := random.IntRange(0, 13)
				weaponData[rollNum].Count = item.Count
				weaponList = append(weaponList, weaponData[rollNum])
				continue
			case "Простое рукопашное оружие":
				rollNum, _ := random.IntRange(0, 9)
				weaponData[rollNum].Count = item.Count
				weaponList = append(weaponList, weaponData[rollNum])
				continue
			case "Воинское оружие":
				rollNum, _ := random.IntRange(14, 36)
				weaponData[rollNum].Count = item.Count
				weaponList = append(weaponList, weaponData[rollNum])
				continue
			case "Воинское рукопашное оружие":
				rollNum, _ := random.IntRange(14, 31)
				weaponData[rollNum].Count = item.Count
				weaponList = append(weaponList, weaponData[rollNum])
				continue
			}
		}
	}

	for _, item := range equipmentList {
		if item.Type == "weapon" {
			for _, weapon := range weaponData {
				if weapon.WeaponName == item.ItemName {
					weapon.Count = item.Count
					weaponList = append(weaponList, weapon)
				}
			}
		}
	}

	for _, item := range equipmentList {
		if item.ItemName == "Короткий лук" || item.ItemName == "Длинный лук" {
			weaponList = append(weaponList, WeaponAnswer{"Колчан стрел", "weapon", "", "Боепр.", 20})
		}
		if item.ItemName == "Лёгкий арбалет" {
			weaponList = append(weaponList, WeaponAnswer{"Арбалетные болты", "weapon", "", "Боепр.", 20})
		}
	}

	return weaponList
}

func setInitiative() string {
	return "D20 + мод_ЛОВ"
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
