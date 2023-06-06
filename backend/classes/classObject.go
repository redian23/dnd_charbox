package classes

type ClassJson struct {
	Classes []Class `json:"classes"`
}
type Class struct {
	ClassName   string   `json:"class_name"`
	Description string   `json:"description"`
	Ability     Ability  `json:"ability"`
	Modifier    Modifier `json:"modifier"`
}

type Ability struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
}

type Modifier struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
}

func (a *Ability) ModifyDexterity(value int) {
	a.Dexterity = value
}
func (a *Ability) ModifyBodyDifficulty(value int) {
	a.BodyDifficulty = value
}
func (a *Ability) ModifyIntelligence(value int) {
	a.Intelligence = value
}
func (a *Ability) ModifyWisdom(value int) {
	a.Wisdom = value
}
func (a *Ability) ModifyCharisma(value int) {
	a.Charisma = value
}
func (a *Ability) ModifyStrength(value int) {
	a.Strength = value
}
