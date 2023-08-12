package races

import (
	"fmt"
	"github.com/mazen160/go-random"
)

var (
	raceType, raceTypeRu, raceSkill string
	raceAge, raceSpeed              int
	raceHeight                      int
	raceHeightFt                    string
	raceWeight, raceWeightLb        int
	raceBodySize                    string
	raceGender                      = setGender()
	firstName, lastName, resist     string
	langs                           []string
	darkvision                      bool
	raceAbil                        []raceAbility
)

func GenerateRaceForCharacter() *RacesAnswer {
	raceData = getRacesFormDB()
	rollNum, _ := random.IntRange(0, len(raceData))

	raceName = raceData[rollNum].RaceName
	raceNameRu := raceData[rollNum].RaceNameRu

	for _, race := range raceData {
		if race.RaceName == raceName {

			rollNum, _ = random.IntRange(0, len(race.Type))
			raceType = race.Type[rollNum].TypeRaceName
			raceTypeRu = race.Type[rollNum].TypeRaceNameRu

			if race.RaceName == raceName && len(race.RaceSkill) > 0 {
				rollNum, _ = random.IntRange(0, len(race.RaceSkill))
				raceSkill = race.RaceSkill[rollNum]
			}

			maxAge := (race.MaxAge * 75) / 100
			raceAge, _ = random.IntRange(race.MinAge, maxAge)

			raceHeight, _ = random.IntRange(race.MinHeight, race.MaxHeight)
			raceHeightFt = fmt.Sprintf("%.1f", float32(raceHeight)*0.0328)

			raceWeight, _ = random.IntRange(race.MinWeight, race.MaxWeight)
			raceWeightLb = (raceWeight * 220) / 100

			raceBodySize = race.BodySize
			raceSpeed = race.Speed

			if raceGender == "Мужской" && len(race.Names.Man) != 0 {
				rollNum, _ = random.IntRange(0, len(race.Names.Man))
				firstName = race.Names.Man[rollNum]
			} else {
				rollNum, _ = random.IntRange(0, len(race.Names.Woman))
				firstName = race.Names.Woman[rollNum]
			}

			if len(race.LastNames) != 0 {
				rollNum, _ = random.IntRange(0, len(race.LastNames))
				lastName = race.LastNames[rollNum].(string)
			}
			langs = race.Langs
			resist = race.Resist
			darkvision = race.Darkvision

			for _, rType := range race.Type {
				if raceTypeRu == rType.TypeRaceNameRu {
					raceAbil = rType.RaceAbilities
				}
			}
		}
	}

	return &RacesAnswer{
		RaceName:      raceName,
		RaceNameRu:    raceNameRu,
		Gender:        raceGender,
		RacePhoto:     setRacePhoto(raceName, raceGender),
		Type:          raceType,
		TypeRu:        raceTypeRu,
		Age:           raceAge,
		Height:        raceHeight,
		HeightFt:      raceHeightFt,
		Weight:        raceWeight,
		WeightLb:      raceWeightLb,
		BodySize:      raceBodySize,
		Eyes:          setEyesColor(),
		Hair:          setHairColor(),
		RaceSkill:     raceSkill,
		FirstName:     firstName,
		LastName:      lastName,
		Speed:         raceSpeed,
		Langs:         langs,
		Resist:        resist,
		Darkvision:    darkvision,
		RaceAbilities: raceAbil,
		Other: other{
			DragonType:       setDragonType(raceTypeRu),
			YuantiAppearance: setAppearanceYuanti(raceTypeRu),
		},
	}
}
