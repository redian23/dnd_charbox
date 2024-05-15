package goblin

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/general"
	"pregen/pkg/races"
	"pregen/pkg/spells"
)

var differentName string
var raceNames = map[string][]string{
	"Мужской": {"Аругет", "Вании", "Ванон", "Вуударадж", "Гудруун", "Даавн", "Дабрак", "Дагии", "Джазаал",
		"Древдуул", "Дуулан", "Каллаад", "Кракуул", "Кроотад", "Мазаан", "Мунта", "Насаар", "Ракари",
		"Рексиит", "Тарик", "Тарууж", "Туун", "Феник", "Четин", "Халуун", "Харуук"},
	"Женский": {"Ааспар", "Агуус", "Белалуур", "Вали", "Вуун", "Грааль", "Гадуул", "Денаал",
		"Джелуум", "Драраар", "Дууша", "Келаал", "Мулаан", "Разу", "Ралин", "Рексин",
		"Сенен", "Таджин", "Тюнер", "Хашак", "Шедруор", "Эхаас", "Элуун"},
}

var raceLangs = []string{"Общий", "Гоблинский"}

var bodyStats = races.BodyStats{
	BodySize:  "Маленький",
	MinAge:    8,
	MaxAge:    60,
	MinHeight: 90,
	MaxHeight: 120,
	MinWeight: 18,
	MaxWeight: 40,
}

var raceResists = []string{}

func getAddictionInfoInterface() string {
	var D8 = general.D8

	var angelicGuideName = map[int]string{
		1: "Вы — шпион, посланный, чтобы уничтожить врагов изнутри.",
		2: "Вы — жертва проклятия или заклинания превращение [polymorph].",
		3: "Вы были воспитаны людьми, эльфами, или дварфами и приняли их культуру.",
		4: "В юном возрасте вы приняли религию людей, и теперь искренне ей служите.",
		5: "Вы получили божественное видение, которое направило вас на этот путь, и иногда вы получаете новые видения, которые вас направляют.",
		6: "Ваш заклятый враг — союзник вашего народа, это вынудило вас оставить ваше племя ради мести.",
		7: "Злая сущность развратила общество Вашего народа.",
		8: "Травма или необычное событие привели к полной потере памяти о своём прошлом, но иногда к Вам приходят обрывки памяти.",
	}

	var addictionInfoTest = "" + angelicGuideName[D8.RollDice()]

	return addictionInfoTest
}

func raceDefaultType() *races.RaceType {
	var typeAbilities = map[string]string{
		"Тёмное зрение": "На расстоянии в 60 футов вы при тусклом освещении можете видеть так, " +
			"как будто это яркое освещение, и в темноте так, как будто это тусклое освещение. " +
			"В темноте вы не можете различать цвета, только оттенки серого.",
		"Разъяренная мелкота": "Когда вы наносите существу урон атакой или заклинанием, " +
			"и размер существа больше чем ваш, вы можете этой атакой или " +
			"заклинанием нанести существу дополнительный урон, равный вашему уровню. " +
			"Использовав эту особенность, вы не можете использовать её повторно, " +
			"пока не завершите короткий или продолжительный отдых.",
		"Шустрый побег": "Вы можете совершать действия Отход или Засада бонусным действием в каждый ваш ход.",
	}

	var AbilityScorePlusMap = map[string]int{
		"BodyDifficulty": 1,
		"Dexterity":      2,
	}

	return &races.RaceType{
		TypeRaceName:     "Goblin",
		TypeRaceNameRu:   "Гоблин",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Volo's guide to monsters",
		Standard:         true,
	}
}
func raceRLWType() *races.RaceType {

	var typeAbilities = map[string]string{
		"Тёмное зрение": "Привыкнув к жизни под землёй, вы обладаете превосходным зрением в темноте и при тусклом освещении. " +
			"На расстоянии в 60 фт. вы при тусклом освещении можете видеть так, как будто это яркое освещение, " +
			"и в темноте так, как будто это тусклое освещение. В темноте вы не можете различать цвета, только оттенки серого.",
		"Ярость мелкого": "Когда вы наносите существу урон атакой или заклинанием, а размер существа больше вашего, " +
			"вы можете этой атакой или заклинанием нанести существу дополнительный урон, равный вашему уровню. " +
			"Как только вы используете эту особенность, вы не можете использовать ее снова, пока не закончите короткий или продолжительный отдых.",
		"Проворное бегство": "Вы можете совершать действие Отход или Засада в качестве бонусного действия в каждом из ваших ходов.",
	}

	var AbilityScorePlusMap = map[string]int{
		"BodyDifficulty": 1,
		"Dexterity":      2,
	}

	return &races.RaceType{
		TypeRaceName:     "Goblin (RLW)",
		TypeRaceNameRu:   "Гоблин (RLW)",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
	}
}
func raceGGRType() *races.RaceType {
	var names = []string{"Аззинакс", "Баболакс", "Бликсаникс", "Вазозав", "Вексин", "Даззаз", "Зиззикс",
		"Калузакс", "Криксизикс", "Лизакса", "Миззикс", "Мизнар", "Никсиспикс", "Паксизаз", "Равиксиз",
		"Стиксил", "Сунникс", "Тозинокс", "Уксивози", "Эстрикс", "Финизикс", "Юзба"}
	randNum, _ := random.IntRange(0, len(names))
	differentName = names[randNum]

	var typeAbilities = map[string]string{
		"Тёмное зрение": "Привыкнув к жизни под землёй, вы обладаете превосходным зрением в темноте и при тусклом освещении. " +
			"На расстоянии в 60 фт. вы при тусклом освещении можете видеть так, как будто это яркое освещение, и в темноте так, " +
			"как будто это тусклое освещение. В темноте вы не можете различать цвета, только оттенки серого.",
		"Разъяренная мелкота": "Когда вы наносите существу урон атакой или заклинанием, " +
			"и размер существа больше чем ваш, вы можете этой атакой или " +
			"заклинанием нанести существу дополнительный урон, равный вашему уровню. " +
			"Использовав эту особенность, вы не можете использовать её повторно, " +
			"пока не завершите короткий или продолжительный отдых.",
		"Шустрый побег": "Вы можете совершать действия Отход или Засада бонусным действием в каждый ваш ход.",
	}

	var AbilityScorePlusMap = map[string]int{
		"BodyDifficulty": 1,
		"Dexterity":      2,
	}

	return &races.RaceType{
		TypeRaceName:     "Goblin (GGR)",
		TypeRaceNameRu:   "Гоблин (GGR)",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Guildmasters’ Guide to Ravnica",
		Standard:         false,
	}
}
func raceDunkwoodType() *races.RaceType {
	var typeAbilities = map[string]string{
		"Тёмное зрение": "На расстоянии в 60 футов вы при тусклом освещении можете видеть так, " +
			"как будто это яркое освещение, и в темноте так, как будто это тусклое освещение. " +
			"В темноте вы не можете различать цвета, только оттенки серого.",
		"Общение с маленькими зверями": " С помощью звуков и жестов вы можете передавать простые понятия Маленьким или " +
			"ещё меньшим зверям. Гоблины Данквуда любят животных и часто держат белок, барсуков, " +
			"кроликов, кротов, дятлов и других животных в качестве питомцев.",
		"Шустрый побег": "Вы можете совершать действия Отход или Засада бонусным действием в каждый ваш ход.",
	}

	var AbilityScorePlusMap = map[string]int{
		"Wisdom":    1,
		"Dexterity": 2,
	}

	return &races.RaceType{
		TypeRaceName:     "Dunkwood Goblins",
		TypeRaceNameRu:   "Гоблин Данквуда",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "",
		Standard:         false,
	}
}

func getRacePersonalization() races.RacePersonalization {
	return races.RacePersonalization{}
}

func GetRace(raceTypeName string, gender string) *races.Race {
	var raceType *races.RaceType
	var yourName string

	switch raceTypeName {
	case "Гоблин":
		raceType = raceDefaultType()
		yourName = races.GetRaceFirstName(gender, raceNames)
	case "Гоблин (RLW)":
		raceType = raceRLWType()
		yourName = differentName
	case "Гоблин (GGR)":
		raceType = raceGGRType()
		yourName = races.GetRaceFirstName(gender, raceNames)
	case "Гоблин Данквуда":
		raceType = raceDunkwoodType()
		yourName = races.GetRaceFirstName(gender, raceNames)
	}

	return &races.Race{
		RaceName:             "Race",
		RaceNameRu:           "Раса",
		AddictionInformation: getAddictionInfoInterface(),
		FirstName:            yourName,
		LastName:             "",
		Gender:               gender,
		Body:                 races.GetBodyStats(bodyStats),
		Language:             raceLangs,
		Resist:               raceResists,
		Type:                 *raceType,
	}
}
