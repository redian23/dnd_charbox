package main

import (
	"pregen/pkg/core"
	"pregen/pkg/races"
	"pregen/pkg/races/Aasimar"
)

func SetRaceInfo() {
	var raceTypeName = "Аасимар–падший"
	var gender = "male"
	races.RaceInfo = Aasimar.GetRace(raceTypeName, gender)
}
func main() {
	core.GetFullCharacter()
}
