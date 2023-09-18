package classes

var (
	HitsCountGlobal   int
	ClassNameGlobal   string
	ClassNameGlobalRu string
)

func GenerateClass(classNameRU string) *ClassAnswer {
	chars = GetClassCharactsFormDB()
	armorData = GetArmorFormDB()
	weaponData = GetWeaponFormDB()

	className, _, classAbility := statAnalyze(classNameRU)
	ClassNameGlobal = className
	ClassNameGlobalRu = classNameRU
	classSkills, _ := setClassSkills()
	modif := setModifiersForClass(classAbility)

	HitsCountGlobal = setHitCount(modif.BodyDifficulty)

	return &ClassAnswer{
		ClassName:        className,
		ClassNameRU:      classNameRU,
		Ability:          classAbility,
		Modifier:         modif,
		Inspiration:      false,
		Proficiencies:    setProficiencies(),
		ProficiencyBonus: 2,
		SkillsOfClass:    classSkills,
		SavingThrows:     setSaveThrowsForClass(modif),
		Hits: hits{
			HitDice:  setHitDice(),
			HitCount: HitsCountGlobal,
		},
		ClassEquipment: setClassEquipmentList(),
		Armor:          setArmor(classNameRU),
		Weapon:         setWeapons(),
		Initiative:     setInitiative(),
		SpellUsing:     getClassSpellBasicCharacteristic(),
		ClassAbilities: getClassAbilities(),
	}
}
