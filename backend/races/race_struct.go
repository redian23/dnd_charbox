package races

type RaceStruct []struct {
	Lang        []interface{} `json:"lang"`
	Speed       int           `json:"speed"`
	MaxAge      int           `json:"max_age"`
	RaceID      int           `json:"race_id"`
	RaceName    string        `json:"race_name"`
	Worldview   string        `json:"worldview"`
	RaceAbility []interface{} `json:"race_ability"`
	BodySizeInc int           `json:"body_size_inc"`
	RaceStatsUp struct {
		Wisdom         int    `json:"wisdom"`
		RaceID         int    `json:"race_id"`
		Charisma       int    `json:"charisma"`
		Strength       int    `json:"strength"`
		Dexterity      int    `json:"dexterity"`
		RaceName       string `json:"race_name"`
		Intelligence   int    `json:"intelligence"`
		Bodydifficulty int    `json:"bodydifficulty"`
	} `json:"race_stats_up"`
	BodySizeMetric int `json:"body_size_metric"`
}
