package classes

func GenerateClass() Class {
	//Start To Generate
	//Default Values
	var class Class
	class.ClassName = "HomeBrew"
	class.Inspiration = true
	class.ProficiencyBonus = 2

	class.ClassName, class.ClassNameRU, class.Ability = statAnalyze(class)
	class.Modifier = setModifiersForClass(class.Ability)
	class.SavingThrows = setSaveThrowsForClass(class.Modifier)
	class.Skills = setSkillsForClass(nil, class.Modifier)
	class.PassiveWisdom = setPassiveWisdom(class.Modifier.Wisdom)

	return class
}
