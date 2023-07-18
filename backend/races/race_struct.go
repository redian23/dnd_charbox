package races

import "go.mongodb.org/mongo-driver/bson/primitive"

type RacesAnswer struct {
	RaceName   string   `json:"race_name"`
	RaceNameRu string   `json:"race_name_ru"`
	Gender     string   `json:"gender"`
	Type       string   `json:"race_type_name"`
	TypeRu     string   `json:"race_type_name_ru"`
	Age        int      `json:"age"`
	Height     int      `json:"height"`
	Weight     int      `json:"weight"`
	BodySize   string   `json:"body_size"`
	Eyes       string   `json:"eyes"`
	Hair       string   `json:"hair"`
	Speed      int      `json:"speed"`
	Langs      []string `json:"langs"`
	RaceSkill  string   `json:"race_skill"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Resist     []string `json:"resist"`
}

type RacesBSON struct {
	ID         primitive.ObjectID `bson:"_id"`
	RaceName   string             `json:"race_name"`
	RaceNameRu string             `json:"race_name_ru"`
	Type       []Type             `json:"type"`
	MinAge     int                `json:"min_age"`
	MaxAge     int                `json:"max_age"`
	MinHeight  int                `json:"min_height"`
	MaxHeight  int                `json:"max_height,omitempty"`
	MinWeight  int                `json:"min_weight"`
	MaxWeight  int                `json:"max_weight,omitempty"`
	BodySize   string             `json:"body_size"`
	Speed      int                `json:"speed"`
	Langs      []string           `json:"langs"`
	RaceSkill  []string           `json:"race_skill"`
	Names      Names              `json:"names"`
	LastNames  []interface{}      `json:"last_names"`
	Resist     []string           `json:"resist"`
}

type RacesJsonStruct []struct {
	RaceName   string        `json:"race_name"`
	RaceNameRu string        `json:"race_name_ru"`
	Type       []Type        `json:"type"`
	MinAge     int           `json:"min_age"`
	MaxAge     int           `json:"max_age"`
	MinHeight  int           `json:"min_height"`
	MaxHeight  int           `json:"max_height,omitempty"`
	MinWeight  int           `json:"min_weight"`
	MaxWeight  int           `json:"max_weight,omitempty"`
	Speed      int           `json:"speed"`
	Langs      []string      `json:"langs"`
	RaceSkill  []interface{} `json:"race_skill"`
	BodySize   string        `json:"body_size"`
	Names      Names         `json:"names"`
	LastNames  []interface{} `json:"last_names"`
	Resist     []string      `json:"resist"`
}
type StatsUp struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
}
type Type struct {
	TypeRaceName   string  `json:"type_race_name"`
	TypeRaceNameRu string  `json:"type_race_name_ru"`
	StatsUp        StatsUp `json:"stats_up"`
}
type Names struct {
	Man   []string `json:"man"`
	Woman []string `json:"woman"`
}
