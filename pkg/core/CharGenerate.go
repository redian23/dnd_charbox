package core

import (
	"encoding/json"
	"fmt"
	"pregen/pkg/backgrounds/artist"
	"pregen/pkg/classes/bard"
	"pregen/pkg/general"
	"pregen/pkg/races/aasimar"
	"pregen/pkg/skills"
)

func GetFullCharacter() {
	var raceTypeName = "Аасимар–падший"
	var gender = "male"

	general.GlobalCharLevel = 1

	var raceInfo = aasimar.GetRace(raceTypeName, gender)
	var backgroundInfo = artist.GetArtistBackground()
	var classInfo = bard.GetBardClass(raceInfo)

	var characterBody = Character{
		Level:            general.GlobalCharLevel,
		Experience:       getExpCount(general.GlobalCharLevel),
		PassiveWisdom:    skills.PassiveWisdom,
		ProficiencyBonus: 0,
		Background:       backgroundInfo,
		Race:             raceInfo,
		Class:            classInfo,
		Skills:           skills.SetSkillsForCharacter(raceInfo, backgroundInfo, classInfo),
		Langs:            raceInfo.Language,
		SpellsList:       classInfo.SpellsList,
	}
	marshaled, _ := json.MarshalIndent(characterBody, "", "   ")
	fmt.Println(string(marshaled))
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
