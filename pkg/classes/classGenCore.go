package classes

import (
	"encoding/json"
	"github.com/mazen160/go-random"
	"log"
	"math"
	"os"
	"pregen/pkg/db"
	"pregen/pkg/dice"
	"pregen/pkg/races"
	"reflect"
	"sort"
)

var (
	chars                   ClassesBSON
	armorData               []ArmorAnswer
	weaponData              []WeaponAnswer
	equipmentList           []Variants
	ModifierGlobal          Modifier
	CharProficienciesGlobal Proficiencies
	CharSavingThrowsGlobal  SavingThrows
)

func GetClassCharactsFormDB() ClassesBSON {
	var results ClassesBSON
	fileContent, err := os.Open(db.DataBasePath + "classes.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile(db.DataBasePath + "classes.json")

	json.Unmarshal(byteResult, &results)
	return results
}

func GetArmorFormDB() []ArmorAnswer {
	var results []ArmorAnswer
	fileContent, err := os.Open(db.DataBasePath + "armor.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile(db.DataBasePath + "armor.json")

	json.Unmarshal(byteResult, &results)
	return results
}

func GetWeaponFormDB() []WeaponAnswer {
	var results []WeaponAnswer
	fileContent, err := os.Open(db.DataBasePath + "weapons.json")
	if err != nil {
		log.Println(err)
	}
	defer fileContent.Close()
	var byteResult, _ = os.ReadFile(db.DataBasePath + "weapons.json")

	json.Unmarshal(byteResult, &results)
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

func AbilitiesWithLevelUp() (string, Ability) {
	classNameEng, classAbilityStats := statAnalyze(ClassNameGlobalRu)
	classAbilityStatsList := map[string]int{ //просто мне надо такая мапа // чтобы синхронить обьект и получать рандом без повторений
		"Strength":       0,
		"Dexterity":      0,
		"BodyDifficulty": 0,
		"Intelligence":   0,
		"Wisdom":         0,
		"Charisma":       0,
	}

	if CharacterLevelGlobal >= 4 && CharacterLevelGlobal < 8 {
		randNum, _ := random.IntRange(0, 2)
		if randNum == 0 {
			for key, _ := range classAbilityStatsList {
				switch key {
				case "Strength":
					classAbilityStats.Strength += 2
				case "Dexterity":
					classAbilityStats.Dexterity += 2
				case "BodyDifficulty":
					classAbilityStats.BodyDifficulty += 2
				case "Intelligence":
					classAbilityStats.Intelligence += 2
				case "Wisdom":
					classAbilityStats.Wisdom += 2
				case "Charisma":
					classAbilityStats.Charisma += 2
				}
				break
			}
		} else {
			var iter int
			for key, _ := range classAbilityStatsList {
				iter++
				switch key {
				case "Strength":
					classAbilityStats.Strength += 1
				case "Dexterity":
					classAbilityStats.Dexterity += 1
				case "BodyDifficulty":
					classAbilityStats.BodyDifficulty += 1
				case "Intelligence":
					classAbilityStats.Intelligence += 1
				case "Wisdom":
					classAbilityStats.Wisdom += 1
				case "Charisma":
					classAbilityStats.Charisma += 1
				}
				if iter > 1 {
					break
				}
			}
		}
	}
	if CharacterLevelGlobal > 7 {
		for i := 0; i < 2; i++ {
			randNum, _ := random.IntRange(0, 2)
			if randNum == 0 {
				for key, _ := range classAbilityStatsList {
					switch key {
					case "Strength":
						classAbilityStats.Strength += 2
					case "Dexterity":
						classAbilityStats.Dexterity += 2
					case "BodyDifficulty":
						classAbilityStats.BodyDifficulty += 2
					case "Intelligence":
						classAbilityStats.Intelligence += 2
					case "Wisdom":
						classAbilityStats.Wisdom += 2
					case "Charisma":
						classAbilityStats.Charisma += 2
					}
					break
				}
			} else {
				var iter int
				for key, _ := range classAbilityStatsList {
					iter++
					switch key {
					case "Strength":
						classAbilityStats.Strength += 1
					case "Dexterity":
						classAbilityStats.Dexterity += 1
					case "BodyDifficulty":
						classAbilityStats.BodyDifficulty += 1
					case "Intelligence":
						classAbilityStats.Intelligence += 1
					case "Wisdom":
						classAbilityStats.Wisdom += 1
					case "Charisma":
						classAbilityStats.Charisma += 1
					}
					if iter > 1 {
						break
					}
				}
			}
		}
	}

	classAbilityStats.Total = classAbilityStats.Strength + classAbilityStats.Dexterity +
		classAbilityStats.BodyDifficulty + classAbilityStats.Intelligence +
		classAbilityStats.Wisdom + classAbilityStats.Charisma

	return classNameEng, classAbilityStats
}

func statAnalyze(needClass string) (string, Ability) {
	var nameEng string
	var abilities Ability

	for i := 0; i < 200; i++ { //если юзать бесконечный for, то cpu взлетает до 100% под нагрузкой
		abilities = rerollClassAbilitiesStats()
		stats := extractStats(abilities)

		for _, char := range chars {
			for _, cla := range char.CharReq {
				if reflect.DeepEqual(stats, cla) && needClass == char.ClassNameRU {
					nameEng = char.ClassName
					return nameEng, abilities
				}
			}
		}
	}
	return getDefaultCharacterStats(needClass)
	//если за 200 траев не смогло найти подходящие статы для класса
	//от отдает дефолт статы для класса по системе (15,14,13,12,10,8)
}

func getDefaultCharacterStats(needClass string) (string, Ability) {

	classNameList := map[string]string{
		"Волшебник":    "Wizard",
		"Воин":         "Fighter",
		"Варвар":       "Barbarian",
		"Паладин":      "Paladin",
		"Монах":        "Monk",
		"Плут":         "Rogue",
		"Следопыт":     "Ranger",
		"Друид":        "Druid",
		"Жрец":         "Cleric",
		"Чародей":      "Sorcerer",
		"Бард":         "Bard",
		"Изобретатель": "Artificer",
		"Колдун":       "Warlock",
	}

	defaultCharactersMap := map[string]Ability{
		"Wizard":    {Strength: 8, Dexterity: 13, BodyDifficulty: 14, Intelligence: 15, Wisdom: 12, Charisma: 10, Total: 72},
		"Barbarian": {Strength: 15, Dexterity: 13, BodyDifficulty: 14, Intelligence: 8, Wisdom: 12, Charisma: 10, Total: 72},
		"Paladin":   {Strength: 15, Dexterity: 10, BodyDifficulty: 13, Intelligence: 8, Wisdom: 12, Charisma: 14, Total: 72},
		"Monk":      {Strength: 8, Dexterity: 15, BodyDifficulty: 13, Intelligence: 10, Wisdom: 14, Charisma: 12, Total: 72},
		"Rogue":     {Strength: 8, Dexterity: 15, BodyDifficulty: 13, Intelligence: 10, Wisdom: 12, Charisma: 14, Total: 72},
		"Ranger":    {Strength: 8, Dexterity: 15, BodyDifficulty: 12, Intelligence: 13, Wisdom: 14, Charisma: 10, Total: 72},
		"Druid":     {Strength: 8, Dexterity: 14, BodyDifficulty: 13, Intelligence: 12, Wisdom: 15, Charisma: 10, Total: 72},
		"Sorcerer":  {Strength: 8, Dexterity: 13, BodyDifficulty: 14, Intelligence: 12, Wisdom: 10, Charisma: 15, Total: 72},
		"Bard":      {Strength: 8, Dexterity: 14, BodyDifficulty: 13, Intelligence: 10, Wisdom: 12, Charisma: 15, Total: 72},
		"Warlock":   {Strength: 8, Dexterity: 14, BodyDifficulty: 13, Intelligence: 12, Wisdom: 10, Charisma: 15, Total: 72},

		"Cleric":    {Strength: 13, Dexterity: 12, BodyDifficulty: 14, Intelligence: 8, Wisdom: 15, Charisma: 10, Total: 72},
		"Artificer": {Strength: 10, Dexterity: 14, BodyDifficulty: 13, Intelligence: 15, Wisdom: 12, Charisma: 8, Total: 72},
		"Fighter":   {Strength: 15, Dexterity: 13, BodyDifficulty: 14, Intelligence: 8, Wisdom: 12, Charisma: 10, Total: 72},
		// перебрать вариации
	}

	classNameEng := classNameList[needClass]

	return classNameEng, defaultCharactersMap[classNameEng]
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

func setSaveThrowsForClass(mod Modifier) *SavingThrows {

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
		if char.ClassNameRU == ClassNameGlobalRu {
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
	CharSavingThrowsGlobal = saveTh
	return &CharSavingThrowsGlobal
}

func setHitDice() string {
	var hitDice string
	for _, char := range chars {
		if char.ClassNameRU == ClassNameGlobalRu {
			hitDice = char.Hits.HitDice
		}
	}
	return hitDice
}

func getHitDiceAverageCount() int {
	var hitDice string
	for _, char := range chars {
		if char.ClassNameRU == ClassNameGlobalRu {
			hitDice = char.Hits.HitDice
		}
	}

	hitDiceMap := map[string]int{
		"1к6":  4,
		"1к8":  5,
		"1к10": 6,
		"1к12": 7,
	}

	hitDiceNum := hitDiceMap[hitDice]
	return hitDiceNum
}

func setHitCount(modBody int) int {
	var hitCount int

	for _, char := range chars {
		for i := 1; i <= CharacterLevelGlobal; i++ {
			if char.ClassNameRU == ClassNameGlobalRu && i == 1 {
				hitCount = char.Hits.HitCount + modBody
			}
			if char.ClassNameRU == ClassNameGlobalRu && i != 1 {
				hitCount += getHitDiceAverageCount() + modBody
			}
		}
	}

	if races.RaceTypeGlobalRu == "Холмовой Дворф" {
		hitCount += 1
	}
	HitsCountGlobal = hitCount
	return hitCount
}

func setProficiencies() Proficiencies {
	for _, char := range chars {
		if char.ClassNameRU == ClassNameGlobalRu {
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
		if char.ClassNameRU == ClassNameGlobalRu {
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
			noArmor.ArmorClassCount = 10 + ModifierGlobal.Dexterity
		case "Чародей":
			noArmor.ArmorClassCount = 10 + ModifierGlobal.Dexterity
		}
		armorAnsList = append(armorAnsList, noArmor)
	}
	return armorAnsList
}

func setClassSkills() ([]string, int) {
	var skillsArray []string
	var randSkillCount int

	for _, char := range chars {
		if char.ClassNameRU == ClassNameGlobalRu {
			skillsArray = char.SkillsInDB.SkillsList
			randSkillCount = char.SkillsInDB.RandomCount
		}
	}

	return skillsArray, randSkillCount
}

func GetAnalyzedSkillSlice(backgroundSkills []string) []string {
	var skillMap = make(map[string]string)
	var skills []string

	raceSkills, raceSkillsCount := races.GetRaceSkill()
	classSkills, classSkillsCount := setClassSkills()

	totalSkillCount := classSkillsCount + raceSkillsCount + CountSkillsToAddToCharacter
	allSkillsSlice := append(classSkills, raceSkills...)

	for _, skillName := range allSkillsSlice {
		skillMap[skillName] = skillName
	} //да я так удаляю дубликаты навыков

	for _, backSkillName := range backgroundSkills {
		delete(skillMap, backSkillName)
	} //удаляю скиллы бекграунда чтобы они были добавлены потом и не дублировались

	var iter int
	for _, skl := range skillMap {
		iter++
		skills = append(skills, skl)
		if iter >= totalSkillCount {
			break
		}
	}
	skills = append(skills, backgroundSkills...)
	return skills
}

func setWeapons() []WeaponAnswer {
	var weaponList []WeaponAnswer
	for _, item := range equipmentList {
		if item.Type == "weapon" {
			switch item.ItemName {
			case "Простое оружие":
				if item.Count == 2 {
					for i := 0; i < 2; i++ {
						rollNum, _ := random.IntRange(0, 13)
						weaponData[rollNum].Count = 1
						weaponList = append(weaponList, weaponData[rollNum])
					}
					break
				} else {
					rollNum, _ := random.IntRange(0, 13)
					weaponData[rollNum].Count = item.Count
					weaponList = append(weaponList, weaponData[rollNum])
					break
				}
			case "Простое рукопашное оружие":
				if item.Count == 2 {
					for i := 0; i < 2; i++ {
						rollNum, _ := random.IntRange(0, 9)
						weaponData[rollNum].Count = 1
						weaponList = append(weaponList, weaponData[rollNum])
					}
					break
				} else {
					rollNum, _ := random.IntRange(0, 9)
					weaponData[rollNum].Count = item.Count
					weaponList = append(weaponList, weaponData[rollNum])
					break
				}
			case "Воинское оружие":
				if item.Count == 2 {
					for i := 0; i < 2; i++ {
						rollNum, _ := random.IntRange(14, 36)
						weaponData[rollNum].Count = 1
						weaponList = append(weaponList, weaponData[rollNum])
					}
					break
				} else {
					rollNum, _ := random.IntRange(14, 36)
					weaponData[rollNum].Count = item.Count
					weaponList = append(weaponList, weaponData[rollNum])
					break
				}
			case "Воинское рукопашное оружие":
				if item.Count == 2 {
					for i := 0; i < 2; i++ {
						rollNum, _ := random.IntRange(14, 31)
						weaponData[rollNum].Count = 1
						weaponList = append(weaponList, weaponData[rollNum])
					}
					break
				} else {
					rollNum, _ := random.IntRange(14, 13)
					weaponData[rollNum].Count = item.Count
					weaponList = append(weaponList, weaponData[rollNum])
					break
				}
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

	if races.RaceNameGlobalRu == "Драконорожденный" {
		dragonBreath := WeaponAnswer{WeaponName: "Дыхание Дракона", WeaponType: "weapon", Damage: races.DragonBornWeapon}
		weaponList = append(weaponList, dragonBreath)
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
