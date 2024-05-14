package classes

import "github.com/mazen160/go-random"

var (
	HitsCountGlobal   int
	ClassNameGlobalRu string
	ArmorInfoGlobal   []ArmorAnswer
	ClassSpecificInfo ClassSpecific
)

func GetRandomBackgroundForClass(classNameRU string) string {
	var classInfo = GetClassCharactsFormDB()
	for _, class := range classInfo {
		randNum, _ := random.IntRange(0, len(class.Background))
		return class.Background[randNum]
	}
	return ""
}

func GenerateClass(classNameRU string) *ClassAnswer {
	ClassNameGlobalRu = ""
	ClassSpecificInfo.Parameter1 = ""
	ClassSpecificInfo.Parameter2 = ""

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
		ClassSpecific:  ClassSpecificInfo,
	}
}
