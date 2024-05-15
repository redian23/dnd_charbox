package sage

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/backgrounds"
	"pregen/pkg/general"
)

var description = "Вы провели много лет, постигая тайны мультивселенной. " +
	"Вы читали рукописи, изучали свитки, и общались с величайшими экспертами в интересующих вас темах. " +
	"Всё это сделало вас знатоком в своей области."

func getBackgroundSpecificType() string {
	var sageTypes = map[int]string{
		1: "Академик с испорченной репутацией",
		2: "Алхимик",
		3: "Астроном",
		4: "Библиотекарь",
		5: "Исследователь",
		6: "Писец",
		7: "Подмастерье волшебника",
		8: "Профессор",
	}

	var sageSpecific string
	sageSpecific = sageTypes[general.D20.RollDice()]

	return sageSpecific
}

func getBackgroundMasteryOfTools() []string {
	return []string{""}
}

var backgroundAbility = map[string]string{
	"Исследователь": "Если вы пытаетесь изучить или вспомнить информацию, которой вы не обладаете, вы часто знаете, где и от кого её можно получить. " +
		"Обычно это библиотека, скрипторий, университет, мудрец или другое образованное существо. " +
		"Мастер может решить, что искомое знание является тайной и хранится в практически недоступном месте, или что оно вообще недоступно. " +
		"Поиски глубочайших тайн вселенной могут потребовать отдельного приключения или даже целой кампании.",
}

func getBackgroundItems() []backgrounds.Item {
	return []backgrounds.Item{
		{
			Name:     "Бутылочка чернил",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Писчее перо",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Небольшой нож",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Письмо от мёртвого коллеги с вопросом (на который вы пока не можете ответит)",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Комплект обычной одежды",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "кошелёк с 10 зм",
			ItemType: "Tools",
			Count:    1,
		},
	}
}

func getBackgroundPersonality() backgrounds.Personalization {
	characterTrait := map[int]string{
		1: "Я использую многосложные слова, создающие впечатление образованности.",
		2: "Я прочитал все книги в величайших библиотеках мира — или, по крайней мере, говорю так.",
		3: "Я привык помогать тем, кто не так умён как я, и терпеливо всем всё объясняю.",
		4: "Больше всего на свете я люблю тайны.",
		5: "Прежде чем принять решение, я выслушаю аргументы обеих спорящих сторон.",
		6: "Я ... говорю ... медленно ... когда разговариваю ... с идиотами, ... то есть ... практически ... всегда.",
		7: "В социальном взаимодействии я ужасно неуклюж.",
		8: "Я убеждён, что остальные постоянно хотят украсть мои тайны.",
	}
	randomTrait := general.D8.RollDice()

	ideals := []backgrounds.Ideal{
		{
			Worldview:   "Lawful Good",
			WorldviewRu: "Законопослушный добрый",
			Text:        "Благородный долг. Я должен защищать тех, кто ниже меня, и заботиться о них.",
		},
		{
			Worldview:   "Neutral Good",
			WorldviewRu: "Нейтральный добрый",
			Text:        "Честь. Все люди, богатые ли они, или бедные, достойны уважения.",
		},
		{
			Worldview:   "Chaotic Good",
			WorldviewRu: "Хаотический добрый",
			Text:        "Уважение. Уважение — мой долг. Кем бы ты ни был, к другим нужно относиться с уважением, невзирая на их происхождение.",
		},
		{
			Worldview:   "Lawful Neutral",
			WorldviewRu: "Законопослушный нейтральный",
			Text:        "Ответственность. Я должен уважать тех, кто выше меня, а те, кто ниже меня, должны уважать меня.",
		},
		{
			Worldview:   "Neutral",
			WorldviewRu: "Истинно нейтральный",
			Text:        "Безмятежность. Моя семья хорошо обеспечена, поэтому я живу в свое удовольствие",
		},
		{
			Worldview:   "Chaotic Neutral",
			WorldviewRu: "Хаотический нейтральный",
			Text:        "Независимость. Я должен доказать, что справлюсь и без заботы своей семьи.",
		},
		{
			Worldview:   "Lawful Evil",
			WorldviewRu: "Законопослушный злой",
			Text:        "Власть. Если соберу много власти, никто не посмеет указывать мне, что делать.",
		},
		{
			Worldview:   "Neutral Evil",
			WorldviewRu: "Нейтральный злой",
			Text:        "Хищность. Земли должны приносить моей семье больше прибыли.",
		},
		{
			Worldview:   "Chaotic Evil",
			WorldviewRu: "Хаотический злой",
			Text:        "Подлость. Мне нравится плести интриги и заговоры.",
		},
		{
			Worldview:   "Any",
			WorldviewRu: "Любой",
			Text:        "Семья. Настоящая кровь гуще.",
		},
	}
	randomIdeal, _ := random.IntRange(0, len(ideals))

	affections := map[int]string{
		1: "Я должен защищать своих учеников.",
		2: "У меня есть древний текст, содержащий ужасную тайну, и который не должен попасть в чужие руки.",
		3: "Я работаю над сохранением библиотеки, университета, скриптория или монастыря.",
		4: "Труд всей моей жизни, это серия томов, посвящённая определённой области знаний.",
		5: "Всю свою жизнь я ищу ответ на один вопрос.",
		6: "Ради знаний я продал душу. Когда-нибудь я надеюсь совершить великое деяние и вернуть её себе.",
	}
	randomAffection := general.D6.RollDice()

	weakness := map[int]string{
		1: "Меня легко отвлечь, пообещав информацию.",
		2: "Увидев демона, большинство закричит и убежит. Я же остановлюсь и буду изучать его анатомию.",
		3: "Для раскрытия древних тайн можно пожертвовать современной цивилизацией.",
		4: "Я избегаю очевидных решений, пользуясь замысловатыми.",
		5: "Я говорю, не обдумывая слова, чем часто оскорбляю других.",
		6: "Ради спасения своей или чьей-либо еще жизни я выболтаю любую тайну.",
	}
	randomWeakness := general.D6.RollDice()

	return backgrounds.Personalization{
		PersonalizationDescription: "Мудрецами становятся после продолжительных занятий, и личность таких существ отражает жизнь в постоянном обучении. " +
			"Будучи в постоянном поиске мудрости, они высоко ценят знания — как для себя, так и в стремлении к идеалам.",
		Advice:         "",
		CharacterTrait: characterTrait[randomTrait],
		Ideal:          ideals[randomIdeal].Text,
		Affection:      affections[randomAffection],
		Weakness:       weakness[randomWeakness],
		Worldview:      ideals[randomIdeal].WorldviewRu,
	}
}

func GetBackground() *backgrounds.Background {
	return &backgrounds.Background{
		BackgroundName:         "Sage",
		BackgroundNameRu:       "Мудрец",
		BackgroundSpecificType: getBackgroundSpecificType(),
		BackgroundAbility:      backgroundAbility,
		Description:            description,
		BackgroundItems:        getBackgroundItems(),
		BackgroundSkills:       []string{"Arcana", "History"},
		MasteryOfTools:         getBackgroundMasteryOfTools(),
		Personalization:        getBackgroundPersonality(),
		Gold:                   10,
		Langs:                  []string{"Один на ваш выбор"},
	}
}
