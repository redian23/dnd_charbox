package spells

import "go.mongodb.org/mongo-driver/bson/primitive"

type SpellsBSON struct {
	ID          primitive.ObjectID `bson:"_id"`
	Source      string             `bson:"source"`
	Classes     []string           `bson:"classes"`
	SpellLevel  int                `bson:"spellLevel"`
	SpellName   string             `bson:"spellNameEng"`
	SpellNameRu string             `bson:"spellNameRu"`
	URL         string             `bson:"url"`
}
