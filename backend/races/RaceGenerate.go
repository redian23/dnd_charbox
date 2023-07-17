package races

func GenerateRace() RacesAnswer {
	var rc RacesAnswer

	rc.RaceName, rc.RaceNameRu = setRaceName()
	rc.Type, rc.TypeRu = setRaceType()
	rc.Gender = setGender()
	rc.Eyes = setEyesColor()
	rc.Hair = setHairColor()
	rc.Age = setAge()
	rc.Height = setHeight()
	rc.Weight = setWeight()
	rc.Speed = setSpeed()
	rc.FirstName = setFirstName()
	rc.LastName = setLastName()
	rc.Langs = setKnowLanguages()
	rc.RaceSkill = setRaceSkills()
	rc.Resist = setResist()

	return rc
}
