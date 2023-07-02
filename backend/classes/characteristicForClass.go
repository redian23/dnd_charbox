package classes

type CharacteristicsForClass struct {
	Data []struct {
		CharReq    [][]string `json:"charReq"`
		ClassName  string     `json:"className"`
		Background []string   `json:"background"`
	} `json:"data"`
}
