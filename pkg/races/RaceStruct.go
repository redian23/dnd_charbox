package races

import "pregen/pkg/spells"

type Race struct {
	RaceName             string      `json:"race_name"`
	RaceNameRu           string      `json:"race_name_ru"`
	Gender               string      `json:"gender"`
	Type                 RaceType    `json:"race_type_name"`
	Body                 *BodyAnswer `json:"body"`
	AddictionInformation string      `json:"addiction_information"`
	RaceSkill            []string    `json:"race_skill"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	Resist               []string    `json:"resist"`
	Language             []string    `json:"language"`
	RacePhoto            racePhoto   `json:"race_photo"`
}

type BodyStats struct {
	BodySize  string `json:"body_size"`
	MinAge    int    `json:"min_age"`
	MaxAge    int    `json:"max_age"`
	MaxHeight int    `json:"max_height"`
	MaxWeight int    `json:"max_weight"`
	MinHeight int    `json:"min_height"`
	MinWeight int    `json:"min_weight"`
}

type BodyAnswer struct {
	Age      int    `json:"age"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
	HeightFt string `json:"height_ft"`
	WeightLb int    `json:"weight_lb"`
	BodySize string `json:"body_size"`
	Eyes     string `json:"eyes"`
	Hair     string `json:"hair"`
}

type racePhoto struct {
	Path     string `json:"path"`
	FileName string `json:"file_name"`
}

type RaceType struct {
	TypeRaceName     string              `json:"type_race_name"`
	TypeRaceNameRu   string              `json:"type_race_name_ru"`
	AbilityScorePlus map[string]int      `json:"stats_up"`
	RaceAbilities    map[string]string   `json:"raceAbilities"`
	Speed            int                 `json:"speed"`
	RaceSpellsList   []spells.SpellsJSON `json:"race_spells_list"`
}
