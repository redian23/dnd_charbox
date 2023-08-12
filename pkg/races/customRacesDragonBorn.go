package races

import (
	"github.com/mazen160/go-random"
)

type dragonType struct {
	Color        string `json:"color"`
	DamageType   string `json:"damage"`
	BreathWeapon string `json:"breath_weapon"`
}

var colorTypeArray = []dragonType{
	{Color: "Белый", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
	{Color: "Зеленый", DamageType: "Яд", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
	{Color: "Красный", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
	{Color: "Синий", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	{Color: "Чёрный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
}

var metalTypeArray = []dragonType{
	{Color: "Бронзовый", DamageType: "Электричество", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	{Color: "Золотой", DamageType: "Огонь", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
	{Color: "Латунный", DamageType: "Огонь", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	{Color: "Медный", DamageType: "Кислота", BreathWeapon: "Линия 5 на 30 футов (спас. ЛОВ)"},
	{Color: "Серебряный", DamageType: "Холод", BreathWeapon: "15-футовый конус (спас. ТЕЛ)"},
}

var gemstoneTypeArray = []dragonType{
	{Color: "Аметистовый", DamageType: "Силовое поле", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
	{Color: "Изумрудный", DamageType: "Психическая энергия", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
	{Color: "Кристальный", DamageType: "Излучение", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
	{Color: "Сапфировый", DamageType: "Звук", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
	{Color: "Топазный", DamageType: "Некротическая энергия", BreathWeapon: "15-футовый конус (спас. ЛОВ)"},
}

func setDragonType(dragonTypeName string) dragonType {
	var typeArray []dragonType

	switch dragonTypeName {
	case "Металлический драконорождённый":
		typeArray = metalTypeArray
	case "Самоцветный драконорождённый":
		typeArray = gemstoneTypeArray
	case "Цветной драконорождённый":
		typeArray = colorTypeArray
	default:
		return dragonType{}
	}

	randElem, _ := random.IntRange(0, len(typeArray))
	dragon := typeArray[randElem]

	return dragon
}
