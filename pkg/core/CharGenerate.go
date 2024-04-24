package core

import (
	"encoding/json"
	"fmt"
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
	"pregen/pkg/classes/bard"
	"pregen/pkg/races"
	"pregen/pkg/races/aasimar"
	"pregen/pkg/races/bugbeat"
	"pregen/pkg/races/gnome"
	"pregen/pkg/races/goblin"
	"pregen/pkg/races/goliath"
	"pregen/pkg/skills"
)

func GetFullCharacter() {
	// Будущие аргументы функции GetFullCharacter
	var level = 8
	var raceName = "Багбир"
	var raceTypeName = "Аасимар–падший"
	var gender = "male"
	var backgroundName = "Артист"
	var className = "Бард"

	var raceInfo = getRace(raceName, raceTypeName, gender)
	var backgroundInfo = getBackground(backgroundName)
	var classInfo = getClass(className, raceInfo, backgroundInfo, level)

	var characterBody = Character{
		Level:            level,
		Experience:       getExpCount(level),
		PassiveWisdom:    skills.PassiveWisdom,
		ProficiencyBonus: 0,
		Background:       backgroundInfo,
		Race:             raceInfo,
		Class:            classInfo,
		Skills:           skills.SetSkillsForCharacter(raceInfo, backgroundInfo, classInfo, level),
		Langs:            getFullLanguagesList(raceInfo, backgroundInfo),
		SpellsList:       classInfo.SpellsList,
	}
	marshaled, _ := json.MarshalIndent(characterBody, "", "   ")
	fmt.Println(string(marshaled))
}

func getRace(raceName, raceTypeName, gender string) *races.Race {
	switch raceName {
	case "Аасимар":
		return aasimar.GetRace(raceTypeName, gender)
	case "Багбир":
		return bugbeat.GetRace(gender)
	case "Гном":
		return gnome.GetGnomeRace(raceTypeName, gender)
	case "Гоблин":
		return goblin.GetRace(raceTypeName, gender)
	case "Голиаф":
		return goliath.GetRace(raceTypeName, gender)
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

func getClass(className string, raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) *classes.Class {
	if raceInfo == nil {
		log.Fatalln("[ERROR] - Race Info is Null / or / RaceName have a mistake.")
	}

	switch className {
	case "Бард":
		return bard.GetBardClass(raceInfo, backgrInfo, lvl)
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
