package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type RaceStruct struct {
	RaceID         int             `json:"race_id"`
	RaceName       string          `json:"race_name"`
	Worldview      string          `json:"worldview"`
	MaxAge         int             `json:"max_age"`
	Speed          int             `json:"speed"`
	BodySizeInc    int             `json:"body_size_inc"`
	BodySizeMetric int             `json:"body_size_metric"`
	Lang           []string        `json:"lang"`
	RaceAbility    []raceAbilities `json:"race_ability"`
	RaceStatsUp    raceStatsUp     `json:"race_stats_up"`
}
type raceAbilities struct {
	RaceName    string `json:"race_name"`
	RaceAbility string `json:"race_ability"`
	Description string `json:"description"`
}

type raceStatsUp struct {
	RaceID         int    `json:"race_id"`
	RaceName       string `json:"race_name"`
	Strength       int    `json:"strength"`
	Dexterity      int    `json:"dexterity"`
	Intelligence   int    `json:"intelligence"`
	BodyDifficulty int    `json:"bodydifficulty"`
	Wisdom         int    `json:"wisdom"`
	Charisma       int    `json:"charisma"`
}

type RacePostgresTable struct {
	ID   string
	Data Data
}

type Data RaceStruct

func (d Data) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *Data) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &d)
}
