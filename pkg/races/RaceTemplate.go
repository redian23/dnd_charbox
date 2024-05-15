package races

//
//import (
//	"pregen/pkg/core"
//	"pregen/pkg/spells"
//)
//
//var raceNames = map[string][]string{
//	"Мужской":   {},
//	"Женский": {},
//}
//
//var raceLangs = []string{}
//
//var bodyStats = races.BodyStats{
//	BodySize:  "",
//	MinAge:    0,
//	MaxAge:    0,
//	MinHeight: 0,
//	MaxHeight: 0,
//	MinWeight: 0,
//	MaxWeight: 0,
//}
//
//var raceResists = []string{}
//
//func getAddictionInfoInterface() string {
//	var D6 = core.D6
//
//	var angelicGuideName = map[int]string{
//		1: " ",
//		2: " ",
//		3: " ",
//		4: " ",
//		5: " ",
//		6: " ",
//	}
//
//	var addictionInfoTest = "" + angelicGuideName[D6.RollDice()]
//
//	return addictionInfoTest
//}
//
//func raceTypeInterface() *races.RaceType {
//	var typeAbilities = map[string]string{
//		"Тёмное зрение": "Благословленное светлым духом, ваше зрение может легко проникать сквозь темноту. На расстоянии в 60 футов вы при тусклом освещении можете видеть так, как будто это яркое освещение, и в темноте так, как будто это тусклое освещение. В темноте вы не можете различать цвета, только оттенки серого.",
//	}
//
//	var AbilityScorePlusMap = map[string]int{
//		"Wisdom":   1,
//		"Charisma": 2,
//	}
//
//	return &races.RaceType{
//		TypeRaceName:     "Type",
//		TypeRaceNameRu:   "Тип",
//		RaceAbilities:    typeAbilities,
//		AbilityScorePlus: AbilityScorePlusMap,
//		Speed:            30,
//		RaceSpellsList: []spells.SpellsJSON{
//			spells.FindSpellInDB("Свет"),
//		},
//		Book:             "",
//		Standard:         false,
//	}
//}
//
//func getRacePersonalization() races.RacePersonalization {
//
//	return races.RacePersonalization{}
//}
//
//func GetRaceInterface(raceTypeName string, gender string) *races.Race {
//	var raceType *races.RaceType
//
//	switch raceTypeName {
//	case "Какой-то архитип":
//		raceType = raceTypeInterface()
//	}
//
//	return &races.Race{
//		RaceName:             "Race",
//		RaceNameRu:           "Раса",
//		AddictionInformation: getAddictionInfoInterface(),
//		FirstName:            races.GetRaceFirstName(gender, raceNames), //внешняя переменная которая прилетает с фронта
//		LastName:             "",
//		Gender:               gender,
//		Body:                 races.GetBodyStats(bodyStats),
//		Language:             raceLangs,
//		Resist:               raceResists,
//		Type:                 *raceType,
//	}
//}
