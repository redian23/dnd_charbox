package races

import (
	"fmt"
	"github.com/mazen160/go-random"
)

var (
	raceType, raceTypeRu     string
	raceName                 string
	raceAge, raceSpeed       int
	raceHeight               int
	raceHeightFt             string
	raceWeight, raceWeightLb int
	raceBodySize             string
	raceGender               string
	raceAbil                 []raceAbility
	firstName, resist        string
)

func GenerateRaceForCharacter(raceNameRu string) *RacesAnswer {
	RaceData = getRacesFormDB()
	rollNum, _ := random.IntRange(0, len(RaceData))
	var lastName string

	for _, race := range RaceData {
		if race.RaceNameRu == raceNameRu {
			raceName = race.RaceName
			RaceNameGlobalRu = raceNameRu

			rollNum, _ = random.IntRange(0, len(race.Type))
			raceType = race.Type[rollNum].TypeRaceName
			raceTypeRu = race.Type[rollNum].TypeRaceNameRu
			raceSpeed = race.Type[rollNum].Speed

			RaceTypeGlobalRu = raceTypeRu

			maxAge := (race.MaxAge * 75) / 100
			raceAge, _ = random.IntRange(race.MinAge, maxAge)

			raceHeight, _ = random.IntRange(race.MinHeight, race.MaxHeight)
			raceHeightFt = fmt.Sprintf("%.1f", float32(raceHeight)*0.0328)

			raceWeight, _ = random.IntRange(race.MinWeight, race.MaxWeight)
			raceWeightLb = (raceWeight * 220) / 100

			raceBodySize = race.BodySize

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

			for _, rType := range race.Type {
				if raceTypeRu == rType.TypeRaceNameRu {
					raceAbil = rType.RaceAbilities
				}
			}
		}
	}

	return &RacesAnswer{
		RaceName:      raceName,
		RaceNameRu:    RaceNameGlobalRu,
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
		RaceSkill:     nil,
		CharacterName: firstName,
		FirstName:     firstName,
		LastName:      lastName,
		Speed:         raceSpeed,
		Resist:        resist,
		RaceAbilities: raceAbil,
		Other: other{
			DragonType:       setDragonType(raceTypeRu),
			YuantiAppearance: setAppearanceYuanti(raceTypeRu),
		},
	}
}
