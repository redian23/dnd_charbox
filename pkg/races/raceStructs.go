package races

import "go.mongodb.org/mongo-driver/bson/primitive"

type RacesAnswer struct {
	CharacterName string        `json:"character_name"`
	RaceName      string        `json:"race_name"`
	RaceNameRu    string        `json:"race_name_ru"`
	Gender        string        `json:"gender"`
	Type          string        `json:"race_type_name"`
	TypeRu        string        `json:"race_type_name_ru"`
	Age           int           `json:"age"`
	Height        int           `json:"height"`
	Weight        int           `json:"weight"`
	HeightFt      string        `json:"height_ft"`
	WeightLb      int           `json:"weight_lb"`
	BodySize      string        `json:"body_size"`
	Eyes          string        `json:"eyes"`
	Hair          string        `json:"hair"`
	Speed         int           `json:"speed"`
	Langs         []string      `json:"langs"`
	RaceSkill     string        `json:"race_skill"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	Resist        string        `json:"resist"`
	Darkvision    bool          `json:"darkvision"`
	Other         other         `json:"other"`
	RaceAbilities []raceAbility `json:"race_abilities"`
	RacePhoto     racePhoto     `json:"race_photo"`
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
	Names      names              `bson:"names"`
	LastNames  []interface{}      `bson:"last_names"`
	Resist     string             `bson:"resist"`
	Darkvision bool               `json:"darkvision"`
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
	Names      names         `json:"names"`
	LastNames  []interface{} `json:"last_names"`
	Resist     []string      `json:"resist"`
}

type raceAbility struct {
	AbilityName string `bson:"ability_name"`
	Description string `bson:"description"`
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
	TypeRaceName   string        `json:"type_race_name"`
	TypeRaceNameRu string        `json:"type_race_name_ru"`
	StatsUp        StatsUp       `json:"stats_up"`
	RaceAbilities  []raceAbility `bson:"raceAbilities"`
}
type names struct {
	Man   []string `json:"man"`
	Woman []string `json:"woman"`
}

type other struct {
	DragonType       dragonType       `json:"dragon_type"`
	YuantiAppearance appearanceYuanti `json:"yuanti_appearance"`
}

type racePhoto struct {
	Path     string `json:"path"`
	FileName string `json:"file_name"`
}
