package rcs

import "github.com/mazen160/go-random"

type appearanceYuanti struct {
	TypeSnakeBody      string `json:"type_snake_body"`
	HumanoidSkinColor  string `json:"humanoid_skin_color"`
	ScaleColor         string `json:"scale_color"`
	ScalePattern       string `json:"scale_pattern"`
	TongueColor        string `json:"tongue_color"`
	EyeColor           string `json:"eye_color"`
	HeadShape          string `json:"head_shape"`
	PurebredsSpecialty string `json:"purebreds_specialty"`
}

func setAppearanceYuanti(race string) appearanceYuanti {
	if race == "Юань-ти" {
		var ap appearanceYuanti
		ap.TypeSnakeBody = setTypeSnakeBody()
		ap.HumanoidSkinColor = setHumanoidSkinColor()
		ap.ScaleColor = setScaleColor()
		ap.ScalePattern = setScalePattern()
		ap.TongueColor = setTongueColor()
		ap.EyeColor = "setEyesColor()"
		ap.HeadShape = setHeadShape()
		ap.PurebredsSpecialty = setPurebredsSpecialty()
		return ap
	} else {
		return appearanceYuanti{}
	}
}

func setTypeSnakeBody() string {
	rollNum, _ := random.IntRange(1, 20)
	switch {
	case rollNum <= 5:
		return "Жирная"
	case rollNum > 5 && rollNum <= 15:
		return "Нормальная"
	case rollNum > 15 && rollNum <= 20:
		return "Стройная"
	}
	return ""
}

func setHumanoidSkinColor() string {
	rollNum, _ := random.IntRange(1, 21)
	switch {
	case rollNum <= 4:
		return "Тёмно-коричневая"
	case rollNum == 5:
		return "Зелёно-коричневая"
	case rollNum >= 6 && rollNum <= 9:
		return "Светло-коричневая"
	case rollNum >= 10 && rollNum <= 15:
		return "Умеренно-коричневая"
	case rollNum == 16:
		return "Бледно-коричневая"
	case rollNum >= 17 && rollNum <= 18:
		return "Красно-коричневая"
	case rollNum >= 19 && rollNum <= 20:
		return "Желто-коричневая"
	}
	return ""
}
func setScaleColor() string {

	rollNum, _ := random.IntRange(1, 100)
	switch {
	case rollNum <= 6:
		return "Чёрный"
	case rollNum == 5:
		return "Чёрный и коричневый"
	case rollNum >= 7 && rollNum <= 12:
		return "Светло-коричневая"
	case rollNum >= 13 && rollNum <= 18:
		return "Чёрный и зелёный"
	case rollNum >= 19 && rollNum <= 23:
		return "Чёрный и красный"
	case rollNum >= 24 && rollNum <= 26:
		return "Чёрный и белый"
	case rollNum >= 27 && rollNum <= 30:
		return "Чёрный и жёлтый"
	case rollNum >= 31 && rollNum <= 36:
		return "Чёрный, золотой и красный"
	case rollNum >= 37 && rollNum <= 42:
		return "Чёрный, красный и белый"
	case rollNum >= 43 && rollNum <= 45:
		return "Голубой"
	case rollNum >= 46 && rollNum <= 48:
		return "Голубой и чёрный"
	case rollNum >= 49 && rollNum <= 51:
		return "Голубой и серый"
	case rollNum >= 52 && rollNum <= 54:
		return "Голубой и жёлтый"
	case rollNum >= 55 && rollNum <= 60:
		return "Коричневый"
	case rollNum >= 61 && rollNum <= 66:
		return "Коричневый и зелёный"
	case rollNum >= 67 && rollNum <= 73:
		return "Зелёный"
	case rollNum >= 74 && rollNum <= 79:
		return "Зелёный и каштановый"
	case rollNum >= 80 && rollNum <= 84:
		return "Зелёный и белый"
	case rollNum >= 85 && rollNum <= 90:
		return "Зелёный и жёлтый"
	case rollNum >= 91 && rollNum <= 96:
		return "Красный и каштановый"
	case rollNum >= 96 && rollNum <= 100:
		return "Альбинос"
	}
	return ""
}

func setScalePattern() string {
	rollNum, _ := random.IntRange(1, 20)
	switch {
	case rollNum <= 5:
		return "Пёстрый"
	case rollNum >= 6 && rollNum <= 7:
		return "Беспорядочный"
	case rollNum >= 8 && rollNum <= 10:
		return "Сетчатый"
	case rollNum >= 11 && rollNum <= 15:
		return "Крапчатый"
	case rollNum >= 16 && rollNum <= 20:
		return "Полосатый"
	}
	return ""
}

func setTongueColor() string {
	rollNum, _ := random.IntRange(1, 7)

	switch {
	case rollNum == 1:
		return "Чёрный"
	case rollNum == 2:
		return "Синий"
	case rollNum == 3:
		return "Оранжевый"
	case rollNum == 4:
		return "Бледный"
	case rollNum >= 5 && rollNum <= 6:
		return "Красный"
	}
	return ""
}

func setHeadShape() string {

	rollNum, _ := random.IntRange(1, 20)
	switch {
	case rollNum <= 5:
		return "Широкая и круглая"
	case rollNum >= 6 && rollNum <= 9:
		return "Сплющенная"
	case rollNum >= 10 && rollNum <= 11:
		return "С капюшоном"
	case rollNum >= 12 && rollNum <= 15:
		return "Вытянутая"
	case rollNum >= 16 && rollNum <= 20:
		return "Треугольная"
	}
	return ""
}

func setPurebredsSpecialty() string {
	rollNum, _ := random.IntRange(1, 20)

	switch {
	case rollNum <= 3:
		return "Клыки"
	case rollNum >= 4 && rollNum <= 5:
		return "Раздвоенный язык"
	case rollNum >= 6 && rollNum <= 9:
		return "Чешуйчатые руки и кисти рук"
	case rollNum >= 10 && rollNum <= 11:
		return "Чешуйчатое лицо"
	case rollNum >= 12 && rollNum <= 15:
		return "Чешуйчатый торс"
	case rollNum >= 16 && rollNum <= 20:
		return "Змеиные глаза"
	}
	return ""

}
