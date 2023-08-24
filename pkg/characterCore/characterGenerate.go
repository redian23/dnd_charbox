package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

var LastCharacter Character

func GenerateFullCharacter() *Character {

	LastCharacter = Character{
		Level:         1,
		Experience:    0,
		Race:          races.GenerateRaceForCharacter(),
		Class:         classes.GenerateClass(),
		Background:    backgrounds.GenerateBackground(),
		Skills:        SetSkillsForCharacter(),
		PassiveWisdom: PassWisdom,
	}
	return &LastCharacter
}
