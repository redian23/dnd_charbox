package classes

type Class struct {
	ClassName        string       `json:"class_name"`
	ClassNameRU      string       `json:"class_name_ru"`
	Description      string       `json:"description"`
	Ability          Ability      `json:"ability"`
	Modifier         Modifier     `json:"modifier"`
	Inspiration      bool         `json:"inspiration"`
	ProficiencyBonus int          `json:"proficiency_bonus"`
	PassiveWisdom    int          `json:"passive_wisdom"`
	Skills           Skills       `json:"skills"`
	SavingThrows     SavingThrows `json:"saving_throws"`
}

type Ability struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
	Total          int `json:"total"`
}

type Modifier struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
}
type Skills struct {
	Acrobatics     skill `json:"acrobatics"`
	AnimalHandling skill `json:"animal_handling"`
	Arcana         skill `json:"arcana"`
	Athletics      skill `json:"athletics"`
	Deception      skill `json:"deception"`
	History        skill `json:"history"`
	Insight        skill `json:"insight"`
	Intimidation   skill `json:"intimidation"`
	Investigation  skill `json:"investigation"`
	Medicine       skill `json:"medicine"`
	Nature         skill `json:"nature"`
	Perception     skill `json:"perception"`
	Performance    skill `json:"performance"`
	Persuasion     skill `json:"persuasion"`
	Religion       skill `json:"religion"`
	SleightOfHand  skill `json:"sleight_of_hand"`
	Stealth        skill `json:"stealth"`
	Survival       skill `json:"survival"`
}

type skill struct {
	SkillName     string `json:"skill_name"`
	ModifierValue int    `json:"modifier_value"`
	Proficiency   bool   `json:"proficiency"`
}

type SavingThrows struct {
	Strength       savingThrows `json:"strength"`
	Dexterity      savingThrows `json:"dexterity"`
	BodyDifficulty savingThrows `json:"body_difficulty"`
	Intelligence   savingThrows `json:"intelligence"`
	Wisdom         savingThrows `json:"wisdom"`
	Charisma       savingThrows `json:"charisma"`
}

type savingThrows struct {
	Name    string `json:"name"`
	Point   int    `json:"point"`
	Mastery bool   `json:"mastery"`
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
