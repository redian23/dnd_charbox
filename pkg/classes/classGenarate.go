package classes

var (
	HitsCountGlobal   int
	ClassNameGlobalRu string
)

func GenerateClass(classNameRU string) *ClassAnswer {
	ClassNameGlobalRu = ""
	chars = GetClassCharactsFormDB()
	armorData = GetArmorFormDB()
	weaponData = GetWeaponFormDB()

	className, _, classAbilityStats := statAnalyze(classNameRU)
	ClassNameGlobalRu = classNameRU

	classSkills, _ := setClassSkills()
	modif := setModifiersForClass(classAbilityStats)

	HitsCountGlobal = setHitCount(modif.BodyDifficulty)

	return &ClassAnswer{
		ClassName:        className,
		ClassNameRU:      classNameRU,
		Ability:          classAbilityStats,
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
