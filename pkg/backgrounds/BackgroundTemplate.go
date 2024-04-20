package backgrounds

//import (
//	"github.com/mazen160/go-random"
//	"pregen/pkg/general"
//)
//
//var description = ""
//
//func getBackgroundMasteryOfTools() []string {
//	var musicalInstruments = []string{}
//	randomInstrument, _ := random.IntRange(0, len(musicalInstruments))
//	return []string{"Набор для грима", "Игра на " + musicalInstruments[randomInstrument]}
//}
//
//var artistAbility = map[string]string{
//}
//
//func getBackgroundItems() []backgrounds.Item {
//	var musicalInstruments = []string{
//	}
//	randomInstrument, _ := random.IntRange(0, len(musicalInstruments))
//
//	var fanPresent = []string{"любовное письмо", "локон волос", "безделушка"}
//	randomPresent, _ := random.IntRange(0, len(fanPresent))
//
//	return []backgrounds.Item{
//		{
//			Name:     "Костюм",
//			ItemType: "Tools",
//			Count:    1,
//		},
//		{
//			Name:     musicalInstruments[randomInstrument],
//			ItemType: "Tools",
//			Count:    1,
//		},
//		{
//			Name:     "Подарок от поклонницы\\поклонника: " + fanPresent[randomPresent],
//			ItemType: "Tools",
//			Count:    1,
//		},
//	}
//}
//
//func getBackgroundSpecificType() string {
//	var artistsTypes = map[int]string{
//		1:  "Акробат",
//		2:  "Актер",
//		3:  "Жонглер",
//		4:  "Музыкант",
//		5:  "Певец",
//		6:  "Пожиратель огня",
//		7:  "Поэт",
//		8:  "Рассказчик",
//		9:  "Танцор",
//		10: "Шут",
//	}
//
//	var artistTricks string
//
//	var iter int
//	tricksCount, _ := random.IntRange(1, 3)
//	for _, truck := range artistsTypes {
//		artistTricks += truck + ", "
//		iter++
//		if iter == tricksCount {
//			break
//		}
//	}
//
//	return artistTricks
//}
//
//func getArtistPersonality() backgrounds.Personalization {
//	characterTrait := map[int]string{
//		1: "Для любой ситуации я найду подходящий рассказ.",
//		2: "Куда бы я ни пришёл, я начинаю собирать местные слухи и распространять сплетни.",
//		3: "Я безнадёжный романтик, всегда ищущий «кого-то особого».",
//		4: "Никто не сердится на меня или возле меня подолгу, так как я могу разрядить любую напряжённую обстановку.",
//		5: "Мне нравятся оскорбления, даже если они направлены на меня.",
//		6: "Мне обидно, если я не нахожусь в центре внимания.",
//		7: "Превыше всего я ценю совершенство.",
//		8: "Моё настроение и намерения меняются так же быстро как ноты в музыке.",
//	}
//	randomTrait := general.D8.RollDice()
//
//	ideals := []backgrounds.Ideal{
//		{
//			Worldview:   "Lawful Good",
//			WorldviewRu: "Законопослушный добрый",
//			Text:        "Патриот. Я люблю выступать в поддержку своей страны.",
//		},
//		{
//			Worldview:   "Neutral Good",
//			WorldviewRu: "Нейтральный добрый",
//			Text:        "Красота. Выступая, я делаю мир лучше.",
//		},
//		{
//			Worldview:   "Chaotic Good",
//			WorldviewRu: "Хаотический добрый",
//			Text:        "Абстракция. Этому миру не хватает ярких красок.",
//		},
//		{
//			Worldview:   "Lawful Neutral",
//			WorldviewRu: "Законопослушный нейтральный",
//			Text:        "Традиции. Рассказы, легенды и песни прошлого не должны забываться, так как они учат нас тому, кто мы такие.",
//		},
//		{
//			Worldview:   "Neutral",
//			WorldviewRu: "Истинно нейтральный",
//			Text:        "Народ. Мне нравится видеть улыбки на лицах во время выступления. В этом-то всё дело.",
//		},
//		{
//			Worldview:   "Chaotic Neutral",
//			WorldviewRu: "Хаотический нейтральный",
//			Text:        "Шут. Главное чтобы весело было мне.",
//		},
//		{
//			Worldview:   "Lawful Evil",
//			WorldviewRu: "Законопослушный злой",
//			Text:        "Жадность. Я занимаюсь всем этим ради денег и славы.",
//		},
//		{
//			Worldview:   "Neutral Evil",
//			WorldviewRu: "Нейтральный злой",
//			Text:        "Нарциссизм. Я прекрасен.",
//		},
//		{
//			Worldview:   "Chaotic Evil",
//			WorldviewRu: "Хаотический злой",
//			Text:        "Творчество. Миру нужны новые идеи и смелые действия.",
//		},
//		{
//			Worldview:   "Any",
//			WorldviewRu: "Любой",
//			Text:        "Честность. Искусство должно отражать душу; оно должно идти изнутри и показывать, чем мы являемся.",
//		},
//	}
//	randomIdeal, _ := random.IntRange(0, len(ideals))
//
//	affections := map[int]string{
//		1: "Мой инструмент — самое драгоценное, что у меня есть, и он напоминает мне о моей любви.",
//		2: "Кто-то украл мой драгоценный инструмент, и когда-нибудь я верну его.",
//		3: "Я хочу быть знаменитым, чего бы это ни стоило.",
//		4: "Я боготворю героя старых рассказов, и всегда сравниваю свои поступки с его.",
//		5: "Я всё сделаю, чтобы доказать превосходство над ненавистным конкурентом.",
//		6: "Я сделаю что угодно для других членов моей старой труппы.",
//	}
//	randomAffection := general.D6.RollDice()
//
//	weakness := map[int]string{
//		1: "Я пойду на всё ради славы и известности.",
//		2: "Не могу устоять перед смазливым личиком.",
//		3: "Я не могу вернуться домой из-за скандала. Неприятности такого рода словно преследуют меня.",
//		4: "Однажды я высмеял дворянина, который всё еще хочет оторвать мне голову. Это была ошибка, но я могу повторить её еще неоднократно.",
//		5: "Я не могу скрывать свои истинные чувства. Острый язык всегда приносит мне неприятности.",
//		6: "Я очень стараюсь исправиться, но друзьям не стоит на меня полагаться.",
//	}
//	randomWeakness := general.D6.RollDice()
//
//	return backgrounds.Personalization{
//		PersonalizationDescription: "Успешные артисты могут овладевать вниманием публики, поэтому у них яркий или напористый характер. " +
//			"Они могут быть романтичными, и в искусстве и красоте часто обращаются к возвышенным идеалам.",
//		Advice:         "",
//		CharacterTrait: characterTrait[randomTrait],
//		Ideal:          ideals[randomIdeal],
//		Affection:      affections[randomAffection],
//		Weakness:       weakness[randomWeakness],
//		Worldview:      ideals[randomIdeal].WorldviewRu,
//	}
//}
//
//func GetArtistBackground() *backgrounds.Background {
//	return &backgrounds.Background{
//		BackgroundName:         "Artist",
//		BackgroundNameRu:       "Артист",
//		BackgroundSpecificType: getBackgroundSpecificType(),
//		BackgroundAbility:      artistAbility,
//		Description:            description,
//		BackgroundItems:        getBackgroundItems(),
//		BackgroundSkills:       []string{"Acrobatics", "Performance"},
//		MasteryOfTools:         getBackgroundMasteryOfTools(),
//		Personalization:        getArtistPersonality(),
//		Gold:                   15,
//	}
//}
