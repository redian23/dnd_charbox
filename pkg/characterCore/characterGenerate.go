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
	}
}
