package classes

var CharacterLevelGlobal int

func GetClassSpells(level int) []string {

	type spellList struct {
		Level int
		List  []string
		Count int
	}
	type classSpell struct {
		ClassName string
		Spells    []spellList
	}

	spellData := []classSpell{
		{ClassName: "Волшебник",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{"Пляшущие огоньки", "Леденящее прикосновение", "Луч холода", "Сообщение", "Брызги кислоты", "Малая иллюзия", "Свет", "Электрошок", "Ядовитые брызги", "Фокусы", "Защита от оружия", "Меткий удар", "Огненный снаряд", "Починка", "Волшебная рука"},
					Count: 3,
				},
				{
					Level: 1,
					List:  []string{"Безмолвный образ", "Ведьмин снаряд", "Волна грома", "Волшебная стрела", "Доспехи мага", "Жуткий смех Таши", "Защита от зла и добра", "Луч болезни", "Маскировка", "Невидимое письмо", "Невидимый слуга", "Обнаружение магии", "Огненные ладони", "Опознание", "Очарование личности", "Падение пёрышком", "Поиск фамильяра", "Понимание языков", "Поспешное отступление", "Прыжок", "Псевдожизнь", "Сверкающие брызги", "Сигнал тревоги", "Скольжение", "Скороход", "Тензеров парящий диск", "Туманное облако", "Усыпление", "Цветной шарик", "Щит"},
					Count: 6,
				},
			},
		},
		{
			ClassName: "Воин",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{},
					Count: 0,
				},
				{
					Level: 1,
					List:  []string{},
					Count: 0,
				},
			},
		},
		{
			ClassName: "Варвар",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{},
					Count: 0,
				},
				{
					Level: 1,
					List:  []string{},
					Count: 0,
				},
			},
		},
		{
			ClassName: "Паладин",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{},
					Count: 0,
				},
				{
					Level: 1,
					List:  []string{},
					Count: 0,
				},
			},
		},
		{
			ClassName: "Монах",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{},
					Count: 0,
				},
				{
					Level: 1,
					List:  []string{},
					Count: 0,
				},
			},
		},
		{
			ClassName: "Плут",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{},
					Count: 0,
				},
				{
					Level: 1,
					List:  []string{},
					Count: 0,
				},
			},
		},
		{
			ClassName: "Следопыт",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{},
					Count: 0,
				},
				{
					Level: 1,
					List:  []string{},
					Count: 0,
				},
			},
		},
		{
			ClassName: "Друид",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{"Сотворение пламени", "Указание", "Сопротивление", "Искусство друидов", "Починка", "Свет", "Ядовитые брызги", "Терновый кнут", "Дубинка"},
					Count: 2,
				},
				{
					Level: 1,
					List:  []string{"Волна грома", "Дружба с животными", "Защита от зла и добра", "Лечащее слово", "Лечение ран", "Обнаружение болезней и яда", "Обнаружение магии", "Огонь фей", "Опутывание", "Очарование личности", "Очищение пищи и питья", "Прыжок", "Разговор с животными", "Скороход", "Сотворение или уничтожение воды", "Туманное облако", "Чудо-ягоды"},
					Count: ModifierGlobal.Wisdom + CharacterLevelGlobal,
				},
			},
		},
		{
			ClassName: "Жрец",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{"Починка", "Свет", "Священное пламя", "Сопротивление", "Указание", "Уход за умирающим", "Чудотворство"},
					Count: 3,
				},
				{
					Level: 1,
					List:  []string{"Благословение", "Защита от зла и добра", "Лечащее слово", "Лечение ран", "Нанесение ран", "Направляющий снаряд", "Обнаружение болезней и яда", "Обнаружение зла и добра", "Обнаружение магии", "Очищение пищи и питья", "Порча", "Приказ", "Сотворение или уничтожение воды", "Убежище", "Щит веры"},
					Count: ModifierGlobal.Wisdom + CharacterLevelGlobal,
				},
			},
		},
		{
			ClassName: "Чародей",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{"Волшебная рука", "Огненный снаряд", "Пляшущие огоньки", "Фокусы", "Брызги кислоты", "Защита от оружия", "Починка", "Свет", "Дружба", "Леденящее прикосновение", "Луч холода", "Меткий удар", "Сообщение", "Малая иллюзия", "Электрошок", "Ядовитые брызги"},
					Count: 4,
				},
				{
					Level: 1,
					List:  []string{"Безмолвный образ", "Ведьмин снаряд", "Волна грома", "Волшебная стрела", "Доспехи мага", "Луч болезни", "Маскировка", "Обнаружение магии", "Огненные ладони", "Очарование личности", "Падение пёрышком", "Понимание языков", "Поспешное отступление", "Прыжок", "Псевдожизнь", "Сверкающие брызги", "Скольжение", "Туманное облако", "Усыпление", "Цветной шарик", "Щит"},
					Count: 2,
				},
			},
		},
		{
			ClassName: "Бард",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{"Защита от оружия", "Злая насмешка", "Малая иллюзия", "Починка", "Сообщение", "Волшебная рука", "Дружба", "Меткий удар", "Пляшущие огоньки", "Свет", "Фокусы"},
					Count: 2,
				},
				{
					Level: 1,
					List:  []string{"Безмолвный образ", "Волна грома", "Героизм", "Диссонирующий шёпот", "Дружба с животными", "Жуткий смех Таши", "Лечащее слово", "Лечение ран", "Маскировка", "Невидимое письмо", "Невидимый слуга", "Обнаружение магии", "Огонь фей", "Опознание", "Очарование личности", "Падение пёрышком", "Понимание языков", "Порча", "Приказ", "Разговор с животными", "Сверкающие брызги", "Скороход", "Усыпление"},
					Count: 4,
				},
			},
		},
		{
			ClassName: "Изобретатель",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{"Свет", "Электрошок", "Сопротивление", "Терновый кнут", "Луч холода", "Пляшущие огоньки", "Уход за умирающим", "Фокусы", "Огненный снаряд", "Сообщение", "Починка", "Указание", "Брызги кислоты", "Волшебная рука", "Ядовитые брызги"},
					Count: 2,
				},
				{
					Level: 1,
					List:  []string{"Лечение ран", "Маскировка", "Обнаружение магии", "Огонь фей", "Опознание", "Очищение пищи и питья", "Падение пёрышком", "Поспешное отступление", "Прыжок", "Псевдожизнь", "Сигнал тревоги", "Скольжение", "Скороход", "Убежище"},
					Count: ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2),
				},
			},
		},
		{
			ClassName: "Колдун",
			Spells: []spellList{
				{
					Level: 0,
					List:  []string{"Малая иллюзия", "Мистический заряд", "Ядовитые брызги", "Волшебная рука", "Леденящее прикосновение", "Меткий удар", "Фокусы", "Дружба", "Защита от оружия"},
					Count: 2,
				},
				{
					Level: 1,
					List:  []string{"Адское возмездие", "Ведьмин снаряд", "Доспех Агатиса", "Защита от зла и добра", "Невидимое письмо", "Невидимый слуга", "Очарование личности", "Понимание языков", "Поспешное отступление", "Руки Хадара", "Сглаз"},
					Count: 2,
				},
			},
		},
	}

	var allClassSpells []spellList
	var classSpellsAnswer []string
	for _, classSpellValue := range spellData {
		if ClassNameGlobalRu == classSpellValue.ClassName {
			allClassSpells = classSpellValue.Spells
		}
	}

	var iter int
	for _, spells := range allClassSpells {
		if spells.Level == level {
			for _, sp := range spells.List {
				classSpellsAnswer = append(classSpellsAnswer, sp)
				iter++
				if iter >= spells.Count {
					break
				}
			}
		}
	}
	return classSpellsAnswer
}

func getClassSpellBasicCharacteristic() spellUsing {
	var su spellUsing

	switch ClassNameGlobalRu {
	case "Друид":
		su.BasicSpellCharacteristics = "Мудрость"
		su.SavingThrowDifficulty = 8 + 2 + ModifierGlobal.Wisdom
		su.SpellDamageModifier = 2 + ModifierGlobal.Wisdom
	case "Жрец":
		su.BasicSpellCharacteristics = "Мудрость"
		su.SavingThrowDifficulty = 8 + 2 + ModifierGlobal.Wisdom
		su.SpellDamageModifier = 2 + ModifierGlobal.Wisdom
	case "Изобретатель":
		su.BasicSpellCharacteristics = "Интеллект"
		su.SavingThrowDifficulty = 8 + 2 + ModifierGlobal.Intelligence
		su.SpellDamageModifier = 2 + ModifierGlobal.Intelligence
	case "Волшебник":
		su.BasicSpellCharacteristics = "Интеллект"
		su.SavingThrowDifficulty = 8 + 2 + ModifierGlobal.Intelligence
		su.SpellDamageModifier = 2 + ModifierGlobal.Intelligence
	case "Колдун":
		su.BasicSpellCharacteristics = "Харизма"
		su.SavingThrowDifficulty = 8 + 2 + ModifierGlobal.Charisma
		su.SpellDamageModifier = 2 + ModifierGlobal.Charisma
	case "Бард":
		su.BasicSpellCharacteristics = "Харизма"
		su.SavingThrowDifficulty = 8 + 2 + ModifierGlobal.Charisma
		su.SpellDamageModifier = 2 + ModifierGlobal.Charisma
	case "Чародей":
		su.BasicSpellCharacteristics = "Харизма"
		su.SavingThrowDifficulty = 8 + 2 + ModifierGlobal.Charisma
		su.SpellDamageModifier = 2 + ModifierGlobal.Charisma
	}

	return su
}
