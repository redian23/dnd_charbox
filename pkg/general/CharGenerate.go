package general

import (
	"encoding/json"
	"fmt"
	"pregen/pkg/classes"
	"pregen/pkg/classes/bard"
	"pregen/pkg/races"
	"pregen/pkg/races/rcs"
)

func GetFullCharacter() {
	var raceTypeName = "Аасимар–падший"
	var gender = "male"
	var aasimar = rcs.GetAasimarRace(raceTypeName, gender)

	races.RaceInfo = *aasimar

	classes.ClassInfo = bard.GetBardClass()
	marshaled, _ := json.MarshalIndent(classes.ClassInfo, "", "   ")
	fmt.Println(string(marshaled))
}
