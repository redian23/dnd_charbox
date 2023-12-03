package classes

var (
	HitsCountGlobal   int
	ClassNameGlobalRu string
	ArmorInfoGlobal   []ArmorAnswer
)

func GenerateClass(classNameRU string) *ClassAnswer {
	ClassNameGlobalRu = ""

	ClassNameGlobalRu = classNameRU
	chars = GetClassCharactsFormDB()
	armorData = GetArmorFormDB()
	weaponData = GetWeaponFormDB()

	classNameEng, classAbilityStats := AbilitiesWithLevelUp()

	armor := setArmor(classNameRU)
	ArmorInfoGlobal = armor
	classSkills, _ := setClassSkills()
	modif := setModifiersForClass(classAbilityStats)

	CharProficienciesGlobal = setProficiencies()

	HitsCountGlobal = setHitCount(modif.BodyDifficulty)

	return &ClassAnswer{
		ClassName:     classNameEng,
		ClassNameRU:   classNameRU,
		Ability:       classAbilityStats,
		Modifier:      modif,
		Inspiration:   false,
		Proficiencies: CharProficienciesGlobal,
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
		Spellcasting:   GetClassSpellBasicCharacteristic(),
		ClassAbilities: getClassAbilities(),
	}
}
