package characterCore

import (
	"pregen/pkg/backgr"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

func GenerateFullCharacter() *Character {

	return &Character{
		Level:      1,
		Experience: 0,
		Race:       races.GenerateRaceForCharacter(),
		Class:      classes.GenerateClass(),
		Background: backgr.GenerateBackground(),
		Skills:     SetSkillsForCharacter(),
	}
}
