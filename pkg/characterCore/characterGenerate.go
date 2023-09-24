package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
	"pregen/pkg/spells"
)

func GenerateFullCharacter(className, raceName string, lvl int) *Character {
	classes.LanguageListGlobal = []string{}

	classes.CharacterLevelGlobal = lvl
	return &Character{
		Level:            lvl,
		Experience:       getExpCount(lvl),
		ProficiencyBonus: getProfBonus(lvl),
		Race:             races.GenerateRaceForCharacter(raceName),
		Class:            classes.GenerateClass(className, lvl),
		Background:       backgrounds.GenerateBackground(),
		Skills:           SetSkillsForCharacter(),
		PassiveWisdom:    PassWisdom,
		Langs:            GetCharLangs(),
		SpellsList: spellsList{
			ZeroLevelSpells: spells.GetSpellsZeroLevelForCharacter(),
			OneLevelSpells:  spells.GetSpellsOneLevelForCharacter(),
		},
	}
}

func getExpCount(lvl int) int {

	var expCountMap = map[int]int{
		1: 0,
		2: 300,
		3: 900,
		4: 2700,
		5: 6500,
		6: 14000,
		7: 23000,
		8: 34000,
	}
	return expCountMap[lvl]
}

func getProfBonus(lvl int) int {
	if lvl <= 4 {
		return 2
	}
	if lvl >= 5 && lvl <= 8 {
		return 3
	} else {
		return 4
	}
}

func GetCharLangs() []string {
	var languageMap = map[string]string{
		"Великаний": "Великаний", "Гномий": "Гномий", "Гоблинский": "Гоблинский",
		"Дварфский": "Дварфский", "Орочий": "Орочий", "Полуросликов": "Полуросликов",
		"Эльфийский": "Эльфийский", "Язык Бездны": "Язык Бездны", "Глубинная Речь": "Глубинная Речь",
		"Драконий": "Драконий", "Инфернальный": "Инфернальный", "Небесный": "Небесный",
		"Первичный": "Первичный", "Подземный": "Подземный", "Сильван": "Сильван",
	}

	var raceLangs []string

	var addictionLangsCount int
	for _, race := range races.RaceData {
		if race.RaceNameRu == races.RaceNameGlobalRu {
			raceLangs = race.RaceLanguages
		}
	}

	if races.RaceTypeGlobalRu == "Человек" || races.RaceTypeGlobalRu == "Высший Эльф" ||
		races.RaceTypeGlobalRu == "Табакси" || races.RaceNameGlobalRu == "Полуэльф" {
		addictionLangsCount += 1
	}
	if backgrounds.BackgroundNameRuGlobal == "Отшельник" || backgrounds.BackgroundNameRuGlobal == "Благородный" ||
		backgrounds.BackgroundNameRuGlobal == "Рыцарь" || backgrounds.BackgroundNameRuGlobal == "Гильдейский ремесленник" ||
		backgrounds.BackgroundNameRuGlobal == "Чужеземец" {
		addictionLangsCount += 1
	}
	if backgrounds.BackgroundNameRuGlobal == "Мудрец" || backgrounds.BackgroundNameRuGlobal == "Прислужник" {
		addictionLangsCount += 2
	}

	for _, lang := range raceLangs {
		delete(languageMap, lang)
	} // удаляю языки от рассы из списка языков, чтобы небыло дублей

	var iter int
	for _, lng := range languageMap {
		if iter >= addictionLangsCount {
			break
		}
		iter++
		raceLangs = append(raceLangs, lng)
	}

	raceLangs = append(raceLangs, classes.LanguageListGlobal...)

	return raceLangs
}
