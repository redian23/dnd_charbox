package backgrounds

type Background struct {
	BackgroundName         string            `json:"background_name"`
	BackgroundNameRu       string            `json:"background_name_ru"`
	BackgroundSpecificType string            `json:"background_specific_type"`
	BackgroundAbility      map[string]string `json:"background_ability"`
	Langs                  []string          `json:"langs"`
	MasteryOfTools         []string          `json:"mastery_of_tools"`
	BackgroundItems        []Item            `json:"equipment"`
	Description            string            `json:"description"`
	Personalization        Personalization   `json:"personalization"`
	BackgroundSkills       []string          `json:"background_skills"`
	Gold                   int               `json:"gold"`
}

type Item struct {
	Name     string `json:"name"`
	ItemType string `json:"item_type"`
	Count    int    `json:"count"`
}
type Ideal struct {
	Worldview   string `json:"worldview"`
	WorldviewRu string `json:"worldview_ru"`
	Text        string `json:"text"`
}

type Personalization struct {
	PersonalizationDescription string `json:"personalization_description"`
	Advice                     string `json:"advice"`
	CharacterTrait             string `json:"character_trait"`
	Ideal                      string `json:"ideal"`
	Affection                  string `json:"affection"`
	Weakness                   string `json:"weakness"`
	Worldview                  string `json:"worldview"`
}
