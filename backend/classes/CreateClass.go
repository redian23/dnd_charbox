package classes

func GenerateClass() ClassAnswer {
	//Start To Generate
	//Default Values
	var class ClassAnswer
	class.Inspiration = true
	class.ProficiencyBonus = 2

	class.ClassName, class.ClassNameRU, class.Ability = statAnalyze()
	class.Modifier = setModifiersForClass()
	class.SavingThrows = setSaveThrowsForClass()
	class.PassiveWisdom = setPassiveWisdom()
	class.Hits.HitDice = setHitDice()
	class.Hits.HitCount = setHitCount()
	class.SkillsOfClass = setClassSkills()

	return class
}
