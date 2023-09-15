package classes

import "pregen/pkg/spells"

var ClassNameGlobal string

func GenerateClass(classNameRU string) *ClassAnswer {
	chars = GetClassCharactsFormDB()
	armorData = GetArmorFormDB()
	weaponData = GetWeaponFormDB()

	className, _, classAbility := statAnalyze(classNameRU)
	ClassNameGlobal = className
	spellCount := getClassZeroSpellCount()
	modif := setModifiersForClass(classAbility)

	return &ClassAnswer{
		ClassName:        className,
		ClassNameRU:      classNameRU,
		Ability:          classAbility,
		Modifier:         modif,
		Inspiration:      true,
		Proficiencies:    setProficiencies(),
		ProficiencyBonus: 2,
		SkillsOfClass:    setClassSkills(),
		SavingThrows:     setSaveThrowsForClass(modif),
		Hits: hits{
			HitDice:  setHitDice(),
			HitCount: setHitCount(modif.BodyDifficulty),
		},
		ClassEquipment: setClassEquipmentList(),
		Armor:          setArmor(classNameRU),
		Weapon:         setWeapons(),
		Initiative:     setInitiative(),
		SpellsLVL0:     spells.GetSpellsZeroLevelForClass(classNameRU, spellCount),
		SpellsLVL1:     nil,
	}
}
