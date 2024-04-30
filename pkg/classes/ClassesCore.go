package classes

import (
	"github.com/mazen160/go-random"
	"math"
)

func GetClassSkillsArray(raceSkills, backSkills []string, availableSkillList []string, count int) []string {

	// Создаем карту для быстрой проверки навыков от рассы и предыстории
	raceSkillsMap := make(map[string]bool)
	backSkillsMap := make(map[string]bool)
	for _, skill := range raceSkills {
		raceSkillsMap[skill] = true
	}
	for _, skill := range backSkills {
		backSkillsMap[skill] = true
	}

	// Выбираем 4 навыка из общего списка, которые не совпадают с навыками от рассы и предыстории
	var chosenSkills []string
	for len(chosenSkills) < count {
		randomIndex, _ := random.IntRange(0, len(availableSkillList))

		skill := availableSkillList[randomIndex]
		// Проверяем, что навык не входит в навыки от рассы и предыстории
		if !raceSkillsMap[skill] && !backSkillsMap[skill] {
			chosenSkills = append(chosenSkills, skill) // Добавляем навык в выбранные, если он подходит
		}
	}
	return chosenSkills
}

func GetModifiersForClass(ab AbilityScore) AbilityModifier {

	abilitiesArray := []int{ab.Strength, ab.Dexterity, ab.BodyDifficulty,
		ab.Intelligence, ab.Wisdom, ab.Charisma}

	var modifier AbilityModifier
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
	return modifier
}

func GetSaveThrowsForClass(mod AbilityModifier, classSavingThrow []string) *SavingThrows {
	var modifierMap = map[string]int{
		"Strength":       mod.Strength,
		"Dexterity":      mod.Dexterity,
		"Intelligence":   mod.Intelligence,
		"BodyDifficulty": mod.BodyDifficulty,
		"Wisdom":         mod.Wisdom,
		"Charisma":       mod.Charisma,
	}

	var saveTh SavingThrows

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

	for _, stat := range classSavingThrow {
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
	return &saveTh
}
