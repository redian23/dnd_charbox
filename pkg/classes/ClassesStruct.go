package classes

import "pregen/pkg/spells"

type Class struct {
	ClassName        string              `json:"class_name"`
	ClassNameRU      string              `json:"class_name_ru"`
	ClassAbilities   []ClassAbility      `json:"class_abilities"`
	AbilityScore     AbilityScore        `json:"ability"`
	AbilityModifier  AbilityModifier     `json:"modifier"`
	SavingThrows     *SavingThrows       `json:"saving_throws"`
	Inspiration      bool                `json:"inspiration"`
	Proficiencies    Proficiencies       `json:"proficiencies"`
	ProficiencyBonus int                 `json:"proficiency_bonus"`
	Hits             Hits                `json:"hits"`
	Caster           bool                `json:"caster"`
	SpellCasting     SpellCasting        `json:"spell_casting"`
	SpellsList       []spells.SpellsJSON `json:"spells_list"`
	Equipment        []Item              `json:"equipment"`
	ArmorInfo        Armor               `json:"armor_info"`
	WeaponInfo       []Weapon            `json:"weapon_info"`
}

type ClassAbility struct {
	Level       int    `json:"level"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SpellCasting struct {
	BasicSpellCharacteristics string      `json:"basic_spell_characteristics"`
	SpellDamageModifier       int         `json:"spell_damage_modifier"`
	SavingThrowDifficulty     int         `json:"saving_throw_difficulty"`
	ZeroLevelSpellsKnownCount int         `json:"zero_level_spells_known_count"`
	OneLevelSpellsKnownCount  int         `json:"one_level_spells_known_count"`
	TwoLevelSpellsKnownCount  int         `json:"two_level_spells_known_count"`
	TreeLevelSpellsKnownCount int         `json:"tree_level_spells_known_count"`
	FourLevelSpellsKnownCount int         `json:"four_level_spells_known_count"`
	TotalSpellCount           int         `json:"total_spell_count"`
	SpellsSlots               SpellsSlots `json:"spells_slots"`
}

type SpellsSlots struct {
	OneLevel  int `json:"one_level"`
	TwoLevel  int `json:"two_level"`
	TreeLevel int `json:"tree_level"`
	FourLevel int `json:"four_level"`
	FiveLevel int `json:"five_level"`
}
type Proficiencies struct {
	Armor         []string `json:"armor"`
	Weapons       []string `json:"weapons"`
	Tools         []string `json:"tools"`
	SavingThrow   []string `json:"saving_throw"`
	SkillsOfClass []string `json:"skills_of_class"`
}

type AbilityScore struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
}

type AbilityModifier struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
}

type SavingThrows struct {
	Strength       savingThrow `json:"strength"`
	Dexterity      savingThrow `json:"dexterity"`
	BodyDifficulty savingThrow `json:"body_difficulty"`
	Intelligence   savingThrow `json:"intelligence"`
	Wisdom         savingThrow `json:"wisdom"`
	Charisma       savingThrow `json:"charisma"`
}

type savingThrow struct {
	Name    string `json:"name"`
	Point   int    `json:"point"`
	Mastery bool   `json:"mastery"`
}

type Hits struct {
	HitDice  string `json:"hit_dice"`
	HitCount int    `json:"hit_count"`
}

type Item struct {
	Name     string `json:"name"`
	ItemType string `json:"item_type"`
	Count    int    `json:"count"`
}
