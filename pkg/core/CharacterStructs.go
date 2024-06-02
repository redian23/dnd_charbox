package core

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/skills"
	"pregen/pkg/spells"
)

type Character struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`

	Background    *backgrounds.Background `json:"background"`
	Race          *races.Race             `json:"race"`
	Class         *classes.Class          `json:"class"`
	Skills        *skills.Skills          `json:"skills"`
	PassiveWisdom *skills.Passive         `json:"passive_info"`
	Langs         []string                `json:"langs"`
	SpellsList    []spells.SpellsJSON     `json:"spells"`
}
