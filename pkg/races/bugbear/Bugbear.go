package bugbear

import (
	"pregen/pkg/general"
	"pregen/pkg/races"
	"pregen/pkg/spells"
)

var raceNames = map[string][]string{
	"male": {"Аругет", "Четин", "Даавн", "Дабрак", "Дагии", "Древдуул", "Дуулан", "Феник", "Гудруун", "Халуун",
		"Харуук", "Джазаал", "Каллаад", "Кракуул", "Кроотад", "Мазаан", "Мунта", "Насаар", "Ракари", "Рексиит",
		"Тарик", "Тарууж", "Туун", "Вании", "Ванон", "Вуударадж"},
	"female": {"Ааспар", "Агуус", "Белалуур", "Денаал", "Драраар", "Дууша", "Эхаас", "Элуун", "Грааль",
		"Гадуул", "Хашак", "Джелуум", "Келаал", "Мулаан", "Насри", "Ралин", "Разу", "Рексин", "Сенен",
		"Шедруор", "Таджин", "Тюнер", "Вали", "Вуун"},
}

var raceLangs = []string{"Общий", "Гоблинский"}

var bodyStats = races.BodyStats{
	BodySize:  "Средний",
	MinAge:    16,
	MaxAge:    80,
	MinHeight: 150,
	MaxHeight: 180,
	MinWeight: 70,
	MaxWeight: 90,
}

var raceResists = []string{"Нет"}

func getAddictionInfoInterface() string {
	var D6 = general.D6

	var variants = map[int]string{
		1: "Вы — шпион, посланный, чтобы уничтожить врагов изнутри.",
		2: "Вы — жертва проклятия или заклинания превращение [polymorph].",
		3: "Вы были воспитаны людьми, эльфами, или дварфами и приняли их культуру.",
		4: "В юном возрасте Вы приняли религию людей, и теперь искренне ей служите.",
		5: "Вы получили божественное видение, которое направило вас на этот путь, и иногда вы получаете новые видения, которые вас направляют.",
		6: "Ваш заклятый враг — союзник вашего народа, это вынудило вас оставить ваше племя ради мести.",
	}

	var addictionInfo = "Происхождение Багбира. Краткая история: " + variants[D6.RollDice()]

	return addictionInfo
}

func raceDefaultType() *races.RaceType {
	var typeAbilities = map[string]string{
		"Тёмное зрение": "Привыкнув к жизни под землёй, вы обладаете превосходным зрением в темноте и при тусклом освещении. " +
			"На расстоянии в 60 футов вы при тусклом освещении можете видеть так, как будто это яркое освещение, " +
			"и в темноте так, как будто это тусклое освещение. В темноте вы не можете различать цвета, только оттенки серого.",
		"Длинные конечности":  "Рукопашную атаку, ваша досягаемость + 5 футов.",
		"Мощное телосложение": "При определении вашей грузоподъёмности, вы существо на один размер больше.",
		"Внезапное нападение": "Если вы застаёте цель врасплох и поражаете её атакой в свой первый ход за бой, то плюс 2к6 того же урона, что и атака. Только один раз за бой.",
	}

	var AbilityScorePlusMap = map[string]int{
		"Dexterity": 1,
		"Strength":  2,
	}

	return &races.RaceType{
		TypeRaceName:     "Bugbear",
		TypeRaceNameRu:   "Багбир",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Volo's guide to monsters",
		Standard:         true,
	}
}

func GetRace(gender string) *races.Race {
	return &races.Race{
		RaceName:             "Bugbear",
		RaceNameRu:           "Багбир",
		AddictionInformation: getAddictionInfoInterface(),
		FirstName:            races.GetRaceFirstName(gender, raceNames), //внешняя переменная которая прилетает с фронта
		LastName:             "",
		Gender:               gender,
		Body:                 races.GetBodyStats(bodyStats),
		Language:             raceLangs,
		Resist:               raceResists,
		Type:                 *raceDefaultType(),
	}
}
