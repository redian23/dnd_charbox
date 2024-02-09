package backgrounds

type BackgroundAnswer struct {
	BackgroundName    string            `json:"background_name"`
	BackgroundNameRu  string            `json:"background_name_ru"`
	BackgroundAbility BackgroundAbility `json:"background_ability"`
	MasteryOfTools    []string          `json:"mastery_of_tools"`
	Equipment         []string          `json:"equipment"`
	Description       string            `json:"description"`
	Advice            string            `json:"advice"`
	Personalization   string            `json:"personalization"`
	SkillMastery      []string          `json:"skill_mastery"`
	Type              string            `json:"type"`
	CharacterTrait    string            `json:"character_trait"`
	Worldview         string            `json:"worldview"`
	Ideal             string            `json:"ideal"`
	Affection         string            `json:"affection"`
	Weakness          string            `json:"weakness"`
	Gold              int               `json:"gold"`
}
type BackgroundBson struct {
	ID                ID                `json:"_id"`
	Advice            string            `json:"advice"`
	Affection         Affection         `json:"affection"`
	BackgroundAbility BackgroundAbility `json:"background_ability"`
	BackgroundName    string            `json:"background_name"`
	BackgroundNameRu  string            `json:"background_name_ru"`
	CharacterTrait    CharacterTrait    `json:"character_trait"`
	Description       string            `json:"description"`
	Equipment         []string          `json:"equipment"`
	Gold              int               `json:"gold"`
	Ideal             Ideal             `json:"ideal"`
	MasteryOfTools    []string          `json:"mastery_of_tools"`
	Personalization   string            `json:"personalization"`
	SkillMastery      []string          `json:"skill_mastery"`
	Type              Type              `json:"type"`
	Weakness          Weakness          `json:"weakness"`
}
type ID struct {
	Oid string `json:"$oid"`
}
type Affection struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type BackgroundAbility struct {
	AbilityName string `json:"ability_name"`
	Description string `json:"description"`
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
type Type struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
type Weakness struct {
	Dice  string   `json:"dice"`
	Value []string `json:"value"`
}
