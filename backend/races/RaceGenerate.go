package races

func GenerateRace() RacesAnswer {
	var rc RacesAnswer

	rc.RaceName, rc.RaceNameRu = setRaceName()
	rc.Type, rc.TypeRu = setRaceType()
	rc.Gender = setGender()
	rc.Eyes = setEyesColor()
	rc.Hair = setHairColor()
	rc.Age = setAge()
	rc.Height, rc.HeightFt = setHeight()
	rc.Weight, rc.WeightLb = setWeight()
	rc.BodySize = setBodySize()
	rc.Speed = setSpeed()
	rc.FirstName = setFirstName()
	rc.LastName = setLastName()
	rc.Langs = setKnowLanguages()
	rc.Resist = setResist()

	return rc
}
