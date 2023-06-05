package classes

type ClassJson struct {
	Classes []Class `json:"classes"`
}
type Class struct {
	ClassName      string `json:"class_name"`
	Description    string `json:"description"`
	Strength       int    `json:"strength"`
	Dexterity      int    `json:"dexterity"`
	BodyDifficulty int    `json:"body_difficulty"`
	Intelligence   int    `json:"intelligence"`
	Wisdom         int    `json:"wisdom"`
	Charisma       int    `json:"charisma"`
}

func (c *Class) ModifyDexterity(value int) {
	c.Dexterity = value
}
func (c *Class) ModifyBodyDifficulty(value int) {
	c.BodyDifficulty = value
}
func (c *Class) ModifyIntelligence(value int) {
	c.Intelligence = value
}
func (c *Class) ModifyWisdom(value int) {
	c.Wisdom = value
}
func (c *Class) ModifyCharisma(value int) {
	c.Charisma = value
}
func (c *Class) ModifyStrength(value int) {
	c.Strength = value
}
