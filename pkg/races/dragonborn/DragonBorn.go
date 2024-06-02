package dragonborn

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/races"
	"pregen/pkg/spells"
)

type dragonType struct {
	Color        string `json:"color"`
	DamageType   string `json:"damage"`
	BreathWeapon string `json:"breath_weapon"`
}

var raceNames = map[string][]string{
	"Мужской": {"Арджхан", "Баласар", "Бхараш", "Гхеш", "Донаар", "Крив", "Медраш",
		"Мехен", "Надарр", "Панджед", "Патрин", "Рхогар", "Тархун", "Торинн", "Хескан", "Шамаш", "Шединн"},
	"Женский": {"Акра", "Бири", "Даар", "Джхери", "Кава", "Коринн", "Мисханн",
		"Нала", "Перра", "Райанн", "Сора", "Сурина", "Тхава", "Уаджит", "Фаридэ", "Хавилар", "Харанн"},
}

var lastNames = []string{"Версисатургиеш", "Даардендриан", "Делмирев", "Драчедандион", "Кепешкмолик", "Керрилон",
	"Кимбатуул", "Клестинсиаллор", "Линксакасендалор", "Мястан", "Неммонис", "Нориксиус", "Офиншталажир",
	"Прексижандилин", "Турнурот", "Фенкенкаьрадон", "Шестенделиат", "Яржерит"}

var raceLangs = []string{"Общий", "Драконий язык"}

var bodyStats = races.BodyStats{
	BodySize:  "Средний",
	MinAge:    15,
	MaxAge:    80,
	MinHeight: 180,
	MaxHeight: 100,
	MinWeight: 90,
	MaxWeight: 120,
}

var raceResists []string

func raceDefaultType() *races.RaceType {
	var drTypeArray = []dragonType{
		{Color: "Белый", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Зеленый", DamageType: "Яд", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Бронзовый", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Золотой", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Красный", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Латунный", DamageType: "Огонь", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Медный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Серебряный", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Синий", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Чёрный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	}

	var numd, _ = random.IntRange(0, len(drTypeArray))
	raceResists = append(raceResists, drTypeArray[numd].DamageType)

	var typeAbilities = map[string]string{
		"Наследие драконов": "Вы получаете драконье наследие. " +
			"Ваш цвет: " + drTypeArray[numd].Color + "" +
			"Вы получаете оружие дыхания (" + drTypeArray[numd].BreathWeapon +
			") и сопротивление урону соответствующего вида (" + drTypeArray[numd].DamageType + ").",
		"Оружие дыхания": "Вы можете действием выдохнуть разрушительную энергию. " +
			"Ваше наследие драконов определяет размер, форму и вид урона вашего выдоха. " +
			"Когда вы используете оружие дыхания, все существа в зоне выдоха должны совершить спасбросок, " +
			"вид которого определяется вашим наследием. " +
			"Сложность этого спасброска равна 8 + ваш модификатор Телосложения + ваш бонус мастерства. " +
			"Существа получают урона 2к6 в случае проваленного спасброска, или половину этого урона, " +
			"если спасбросок был успешен. Урон повышается до 3к6 на 6-м уровне, до 4к6 на 11-м, " +
			"и до 5к6 на 16-м уровне. После использования оружия дыхания вы не можете использовать его повторно, " +
			"пока не завершите короткий либо продолжительный отдых.",
	}

	var AbilityScorePlusMap = map[string]int{
		"Strength": 2,
		"Charisma": 1,
	}

	return &races.RaceType{
		TypeRaceName:     "Default",
		TypeRaceNameRu:   "Стандартный",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Player's handbook",
		Standard:         true,
	}
}

func raceDragonbloodType() *races.RaceType {
	var drTypeArray = []dragonType{
		{Color: "Белый", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Зеленый", DamageType: "Яд", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Бронзовый", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Золотой", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Красный", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Латунный", DamageType: "Огонь", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Медный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Серебряный", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Синий", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Чёрный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	}

	var numd, _ = random.IntRange(0, len(drTypeArray))
	raceResists = append(raceResists, drTypeArray[numd].DamageType)

	var typeAbilities = map[string]string{
		"Наследие драконов": "Вы получаете драконье наследие. " +
			"Ваш цвет: " + drTypeArray[numd].Color + "" +
			"Вы получаете оружие дыхания (" + drTypeArray[numd].BreathWeapon +
			") и сопротивление урону соответствующего вида (" + drTypeArray[numd].DamageType + ").",
		"Оружие дыхания": "Вы можете действием выдохнуть разрушительную энергию. " +
			"Ваше наследие драконов определяет размер, форму и вид урона вашего выдоха. " +
			"Когда вы используете оружие дыхания, все существа в зоне выдоха должны совершить спасбросок, " +
			"вид которого определяется вашим наследием. " +
			"Сложность этого спасброска равна 8 + ваш модификатор Телосложения + ваш бонус мастерства. " +
			"Существа получают урона 2к6 в случае проваленного спасброска, или половину этого урона, " +
			"если спасбросок был успешен. Урон повышается до 3к6 на 6-м уровне, до 4к6 на 11-м, " +
			"и до 5к6 на 16-м уровне. После использования оружия дыхания вы не можете использовать его повторно, " +
			"пока не завершите короткий либо продолжительный отдых.",
		"Тёмное зрение": "Вы можете видеть в тусклом свете на 60 футов, как если бы это был яркий свет, а также в темноте, если бы это был тусклый свет. " +
			"Вы не можете различать цвета в темноте, вам видны лишь оттенки серого.",
		"Довлеющее присутствие": "Вы можете использовать собственное толкование изобретательной дипломатии или запугивания, " +
			"чтобы изменить ход дискуссии в свою пользу. Когда вы совершаете проверку Харизмы (Запугивание или Убеждение), " +
			"вы можете совершить её с преимуществом. Как только вы воспользуетесь этой особенностью , вы не можете воспользоваться ею снова, пока не закончите короткий или продолжительный отдых.",
	}

	var AbilityScorePlusMap = map[string]int{
		"Intelligence": 2,
		"Charisma":     1,
	}

	return &races.RaceType{
		TypeRaceName:     "Dragonblood",
		TypeRaceNameRu:   "Драконокровный",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Explorer’s Guide to Wildemount",
		Standard:         true,
	}
}

func raceRaveniteType() *races.RaceType {
	var drTypeArray = []dragonType{
		{Color: "Белый", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Зеленый", DamageType: "Яд", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Бронзовый", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Золотой", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Красный", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Латунный", DamageType: "Огонь", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Медный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Серебряный", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
		{Color: "Синий", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Чёрный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	}

	var numd, _ = random.IntRange(0, len(drTypeArray))
	raceResists = append(raceResists, drTypeArray[numd].DamageType)

	var typeAbilities = map[string]string{
		"Наследие драконов": "Вы получаете драконье наследие. " +
			"Ваш цвет: " + drTypeArray[numd].Color + "" +
			"Вы получаете оружие дыхания (" + drTypeArray[numd].BreathWeapon +
			") и сопротивление урону соответствующего вида (" + drTypeArray[numd].DamageType + ").",
		"Оружие дыхания": "Вы можете действием выдохнуть разрушительную энергию. " +
			"Ваше наследие драконов определяет размер, форму и вид урона вашего выдоха. " +
			"Когда вы используете оружие дыхания, все существа в зоне выдоха должны совершить спасбросок, " +
			"вид которого определяется вашим наследием. " +
			"Сложность этого спасброска равна 8 + ваш модификатор Телосложения + ваш бонус мастерства. " +
			"Существа получают урона 2к6 в случае проваленного спасброска, или половину этого урона, " +
			"если спасбросок был успешен. Урон повышается до 3к6 на 6-м уровне, до 4к6 на 11-м, " +
			"и до 5к6 на 16-м уровне. После использования оружия дыхания вы не можете использовать его повторно, " +
			"пока не завершите короткий либо продолжительный отдых.",
		"Тёмное зрение":         "Вы можете видеть в тусклом свете на 60 футов, как если бы это был яркий свет, а также в темноте, если бы это был тусклый свет. Вы не можете различать цвета в темноте, вам видны лишь оттенки серого.",
		"Мстительное нападение": "Когда вы получаете урон от существа, находящегося в пределах досягаемости вашего оружия, которое вы держите в ваших руках, вы можете реакцией совершить атаку этим оружием по этому существу.\n\nКак только вы используете эту особенность, вы не можете воспользоваться ею снова, пока не закончите короткий или продолжительный отдых.",
	}

	var AbilityScorePlusMap = map[string]int{
		"Strength":       2,
		"BodyDifficulty": 1,
	}

	return &races.RaceType{
		TypeRaceName:     "Ravenite",
		TypeRaceNameRu:   "Равенит",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: AbilityScorePlusMap,
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Volo's guide to monsters",
		Standard:         true,
	}
}

func raceMetalType() *races.RaceType {
	var drTypeArray = []dragonType{
		{Color: "Бронзовый", DamageType: "Электричество", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Золотой", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Латунный", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Медный", DamageType: "Кислота", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Серебряный", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
	}

	var numd, _ = random.IntRange(0, len(drTypeArray))

	raceResists = append(raceResists, drTypeArray[numd].DamageType)
	var typeAbilities = map[string]string{
		"Наследие драконов": "Вы получаете драконье наследие. " +
			"Ваш метелл: " + drTypeArray[numd].Color + "" +
			"Вы получаете оружие дыхания (" + drTypeArray[numd].BreathWeapon +
			") и сопротивление урону соответствующего вида (" + drTypeArray[numd].DamageType + ").",
		"Оружие дыхания": "Когда в свой ход вы совершаете действие Атака, " +
			"вы можете заменить одну из ваших атак на выдыхание 15-футового конуса магической энергии. " +
			"Все существа в области выдоха должны совершить спасбросок Ловкости " +
			"<strong>(Сл = 8 + ваш модификатор Телосложения + ваш бонус мастерства).</strong> " +
			"При провале существа получают 1к10 урона, вид которого определяется " +
			"вашим «Наследием металлического дракона»." +
			"При успешном спасброске существо получает вдвое меньше урона. " +
			"Урон повышается на 1к10 когда вы достигаете 5-го уровня (2к10), " +
			"11-го уровня (3к10) и 17-го уровня (4к10)." +
			"Вы можете использовать это Оружие дыхания количество раз, равное вашему бонусу мастерства, " +
			"и вы восстанавливаете все израсходованные применения после окончания продолжительного отдыха.",
		"Металлическое оружие дыхания": "На 5-м уровне вы получаете второе оружие дыхания. Когда в свой ход вы совершаете действие Атака, вы можете заменить одну из ваших атак на выдыхание 15-футового конуса. Сл спасброска = 8 + ваш модификатор Телосложения + ваш бонус мастерства. Всякий раз, когда вы используете эту особенность, выбирайте одно из:\n\nИзнуряющее дыхание. Каждое существо в конусе должно преуспеть в спасброске Телосложения или стать недееспособным до начала вашего следующего хода.\nОтталкивающее дыхание. Каждое существо в конусе должно преуспеть в спасброске Силы. При провале существо отталкивается от вас на 20 футов от вас и падает ничком.\nКак только вы используете своё Металлическое оружие дыхания, вы не можете использовать его повторно, пока не завершите продолжительный отдых.",
	}

	raceLangs = []string{"Общий", "Один язык на ваш выбор"}

	var num, _ = random.IntRange(2, 4)
	return &races.RaceType{
		TypeRaceName:     "Metal",
		TypeRaceNameRu:   "Металлический",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: races.GenerateRandomAbilityScoreMap(num),
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Player's handbook",
		Standard:         true,
	}
}

func raceGemType() *races.RaceType {
	var drTypeArray = []dragonType{
		{Color: "Аметистовый", DamageType: "Силовое поле", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Изумрудный", DamageType: "Психическая энергия", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Кристальный", DamageType: "Излучение", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Сапфировый", DamageType: "Звук", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
		{Color: "Топазный", DamageType: "Некротическая энергия", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
	}

	var numd, _ = random.IntRange(0, len(drTypeArray))

	raceResists = append(raceResists, drTypeArray[numd].DamageType)
	var typeAbilities = map[string]string{
		"Наследие драконов": "Вы получаете драконье наследие. " +
			"Ваш самоцвет: " + drTypeArray[numd].Color + "" +
			"Вы получаете оружие дыхания (" + drTypeArray[numd].BreathWeapon +
			") и сопротивление урону соответствующего вида (" + drTypeArray[numd].DamageType + ").",
		"Оружие дыхания": "Когда в свой ход вы совершаете действие Атака, вы можете заменить одну из ваших атак на выдыхание 15-футового конуса магической энергии. Все существа в области выдоха должны совершить спасбросок Ловкости (Сл = 8 + ваш модификатор Телосложения + ваш бонус мастерства). При провале существа получают 1к10 урона, вид которого определяется вашим «Наследием самоцветного дракона»." +
			"При успешном спасброске существо получает вдвое меньше урона. Урон повышается на 1к10 когда вы достигаете 5-го уровня (2к10), 11-го уровня (3к10) и 17-го уровня (4к10)." +
			"Вы можете использовать это Оружие дыхания количество раз, равное вашему бонусу мастерства, и вы восстанавливаете все израсходованные применения после окончания продолжительного отдыха.",
		"Псионический разум":         "Вы можете посылать телепатические сообщения любому существу, которое можете видеть в пределах 30 футов. Вам не надо знать язык существа, чтобы оно понимало ваши телепатические сообщения, но существо должно быть в состоянии понимать по крайней мере один язык.",
		"Полёт самоцветного дракона": "Начиная с 5-го уровня, вы можете бонусным действием проявить на своем теле спектральные крылья. Эти крылья существуют 1 минуту. В течении этого времени вы получаете скорость полёта, равную вашей скорости ходьбы, и способность парить. Как только вы используете эту особенность, вы не можете использовать её снова, пока не закончите продолжительный отдых.",
	}

	raceLangs = []string{"Общий", "Один язык на ваш выбор"}

	var num, _ = random.IntRange(2, 4)
	return &races.RaceType{
		TypeRaceName:     "Gem",
		TypeRaceNameRu:   "Самоцветный",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: races.GenerateRandomAbilityScoreMap(num),
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Player's handbook",
		Standard:         true,
	}
}

func raceColorType() *races.RaceType {
	var drTypeArray = []dragonType{
		{Color: "Белый", DamageType: "Холод", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Зеленый", DamageType: "Яд", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Красный", DamageType: "Огонь", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Синий", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
		{Color: "Чёрный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	}

	var numd, _ = random.IntRange(0, len(drTypeArray))

	raceResists = append(raceResists, drTypeArray[numd].DamageType)
	var typeAbilities = map[string]string{
		"Наследие драконов": "Вы получаете драконье наследие. " +
			"Ваш цвет: " + drTypeArray[numd].Color + "" +
			"Вы получаете оружие дыхания (" + drTypeArray[numd].BreathWeapon +
			") и сопротивление урону соответствующего вида (" + drTypeArray[numd].DamageType + ").",
		"Оружие дыхания":          "Когда в свой ход вы совершаете действие Атака, вы можете заменить одну из ваших атак на выдох магической энергии, формируя 30-футовую линию шириной 5 футов. Все существа в области выдоха должны совершить спасбросок Ловкости (Сл = 8 + ваш модификатор Телосложения + ваш бонус мастерства). При провале существа получают 1к10 урона, вид которого определяется вашим «Наследием цветного дракона».\n\nПри успешном спасброске существо получает вдвое меньше урона. Урон повышается на 1к10  когда вы достигаете 5-го уровня (2к10), 11-го уровня (3к10) и 17-го уровня (4к10).\n\nВы можете использовать это Оружие дыхания количество раз, равное вашему бонусу мастерства, и вы восстанавливаете все израсходованные применения после окончания продолжительного отдыха.",
		"Защита цветного дракона": "Начиная с 5-го уровня, действием, вы можете направить драконью энергию, на свою защиту. На 1 минуту вы получаете иммунитет к урону того вида, который указан в вашем «Наследии цветного дракона». Как только вы используете эту особенность, вы не можете использовать её снова, пока не закончите продолжительный отдых.",
	}

	raceLangs = []string{"Общий", "Один язык на ваш выбор"}

	var num, _ = random.IntRange(2, 4)
	return &races.RaceType{
		TypeRaceName:     "Color",
		TypeRaceNameRu:   "Цветной",
		RaceAbilities:    typeAbilities,
		AbilityScorePlus: races.GenerateRandomAbilityScoreMap(num),
		Speed:            30,
		RaceSpellsList:   []spells.SpellsJSON{},
		Book:             "Player's handbook",
		Standard:         true,
	}
}

func getRaceType(raceTypeName string) *races.RaceType {
	switch raceTypeName {
	case "Стандартный":
		return raceDefaultType()
	case "Драконокровный":
		return raceDragonbloodType()
	case "Равенит":
		return raceRaveniteType()
	case "Металлический":
		return raceMetalType()
	case "Самоцветный":
		return raceGemType()
	case "Цветной":
		return raceColorType()
	}
	return nil
}

func getRacePersonalization() races.RacePersonalization {
	return races.RacePersonalization{}
}

func GetRace(raceTypeName, gender string) *races.Race {
	randNum, _ := random.IntRange(0, len(lastNames))
	return &races.Race{
		RaceName:             "Dragonborn",
		RaceNameRu:           "Драконорождённый",
		AddictionInformation: "",
		FirstName:            races.GetRaceFirstName(gender, raceNames),
		LastName:             "из клана " + lastNames[randNum],
		Gender:               gender,
		Body:                 races.GetBodyStats(bodyStats),
		Language:             raceLangs,
		Resist:               raceResists,
		Type:                 *getRaceType(raceTypeName),
		RaceSkill:            []string{""},
	}
}
