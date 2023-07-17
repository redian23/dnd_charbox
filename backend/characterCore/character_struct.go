package characterCore

import (
	"pregen/backend/backgr"
	"pregen/backend/classes"
	"pregen/backend/races"
)

type Character struct {
	CharacterName string                  `json:"character_name"`
	Level         int                     `json:"level"`
	Experience    int                     `json:"experience"`
	Background    backgr.BackgroundAnswer `json:"background"`
	Race          races.RacesAnswer       `json:"race"`
	Class         classes.ClassAnswer     `json:"class"`
}
