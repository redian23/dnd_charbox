package spells

import "go.mongodb.org/mongo-driver/bson/primitive"

type SpellsBSON struct {
	ID          primitive.ObjectID `json:"_id"`
	Source      string             `json:"source"`
	Classes     []string           `json:"classes"`
	SpellLevel  int                `json:"spellLevel"`
	SpellName   string             `json:"spellNameEng"`
	SpellNameRu string             `json:"spellNameRu"`
	URL         string             `json:"url"`
}
