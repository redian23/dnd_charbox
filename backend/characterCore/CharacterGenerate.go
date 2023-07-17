package characterCore

import (
	"pregen/backend/backgr"
	"pregen/backend/classes"
	"pregen/backend/races"
)

func GenerateFullCharacter() Character {
	var char Character

	char.Level = 1
	char.Experience = 0
	char.Race = races.GenerateRace()
	char.Class = classes.GenerateClass()
	char.Background = backgr.GenerateBackground(char.Class.ClassName)
	char.Class.Skills = classes.SetSkillsForCharacter(char.Background.SkillMastery)

	return char
}
