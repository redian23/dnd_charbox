package general

import (
	"fmt"
	"pregen/pkg/backrounds"
	"pregen/pkg/classes"
	"pregen/pkg/classes/bard"
	"pregen/pkg/races"
	"pregen/pkg/races/rcs"
	"pregen/pkg/skills"
)

func GetFullCharacter() {
	var raceTypeName = "Аасимар–падший"
	var gender = "male"
	var aasimar = rcs.GetAasimarRace(raceTypeName, gender)

	var artist = backrounds.GetArtistBackground()

	races.RaceInfo = *aasimar
	backrounds.BackgroundInfo = *artist

	classes.ClassInfo = bard.GetBardClass()

	var skills = skills.SetSkillsForCharacter()
	fmt.Println(skills)
	//marshaled, _ := json.MarshalIndent(classes.ClassInfo, "", "   ")
	//fmt.Println(string(marshaled))
}
