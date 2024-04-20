package core

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/classes/bard"
)

func GetFullCharacter() {
	backgrounds.BackgroundInfo = backgrounds.GetArtistBackground()
	classes.ClassInfo = bard.GetBardClass()
	//
	//var characterBody = Character{
	//	Level:            general.GlobalCharLevel,
	//	Experience:       getExpCount(general.GlobalCharLevel),
	//	PassiveWisdom:    skills.PassiveWisdom,
	//	ProficiencyBonus: 0,
	//	Background:       backgrounds.BackgroundInfo,
	//	Race:             general.RaceInfo,
	//	Class:            classes.ClassInfo,
	//	Skills:           skills.SetSkillsForCharacter(),
	//	Langs:            general.RaceInfo.Language,
	//	SpellsList:       classes.ClassInfo.SpellsList,
	//}
	//marshaled, _ := json.MarshalIndent(characterBody, "", "   ")
	//fmt.Println(string(marshaled))
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
