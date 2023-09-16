package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

func GenerateFullCharacter(className, raceName string) *Character {
	return &Character{
		Level:         1,
		Experience:    0,
		Race:          races.GenerateRaceForCharacter(raceName),
		Class:         classes.GenerateClass(className),
		Background:    backgrounds.GenerateBackground(),
		Skills:        SetSkillsForCharacter(),
		PassiveWisdom: PassWisdom,
		Langs:         GetCharLangs(),
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
		if race.RaceName == races.RaceNameGlobal {
			raceLangs = race.RaceLanguages
		}
	}

	if races.RaceTypeGlobalRu == "Человек" || races.RaceTypeGlobalRu == "Высший Эльф" ||
		races.RaceTypeGlobalRu == "Табакси" || races.RaceNameGlobal == "Полуэльф" {
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
	return raceLangs
}
