package classes

var (
	HitsCountGlobal   int
	ClassNameGlobalRu string
	ArmorInfoGlobal   []ArmorAnswer
)

func GenerateClass(classNameRU string, lvl int) *ClassAnswer {
	ClassNameGlobalRu = ""
	chars = GetClassCharactsFormDB()
	armorData = GetArmorFormDB()
	weaponData = GetWeaponFormDB()

	className, _, classAbilityStats := statAnalyze(classNameRU)
	ClassNameGlobalRu = classNameRU

	armor := setArmor(classNameRU)
	ArmorInfoGlobal = armor
	classSkills, _ := setClassSkills()
	modif := setModifiersForClass(classAbilityStats)

	HitsCountGlobal = setHitCount(modif.BodyDifficulty)

	return &ClassAnswer{
		ClassName:     className,
		ClassNameRU:   classNameRU,
		Ability:       classAbilityStats,
		Modifier:      modif,
		Inspiration:   false,
		Proficiencies: setProficiencies(),
		SkillsOfClass: classSkills,
		SavingThrows:  setSaveThrowsForClass(modif),
		Hits: hits{
			HitDice:  setHitDice(),
			HitCount: HitsCountGlobal,
		},
		ClassEquipment: setClassEquipmentList(),
		Armor:          setArmor(classNameRU),
		Weapon:         setWeapons(),
		Initiative:     setInitiative(),
		Spellcasting:   getClassSpellBasicCharacteristic(),
		ClassAbilities: getClassAbilities(),
	}
}
