package core

import (
	"log"
	"pregen/pkg/backgrounds"
	"pregen/pkg/backgrounds/artist"
	"pregen/pkg/backgrounds/guildArtisan"
	"pregen/pkg/backgrounds/knight"
	"pregen/pkg/backgrounds/noble"
	"pregen/pkg/backgrounds/pirate"
	"pregen/pkg/backgrounds/sage"
	"pregen/pkg/backgrounds/sailor"
	"pregen/pkg/backgrounds/urchin"
	"pregen/pkg/classes"
	"pregen/pkg/classes/barbarian"
	"pregen/pkg/classes/bard"
	"pregen/pkg/classes/cleric"
	"pregen/pkg/races"
	"pregen/pkg/races/aasimar"
	"pregen/pkg/races/bugbear"
	"pregen/pkg/races/dragonborn"
	"pregen/pkg/races/gnome"
	"pregen/pkg/races/goblin"
	"pregen/pkg/races/goliath"
	"pregen/pkg/races/kenku"
	"pregen/pkg/races/owlin"
	"pregen/pkg/skills"
)

func GetFullCharacter(className, classArchetypeName, raceName, raceTypeName, backgroundName, gender string, level int) *Character {

	var raceInfo = getRace(raceName, raceTypeName, gender)
	var backgroundInfo = getBackground(backgroundName)
	var classInfo = getClass(className, classArchetypeName, raceInfo, backgroundInfo, level)

	return &Character{
		Level:         level,
		Experience:    getExpCount(level),
		PassiveWisdom: skills.GetPassive(),
		Background:    backgroundInfo,
		Race:          raceInfo,
		Class:         classInfo,
		Skills:        skills.SetSkillsForCharacter(raceInfo, backgroundInfo, classInfo, level),
		Langs:         getFullLanguagesList(raceInfo, backgroundInfo),
		SpellsList:    classInfo.SpellsList,
	}
	//marshaled, _ := json.MarshalIndent(characterBody, "", "   ")
	//fmt.Println(string(marshaled))
}

func getRace(raceName, raceTypeName, gender string) *races.Race {
	switch raceName {
	case "Аасимар":
		return aasimar.GetRace(raceTypeName, gender)
	case "Багбир":
		return bugbear.GetRace(gender)
	case "Гном":
		return gnome.GetRace(raceTypeName, gender)
	case "Гоблин":
		return goblin.GetRace(raceTypeName, gender)
	case "Голиаф":
		return goliath.GetRace(gender)
	case "Кенку":
		return kenku.GetRace(gender)
	case "Совлин":
		return owlin.GetRace(gender)
	case "Драконорожденный":
		return dragonborn.GetRace(raceTypeName, gender)
	}
	return nil
}

func getBackground(backgroundName string) *backgrounds.Background {
	switch backgroundName {
	case "Артист":
		return artist.GetBackground()
	case "Беспризорник":
		return urchin.GetBackground()
	case "Благородный":
		return noble.GetBackground()
	case "Гильдейский ремесленник":
		return guildArtisan.GetBackground()
	case "Моряк":
		return sailor.GetBackground()
	case "Мудрец":
		return sage.GetBackground()
	case "Пират":
		return pirate.GetBackground()
	case "Рыцарь":
		return knight.GetBackground()
	}
	return nil
}

func getClass(className, classArchetypeName string, raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) *classes.Class {
	if raceInfo == nil {
		log.Println("[ERROR] - Race Info is Null / or / RaceName have a mistake.")
	}
	if backgrInfo == nil {
		log.Println("[ERROR] - backgrInfo is Null / or / backgrInfo have a mistake.")
	}

	switch className {
	case "Бард":
		return bard.GetClass(raceInfo, backgrInfo, lvl, classArchetypeName)
	//case "Воин":
	//	return fighter.GetClass(raceInfo, backgrInfo, lvl, classArchetypeName)
	case "Варвар":
		return barbarian.GetClass(raceInfo, backgrInfo, lvl, classArchetypeName)
	case "Жрец":
		return cleric.GetClass(raceInfo, backgrInfo, lvl, classArchetypeName)
	}
	return nil
}

func getFullLanguagesList(raceInfo *races.Race, backgroundInfo *backgrounds.Background) []string {
	var langs []string
	langs = append(langs, raceInfo.Language...)
	langs = append(langs, backgroundInfo.Langs...)

	return langs
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
