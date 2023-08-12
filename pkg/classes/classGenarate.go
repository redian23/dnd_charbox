package classes

var ClassNameGlobal string

func GenerateClass() *ClassAnswer {
	chars = GetClassCharactsFormDB()
	armorData = GetArmorFormDB()
	weaponData = GetWeaponFormDB()

	className, classNameRU, classAbility := statAnalyze()
	ClassNameGlobal = className
	modif := setModifiersForClass(classAbility)

	return &ClassAnswer{
		ClassName:        className,
		ClassNameRU:      classNameRU,
		Ability:          classAbility,
		Modifier:         modif,
		Inspiration:      true,
		Proficiencies:    setProficiencies(),
		ProficiencyBonus: 2,
		PassiveWisdom:    setPassiveWisdom(modif.Wisdom),
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
	}
}