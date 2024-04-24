package noble

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/backgrounds"
	"pregen/pkg/general"
)

var description = "Ваша семья, древняя и прославленная, носит титул графа или графини, " +
	"обладая обширными владениями, собирая налоги и имея значительное политическое влияние. " +
	"Вы, возможно, наследник главы семейства, чувствуя на себе ответственность за его сохранение и процветание, " +
	"хотя ваше желание исследовать мир может вызывать разногласия среди родственников. " +
	"Герб, кольцо с печатью, и специальные цвета служат символами вашего рода, " +
	"а возможно, у вас есть даже животное-тотем, которое вы считаете своим духовным наставником."

func getBackgroundMasteryOfTools() []string {
	return []string{"Один игровой набор"}
}

var backgroundAbility = map[string]string{
	"Привилегированность": "Благодаря знатному происхождению, другие хорошо к вам относятся. " +
		"Вас принимают в высшем обществе, и считается, что у вас есть право посещать любые места. " +
		"Обыватели изо всех сил стараются сделать вам приятно и избежать вашего гнева, " +
		"а другие высокородные считают вас своей ровней. Если нужно, вы можете получить аудиенцию местного дворянина.",
}

func getBackgroundItems() []backgrounds.Item {
	return []backgrounds.Item{
		{
			Name:     "Комплект отличной одежды",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Кольцо-печатка",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Свиток с генеалогическим древом",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "кошелёк с 25 зм",
			ItemType: "Tools",
			Count:    1,
		},
	}
}

func getBackgroundPersonality() backgrounds.Personalization {
	characterTrait := map[int]string{
		1: "Я применяю так много лести, что все, с кем я разговариваю, чувствуют себя самыми чудесными и важными персонами в мире.",
		2: "Обыватели любят меня за доброту и великодушие.",
		3: "Один лишь взгляд на мои регалии заставляет перестать сомневаться в том, что я выше немытого отребья.",
		4: "Я много трачу на то, чтобы выглядеть потрясающе, и по последнему слову моды.",
		5: "Мне не нравится марать руки, и я не хочу умереть в каком-нибудь неподобающем месте.",
		6: "Несмотря на благородное рождение, я не ставлю себя выше народа. У всех нас течёт одинаковая кровь.",
		7: "Потеряв мою милость, обратно её не вернёшь.",
		8: "Оскорбишь меня, и я сотру тебя в порошок, уничтожу твоё имя, и сожгу твои поля.",
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
		1: "Я пойду на любой риск, лишь бы заслужить признание семьи.",
		2: "Союз моего дома с другой благородной семьёй нужно поддерживать любой ценой.",
		3: "Нет никого важнее других членов моей семьи.",
		4: "Я влюблён в наследницу семейства, презираемого моей роднёй.",
		5: "Моя преданность правителю непоколебима.",
		6: "Обыватели должны считать меня своим героем.",
	}
	randomAffection := general.D6.RollDice()

	weakness := map[int]string{
		1: "Я втайне считаю всех ниже себя.",
		2: "Я скрываю позорную тайну, которая может уничтожить мою семью.",
		3: "Я слишком часто слышал завуалированные оскорбления и угрозы, и потому быстро впадаю в гнев.",
		4: "У меня неустанная страсть к плотским удовольствиям.",
		5: "Весь мир вращается вокруг меня.",
		6: "Я часто навлекаю позор на свою семью словами и действиями.",
	}
	randomWeakness := general.D6.RollDice()

	return backgrounds.Personalization{
		PersonalizationDescription: "Благородные рождаются и растут в особом окружении, и их личность несёт отпечаток этого воспитания. " +
			"Благородный титул идёт вкупе с множеством обязательств — ответственностью перед семьёй, другими благородными (включая короля)," +
			" народом, доверенным вашей семье, и даже перед самим титулом. Все эти обязательства другие могут использовать против вас.",
		Advice:         "",
		CharacterTrait: characterTrait[randomTrait],
		Ideal:          ideals[randomIdeal],
		Affection:      affections[randomAffection],
		Weakness:       weakness[randomWeakness],
		Worldview:      ideals[randomIdeal].WorldviewRu,
	}
}

func GetBackground() *backgrounds.Background {
	return &backgrounds.Background{
		BackgroundName:         "Noble",
		BackgroundNameRu:       "Благородный",
		BackgroundSpecificType: "Стандартный",
		BackgroundAbility:      backgroundAbility,
		Description:            description,
		BackgroundItems:        getBackgroundItems(),
		BackgroundSkills:       []string{"Acrobatics", "History"},
		MasteryOfTools:         getBackgroundMasteryOfTools(),
		Personalization:        getBackgroundPersonality(),
		Gold:                   15,
		Langs:                  []string{"Один на ваш выбор"},
	}
}
