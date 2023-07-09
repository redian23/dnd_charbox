package classes

import "go.mongodb.org/mongo-driver/bson/primitive"

type RaceCharacteristicsBSON []struct {
	ID          primitive.ObjectID `bson:"_id"`
	CharReq     [][]string         `json:"charReq"`
	ClassName   string             `json:"className"`
	ClassNameRU string             `json:"classNameRU"`
	Background  []string           `json:"background"`
}
