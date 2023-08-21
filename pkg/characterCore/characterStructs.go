package characterCore

import (
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/races"
)

type Character struct {
	Level      int                           `json:"level"`
	Experience int                           `json:"experience"`
	Background *backgrounds.BackgroundAnswer `json:"background"`
	Race       *races.RacesAnswer            `json:"race"`
	Class      *classes.ClassAnswer          `json:"class"`
	Skills     Skills                        `json:"skills"`
}
