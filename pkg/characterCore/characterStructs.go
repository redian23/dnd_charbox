package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

type Character struct {
	Level            int                           `json:"level"`
	Experience       int                           `json:"experience"`
	PassiveWisdom    int                           `json:"passive_wisdom"`
	ProficiencyBonus int                           `json:"proficiency_bonus"`
	Background       *backgrounds.BackgroundAnswer `json:"background"`
	Race             *races.RacesAnswer            `json:"race"`
	Class            *classes.ClassAnswer          `json:"class"`
	Skills           Skills                        `json:"skills"`
	Langs            []string                      `json:"langs"`
	SpellsList       spellsList                    `json:"spells"`
}

type spellsList struct {
	ZeroLevelSpells []string `json:"zero_level_spells"`
	OneLevelSpells  []string `json:"one_level_spells"`
	TwoLevelSpells  []string `json:"two_level_spells"`
	TreeLevelSpells []string `json:"tree_level_spells"`
	FourLevelSpells []string `json:"four_level_spells"`
	FiveLevelSpells []string `json:"five_level_spells"`
}
