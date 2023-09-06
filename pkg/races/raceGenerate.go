package races

import (
	"fmt"
	"github.com/mazen160/go-random"
)

var (
	raceType, raceTypeRu     string
	raceAge, raceSpeed       int
	raceHeight               int
	raceHeightFt             string
	raceWeight, raceWeightLb int
	raceBodySize             string
	raceGender               string
	langs                    []string
	darkvision               bool
	raceAbil                 []raceAbility
	firstName, resist        string
)

func GenerateRaceForCharacter(raceNameRu string) *RacesAnswer {
	raceData = getRacesFormDB()
	rollNum, _ := random.IntRange(0, len(raceData))
	var lastName string

	for _, race := range raceData {
		if race.RaceNameRu == raceNameRu {
			raceNameGlobal = race.RaceName

			rollNum, _ = random.IntRange(0, len(race.Type))
			raceType = race.Type[rollNum].TypeRaceName
			raceTypeRu = race.Type[rollNum].TypeRaceNameRu

			maxAge := (race.MaxAge * 75) / 100
			raceAge, _ = random.IntRange(race.MinAge, maxAge)

			raceHeight, _ = random.IntRange(race.MinHeight, race.MaxHeight)
			raceHeightFt = fmt.Sprintf("%.1f", float32(raceHeight)*0.0328)

			raceWeight, _ = random.IntRange(race.MinWeight, race.MaxWeight)
			raceWeightLb = (raceWeight * 220) / 100

			raceBodySize = race.BodySize
			raceSpeed = race.Speed

			genders := [2]string{"Мужской", "Женский"} //шах и мат феминистки
			count, _ := random.IntRange(0, 1)
			raceGender = genders[count]

			if raceGender == "Мужской" {
				rollNum, _ = random.IntRange(0, len(race.Names.Man)-1)
				firstName = race.Names.Man[rollNum]
			} else {
				rollNum, _ = random.IntRange(0, len(race.Names.Woman)-1)
				firstName = race.Names.Woman[rollNum]
			}

			if len(race.LastNames) != 0 {
				rollNum, _ = random.IntRange(0, len(race.LastNames))
				lastName = race.LastNames[rollNum].(string)
			}

			resist = race.Resist
			langs = race.Langs

			darkvision = race.Darkvision

			for _, rType := range race.Type {
				if raceTypeRu == rType.TypeRaceNameRu {
					raceAbil = rType.RaceAbilities
				}
			}
		}
	}

	return &RacesAnswer{
		RaceName:      raceNameGlobal,
		RaceNameRu:    raceNameRu,
		Gender:        raceGender,
		RacePhoto:     setRacePhoto(raceNameGlobal, raceGender),
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
		RaceSkill:     nil,
		CharacterName: firstName,
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
