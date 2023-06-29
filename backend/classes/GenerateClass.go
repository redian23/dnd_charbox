package classes

func GetNewClass() Class {
	//Start To Generate
	//Default Values
	var class Class
	class.ClassName = "HomeBrew"
	class.Inspiration = true
	class.ProficiencyBonus = 2
	class.Ability = rerollClassAbilitiesStats()

	class.ClassName = statAnalyze(class)
	class.Modifier = setModifiersForClass(class.Ability)
	class.SavingThrows = setSaveThrowsForClass(class.Modifier)
	class.Skills = setSkillsForClass(nil, class.Modifier)

	class.Appearance.Gender = setGender()
	class.Appearance.Age = setAge()
	class.Appearance.Eyes = setEyesColor()
	class.Appearance.Height = setHeight()
	class.Appearance.Weight = setWeight(class.Appearance.Height)
	class.Appearance.Hair = setHairColor()
	return class
}
