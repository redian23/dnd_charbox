package core

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
	"pregen/pkg/skills"
	"pregen/pkg/spells"
)

type Character struct {
	Level            int                     `json:"level"`
	Experience       int                     `json:"experience"`
	PassiveWisdom    int                     `json:"passive_wisdom"`
	ProficiencyBonus int                     `json:"proficiency_bonus"`
	Background       *backgrounds.Background `json:"background"`
	Race             *races.Race             `json:"race"`
	Class            *classes.Class          `json:"class"`
	Skills           *skills.Skills          `json:"skills"`
	Langs            []string                `json:"langs"`
	SpellsList       []spells.SpellsJSON     `json:"spells"`
}
