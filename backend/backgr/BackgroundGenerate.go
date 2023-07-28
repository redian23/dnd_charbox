package backgr

func GenerateBackground(className string) BackgroundAnswer {
	var bg BackgroundAnswer
	bg.BackgroundName = backgroundAnalyze(className)
	bg.BackgroundNameRu = setBackgroundNameRU()
	bg.Type = setBackgroundType()
	bg.Description = setDescription()
	bg.Advice = setAdvice()
	bg.SkillMastery = setSkillMastery()
	bg.Personalization = setPersonalization()
	bg.CharacterTrait = setCharacterTrait()
	bg.Ideal, bg.Worldview = setIdeal()
	bg.Affection = setAffection()
	bg.Weakness = setWeakness()
	return bg
}
