package backgr

import "go.mongodb.org/mongo-driver/bson/primitive"

type Type struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type CharacterTrait struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type Value struct {
	Worldview   string `json:"worldview"`
	WorldviewRu string `json:"worldview_ru"`
	Text        string `json:"text"`
}
type Ideal struct {
	Dice  string  `json:"dice"`
	Value []Value `json:"value"`
}
type Affection struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type Weakness struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}

type BackgroundJson struct {
	BackgroundName   string         `json:"background_name"`
	BackgroundNameRu string         `json:"background_name_ru"`
	Description      string         `json:"description"`
	Personalization  string         `json:"personalization"`
	Advice           string         `json:"advice"`
	SkillMastery     []string       `json:"skill_mastery"`
	Type             Type           `json:"type,omitempty"`
	CharacterTrait   CharacterTrait `json:"character_trait"`
	Ideal            Ideal          `json:"ideal"`
	Affection        Affection      `json:"affection"`
	Weakness         Weakness       `json:"weakness"`
}
type BackgroundBson struct {
	ID               primitive.ObjectID `bson:"_id"`
	BackgroundName   string             `json:"background_name"`
	BackgroundNameRu string             `json:"background_name_ru"`
	Description      string             `json:"description"`
	Personalization  string             `json:"personalization"`
	Advice           string             `json:"advice"`
	SkillMastery     []string           `json:"skill_mastery"`
	Type             Type               `json:"type,omitempty"`
	CharacterTrait   CharacterTrait     `json:"character_trait"`
	Ideal            Ideal              `json:"ideal"`
	Affection        Affection          `json:"affection"`
	Weakness         Weakness           `json:"weakness"`
}

type BackgroundAnswer struct {
	BackgroundName   string   `json:"background_name"`
	BackgroundNameRu string   `json:"background_name_ru"`
	Description      string   `json:"description"`
	Advice           string   `json:"advice"`
	Personalization  string   `json:"personalization"`
	SkillMastery     []string `json:"skill_mastery"`
	Type             string   `json:"type"`
	CharacterTrait   string   `json:"character_trait"`
	Worldview        string   `json:"worldview"`
	Ideal            string   `json:"ideal"`
	Affection        string   `json:"affection"`
	Weakness         string   `json:"weakness"`
}
