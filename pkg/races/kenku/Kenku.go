package kenku

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
var nickNames = map[string][]string{
	"Мужской": {"Крысиный Скрежет", "Свистун", "Мышатник", "Рычун", "Лязг Молота", "Хлопок Паруса", "Скрипучий Пол",
		"Ворон Каркающий", "Шорох Листьев", "Пыхтящий Котел", "Звон Колокола", "Треск Пламени", "Шелест Травы",
		"Гул Ветра", "Рёв Мотор", "Гром Грозы", "Плеск Волны", "Стон Дерева", "Шум Прибоя", "Скрежет Перьев"},
	"Женский": {"Звон Монет", "Шелест Шёлка", "Песнь Птицы", "Хлопок Крыльев", "Стон Ветра", "Шёпот Ночи",
		"Мурлыканье Кошки", "Писк Мыши", "Трель Соловья", "Вздох Листьев", "Бульканье Воды", "Чириканье",
		"Звук Песка", "Шёпот Воды", "Капель Роса", "Шелест Бумаги", "Стук Сердца", "Звяканье Цепи", "Шорох Камней"},
}

var raceLangs = []string{"Общий", "Ауран (воздушный диалект Первичного)"}

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
		"Мастер подделок": "Вы можете копировать почерк и ремесло других существ. " +
			"Вы получаете преимущество на все проверки, проводимые для создания подделки или копии существующего объекта",
		"Имитирование": "Вы можете подражать звукам, которые вы слышали, включая голоса. " +
			"Существо, которое слышит звуки, которые вы издаете, может определить " +
			"что это имитация при успешной проверке Мудрости (Проницательность), против вашей проверки Харизмы (Обман).",
	}

	var AbilityScorePlusMap = map[string]int{
		"Dexterity": 2,
		"Wisdom":    1,
	}

	return &races.RaceType{
		TypeRaceName:     "Kenku",
		TypeRaceNameRu:   "Кенку",
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

func getRaceSkills() []string {
	var raceSkillsVariants = []string{"Acrobatics", "Deception", "Stealth", "SleightOfHand"}
	var raceSkills []string
	var iter int

	for _, value := range raceSkillsVariants {
		raceSkills = append(raceSkills, value)
		iter++
		if iter == 2 {
			break
		}
	}
	return raceSkills
}

func GetRace(gender string) *races.Race {
	return &races.Race{
		RaceName:             "Kenku",
		RaceNameRu:           "Кенку",
		AddictionInformation: "",
		FirstName:            getFirstName(gender),
		LastName:             "",
		Gender:               gender,
		Body:                 races.GetBodyStats(bodyStats),
		Language:             raceLangs,
		Resist:               []string{"Нету"},
		Type:                 *raceDefaultType(),
		RaceSkill:            getRaceSkills(),
	}
}
