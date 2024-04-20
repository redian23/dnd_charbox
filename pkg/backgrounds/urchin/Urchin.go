package urchin

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/backgrounds"
	"pregen/pkg/general"
)

var description = "Вы выросли на улице в бедности и одиночестве, лишённые родителей. " +
	"Никто не присматривал и не заботился о вас, и вам пришлось научиться делать это самому. " +
	"Вам приходилось постоянно бороться за еду и следить за другими неприкаянными душами, способными обокрасть вас. " +
	"Вы спали на чердаках и в переулках, мокли под дождём и боролись с болезнями, не получая медицинской помощи или приюта. " +
	"Вы выжили, невзирая на все трудности, и сделали это благодаря своей сноровке, силе или скорости."

func getBackgroundMasteryOfTools() []string {
	return []string{"Набор для грима", "Воровские инструменты"}
}

func getBackgroundItems() []backgrounds.Item {
	return []backgrounds.Item{
		{
			Name:     "Маленький нож",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Карта города (в котором вы выросли)",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Ручная мышь",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Безделушка в память о родителях",
			ItemType: "Tools",
			Count:    1,
		},
		{
			Name:     "Комплект обычной одежды",
			ItemType: "Tools",
			Count:    1,
		},
	}
}

var urchinAbility = map[string]string{
	"Городские тайны": "Вы знаете тайные лазы и проходы городских улиц, позволяющие пройти там, где другие не увидят пути. " +
		"Вне боя вы (и ведомые вами союзники) можете перемещаться по городу вдвое быстрее обычного.",
}

func geturchinPersonality() backgrounds.Personalization {
	characterTrait := map[int]string{
		1: "В моих карманах полно побрякушек и объедков.",
		2: "Я задаю очень много вопросов.",
		3: "Я часто забиваюсь в узкие закутки, где никто не сможет добраться до меня.",
		4: "Я всегда сплю, прижавшись спиной к стене или дереву, сжимая узелок со всеми своими пожитками в руках.",
		5: "Я не воспитан, и ем как свинья.",
		6: "Я убеждён, что все, кто проявляют доброту ко мне, на самом деле скрывают злые намерения.",
		7: "Я не люблю мыться.",
		8: "Я прямо говорю о том, на что прочие предпочитают лишь намекнуть, или промолчать.",
	}
	randomTrait := general.D8.RollDice()

	ideals := []backgrounds.Ideal{
		{
			Worldview:   "Lawful Good",
			WorldviewRu: "Законопослушный добрый",
			Text:        "Рвение. Я сделаю эти трущобы процветающим местом.",
		},
		{
			Worldview:   "Neutral Good",
			WorldviewRu: "Нейтральный добрый",
			Text:        "Уважение. Все люди, богатые ли они, или бедные, достойны уважения.",
		},
		{
			Worldview:   "Chaotic Good",
			WorldviewRu: "Хаотический добрый",
			Text:        "Справедливость. Я граблю богатых и отдаю все награбленное бедным.",
		},
		{
			Worldview:   "Lawful Neutral",
			WorldviewRu: "Законопослушный нейтральный",
			Text:        "Общность. Вы должны заботиться друг о друге, ведь никто другой этого не сделает.",
		},
		{
			Worldview:   "Neutral",
			WorldviewRu: "Истинно нейтральный",
			Text:        "Люди. Я помогаю тем, кто помогает мне. Это позволяет нам выжить.",
		},
		{
			Worldview:   "Chaotic Neutral",
			WorldviewRu: "Хаотический нейтральный",
			Text:        "Перемены. Убогие возвысятся, а великие падут. Перемены в природе вещей.",
		},
		{
			Worldview:   "Lawful Evil",
			WorldviewRu: "Законопослушный злой",
			Text:        "Возмездие. Нужно показать богачам, чего стоит жизнь и смерть в трущобах.",
		},
		{
			Worldview:   "Neutral Evil",
			WorldviewRu: "Нейтральный злой",
			Text:        "Безразличие. Мне плевать на моих соседей, главное чтобы мне было что поесть",
		},
		{
			Worldview:   "Chaotic Evil",
			WorldviewRu: "Хаотический злой",
			Text:        "Ненависть. Этот мир несправедливо со мной обошелся.",
		},
		{
			Worldview:   "Any",
			WorldviewRu: "Любой",
			Text:        "Устремление. Я готов доказать, что достоин лучшей жизни.",
		},
	}
	randomIdeal, _ := random.IntRange(0, len(ideals))

	affections := map[int]string{
		1: "Мой город или деревня это мой дом, и я буду защищать его.",
		2: "Я спонсирую приют, чтобы оградить других от того, что пришлось пережить мне.",
		3: "Я выжил лишь благодаря другому беспризорнику, что передал мне знания, как вести себя на улицах.",
		4: "Я в неоплатном долгу перед человеком, что сжалился и помог мне.",
		5: "Я избавился от нищеты, ограбив важного человека, и меня разыскивают.",
		6: "Никто не должен пережить те трудности, через которые пришлось пройти мне.",
	}
	randomAffection := general.D6.RollDice()

	weakness := map[int]string{
		1: "Если я в меньшинстве, то я убегу из битвы.",
		2: "Золото в любом виде выглядит для меня как большая куча денег, и я сделаю всё, чтобы его у меня стало больше.",
		3: "Я никогда не доверяю полностью кому бы то ни было, кроме себя.",
		4: "Я предпочту убить кого-нибудь во сне, нежели в честном поединке.",
		5: "Это не воровство, если я взял то, в чём нуждаюсь больше, чем кто-либо другой.",
		6: "Те, кто не могут позаботиться о себе, получат то, что заслужили.",
	}
	randomWeakness := general.D6.RollDice()

	return backgrounds.Personalization{
		PersonalizationDescription: "Жизнь в нищете оставляет свой отпечаток на беспризорниках. " +
			"В них, как правило, сильна привязанность к людям, с которыми они делили тяготы уличной жизни, " +
			"или они исполнены желанием добиться лучшей доли, и, возможно, расквитаться с богачами, от которых они натерпелись.",
		Advice:         "",
		CharacterTrait: characterTrait[randomTrait],
		Ideal:          ideals[randomIdeal],
		Affection:      affections[randomAffection],
		Weakness:       weakness[randomWeakness],
		Worldview:      ideals[randomIdeal].WorldviewRu,
	}
}

func GetUrchinBackground() *backgrounds.Background {
	return &backgrounds.Background{
		BackgroundName:         "Urchin",
		BackgroundNameRu:       "Беспризорник",
		BackgroundSpecificType: "Стандарт",
		BackgroundAbility:      urchinAbility,
		Description:            description,
		BackgroundItems:        getBackgroundItems(),
		BackgroundSkills:       []string{"Sleight Of Hand", "Stealth"},
		MasteryOfTools:         getBackgroundMasteryOfTools(),
		Personalization:        geturchinPersonality(),
		Gold:                   10,
	}
}
