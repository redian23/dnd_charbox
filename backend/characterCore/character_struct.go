package characterCore

import (
	"pregen/backend/backgr"
	"pregen/backend/classes"
)

type Character struct {
	CharacterName string            `json:"character_name"`
	Level         int               `json:"level"`
	Experience    int               `json:"experience"`
	Background    backgr.Background `json:"background"`
	Class         classes.Class     `json:"class"`
	Appearance    Appearance        `json:"appearance"`
}

type Appearance struct {
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Eyes   string `json:"eyes"`
	Height int    `json:"height"`
	Skin   string `json:"skin"`
	Weight int    `json:"weight"`
	Hair   string `json:"hair"`
}
