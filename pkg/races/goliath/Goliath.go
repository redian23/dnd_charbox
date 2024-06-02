package goliath

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/spells"
)

var raceNames = map[string][]string{
	"Мужской": {"Аукан", "Ваунеа", "Вимак", "Гае-Ал", "Иглат", "Иликан", "Кеотхи", "Куори",
		"Ло-Каг", "Мавейт", "Маннео", "Налла", "Орило", "Пааву", "Петани", "Талай", "Тхотам", "Утхал"},
	"Женский": {"Аукан", "Ваунеа", "Вимак", "Гае-Ал", "Иглат", "Иликан", "Кеотхи", "Куори",
		"Ло-Каг", "Мавейт", "Маннео", "Налла", "Орило", "Пааву", "Петани", "Талай", "Тхотам", "Утхал"},
}

var nickNames = map[string][]string{
	"Мужской": {"Бесстрашный", "Встающий с рассветом", "Дважды осиротевший", "Зоркий глаз", "Искатель кремня",
		"Красиво говорящий", "Кривые руки", "Одинокий охотник", "Прядильщик", "Разрубивший корни", "Резчик рогов",
		"Смотрящий в небо", "Идущий к реке", "Твёрдая рука", "Убийца медведей", "Умелый прыгун"},
	"Женский": {"Бесстрашная", "Встающяя с рассветом", "Дважды осиротевшая", "Зоркий глаз", "Искательница кремня",
		"Красиво говорящая", "Кривые руки", "Одинокая охотницы", "Мастерица пряжи", "Разрубившая корни", "Резчица рогов",
		"Смотрящяя в небо", "Идущая к реке", "Твёрдая рука", "Убийца медведей", "Умелая прыгунья"},
}

var lastNames = []string{"Анакалатай", "Ваймей-Лага", "Гатаканати", "Калагиано", "Като-Олави",
	"Кола-Гилеана", "Тулиага", "Тунукалати", "Эланитино"}
var raceLangs = []string{"Общий", "Язык Великанов"}

var bodyStats = races.BodyStats{
	BodySize:  "Средний",
	MinAge:    22,
	MaxAge:    90,
	MinHeight: 210,
	MaxHeight: 250,
	MinWeight: 125,
	MaxWeight: 160,
}

var raceResists = []string{"Холоду"}

func raceDefaultType() *races.RaceType {
	var typeAbilities = map[string]string{
		"Выносливость камня": "Вы можете сосредоточиться, чтобы уменьшить полученные повреждения. " +
			"При получении урона вы можете реакцией бросить к12 и прибавить к результату модификатор Телосложения, " +
			"а после уменьшить полученный урон на итоговую сумму. Вы не сможете вновь воспользоваться этой особенностью, " +
			"пока не закончите короткий или продолжительный отдых.",
		"Сильные руки": "При определении вашей грузоподъёмности, а также веса, который вы можете толкать, " +
			"тянуть и поднимать, вы считаетесь существом на один размер больше.",
		"Рождённый в горах": "Вы получаете сопротивление урону холодом. Кроме того, вы акклиматизированы к большой высоте, " +
			"включая высоту более 20 000 футов (6 километров). " +
			"Вы также адаптированы к холодным климатическим условиям, как описано в главе 5 «Руководства Мастера».",
	}

	var AbilityScorePlusMap = map[string]int{
		"BodyDifficulty": 1,
		"Strength":       2,
	}

	return &races.RaceType{
		TypeRaceName:     "Goliath",
		TypeRaceNameRu:   "Голиаф",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Volo's guide to monsters",
		Standard:         true,
	}
}

func getRacePersonalization() races.RacePersonalization {
	return races.RacePersonalization{}
}

func getFirstName(gender string) string {
	var numd, _ = random.IntRange(0, 2)
	if numd == 0 {
		return races.GetRaceFirstName(gender, raceNames)
	} else {
		return races.GetRaceFirstName(gender, nickNames)
	}
}

func GetRace(gender string) *races.Race {

	randNum, _ := random.IntRange(0, len(lastNames))
	return &races.Race{
		RaceName:             "Goliath",
		RaceNameRu:           "Голиаф",
		AddictionInformation: "",
		FirstName:            getFirstName(gender),
		LastName:             "из клана " + lastNames[randNum],
		Gender:               gender,
		Body:                 races.GetBodyStats(bodyStats),
		Language:             raceLangs,
		Resist:               raceResists,
		Type:                 *raceDefaultType(),
		RaceSkill:            []string{"Athletics"},
	}
}
