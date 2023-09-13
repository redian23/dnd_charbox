package spells

import "go.mongodb.org/mongo-driver/bson/primitive"

type SpellsBSON struct {
	ID          primitive.ObjectID `bson:"_id"`
	Archetypes  []interface{}      `json:"archetypes"`
	Classes     []string           `json:"classes"`
	Components  string             `json:"components"`
	Description string             `json:"description"`
	Duration    string             `json:"duration"`
	OverlapTime string             `json:"overlapTime"`
	Range       string             `json:"range"`
	Source      string             `json:"source"`
	SpellName   string             `json:"spellName"`
	SpellNameRu string             `json:"spellNameRu"`
	Type        string             `json:"type"`
}
