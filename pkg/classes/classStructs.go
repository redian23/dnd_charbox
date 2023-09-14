package classes

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClassAnswer struct {
	ClassName        string         `json:"class_name"`
	ClassNameRU      string         `json:"class_name_ru"`
	Description      string         `json:"description"`
	Ability          Ability        `json:"ability"`
	Modifier         Modifier       `json:"modifier"`
	Inspiration      bool           `json:"inspiration"`
	Proficiencies    Proficiencies  `json:"proficiencies"`
	ProficiencyBonus int            `json:"proficiency_bonus"`
	SkillsOfClass    []string       `json:"skills_of_class"`
	SavingThrows     SavingThrows   `json:"saving_throws"`
	Hits             hits           `json:"hits"`
	ClassEquipment   []Variants     `json:"class_equipment"`
	Armor            []ArmorAnswer  `json:"armor"`
	Weapon           []WeaponAnswer `json:"weapon"`
	Initiative       string         `json:"initiative"`
	SpellsLVL0       []string       `json:"spells_lvl_0"`
	SpellsLVL1       []string       `json:"spells_lvl_1"`
}

type Proficiencies struct {
	Armor       string `json:"armor"`
	Weapons     string `json:"weapons"`
	Instruments string `json:"instruments"`
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
	ID             primitive.ObjectID `bson:"_id"`
	ClassName      string             `json:"className"`
	ClassNameRU    string             `json:"classNameRU"`
	CharReq        [][]string         `json:"charReq"`
	Background     []string           `json:"background"`
	Hits           hits               `json:"hits"`
	SkillsInDB     skillsInDB         `bson:"skills"`
	SavingThrows   []string           `json:"saving_throws"`
	Proficiencies  Proficiencies      `json:"proficiencies"`
	PicksEquipment equipPicks         `bson:"equipPicks"`
	BasicEquipment equipBasic         `bson:"equipBasic"`
	Spells         Spells             `bson:"spells"`
}

type hits struct {
	HitDice  string `json:"hit_dice"`
	HitCount int    `json:"hit_count"`
}

type skillsInDB struct {
	RandomCount int      `json:"random_count"`
	SkillsList  []string `bson:"skilllist"`
}

type ClassWriteToBD []struct {
	CharReq [][]string `json:"charReq"`
	Hits    struct {
		HitDice  string `json:"hit_dice"`
		HitCount int    `json:"hit_count"`
	} `json:"hits"`
	SavingThrows []string `json:"saving_throws"`
	Skills       struct {
		RandomCount int      `json:"random_count"`
		SkillList   []string `json:"skill_list"`
	} `json:"skills"`
	ClassName   string   `json:"className"`
	ClassNameRU string   `json:"classNameRU"`
	Background  []string `json:"background"`
}

type Spells []struct {
	Level string   `json:"level"`
	List  []string `json:"list"`
	Count int      `json:"count"`
}
