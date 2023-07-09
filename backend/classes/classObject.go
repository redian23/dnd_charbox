package classes

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClassAnswer struct {
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
	Hits             hits         `json:"hits"`
	Armor            string       `json:"armor"`
	Instruments      []string     `json:"instruments"`
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

type ClassesBSON []struct {
	ID           primitive.ObjectID `bson:"_id"`
	ClassName    string             `json:"className"`
	ClassNameRU  string             `json:"classNameRU"`
	CharReq      [][]string         `json:"charReq"`
	Background   []string           `json:"background"`
	Armor        []string           `json:"armor"`
	Hits         hits               `json:"hits"`
	Instruments  []string           `json:"instruments"`
	Skills       skillsInDB         `json:"skills"`
	SavingThrows []string           `json:"saving_throws"`
	Weapon       weapon             `json:"weapon"`
}

type hits struct {
	HitDice  string
	HitCount int
}

type skillsInDB struct {
	RandomCount int      `json:"random_count"`
	Skills      []string `json:"skills"`
}

type weapon struct {
	WeaponType string `json:"weapon_type"`
	WeaponName string `json:"weapon_name"`
}
