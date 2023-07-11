package characterCore

import (
	"pregen/backend/backgr"
	"pregen/backend/classes"
)

func StartCharacterGenerate() Character {
	var char Character

	char.Level = 1
	char.Experience = 0
	char.Class = classes.GenerateClass()
	char.Appearance = GenerateAppearance()
	char.Background = backgr.GenerateBackground(char.Class.ClassName)
	char.Class.Skills = classes.SetSkillsForClass(char.Background.SkillMastery, char.Class.SkillsOfClass, char.Class.Modifier)

	return char
}
