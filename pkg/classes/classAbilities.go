package classes

import (
	"github.com/mazen160/go-random"
)

var ClassSpecialSpells []string

type ClassAbility struct {
	Level       int    `json:"level"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func getClassAbilities() []ClassAbility {
	switch ClassNameGlobalRu {
	case "Воин":
		return setFighterClassAbilities()
	case "Варвар":
		return setBarbarianClassAbilities()
	case "Паладин":
		return setPaladinClassAbilities()
	case "Монах":
		return setMonkClassAbilities()
	case "Плут":
		return setRogueClassAbilities()
	case "Следопыт":
		return setRangerClassAbilities()
	case "Волшебник":
		return setWizardClassAbilities()
	case "Друид":
		return setDruidClassAbilities()
	case "Жрец":
		return setClericClassAbilities()
	case "Чародей":
		return setSorcererClassAbilities()
	case "Бард":
		return setBardClassAbilities()
	case "Изобретатель":
		return setArtificerClassAbilities()
	case "Колдун":
		return setWarlockClassAbilities()
	}
	return nil
}

func getLevelAbilities(classAbilityList []ClassAbility) []ClassAbility {
	var classAbilityListAnswer []ClassAbility
	for _, ability := range classAbilityList {
		if ability.Level <= CharacterLevelGlobal {
			classAbilityListAnswer = append(classAbilityListAnswer, ability)
		}
	}
	return classAbilityListAnswer
}

func setFighterClassAbilities() []ClassAbility {

	type battleStyleVariant struct {
		Name        string
		Description string
	}

	battelStyleVariants := []battleStyleVariant{
		{Name: "Дуэлянт", Description: "Пока вы держите рукопашное оружие в одной руке и не используете другого оружия, вы получаете бонус +2 к броскам урона этим оружием."},
		{Name: "Защита", Description: "Если существо, которое вы видите, атакует не вас, а другое существо, находящееся в пределах 5 футов от вас, вы можете реакцией создать помеху его броску атаки. Для этого вы должны использовать щит."},
		{Name: "Оборона", Description: "Пока вы носите доспехи, вы получаете бонус +1 к КД."},
		{Name: "Сражение большим оружием", Description: "Если у вас выпало «1» или «2» на кости урона оружия при атаке, которую вы совершали рукопашным оружием, удерживая его двумя руками, то вы можете перебросить эту кость, и должны использовать новый результат, даже если снова выпало «1» или «2». Чтобы воспользоваться этим преимуществом, ваше оружие должно иметь свойство «двуручное» или «универсальное»."},
		{Name: "Сражение двумя оружиями", Description: "Если вы сражаетесь двумя оружиями, вы можете добавить модификатор характеристики к урону от второй атаки."},
		{Name: "Стрельба", Description: "Вы получаете бонус +2 к броску атаки, когда атакуете дальнобойным оружием."},
	}

	randNum, _ := random.IntRange(0, len(battelStyleVariants))

	classAbilityList := []ClassAbility{
		{
			Level:       1,
			Name:        "Боевой стиль",
			Description: battelStyleVariants[randNum].Name + " - " + battelStyleVariants[randNum].Description,
		},
		{
			Level: 1,
			Name:  "Второе дыхание",
			Description: "Вы обладаете ограниченным источником выносливости, которым можете воспользоваться, чтобы уберечь себя. В свой ход вы можете бонусным действием восстановить хиты в размере 1к10 + ваш уровень воина.\n\n" +
				"Использовав это умение, вы должны завершить короткий либо продолжительный отдых, чтобы получить возможность использовать его снова.",
		},
		{
			Level: 2,
			Name:  "Всплеск действий",
			Description: "Вы получаете возможность на мгновение преодолеть обычные возможности. В свой ход вы можете совершить одно дополнительное действие помимо обычного и бонусного действий. " +
				"Использовав это умение, вы должны завершить короткий или продолжительный отдых, чтобы получить возможность использовать его снова. Начиная с 17-го уровня, вы можете использовать это умение дважды, прежде чем вам понадобится отдых, но в течение одного хода его всё равно можно использовать лишь один раз.",
		},
		{
			Level:       3,
			Name:        "Воинский Архитип",
			Description: "Выбранный вами архетип предоставляет вам умения на 3-м, 7-м, 10-м, 15-м и 18-м уровнях.\n\n",
		},
	}

	classAbilityListAnswer := getLevelAbilities(classAbilityList)

	return classAbilityListAnswer
}

func setBarbarianClassAbilities() []ClassAbility {

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Защита без доспехов",
		Description: "Если вы не носите доспехов, ваш Класс Доспеха равен 10 + модификатор Ловкости + модификатор Телосложения." +
			"Вы можете использовать щит, не теряя этого преимущества.",
	}

	abil2 := ClassAbility{
		Level: 1,
		Name:  "Ярость",
		Description: "В бою вы сражаетесь с первобытной свирепостью. В свой ход вы можете бонусным действием войти в состояние ярости.\n\n" +
			"В состоянии ярости вы получаете следующие преимущества, если не носите тяжёлую броню:\n\n" +
			"Вы совершаете с преимуществом проверки и спасброски Силы.\n" +
			"Если вы совершаете рукопашную атаку оружием, используя Силу, вы получаете бонус к броску урона, соответствующий вашему уровню варвара, как показано в колонке «урон ярости» таблицы «Варвар».\n" +
			"Вы получаете сопротивление дробящему, колющему и рубящему урону.\nЕсли вы способны накладывать заклинания, то вы не можете накладывать или концентрироваться на заклинаниях, пока находитесь в состоянии ярости.\n\n" +
			"Ваша ярость длится 1 минуту. Она прекращается раньше, если вы потеряли сознание или если вы закончили свой ход, не получив урон или не атаковав враждебное по отношению к вам существо с момента окончания вашего прошлого хода. Также вы можете прекратить свою ярость бонусным действием.\n\n" +
			"Если вы впадали в состояние ярости максимальное для вашего уровня количество раз (смотрите колонку «ярость»), то вы должны совершить продолжительный отдых, прежде чем сможете использовать ярость ещё раз.",
	}

	classAbilityList := []ClassAbility{abil1, abil2}

	return classAbilityList
}

func setPaladinClassAbilities() []ClassAbility {

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Божественное чувство",
		Description: "Присутствие сильного зла воспринимается вашими чувствами как неприятный запах, а могущественное добро звучит как небесная музыка в ваших ушах. Вы можете действием открыть своё сознание для обнаружения таких сил. " +
			"Вы до конца своего следующего хода знаете местоположение всех Исчадий, Небожителей и Нежити в пределах 60 футов, не имеющих полного укрытия. " +
			"Вы знаете вид (Исчадие, Небожитель, Нежить) любого существа, чьё присутствие вы чувствуете, но не можете определить, кто это конкретно (например, вампир граф Страд фон Зарович). " +
			"В этом же радиусе вы также обнаруживаете присутствие мест и предметов, которые были освящены или осквернены, например, заклинанием святилище [hallow]." +
			"Вы можете использовать это умение количество раз, равное 1 + модификатор Харизмы. Когда вы заканчиваете продолжительный отдых, вы восстанавливаете все потраченные использования.",
	}

	abil2 := ClassAbility{
		Level: 1,
		Name:  "Наложение рук",
		Description: "Ваше благословенное касание может лечить раны. У вас есть запас целительной силы, который восстанавливается после продолжительного отдыха. " +
			"При помощи этого запаса вы можете восстанавливать количество хитов, равное уровню паладина, умноженному на 5.\n\n" +
			"Вы можете действием коснуться существа и, зачерпнув силу из запаса, восстановить количество хитов этого существа на любое число, вплоть до максимума, оставшегося в вашем запасе.\n\n" +
			"В качестве альтернативы, вы можете потратить 5 хитов из вашего запаса хитов для излечения цели от одной болезни или одного действующего на неё яда. " +
			"Вы можете устранить несколько эффектов болезни и ядов одним использованием «Наложения рук», тратя хиты отдельно для каждого эффекта.\n\n" +
			"Это умение не оказывает никакого эффекта на Нежить и Конструктов.",
	}

	classAbilityList := []ClassAbility{abil1, abil2}

	return classAbilityList
}

func setMonkClassAbilities() []ClassAbility {

	abil1 := ClassAbility{
		Level:       1,
		Name:        "Защита без доспехов",
		Description: "Если вы не носите ни доспех, ни щит, ваш Класс Доспеха равен 10 + модификатор Ловкости + модификатор Мудрости.",
	}

	abil2 := ClassAbility{
		Level: 1,
		Name:  "Боевые искусства",
		Description: "Ваше знание боевых искусств позволяет вам эффективно использовать в бою безоружные удары и монашеское оружие — короткие мечи, а также любое простое рукопашное оружие, не имеющее свойств «двуручное» и «тяжёлое»." +
			"Если вы безоружны или используете только монашеское оружие, и не носите ни доспехов, ни щита, вы получаете следующие преимущества:" +
			"Вы можете использовать Ловкость вместо Силы для бросков атак и урона ваших безоружных ударов и атак монашеским оружием." +
			"Вы можете использовать к4 вместо обычной кости урона ваших безоружных ударов или атак монашеским оружием. Эта кость увеличивается с вашим уровнем, как показано в колонке «боевые искусства»." +
			"Если в свой ход вы используете действие Атака для безоружного удара или атаки монашеским оружием, вы можете бонусным действием совершить ещё один безоружный удар. " +
			"Например, если вы совершили действие Атака и атаковали боевым посохом, вы можете совершить бонусным действием безоружный удар, при условии, что в этом ходу вы еще не совершали бонусное действие." +
			"Некоторые монастыри используют особые виды монашеского оружия. Например, вы можете использовать дубинку в виде двух деревянных брусков, соединённых короткой цепью (такое оружие называется нунчаками), или серп с более коротким и прямым лезвием (называется камой). " +
			"Как бы ни называлось ваше монашеское оружие, вы используете характеристики, соответствующие этому оружию.",
	}

	classAbilityList := []ClassAbility{abil1, abil2}

	return classAbilityList
}

func setRogueClassAbilities() []ClassAbility {

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Компетентность",
		Description: "Выберите два ваших владения в навыках или одно владение навыком и владение воровскими инструментами." +
			"Ваш бонус мастерства удваивается для всех проверок характеристик, которые вы совершаете, используя любое из выбранных владений.\n\n" +
			"На 6-м уровне вы можете выбрать ещё два владения (навыки или воровские инструменты), чтобы получить эту же выгоду.",
	}

	abil2 := ClassAbility{
		Level: 1,
		Name:  "Скрытая атака",
		Description: "Вы знаете, как точно наносить удар и использовать отвлечение врага. Один раз в ход вы можете причинить дополнительный урон 1к6 одному из существ, по которому вы попали атакой, совершённой с преимуществом к броску атаки. " +
			"Атака должна использовать дальнобойное оружие или оружие со свойством «фехтовальное». " +
			"Вам не нужно иметь преимущество при броске атаки, если другой враг цели находится в пределах 5 футов от неё. Этот враг не должен быть недееспособным, и у вас не должно быть помехи для броска атаки.\n\n" +
			"Дополнительный урон увеличивается, когда вы получаете уровни в этом классе, как показано в колонке «скрытая атака».",
	}

	abil3 := ClassAbility{
		Level: 1,
		Name:  "Воровской жаргон",
		Description: "Во время плутовского обучения вы выучили воровской жаргон, тайную смесь диалекта, жаргона и шифра, который позволяет скрывать сообщения в, казалось бы, обычном разговоре. " +
			"Только другое существо, знающее воровской жаргон, понимает такие сообщения. " +
			"Это занимает в четыре раза больше времени, нежели передача тех же слов прямым текстом.\n\n" +
			"Кроме того, вы понимаете набор секретных знаков и символов, используемый для передачи коротких и простых сообщений. " +
			"Например, является ли область опасной или территорией гильдии воров, находится ли поблизости добыча, простодушны ли люди в округе, и предоставляют ли здесь безопасное убежище для воров в бегах.",
	}

	classAbilityList := []ClassAbility{abil1, abil2, abil3}

	return classAbilityList
}

func setRangerClassAbilities() []ClassAbility {

	chosenEnemyList := []string{"Аберрации", "Великаны", "Драконы", "Звери", "Исчадия", "Конструкты",
		"Монстры", "Небожители", "Нежить", "Растения", "Слизи", "Феи", "Элементали",
	}
	randNum, _ := random.IntRange(0, len(chosenEnemyList))

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Избранный враг",
		Description: "У вас есть значительный опыт изучения, отслеживания, охоты и даже общения с определённым видом врагов.\n\n" +
			"Ваш избранный враг:" + chosenEnemyList[randNum] + ". Вы совершаете с преимуществом проверки Мудрости (Выживание) для выслеживания избранных врагов, а также проверки Интеллекта для вспоминания информации о них.\n\n" +
			"Когда вы получаете это умение, вы также обучаетесь одному из языков, на котором говорит ваш избранный враг, если он вообще умеет говорить.\n\n" +
			"Вы дополнительно выбираете по одному избранному врагу и языку, связанному с ним, на 6-м и 14-м уровнях. Получая уровни, ваш выбор должен отражать чудовищ, с которыми вы уже сталкивались во время приключений.",
	}

	groundList := []string{"Арктика", "Болота", "Горы", "Леса", "Луга", "Побережье", "Подземье", "Пустыня"}
	randNum2, _ := random.IntRange(0, len(groundList))

	abil2 := ClassAbility{
		Level: 1,
		Name:  "Исследователь природы",
		Description: "Вы очень хорошо знакомы с одним видом природной среды и имеете большой опыт путешествий и выживания в регионах с таким климатом. " +
			"Вид известной местности: " + groundList[randNum2] + ". " +
			"Когда вы совершаете проверку Интеллекта или Мудрости, связанную с известной вам местностью, ваш бонус мастерства удваивается, если вы используете навык, которым владеете.\n\n" +
			"Путешествуя по избранной вами местности в течение часа или более, вы получаете следующие преимущества:\n\n" +
			"Труднопроходимая местность не замедляет путешествие группы.\nВаша группа не может заблудиться, кроме как по вине магии.\n" +
			"Даже когда вы занимаетесь другой деятельностью во время путешествия (например, ищете пищу, ориентируетесь или занимаетесь выслеживанием), вы остаётесь готовы к опасности.\n" +
			"Если вы путешествуете в одиночку, вы можете перемещаться скрытно в нормальном темпе.\n" +
			"Когда вы ищете пищу, то находите в два раза больше, чем обычно.\nКогда вы выслеживаете других существ, вы также узнаёте их точное количество, их размеры, и как давно они прошли через область. \n" +
			"Вы можете выбрать дополнительную известную местность на 6-м и 10-м уровнях.",
	}

	abil3 := ClassAbility{
		Level: 1,
		Name:  "Предпочтительный противник (TCE)",
		Description: "Когда вы попадаете атакой по существу, вы можете призвать силы природы, чтобы отметить существо и сделать его своим избранным врагом на 1 минуту или до тех пор, пока не потеряете концентрацию (как если бы вы концентрировались на заклинании).\n\n" +
			"Первый раз в каждый свой ход, когда вы попадаете атакой по избранному врагу и наносите ему урон, в том числе и когда вы отмечаете его, вы можете дополнительно нанести 1к4 урона того же типа.\n\n" +
			"Вы можете использовать это умение для отметки избранного врага количество раз, равное вашему бонусу мастерства. Вы восстанавливаете все потраченные использования после окончания продолжительного отдыха.\n\n" +
			"Дополнительный урон этого умения увеличивается, когда вы достигаете определённых уровней в этом классе: 1к6 на 6-м уровне и 1к8 на 14-м.",
	}

	abil4 := ClassAbility{
		Level: 1,
		Name:  "Искусный исследователь",
		Description: "Вы — непревзойдённый исследователь и отлично выживаете в природе, будучи способным как самостоятельно странствовать по дикой местности, так и помогать в этом другим. " +
			"Вы получаете описанное ниже преимущество «Хитрец», а также дополнительные преимущества, описанные ниже, на 6-м и 10-м уровнях следопыта соответственно.\n\n" +
			"Хитрец (1-й уровень)\n" +
			"Выберите один из навыков, которым вы владеете. Ваш бонус мастерства удваивается для всех проверок характеристик, которые вы совершаете, используя выбранный навык.\n\n" +
			"Вы также можете говорить, читать и писать на двух дополнительных языках по вашему выбору.\n\n" +
			"Бродяга (6-й уровень)\n" +
			"Ваша скорость ходьбы увеличивается на 5 футов, и вы получаете скорость лазания и плавания, равную вашей скорости ходьбы.\n\n" +
			"Неутомимый (10-й уровень)\n" +
			"Действием вы можете дать себе количество временных хитов, равное 1к8 + модификатор Мудрости (минимум 1 хит). " +
			"Вы можете использовать это умение количество раз, равное вашему бонусу мастерства. Вы восстанавливаете все потраченные использования после окончания продолжительного отдыха.\n\n" +
			"Помимо этого, в конце короткого отдыха ваша степень истощения уменьшается на 1.",
	}

	randNum3, _ := random.IntRange(0, 1)
	if randNum3 == 0 {
		classAbilityList := []ClassAbility{abil1, abil2}
		return classAbilityList
	} else {
		classAbilityList := []ClassAbility{abil3, abil4}
		return classAbilityList
	}
}

func setWizardClassAbilities() []ClassAbility {

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Магическое восстановление",
		Description: "Вы знаете как восстанавливать часть магической энергии, изучая книгу заклинаний. Один раз в день, когда вы заканчиваете короткий отдых, вы можете восстановить часть использованных ячеек заклинаний. " +
			"Ячейки заклинаний могут иметь суммарный уровень, который не превышает половину уровня вашего волшебника (округляя в большую сторону), и ни одна из ячеек не может быть шестого уровня или выше." +
			"Например, если вы волшебник 4-го уровня, вы можете восстановить ячейки заклинаний с суммой уровней не больше двух. Вы можете восстановить одну ячейку заклинаний 2-го уровня, или две ячейки заклинаний 1-го уровня.",
	}

	classAbilityList := []ClassAbility{abil1}

	return classAbilityList
}

func setDruidClassAbilities() []ClassAbility {

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Друидический язык",
		Description: "Вы знаете Друидический язык — тайный язык друидов. Вы можете на нём говорить и оставлять тайные послания. Вы и все, кто знают этот язык, автоматически замечаете эти послания. " +
			"Другие замечают присутствие послания при успешной проверке Мудрости (Восприятие) Сл 15, но без помощи магии не могут расшифровать его.",
	}

	classAbilityList := []ClassAbility{abil1}

	return classAbilityList
}

func setClericClassAbilities() []ClassAbility {

	domainList := []string{"Домен бури", "Домен войны", "Домен жизни", "Домен знаний", "Домен обмана", "Домен природы",
		"Домен света", "Домен смерти", "Домен магии", "Домен кузни", "Домен упокоения", "Домен мира", "Домен порядка", "Домен сумерек",
	}
	randNum, _ := random.IntRange(0, len(domainList))

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Божественный домен",
		Description: domainList[randNum] + ", связанный с вашим божеством. " +
			"Все домены детально рассмотрены в конце описания класса, и каждый содержит примеры божеств, связанных с ним. " +
			"Ваш выбор даёт вам заклинания домена и другие умения. " +
			"Он также даёт вам дополнительные способы использования «Божественного канала», когда вы получаете это умение на 2-м уровне, и дополнительные умения на 6-м, 8-м и 17-м уровнях.\n\n" +
			"ЗАКЛИНАНИЯ ДОМЕНА\n" +
			"У каждого домена есть список заклинаний, которые вы получаете на новых уровнях жреца. " +
			"Как только вы получаете заклинание домена, оно всегда считается подготовленным и не учитывается при подсчёте количества заклинаний, которые можно подготовить.\n\n" +
			"Если вы получаете доступ к заклинанию, отсутствующему в списке заклинаний жреца, оно всё равно будет считаться для вас заклинанием жреца.",
	}

	classAbilityList := []ClassAbility{abil1}

	return classAbilityList
}

func setSorcererClassAbilities() []ClassAbility {
	var specAbilities []ClassAbility
	var originsDescription string

	origins := []string{
		"Наследие драконьей крови",
		"Дикая магия",
	}
	randNum, _ := random.IntRange(0, len(origins))
	specAbilities, originsDescription = SetSorcererOrigin(origins[randNum])
	abil1 := ClassAbility{
		Level: 1,
		Name:  "Происхождение чародея",
		Description: "Ваш источник - " + origins[randNum] + ", из которого ваш персонаж черпает свою силу. " +
			originsDescription +
			".<br>Ваш выбор предоставляет вам умения на 1-м, 6-м, 14-м и 18-м уровнях.",
	}

	classAbilityList := []ClassAbility{abil1}
	classAbilityList = append(classAbilityList, specAbilities...)
	return classAbilityList
}

func setBardClassAbilities() []ClassAbility {
	collegiumName, collegiumDescription, collegiumAbilities := choiceBardCollegium()

	classAbilityList := []ClassAbility{
		{
			Level: 1,
			Name:  "Вдохновение барда",
			Description: "Своими словами или музыкой вы можете вдохновлять других. " +
				"Для этого вы должны бонусным действием выбрать одно существо, отличное от вас, в пределах 60 футов, которое может вас слышать. " +
				"Это существо получает кость бардовского вдохновения — к6.<br>" +
				"В течение следующих 10 минут это существо может один раз бросить эту кость и добавить результат к проверке характеристики, броску атаки или спасброску, который оно совершает. " +
				"Существо может принять решение о броске кости вдохновения уже после броска к20, но должно сделать это прежде, чем Мастер объявит результат броска. " +
				"Как только кость бардовского вдохновения брошена, она исчезает. Существо может иметь только одну такую кость одновременно.<br>" +
				"Вы можете использовать это умение количество раз, равное модификатору вашей Харизмы, но как минимум один раз. " +
				"Потраченные использования этого умения восстанавливаются после продолжительного отдыха.<br>" +
				"Ваша кость бардовского вдохновения изменяется с ростом вашего уровня в этом классе. Она становится к8 на 5-м уровне, к10 на 10-м уровне и к12 на 15-м уровне.",
		},
		{
			Level:       2,
			Name:        "Мастер на все руки",
			Description: "Вы можете добавлять половину бонуса мастерства, округлённую в меньшую сторону, ко всем проверкам характеристики, куда этот бонус еще не включён.<br>",
		},
		{
			Level: 2,
			Name:  "Песнь отдыха",
			Description: "Вы с помощью успокаивающей музыки или речей можете помочь своим раненым союзникам восстановить их силы во время короткого отдыха. " +
				"Если вы или любые союзные существа, способные слышать ваше исполнение, восстанавливаете хиты в конце короткого отдыха, используя Кости Хитов, " +
				"каждое потратившее Кость Хитов существо восстанавливает дополнительно 1к6 хитов.<br>" +
				"Количество дополнительно восстанавливаемых хитов растёт с вашим уровнем в этом классе: 1к8 на 9-м уровне, 1к10 на 13 уровне и 1к12 на 17 уровне.<br>",
		},
		{
			Level: 3,
			Name:  "Коллегия бардов",
			Description: "Вы углубляетесь в традиции выбранной вами коллегии бардов. Ваша коллегия - " + collegiumName + "." +
				collegiumDescription + ".<br>",
		},
		{
			Level:       3,
			Name:        "Компетентность",
			Description: "Выберите 2 навыка из тех, которыми вы владеете. Ваш бонус мастерства для этих навыков удваивается.<br>",
		},
		{
			Level:       5,
			Name:        "Источник вдохновения",
			Description: "Вы восстанавливаете истраченные вдохновения барда и после короткого и после продолжительного отдыха.<br>",
		},
		{
			Level: 6,
			Name:  "Контрочарование",
			Description: "Вы получаете возможность использовать звуки или слова силы для разрушения воздействующих на разум эффектов. " +
				"Вы можете действием начать исполнение, которое продлится до конца вашего следующего хода. " +
				"В течение этого времени вы и все дружественные существа в пределах 30 футов от вас совершают спасброски от запугивания и очарования с преимуществом. " +
				"Чтобы получить это преимущество, существа должны слышать вас. " +
				"Исполнение заканчивается преждевременно, если вы оказываетесь недееспособны, теряете способность говорить, или прекращаете исполнение добровольно (на это не требуется действие).<br>",
		},
	}

	classAbilityListAnswer := getLevelAbilities(classAbilityList)

	classAbilityListAnswer = append(classAbilityListAnswer, collegiumAbilities...)

	return classAbilityListAnswer
}

func setArtificerClassAbilities() []ClassAbility {

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Магический мастеровой",
		Description: "Вы научились вкладывать искру магии в обычные предметы. " +
			"Чтобы использовать это умение, вы должны держать в руках воровские инструменты или инструменты ремесленника. " +
			"Затем действием вы касаетесь Крошечного немагического объекта и наделяете его одним из следующих магических свойств на ваш выбор:\n\n" +
			"Зачарованный объект излучает яркий свет в радиусе 5 футов и тусклый свет в радиусе еще 5 футов.\n" +
			"Объект проигрывает записанное сообщение, которое можно услышать в пределах 10 футов каждый раз, когда до него дотрагивается существо. " +
			"Вы произносите это сообщение, когда наделяете объект данным свойством, а сама запись не может быть длиннее 6 секунд.\n" +
			"Объект непрерывно испускает запах или издаёт звук на ваш выбор (ветер, волны, стрекотание и прочее). " +
			"Выбранное явление можно ощутить на расстоянии 10 футов.\nСтатичный визуальный эффект появляется на одной из поверхностей объекта. " +
			"Этот эффект может быть изображением, текстом до 25 слов, линиями и формами или совмещением этих элементов по вашему выбору.\n" +
			"Выбранное свойство навсегда остается присущим объекту. Действием вы можете коснуться объекта и лишить его этого свойства.\n\n" +
			"Таким образом можно наделить магическими свойствами несколько предметов, касаясь одного предмета каждый раз, когда вы используете это умение, но не больше, чем одно свойство на один предмет. " +
			"Максимальное количество объектов, которые вы можете наделить магическими свойствами этим умением одновременно, равно вашему модификатору Интеллекта (минимум один объект). " +
			"Если вы пытаетесь превысить свой максимум, самое старое свойство немедленно заканчивается, а затем начинает действовать новое свойство.",
	}

	classAbilityList := []ClassAbility{abil1}

	return classAbilityList
}

func setWarlockClassAbilities() []ClassAbility {
	var specAbilities []ClassAbility
	var partnerDescription string

	otherSidePatrons := []string{"Архифея", "Исчадие", "Великий древний", "Бездонный (TCE)", "Нежить (VRGR)"}
	randNum, _ := random.IntRange(0, len(otherSidePatrons))
	partner := otherSidePatrons[randNum]
	specAbilities, partnerDescription, ClassSpecialSpells = SetOtherSidePatrons(partner)

	abil1 := ClassAbility{
		Level: 1,
		Name:  "Потусторонний покровитель",
		Description: "Ваш потусторонний покровитель -  <strong>" + partner +
			"</strong>. " + partnerDescription + "<br>" +
			"Ваш выбор определит умения, предоставляемые вам на 1-м, 6-м, 10-м и 14-м уровнях.",
	}

	classAbilityList := []ClassAbility{abil1}
	classAbilityList = append(classAbilityList, specAbilities...)

	return classAbilityList
}
