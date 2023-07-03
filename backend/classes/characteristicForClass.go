package classes

type CharacteristicsForClass []struct {
	CharReq     [][]string `json:"charReq"`
	ClassName   string     `json:"className"`
	ClassNameRU string     `json:"classNameRU"`
	Background  []string   `json:"background"`
}
