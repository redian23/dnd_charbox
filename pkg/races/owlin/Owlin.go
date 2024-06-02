package owlin

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/spells"
)

var raceNames = map[string][]string{
	"Мужской": {"Карак", "Аманак", "Ракуун", "Паштуук", "Кертиу", "Дахар", "Менацку",
		"Сажин", "Нашти", "Ресчу", "Падри", "Вазир", "Камтиш"},
	"Женский": {"Ката", "Курше", "Васши", "Керири", "Кентра", "Узашва", "Флая",
		"Курури", "Генра", "Мишита", "Цурка", "Фриса", "Зашра", "Лурана", "Квария"},
}

var raceLangs = []string{"Общий", "Один язык на ваш выбор"}

var bodyStats = races.BodyStats{
	BodySize:  "Средний",
	MinAge:    12,
	MaxAge:    60,
	MinHeight: 120,
	MaxHeight: 150,
	MinWeight: 40,
	MaxWeight: 55,
}

func raceDefaultType() *races.RaceType {
	var typeAbilities = map[string]string{
		"Тёмное зрение": " На расстоянии 120 футов вы при тусклом освещении можете видеть так, " +
			"как будто это яркое освещение, и в темноте так, как будто это тусклое освещение. " +
			"В темноте вы не можете различать цвета, только оттенки серого.",
		"Полёт": "Благодаря вашим крыльям вы обладаете скоростью полёта, равной вашей скорости ходьбы. " +
			"Вы не можете использовать эту скорость полёта, если вы облачены в средние или тяжёлые доспехи.",
	}
	var numd, _ = random.IntRange(2, 4)
	return &races.RaceType{
		TypeRaceName:     "Owlin",
		TypeRaceNameRu:   "Совлин",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: races.GenerateRandomAbilityScoreMap(numd),
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Strixhaven: A Curriculum of Chaos",
		Standard:         false,
	}
}

func getRacePersonalization() races.RacePersonalization {
	return races.RacePersonalization{}
}

func GetRace(gender string) *races.Race {
	return &races.Race{
		RaceName:             "Owlin",
		RaceNameRu:           "Совлин",
		AddictionInformation: "",
		FirstName:            races.GetRaceFirstName(gender, raceNames),
		LastName:             "",
		Gender:               gender,
		Body:                 races.GetBodyStats(bodyStats),
		Language:             raceLangs,
		Resist:               []string{"Нету"},
		Type:                 *raceDefaultType(),
		RaceSkill:            []string{"Stealth"},
	}
}
