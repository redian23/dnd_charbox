package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
	"pregen/pkg/spells"
)

func GenerateFullCharacter(className, raceName string, lvl int, gender string) *Character {
	classes.LanguageListGlobal = []string{} //очистка глоб переменной перед генерацией
	classes.CharacterLevelGlobal = lvl
	return &Character{
		Level:            lvl,
		Experience:       getExpCount(lvl),
		ProficiencyBonus: getProfBonus(lvl),
		Race:             races.GenerateRaceForCharacter(raceName, gender),
		Class:            classes.GenerateClass(className),
		Background:       backgrounds.GenerateBackground(),
		Skills:           classes.SetSkillsForCharacter(backgrounds.BackgroundSkills),
		PassiveWisdom:    classes.PassWisdom,
		Langs:            GetCharLangs(),
		SpellsList: spellsList{
			ZeroLevelSpells: spells.SetSpellsForCharacter(0),
			OneLevelSpells:  spells.SetSpellsForCharacter(1),
			TwoLevelSpells:  spells.SetSpellsForCharacter(2),
			TreeLevelSpells: spells.SetSpellsForCharacter(3),
			FourLevelSpells: spells.SetSpellsForCharacter(4),
		},
	}
}
