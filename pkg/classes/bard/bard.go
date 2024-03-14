package bard

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/classes"
	"pregen/pkg/core"
)

var musicalInstruments = []string{
	"Барабан", "Виола", "Волынка", "Лира", "Лютня",
	"Рожок", "Свирель", "Флейта", "Цимбалы", "Шалмей",
}

var lvl = classes.Level

func getBardHits(lvl int) classes.Hits {
	var hitCount int
	var D6 = core.D6

	for i := 1; i < lvl; i++ {
		if i == 1 {
			hitCount = D6.GetMaxRange()
		} else {
			hitCount += D6.RollDice()
		}
	}

	return classes.Hits{
		HitDice:  D6.GetDiceName(),
		HitCount: hitCount,
	}
}

func getBardProficiencies() classes.Proficiencies {
	var instrumentsArray []string
	var iter int
	for _, value := range musicalInstruments {
		instrumentsArray = append(instrumentsArray, value)
		iter++
		if iter == 3 {
			break
		}
	}
	//var skill_list = []string{"Acrobatics", "Animal Handling", "Arcana", "Athletics", "Deception", "History", "Insight", "Intimidation", "Investigation", "Medicine", "Nature", "Perception", "Performance", "Persuasion", "Religion", "Sleight Of Hand", "Stealth", "Survival"}

	return classes.Proficiencies{
		Weapons:       []string{"Лёгкие доспехи"},
		Armor:         []string{"Простое оружие", "Длинные мечи", "Короткие мечи", "Рапиры", "Ручные арбалеты"},
		Tools:         instrumentsArray,
		SavingThrow:   []string{"Dexterity", "Charisma"},
		SkillsOfClass: []string{},
	}
	// Навыки: Выберите три любых
	// прощитываеть не в моменте, а ждать когда будут известны навыки от рассы и предыстории
}

func getBardEquipment() []classes.Item {
	var equipment = []classes.Item{
		{
			Name:     "Кожаный доспех",
			ItemType: "Armor",
			Count:    1,
		},
		{
			Name:     "Кинжал",
			ItemType: "Weapon",
			Count:    1,
		},
	}

	variant, _ := random.IntRange(0, 1)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Набор дипломата",
				ItemType: "Tool",
				Count:    1,
			})
	case 1:
		equipment = append(equipment,
			classes.Item{
				Name:     "Набор артиста",
				ItemType: "Tool",
				Count:    1,
			})
	}

	variant, _ = random.IntRange(0, 2)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Рапира",
				ItemType: "Weapon",
				Count:    1,
			})
	case 1:
		equipment = append(equipment,
			classes.Item{
				Name:     "Длинный меч",
				ItemType: "Weapon",
				Count:    1,
			})
	case 2:
		var weaponsList []string
		for _, weapon := range classes.WeaponData {
			if weapon.WeaponType == "Простое рукопашное оружие" ||
				weapon.WeaponType == "Простое дальнобойное оружие" ||
				weapon.Name != "Рапира" && weapon.Name != "Длинный меч" {
				weaponsList = append(weaponsList, weapon.Name)
			}
		}
		randNum, _ := random.IntRange(0, len(weaponsList))

		equipment = append(equipment,
			classes.Item{
				Name:     weaponsList[randNum],
				ItemType: "Weapon",
				Count:    1,
			})
	}

	randNum, _ := random.IntRange(0, len(musicalInstruments))
	equipment = append(equipment,
		classes.Item{
			Name:     musicalInstruments[randNum],
			ItemType: "Tool",
			Count:    1,
		})

	return equipment
}

func getBardAbilitiesScore() classes.AbilityScore {
	var classAbilityPriority = []string{"Charisma", "Dexterity"}

	return core.GetClassAbilitiesScore(classAbilityPriority)
}

func GetBardClass() classes.Class {

	return classes.Class{
		ClassName:   "Bard",
		ClassNameRU: "Бард",
		Hits:        getBardHits(lvl),
		SpellCasting: classes.GetClassSpellBasicCharacteristic(
			"Бард", lvl, classes.AbilityModifier{Strength: 0},
		),
		Proficiencies: getBardProficiencies(),
		Equipment:     getBardEquipment(),
		AbilityScore:  getBardAbilitiesScore(),
	}
}
