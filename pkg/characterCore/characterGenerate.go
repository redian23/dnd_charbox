package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
	"pregen/pkg/spells"
)

func GenerateFullCharacter(className, raceName string, lvl int) *Character {
	classes.LanguageListGlobal = []string{} //очистка глоб переменной перед генерацией

	classes.CharacterLevelGlobal = lvl
	return &Character{
		Level:            lvl,
		Experience:       getExpCount(lvl),
		ProficiencyBonus: getProfBonus(lvl),
		Race:             races.GenerateRaceForCharacter(raceName),
		Class:            classes.GenerateClass(className),
		Background:       backgrounds.GenerateBackground(),
		Skills:           SetSkillsForCharacter(),
		PassiveWisdom:    PassWisdom,
		Langs:            GetCharLangs(),
		SpellsList: spellsList{
			ZeroLevelSpells: spells.GetSpellsForCharacter(0),
			OneLevelSpells:  spells.GetSpellsForCharacter(1),
			TwoLevelSpells:  spells.GetSpellsForCharacter(2),
			TreeLevelSpells: spells.GetSpellsForCharacter(3),
			FourLevelSpells: spells.GetSpellsForCharacter(4),
		},
	}
}
