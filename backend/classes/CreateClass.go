package classes

var ClassNameGlobal string

func GenerateClass() ClassAnswer {
	//Start To Generate
	//Default Values
	var class ClassAnswer
	class.Inspiration = true
	class.ProficiencyBonus = 2

	class.ClassName, class.ClassNameRU, class.Ability = statAnalyze()
	ClassNameGlobal = class.ClassName

	class.Modifier = setModifiersForClass(class.Ability)
	class.SavingThrows = setSaveThrowsForClass(class.Modifier)
	class.PassiveWisdom = setPassiveWisdom(class.Modifier.Wisdom)
	class.Hits.HitDice = setHitDice()
	class.Hits.HitCount = setHitCount(class.Modifier.BodyDifficulty)
	class.SkillsOfClass = setClassSkills()
	class.Proficiencies = setProficiencies()

	return class
}
