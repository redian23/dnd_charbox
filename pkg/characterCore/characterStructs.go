package characterCore

import (
	"pregen/pkg/backgr"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

type Character struct {
	CharacterName string                   `json:"character_name"`
	Level         int                      `json:"level"`
	Experience    int                      `json:"experience"`
	Background    *backgr.BackgroundAnswer `json:"background"`
	Race          *races.RacesAnswer       `json:"race"`
	Class         *classes.ClassAnswer     `json:"class"`
	Skills        Skills                   `json:"skills"`
}
