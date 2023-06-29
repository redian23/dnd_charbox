package classes

type CharacteristicsForClass struct {
	Data []struct {
		ClassName string     `json:"className"`
		CharReq   [][]string `json:"charReq"`
	} `json:"data"`
}
