package cleric

import (
	"github.com/mazen160/go-random"
	"math/rand"
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/general"
	"pregen/pkg/races"
	"pregen/pkg/spells"
	"time"
)

var (
	spellList  []spells.SpellsJSON
	skillCount = 2
)

func GetClass(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int, classArchetypeName string) *classes.Class {
	var proficiencyBonus = classes.GetProficiencyBonus(lvl)
	var statsPriority = []string{"Wisdom", "Charisma"}
	var abilitiesScore = general.GetClassAbilitiesScore(statsPriority, raceInfo, lvl)
	var modifier = classes.GetModifiersForClass(abilitiesScore)
	var savingThrows = classes.GetSaveThrowsForClass(modifier, statsPriority)

	var classAbilities, proficiencies = getClassAbilities(raceInfo, backgrInfo, lvl, classArchetypeName)
	var spellCastingInfo = classes.GetClassSpellBasicCharacteristic("Жрец", lvl, modifier, proficiencyBonus)

	var equip = getEquipment(proficiencies)

	return &classes.Class{
		ClassName:          "Cleric",
		ClassNameRU:        "Жрец",
		ClassArchetypeName: classArchetypeName,
		ClassAbilities:     classAbilities,
		AbilityScore:       abilitiesScore,
		AbilityModifier:    modifier,
		SavingThrows:       savingThrows,
		Inspiration:        false,
		Proficiencies:      *proficiencies, //need to fix
		ProficiencyBonus:   proficiencyBonus,
		Hits:               general.GetHits("d8", modifier, lvl),
		Caster:             true,
		SpellCasting:       spellCastingInfo,
		SpellsList:         getSpells(modifier, lvl),
		Equipment:          equip,
		ArmorInfo:          GetArmorInfo(equip, classArchetypeName, lvl),
		WeaponInfo:         GetWeaponInfo(equip),
	}
}

func getBasicProficiencies(raceInfo *races.Race, backgrInfo *backgrounds.Background) *classes.Proficiencies {
	var availableSkillList = []string{"History", "Medicine", "Perception", "Religion", "Persuasion"}
	classSkills := classes.GetClassSkillsArray(
		raceInfo.RaceSkill,
		backgrInfo.BackgroundSkills,
		availableSkillList,
		skillCount,
	)

	return &classes.Proficiencies{
		Armor:         []string{"Лёгкие доспехи, средние доспехи, щиты"},
		Weapons:       []string{"Простое оружие"},
		Tools:         []string{"Нет"},
		SavingThrow:   []string{"Wisdom", "Charisma"},
		SkillsOfClass: classSkills,
	}
}

func getEquipment(prof *classes.Proficiencies) []classes.Item {
	var holySymbols = []string{"Амулет", "Реликварий", "Эмблема"}
	randSymbol, _ := random.IntRange(0, len(holySymbols))
	var equipment = []classes.Item{
		{
			Name:     "Щит",
			ItemType: "Armor",
			Count:    1,
		},
		{
			Name:     "Священный символ: " + holySymbols[randSymbol],
			ItemType: "Tool",
			Count:    1,
		},
	}

	variant, _ := random.IntRange(0, 1)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Набор священника",
				ItemType: "Tool",
				Count:    1,
			})
	case 1:
		equipment = append(equipment,
			classes.Item{
				Name:     "Набор путешественника",
				ItemType: "Tool",
				Count:    1,
			})
	}

	variant, _ = random.IntRange(0, 2)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Лёгкий арбалет",
				ItemType: "Weapon",
				Count:    1,
			})
	case 1:
		var weaponsList []string
		for _, weapon := range classes.WeaponData {
			if weapon.WeaponType == "Простое рукопашное оружие" ||
				weapon.Name != "Боевой молот" && weapon.Name != "Булава" {
				weaponsList = append(weaponsList, weapon.Name)
			}
		}
		randNum, _ := random.IntRange(0, len(weaponsList))
		equipment = append(equipment,
			classes.Item{
				Name:     weaponsList[randNum],
				ItemType: "Weapon",
				Count:    1,
			})
	}

	variant, _ = random.IntRange(0, 2)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Чешуйчатый доспех",
				ItemType: "Armor",
				Count:    1,
			})
	case 1:
		equipment = append(equipment,
			classes.Item{
				Name:     "Кожаный доспех",
				ItemType: "Armor",
				Count:    1,
			})
	case 2:
		for _, armor := range prof.Armor {
			if armor == "Кольчуга" || armor == "Тяжелый доспех" {
				equipment = append(equipment,
					classes.Item{
						Name:     "Кольчуга",
						ItemType: "Armor",
						Count:    1,
					})
			} else {
				// HomeBrew от Разраба
				equipment = append(equipment,
					classes.Item{
						Name:     "Кольчужная рубаха",
						ItemType: "Armor",
						Count:    1,
					})
			}
		}
	}

	variant, _ = random.IntRange(0, 2)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Булава",
				ItemType: "Weapon",
				Count:    1,
			})
	case 1:
		for _, weapon := range prof.Weapons {
			if weapon == "Боевой молот" || weapon == "Воинское рукопашное оружие" {
				equipment = append(equipment,
					classes.Item{
						Name:     "Боевой молот",
						ItemType: "Weapon",
						Count:    1,
					})
			} else {
				equipment = append(equipment,
					classes.Item{
						Name:     "Булава",
						ItemType: "Weapon",
						Count:    1,
					})
			}
		}
	}
	return equipment
}

func getClassAbilities(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int, classArchetypeName string) ([]classes.ClassAbility, *classes.Proficiencies) {
	var proficiencies = getBasicProficiencies(raceInfo, backgrInfo)

	var dangerousPoint string
	if lvl >= 5 && lvl <= 7 {
		dangerousPoint = "1/2 или ниже."
	}
	if lvl == 8 {
		dangerousPoint = "1 или ниже."
	}
	var clericClassAbilities = []classes.ClassAbility{
		{
			Level: 1,
			Name:  "Использование заклинаний",
			Description: "Будучи проводником божественной силы, вы можете накладывать заклинания жреца. " +
				"Вы найдёте список заклинаний, доступных жрецу в этом разделе: заклинания жреца.",
		},
		{
			Level: 1,
			Name:  "Божественный домен",
			Description: "Выберите один домен, связанный с вашим божеством. " +
				"Все домены детально рассмотрены в конце описания класса, " +
				"и каждый содержит примеры божеств, связанных с ним.",
		},
		{
			Level: 2,
			Name:  "Божественный канал",
			Description: "Вы получаете возможность направлять божественную энергию непосредственно от своего божества, " +
				"используя её для подпитки магических эффектов. " +
				"Вы начинаете с двумя такими эффектами: «Изгнание Нежити» и эффектом, определяемым вашим доменом. " +
				"Некоторые домены дают вам дополнительные эффекты, как только вы получите новые уровни.",
		},
		{
			Level: 2,
			Name:  "Божественный канал: ИЗГНАНИЕ НЕЖИТИ",
			Description: "Вы действием демонстрируете свой священный символ и читаете молитву, изгоняющую Нежить. " +
				"Вся Нежить, которая может видеть или слышать вас в пределах 30 футов, должна совершить спасбросок Мудрости. " +
				"Если существо провалило спасбросок, оно изгоняется на 1 минуту, или пока не получит урон. " +
				"Изгнанное существо должно тратить свои ходы, пытаясь уйти от вас как можно дальше, " +
				"и не может добровольно переместиться в пространство, находящееся в пределах 30 футов от вас. " +
				"Оно также не может совершать реакции. " +
				"Действием существо может использовать только Рывок или пытаться освободиться от эффекта, " +
				"препятствующего его передвижению. Если двигаться некуда, существо может использовать действие Уклонение.",
		},
		{
			Level: 5,
			Name:  "Уничтожение нежити",
			Description: "Когда Нежить проваливает спасбросок от вашего умения «Изгнание Нежити», " +
				"существо мгновенно уничтожается, если его показатель опасности не превышает значения " + dangerousPoint,
		},
	}
	switch classArchetypeName {
	case "Домен бури":
		proficiencies.Armor = append(proficiencies.Armor, "Тяжелый доспех")
		proficiencies.Weapons = append(proficiencies.Weapons, "Воинское рукопашное оружие")

		var addSpellMap = map[int][]string{
			1: []string{"Волна грома", "Туманное облако"},
			3: []string{"Дребезги", "Порыв ветра"},
			5: []string{"Метель", "Призыв молнии"},
			7: []string{"Пласть над водами", "Град"},
			9: []string{"Нашествие насекомых", "Разрушительная волна"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение воинским оружием и тяжёлыми доспехами.",
			},
			{
				Level: 1,
				Name:  "Гнев бури",
				Description: "Вы можете громогласно покарать атакующих. " +
					"Если существо в пределах 5 футов от вас, которое вы можете видеть, успешно попадает по вам атакой, " +
					"вы можете реакцией заставить существо совершить спасбросок Ловкости. " +
					"Существо получает 2к8 урона звуком или электричеством (по вашему выбору), если провалит спасбросок, " +
					"и половину этого урона если преуспеет. Вы можете использовать это умение количество раз, " +
					"равное вашему модификатору Мудрости (минимум 1 раз). " +
					"Вы восстанавливаете все потраченные применения, когда заканчиваете продолжительный отдых.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Разрушительный гнев",
				Description: "Вы можете использовать «Божественный канал», чтобы овладеть могуществом бури с необузданной свирепостью. " +
					"Когда вы совершаете бросок урона звуком или электричеством, " +
					"вы можете использовать «Божественный канал», чтобы нанести максимальный урон вместо броска.",
			},
			{
				Level:       6,
				Name:        "Удар грома",
				Description: "Когда вы наносите урон электричеством существу с размером Большое или меньше, вы также можете оттолкнуть его на 10 футов от себя.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наполнять удары своего оружия божественной энергией. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, вы можете причинить цели дополнительный урон звуком 1к8. " +
					"Когда вы достигаете 14-го уровня, дополнительный урон возрастает до 2к8.",
			},
		}...)
	case "Домен войны":
		proficiencies.Armor = append(proficiencies.Armor, "Тяжелый доспех")
		proficiencies.Weapons = append(proficiencies.Weapons, "Воинское рукопашное оружие")

		var addSpellMap = map[int][]string{
			1: []string{"Божественное благоволение", "Щит веры"},
			3: []string{"Божественное оружие", "Магическое оружие"},
			5: []string{"Духовные стражи", "Мантия крестоносца"},
			7: []string{"Каменная кожа", "Свобода перемещения"},
			9: []string{"Небесный огонь", "Удержание чудовища"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение воинским оружием и тяжёлыми доспехами.",
			},
			{
				Level: 1,
				Name:  "Боевой священник",
				Description: "Ваш бог наделяет вас воодушевлением, когда вы вступаете в битву. " +
					"Когда вы используете действие Атака, вы можете совершить одну атаку оружием бонусным действием. " +
					"Вы можете использовать это умение количество раз, равное вашему модификатору Мудрости (минимум 1 раз). " +
					"Вы восстанавливаете все потраченные использования, когда завершаете продолжительный отдых.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Направленный удар",
				Description: "Вы можете использовать «Божественный канал», чтобы нанести удар со сверхъестественной точностью. " +
					"Когда вы совершаете бросок атаки, вы можете использовать «Божественный канал», чтобы получить бонус +10 к броску. " +
					"Вы можете решить, применять ли это умение, после того как увидите результат броска, " +
					"но до того как Мастер скажет, попала атака или промахнулась.",
			},
			{
				Level:       6,
				Name:        "Божественный канал: БЛАГОСЛОВЕНИЕ БОГА ВОЙНЫ",
				Description: "Когда вы наносите урон электричеством существу с размером Большое или меньше, вы также можете оттолкнуть его на 10 футов от себя.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наполнять удары своего оружия божественной энергией. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, " +
					"вы можете причинить цели дополнительный урон 1к8 того же вида, что и у оружия. " +
					"Когда вы достигаете 14 уровня, дополнительный урон увеличивается до 2к8.",
			},
		}...)
	case "Домен жизни":
		proficiencies.Armor = append(proficiencies.Armor, "Тяжелый доспех")

		var addSpellMap = map[int][]string{
			1: []string{"Благословение", "Лечение ран"},
			3: []string{"Божественное оружие", "Малое восстановление"},
			5: []string{"Возрождение", "Маяк надежды"},
			7: []string{"Защита от смерти", "Страж веры"},
			9: []string{"Множественное лечение ран", "Оживление"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение тяжёлыми доспехами.",
			},
			{
				Level: 1,
				Name:  "Поборник жизни",
				Description: "Ваши лечащие заклинания становятся более эффективными. " +
					"Каждый раз, когда вы используете заклинание 1-го уровня или выше, восстанавливающее хиты существу, " +
					"это существо восстанавливает дополнительно число хитов, равное 2 + уровень заклинания.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Сохранение жизни",
				Description: "Вы можете использовать «Божественный канал», чтобы нанести удар со сверхъестественной точностью. " +
					"Когда вы совершаете бросок атаки, вы можете использовать «Божественный канал», чтобы получить бонус +10 к броску. " +
					"Вы можете решить, применять ли это умение, после того как увидите результат броска, " +
					"но до того как Мастер скажет, попала атака или промахнулась.",
			},
			{
				Level: 6,
				Name:  "Благословенный целитель",
				Description: "Заклинания лечения, которое вы накладываете на других, также лечат и вас. " +
					"Когда вы накладываете заклинание 1-го уровня или выше, которое восстанавливает хиты другому существу, " +
					"вы восстанавливаете себе число хитов, равное 2 + уровень заклинания.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наполнять удары своего оружия божественной энергией. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, " +
					"вы можете причинить цели дополнительный урон излучением 1к8. " +
					"Когда вы достигаете 14-го уровня, дополнительный урон возрастает до 2к8.",
			},
		}...)
	case "Домен знаний":
		raceInfo.Language = append(raceInfo.Language, "Два языка на свой выбор")

		var availableSkillList = []string{"History", "Nature", "Arcana", "Religion"}
		// Seed the random number generator to produce different results each time
		rand.Seed(time.Now().UnixNano())

		// Shuffle the slice
		rand.Shuffle(len(availableSkillList), func(i, j int) {
			availableSkillList[i], availableSkillList[j] = availableSkillList[j], availableSkillList[i]
		})

		// Get the first two elements from the shuffled slice
		randomSkills := availableSkillList[:2]

		newClassSkills := classes.GetClassSkillsArray(
			raceInfo.RaceSkill,
			backgrInfo.BackgroundSkills,
			randomSkills,
			skillCount,
		)
		proficiencies.SkillsOfClass = append(proficiencies.SkillsOfClass, newClassSkills...)

		// TODO
		//bag-0002 Нужно как-то добавить +2 к выбранным скиллам из randomSkills
		//слишком далеко лесть в код скиллов, чтобы пробрасывать ради одного домена
		//и менять обработчик скиллов

		var addSpellMap = map[int][]string{
			1: []string{"Опознание", "Приказ"},
			3: []string{"Внушение", "Гадание"},
			5: []string{"Необнаружимость", "Разговор с мёртвыми"},
			7: []string{"Магический глаз", "Смятение"},
			9: []string{"Знание легенд", "Наблюдение"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение тяжёлыми доспехами.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Знания веков",
				Description: "Вы можете использовать «Божественный канал», чтобы получить доступ к источнику знаний. " +
					"Вы действием выбираете навык или инструмент. На 10 минут вы осваиваете владение выбранным навыком или инструментом.",
			},
			{
				Level: 6,
				Name:  "Божественный канал: Чтение мыслей",
				Description: "Вы можете использовать свой «Божественный канал», чтобы читать мысли существ. " +
					"Затем вы можете использовать доступ к разуму существа, чтобы командовать им. <br>" +
					"Выберите действием одно существо, которое вы можете видеть, находящееся в пределах 60 футов от вас. " +
					"Это существо должно совершить спасбросок Мудрости. " +
					"Если существо преуспело, вы не можете использовать это умение на нём, пока не завершите продолжительный отдых. " +
					" существо проваливает спасбросок, вы можете прочитать его поверхностные мысли " +
					"(то, что у него на уме, его текущие эмоции и то, о чём оно активно думает), " +
					"когда оно находится в пределах 60 футов от вас. Этот эффект длится в течение 1 минуты." +
					"В течение этого времени вы можете действием окончить этот эффект, накладывая на существо заклинание внушение [suggestion], " +
					"не тратя ячейку заклинания. Цель автоматически проваливает спасбросок от этого заклинания.",
			},
			{
				Level:       8,
				Name:        "Могущественное колдовство",
				Description: "Вы добавляете модификатор Мудрости к урону, который причиняете заговорами жреца.",
			},
		}...)
	case "Домен обмана":
		var addSpellMap = map[int][]string{
			1: []string{"Маскировка", "Очарование личности"},
			3: []string{"Бесследное передвижение", "Отражения"},
			5: []string{"Мерцание", "Рассеивание магии"},
			7: []string{"Переносящая дверь", "Превращение"},
			9: []string{"Изменение памяти", "Подчинение личности"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level: 1,
				Name:  "Благословение обманщика",
				Description: "Вы можете действием коснуться согласного существа, отличного от вас, " +
					"чтобы позволить ему совершать с преимуществом проверки Ловкости (Скрытность). " +
					"Это благословение длится 1 час, или пока вы не используете это умение снова.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Двуличность",
				Description: "Вы можете использовать «Божественный канал», чтобы создать иллюзорную копию себя. " +
					"Вы действием создаёте идеальную иллюзию себя, которая существует в течение 1 минуты, " +
					"или пока вы не потеряете концентрацию (как если бы вы концентрировались на заклинании). " +
					"Иллюзия появляется в свободном пространстве, которое вы можете видеть в пределах 30 футов от себя. " +
					"Бонусным действием в свой ход вы можете переместить иллюзию на расстояние до 30 футов в пространство, " +
					"которое вы можете видеть, но иллюзия должна оставаться в пределах 120 футов от вас. " +
					"На протяжении действия умения вы можете накладывать заклинания, как если бы вы находились в пространстве иллюзии, " +
					"но вы должны использовать собственные чувства. Кроме того, когда и вы и ваша иллюзия находитесь в пределах 5 футов от существа, " +
					"которое может видеть иллюзию, вы совершаете броски атаки по этому существу с преимуществом, так как иллюзия отвлекает его.",
			},
			{
				Level: 6,
				Name:  "Божественный канал: Плащ теней",
				Description: "Вы можете использовать «Божественный канал», чтобы исчезать. " +
					"Вы действием становитесь невидимым до конца своего следующего хода. " +
					"Вы становитесь видимым, если атакуете или накладываете заклинание.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наполнять удары своего оружия ядом — это дар вашего божества. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, вы можете причинить цели дополнительный урон ядом 1к8. " +
					"Когда вы достигаете 14-го уровня, дополнительный урон возрастает до 2к8.",
			},
		}...)
	case "Домен природы":
		proficiencies.Armor = append(proficiencies.Armor, "Тяжелый доспех")

		druidZeroLevelSpells := spells.GetAllSpellForClass("Друид", 0)
		rand.Seed(time.Now().UnixNano())
		// Shuffle the slice
		rand.Shuffle(len(druidZeroLevelSpells), func(i, j int) {
			druidZeroLevelSpells[i], druidZeroLevelSpells[j] = druidZeroLevelSpells[j], druidZeroLevelSpells[i]
		})

		for _, druidSpell := range druidZeroLevelSpells {
			spellList = append(spellList, druidSpell)
			break
		}

		var availableSkillList = []string{"Survival", "Nature", "AnimalHandling"}
		// Seed the random number generator to produce different results each time
		rand.Seed(time.Now().UnixNano())

		// Shuffle the slice
		rand.Shuffle(len(availableSkillList), func(i, j int) {
			availableSkillList[i], availableSkillList[j] = availableSkillList[j], availableSkillList[i]
		})

		// Get the first two elements from the shuffled slice
		randomSkills := availableSkillList[:2]

		newClassSkills := classes.GetClassSkillsArray(
			raceInfo.RaceSkill,
			backgrInfo.BackgroundSkills,
			randomSkills,
			skillCount,
		)
		proficiencies.SkillsOfClass = append(proficiencies.SkillsOfClass, newClassSkills...)

		var addSpellMap = map[int][]string{
			1: []string{"Дружба с животными", "Разговор с животными"},
			3: []string{"Дубовая кора", "Шипы"},
			5: []string{"Рост растений", "Стена ветров"},
			7: []string{"Подчинение зверя", "Цепкая лоза"},
			9: []string{"Древесный путь", "Нашествие насекомых"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level: 1,
				Name:  "Послушник природы",
				Description: "Вы изучаете один заговор друида на свой выбор. " +
					"Этот заговор считается заговором жреца для вас, но он не учитывается в общем количестве известных вам заговоров жреца. " +
					"Вы также получаете владение одним из следующих навыков: Выживание, Природа, Уход за животными.",
			},
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение тяжёлыми доспехами.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Очарование животных и растений",
				Description: "Вы можете использовать «Божественный канал», чтобы очаровать животных и растения. " +
					"Вы действием демонстрируете свой священный символ и называете имя своего божества. " +
					"Все Звери и Растения, которые могут видеть вас, и находятся в пределах 30 футов от вас, должны совершить спасбросок Мудрости. " +
					"Если существо провалит спасбросок, оно становится очарованным вами на 1 минуту, или пока не получит урон. " +
					"Пока существо очаровано, оно дружелюбно к вам и другим существам, которых вы укажете.",
			},
			{
				Level: 6,
				Name:  "Сдерживание стихий",
				Description: "Когда вы или существо в пределах 30 футов от вас получает урон звуком, кислотой, огнём, " +
					"холодом или электричеством, вы можете реакцией предоставить существу сопротивление этому урону.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наполнять удары своего оружия божественной энергией. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, " +
					"вы можете причинить цели дополнительный урон огнём, холодом или электричеством (на ваш выбор) 1к8. " +
					"Когда вы достигаете 14-го уровня, дополнительный урон возрастает до 2к8.",
			},
		}...)
	case "Домен света":
		spellList = append(spellList, spells.FindSpellInDB("Cвет"))

		var addSpellMap = map[int][]string{
			1: []string{"Огненные ладони", "Огонь фей"},
			3: []string{"Палящий луч", "Пылающий шар"},
			5: []string{"Дневной свет", "Огненный шар"},
			7: []string{"Огненная стена", "Страж веры"},
			9: []string{"Наблюдение", "Небесный огонь"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level: 1,
				Name:  "Дополнительный заговор",
				Description: "Вы изучаете заговор свет [light], если еще не имели его раньше. " +
					"Этот заговор не учитывается в количестве известных вам заговоров жреца.",
			},
			{
				Level: 1,
				Name:  "Защищающая вспышка",
				Description: "Вы можете создать божественный свет между собой и атакующим противником. " +
					"Если вас атакует существо, которое вы видите в пределах 30 футов, " +
					"вы можете реакцией создать помеху его броску атаки, " +
					"вызывая перед атакующим вспышку света, до того как он попадёт или промажет. " +
					"Существа, которые не могут быть ослеплены, обладают иммунитетом к этому умению. " +
					"Вы можете использовать это умение количество раз, равное вашему модификатору Мудрости (минимум один раз). " +
					"Вы восстанавливаете все потраченные использования, когда завершаете продолжительный отдых.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Сияние рассвета",
				Description: "Вы можете использовать «Божественный канал», чтобы призывать солнечный свет, прогоняющий тьму и причиняющий урон излучением врагам. " +
					"Вы действием демонстрируете свой священный символ, и вся магическая тьма в пределах 30 футов от вас рассеивается. " +
					"Кроме того, все враждебные существа в пределах 30 футов от вас должны совершить спасбросок Телосложения. " +
					"Существа получают урон излучением, равный 2к10 + ваш уровень жреца в случае провала, и половину этого урона в случае успешного спасброска. " +
					"Существа с полным укрытием от вас не подвержены воздействию. ",
			},
			{
				Level: 6,
				Name:  "Улучшенная вспышка",
				Description: "Вы можете использовать умение «Защищающая вспышка», когда существо, " +
					"которое вы можете видеть в пределах 30 футов от себя, атакует не вас, а другое существо.",
			},
			{
				Level:       8,
				Name:        "Могущественное колдовство",
				Description: "Вы добавляете модификатор Мудрости к урону, который причиняете заговорами жреца.",
			},
		}...)
	case "Домен смерти":
		proficiencies.Weapons = append(proficiencies.Weapons, "Воинское рукопашное оружие")

		var addSpellMap = map[int][]string{
			1: []string{"Псевдожизнь", "Луч болезни"},
			3: []string{"Глухота/слепота", "Луч слабости"},
			5: []string{"Восставший труп", "Прикосновение вампира"},
			7: []string{"Усыхание", "Защита от смерти"},
			9: []string{"Преграда жизни", "Облако смерти"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение воинским оружием",
			},
			{
				Level: 1,
				Name:  "Жнец",
				Description: "Вы изучаете один заговор из школы Некромантии из списка любого класса. " +
					"Когда вы накладываете заговор школы Некромантии, в обычных условиях нацеливающийся только на одно существо, " +
					"вы может нацелиться на двух существ, находящихся в пределах дистанции и при этом в пределах 5 футов друг от друга.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Касание смерти",
				Description: "Вы можете использовать «Божественный канал» для уничтожения касанием чужой жизненной силы." +
					"Когда вы попадаете по существу рукопашной атакой, вы можете использовать «Божественный канал», " +
					"чтобы нанести цели дополнительный урон некротической энергией, равный 5 + ваш удвоенный уровень Жреца.",
			},
			{
				Level: 6,
				Name:  "Неизбежное разрушение",
				Description: "Способность жреца проводить негативную энергию улучшается. " +
					"Урон некротической энергией, причиняемый вашими заклинаниями Жреца и «Божественным каналом», " +
					"игнорируют сопротивление урону некротической энергией.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наполнять удары своего оружия некротической энергией. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, " +
					"вы можете причинить цели дополнительный урон некротической энергией 1к8. " +
					"Когда вы достигаете 14-го уровня, дополнительный урон возрастает до 2к8.",
			},
		}...)
	case "Домен магии":
		wizardZeroLevelSpells := spells.GetAllSpellForClass("Волшебник", 0)
		// Seed the random number generator to produce different results each time
		rand.Seed(time.Now().UnixNano())

		// Shuffle the slice
		rand.Shuffle(len(wizardZeroLevelSpells), func(i, j int) {
			wizardZeroLevelSpells[i], wizardZeroLevelSpells[j] = wizardZeroLevelSpells[j], wizardZeroLevelSpells[i]
		})
		for _, wizardSpell := range wizardZeroLevelSpells {
			spellList = append(spellList, wizardSpell)
			break
		}

		var addSpellMap = map[int][]string{
			1: []string{"Обнаружение магии", "Волшебная стрела"},
			3: []string{"Магическое оружие", "Нистулова ложная магия"},
			5: []string{"Рассеивание магии", "Магический круг"},
			7: []string{"Магический глаз", "Леомундов потайной сундук"},
			9: []string{"Планарные узы", "Круг телепортации"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}
		var dangerousPoint string
		if lvl >= 5 && lvl <= 7 {
			dangerousPoint = "1/2 или ниже."
		}
		if lvl == 8 {
			dangerousPoint = "1 или ниже."
		}
		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level: 1,
				Name:  "Начинающий маг",
				Description: "Вы получаете владение навыком Магия и узнаёте два заговора на ваш выбор из списка заклинаний волшебника. " +
					"Для вас эти заговоры считаются заговорами жреца.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Ограждение магией",
				Description: "Вы можете использовать «Божественный канал», чтобы оградиться от существ из иных миров." +
					"Вы действием демонстрируете свой священный символ, и один выбранный вами Небожитель, Элементаль, Фея или Исчадие " +
					"в 30 футах от вас должен совершить спасбросок Мудрости, при условии, что существо вас видит и слышит. " +
					"Если существо проваливает спасбросок, оно изгоняется на 1 минуту или пока не получит любой урон. " +
					"Изгнанное существо должно тратить свои ходы, пытаясь уйти от вас как можно дальше, " +
					"и не может добровольно переместиться в пространство, находящееся в пределах 30 футов от вас. " +
					"Оно также не может совершать реакции. " +
					"Действием существо может использовать только Рывок или пытаться освободиться от эффекта, препятствующее его передвижению. " +
					"Если двигаться некуда, существо может использовать действие Уклонение. " +
					"После 5-го уровня, провалив спасбросок против вашей способности «Ограждения магией», " +
					"существо изгоняется на 1 минуту (как заклинание изгнание [banishment], концентрация не требуется), " +
					"если он не находится на своём родном плане и его показатель опасности равен или ниже порога " + dangerousPoint + " .",
			},
			{
				Level: 6,
				Name:  "Разрушитель заклинаний",
				Description: "При восстановлении хитов союзнику с помощью заклинания 1-го уровня или выше, " +
					"вы можете так же развеять одно выбранное вами заклинание, наложенное на это существо. " +
					"Уровень прерываемого заклинания должен быть равен или ниже уровню ячейки целебного заклинания.",
			},
			{
				Level:       8,
				Name:        "Могущественное колдовство",
				Description: "Вы добавляете модификатор Мудрости к урону, который причиняете заговорами жреца.",
			},
		}...)
	case "Домен кузни":
		proficiencies.Armor = append(proficiencies.Armor, "Тяжёлые доспехи")
		proficiencies.Tools = append(proficiencies.Tools, "Инструментами кузнеца")

		if lvl >= 6 {
			raceInfo.Resist = append(raceInfo.Resist, "Огненному урону")
		}

		var addSpellMap = map[int][]string{
			1: []string{"Опознание", "Палящая кара"},
			3: []string{"Раскалённый металл", "Магическое оружие"},
			5: []string{"Стихийное оружие", "Защита от энергии"},
			7: []string{"Изготовление", "Огненная стена"},
			9: []string{"Оживление вещей", "Сотворение"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение тяжёлыми доспехами и инструментами кузнеца.",
			},
			{
				Level: 1,
				Name:  "Благословление кузницы",
				Description: "Вы получаете способность наделять магией оружие или доспехи. " +
					"В конце продолжительного отдыха вы можете коснуться одного немагического предмета, " +
					"который является доспехом, простым или воинским оружием. " +
					"До конца вашего следующего продолжительного отдыха или до тех пор пока вы не умрёте, " +
					"этот объект считается магическим предметом, дающей бонус +1 к КД, если это доспех, " +
					"или бонус +1 к броскам атаки и урона, если это оружие." +
					"Использовав это умение, вы не можете использовать его повторно до окончания продолжительного отдыха.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Благословение ремесленника",
				Description: "Вы можете использовать свой «Божественный канал» для создания простых вещей. " +
					"Вы проводите ритуал длительностью в час, который создаёт немагический предмет, содержащий немного металла: " +
					"простое или воинское оружие, доспехи, десять боеприпасов, набор инструментов или другой металлический объект. " +
					"Создание завершается через час, предмет появляется в свободном пространстве по вашему выбору на поверхности в пределах 5 футов от вас. " +
					"То, что вы создаёте, может быть чем-то, что стоит не более 100 зм. " +
					"В рамках этого ритуала вы должны выложить металл, который может содержать монеты, со стоимостью, равной стоимости создаваемого предмета. " +
					"Металл безвозвратно сливается и превращается в творение в конце ритуала, магически формируя даже неметаллические части творения." +
					"Ритуал может создать дубликат немагического предмета, который содержит металл, такой как ключ, если вы обладаете оригиналом во время ритуала.",
			},
			{
				Level: 6,
				Name:  "Душа кузницы",
				Description: "Ваше мастерство жреца кузни предоставляет вам особые способности: <br>" +
					"* Вы получаете сопротивление огненному урону." +
					"* Когда вы носите тяжёлый доспех, то получаете +1 к КД",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете возможность наполнять свои удары оружием огненной силой кузницы. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, вы можете причинить цели дополнительный урон огнём 1к8. " +
					"Когда вы достигаете 14-го уровня, дополнительный урон увеличивается до 2к8.",
			},
		}...)
	case "Домен упокоения":
		spellList = append(spellList, spells.FindSpellInDB("Уход за умирающим"))

		var addSpellMap = map[int][]string{
			1: []string{"Порча", "Псевдожизнь"},
			3: []string{"Нетленные останки", "Луч слабости"},
			5: []string{"Возрождение", "Прикосновение вампира"},
			7: []string{"Усыхание", "Защита от смерти"},
			9: []string{"Преграда жизни", "Оживление"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level: 1,
				Name:  "Круг смерти",
				Description: "Вы получаете возможность управлять гранью между жизнью и смертью. " +
					"Если обычно вы должны бросить одну или несколько костей, чтобы заклинанием восстановить хиты существу с 0 хитов, " +
					"то вместо этого вы используете максимально возможное число, для каждой кости. " +
					"Кроме того, вы изучаете заговор уход за умирающим [spare the dying], " +
					"который не учитывается при подсчёте известных вам заговоров жреца. " +
					"Для вас он имеет дистанцию 30 футов, и вы можете накладывать его бонусным действием.",
			},
			{
				Level: 1,
				Name:  "Глаза могилы",
				Description: "Вы получаете возможность иногда ощущать присутствие нежити, существование которой является оскорблением естественного цикла жизни. <br>" +
					"Действием вы можете открыть свой разум, чтобы магическим образом обнаружить Нежить. " +
					"До конца вашего следующего хода вы знаете местоположение любой Нежити в пределах 60 футов от вас, " +
					"которая не находится за полным укрытием и не защищена от магии прорицания. <br>" +
					"Это чувство ничего не позволяет вам узнать о возможностях существа или идентифицировать его. " +
					"Вы можете использовать это умение количество раз, равное вашему модификатору Мудрости (минимум 1 раз). <br>" +
					"Вы восстанавливаете все потраченные применения после окончания продолжительного отдыха.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Путь к могиле",
				Description: "Вы можете использовать свой «Божественный канал», " +
					"чтобы пометить жизненную силу другого существа как подлежащую уничтожению." +
					"Действием вы выбираете одно существо, которое вы можете видеть, в пределах 30 футов от вас, " +
					"проклиная его до конца вашего следующего хода. " +
					"Когда вы или союзник попадете по проклятому существу атакой, " +
					"у существа будет уязвимость ко всему урону от атаки, а затем проклятие заканчивается.",
			},
			{
				Level: 6,
				Name:  "Страж на пороге смерти",
				Description: "Вы получаете возможность препятствовать смерти." +
					"Когда вы или существо, которое вы можете видеть в пределах 30 футов, подвергаются критическому удару, " +
					"вы можете реакцией превратить этот удар в нормальный удар. Любые эффекты, вызванные критическим ударом, будут отменены. " +
					"Вы можете использовать это умение количество раз, равное вашему модификатору Мудрости (минимум 1 раз). " +
					"Вы восстанавливаете все потраченные применения, после окончания продолжительного отдыха.",
			},
			{
				Level:       8,
				Name:        "Могущественное колдовство",
				Description: "Вы добавляете модификатор Мудрости к урону, который вы наносите любым заговором жреца.",
			},
		}...)
	case "Домен мира":
		var availableSkillList = []string{"Perception", "Performance", "Persuasion"}
		// Seed the random number generator to produce different results each time
		rand.Seed(time.Now().UnixNano())

		// Shuffle the slice
		rand.Shuffle(len(availableSkillList), func(i, j int) {
			availableSkillList[i], availableSkillList[j] = availableSkillList[j], availableSkillList[i]
		})

		// Get the first two elements from the shuffled slice
		randomSkills := availableSkillList[:1]

		newClassSkills := classes.GetClassSkillsArray(
			raceInfo.RaceSkill,
			backgrInfo.BackgroundSkills,
			randomSkills,
			skillCount,
		)
		proficiencies.SkillsOfClass = append(proficiencies.SkillsOfClass, newClassSkills...)

		var addSpellMap = map[int][]string{
			1: []string{"Героизм", "Убежище"},
			3: []string{"Охраняющая связь", "Подмога"},
			5: []string{"Маяк надежды", "Послание"},
			7: []string{"Аура очищения", "Отилюков упругий шар"},
			9: []string{"Высшее восстановление", "Ментальная связь Рэри"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Установление мира",
				Description: "Вы получаете владение навыком Проницательность, Выступление или Убеждение.",
			},
			{
				Level: 1,
				Name:  "Ободряющая связь",
				Description: "Вы можете создать укрепляющую связь между существами, которые находятся в мире друг с другом. " +
					"Действием вы можете выбрать количество согласных существ (целью можете быть и вы), " +
					"равное вашему бонусу мастерства, которых вы можете видеть в пределах 30 футов от вас. " +
					"Вы создаёте между этими существами магическую связь, которая длится 10 минут или " +
					"до тех пор, пока не используете это умение повторно. " +
					"Пока любое из связанных существ находится в пределах 30 футов от другого," +
					" каждое существо может бросить к4 и добавить выпавшее значение к своим броску атаки, проверке характеристики или спасброску. " +
					"Каждое существо может добавлять к4 не более одного раза за ход. " +
					"Вы можете использовать это умение количество раз, равное вашему бонусу мастерства. " +
					"Вы восстанавливаете все потраченные использования после окончания продолжительного отдыха.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Бальзам мира",
				Description: "Вы можете использовать свой «Божественный канал», чтобы сделать само своё присутствие умиротворяющим бальзамом. " +
					"Действием вы можете передвинуться на любое расстояние вплоть до своей скорости, не провоцируя атак. " +
					"Когда вы оказываетесь в пределах 5 футов от любого другого существа во время этого действия, " +
					"вы можете восстановить существу количество хитов, равное 2к6 + ваш модификатор Мудрости (минимум 1 хит). " +
					"Существо может получить это исцеление только один раз, пока вы совершаете это действие.",
			},
			{
				Level: 6,
				Name:  "Защитная связь",
				Description: "Связь, которую вы устанавливаете между личностями, помогает им защищать друг друга. " +
					"Когда существо, связанное вашим умением «Ободряющая связь», должно получить урон, другое связанное с ним существо," +
					" находящееся в пределах 30 футов от него, может реакцией телепортироваться в свободное пространство в пределах 5 футов от первого существа, " +
					"после чего оно получает весь урон, который был предназначен первому существу.",
			},
			{
				Level:       8,
				Name:        "Могущественное колдовство",
				Description: "Вы добавляете свой модификатор Мудрости к урону, который наносите любыми заговорами жреца.",
			},
		}...)
	case "Домен порядка":
		proficiencies.Armor = append(proficiencies.Armor, "Тяжёлые доспехи")

		var availableSkillList = []string{"Intimidation", "Persuasion"}
		// Seed the random number generator to produce different results each time
		rand.Seed(time.Now().UnixNano())

		// Shuffle the slice
		rand.Shuffle(len(availableSkillList), func(i, j int) {
			availableSkillList[i], availableSkillList[j] = availableSkillList[j], availableSkillList[i]
		})

		// Get the first two elements from the shuffled slice
		randomSkills := availableSkillList[:1]

		newClassSkills := classes.GetClassSkillsArray(
			raceInfo.RaceSkill,
			backgrInfo.BackgroundSkills,
			randomSkills,
			skillCount,
		)
		proficiencies.SkillsOfClass = append(proficiencies.SkillsOfClass, newClassSkills...)

		var addSpellMap = map[int][]string{
			1: []string{"Героизм", "Приказ"},
			3: []string{"Область истины", "Удержание личности"},
			5: []string{"Замедление", "Множественное лечащее слово"},
			7: []string{"Поиск существа", "Принуждение"},
			9: []string{"Общение", "Подчинение личности"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусные владения",
				Description: "Вы получаете владение тяжёлыми доспехами. Вы также получаете владение навыком Запугивание или Убеждение.",
			},
			{
				Level: 1,
				Name:  "Голос авторитета",
				Description: "Вы можете призвать силу закона, чтобы повести союзников в атаку. " +
					"Когда вы накладываете заклинание на союзника, используя ячейку 1-го уровня и выше, " +
					"этот союзник может реакцией совершить одну атаку оружием по существу, которое вы можете видеть, по вашему выбору." +
					"Если заклинание затрагивает больше чем одного союзника, вы выбираете, кто из союзников может совершать атаку в данный момент.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Требование порядка",
				Description: "Вы можете использовать свой «Божественный канал», чтобы надавить на кого-то. " +
					"Действием вы демонстрируете священный символ. " +
					"Каждое существо по вашему выбору в пределах 30 футов от вас, которое может слышать или видеть вас, " +
					"должно преуспеть в спасброске Мудрости, иначе будет очаровано вами до окончания вашего следующего хода или до того, " +
					"как очарованное существо получит урон. " +
					"При провале спасброска вы также можете заставить любое очарованное существо выронить то, что у него в руках.",
			},
			{
				Level: 6,
				Name:  "Воплощение закона",
				Description: "Вы становитесь невероятно сильны в применении магической энергии для подчинения других. " +
					"Когда вы накладываете заклинание школы Очарования, используя ячейку заклинаний 1-го уровня или выше, " +
					"со временем накладывания «1 действие», вы можете сменить время накладывания этого заклинания на «1 бонусное действие»." +
					"Вы можете использовать это умение количество раз, равное вашему модификатору Мудрости (минимум 1 раз). " +
					"Вы восстанавливаете все потраченные использования после окончания продолжительного отдыха.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наделять удары своего оружия божественной энергией. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, " +
					"вы можете дополнительно нанести цели 1к8 урона психической энергией. " +
					"Когда вы достигаете 14-го уровня, этот урон увеличивается до 2к8.",
			},
		}...)
	case "Домен сумерек":
		proficiencies.Armor = append(proficiencies.Armor, "Тяжелый доспех")
		proficiencies.Weapons = append(proficiencies.Weapons, "Воинское рукопашное оружие")

		var addSpellMap = map[int][]string{
			1: []string{"Огонь фей", "Усыпление"},
			3: []string{"Видение невидимого", "Лунный луч"},
			5: []string{"Фура живучести", "Леомундова хижина"},
			7: []string{"Аура жизни", "Высшая невидимость"},
			9: []string{"Круг силы", "Фальшивый двойник"},
		}

		for i, spellsStormDomain := range addSpellMap {
			if i <= lvl {
				for _, spell := range spellsStormDomain {
					spellList = append(spellList, spells.FindSpellInDB(spell))
				}
			}
		}

		clericClassAbilities = append(clericClassAbilities, []classes.ClassAbility{
			{
				Level:       1,
				Name:        "Бонусное владение",
				Description: "Вы получаете владение воинским оружием и тяжёлыми доспехами.",
			},
			{
				Level: 1,
				Name:  "Глаза ночи",
				Description: "Вы способны видеть сквозь непроглядную тьму. Вы получаете тёмное зрение в пределах 300 футов. " +
					"На этой дистанции вы можете видеть в тусклом свете так же, как при ярком свете, а в полной темноте — как при тусклом свете. " +
					"Действием вы можете магическим образом наделить этим тёмным зрением определённое количество согласных существ, " +
					"которых вы можете видеть в пределах 10 футов. Количество существ равно вашему модификатору Мудрости (минимум одно существо). " +
					"Эффект общего тёмного зрения длится 1 час. После того, как вы поделились своим тёмным зрением, вы не можете сделать это вновь, " +
					"пока не закончите продолжительный отдых или не потратите ячейку заклинаний любого уровня на повторное использование этого умения.",
			},
			{
				Level: 1,
				Name:  "Благословение бдительности",
				Description: "Ночь научила вас бдительности. " +
					"Действием вы касаетесь одного существа (включая себя) и даруете ему преимущество на совершение следующего броска инициативы. " +
					"Это преимущество пропадает сразу после совершения броска или при повторном использовании этого умения.",
			},
			{
				Level: 2,
				Name:  "Божественный канал: Сумеречное святилище",
				Description: "Вы можете использовать «Божественный канал», чтобы поддержать ваших союзников спокойствием сумерек. " +
					"Действием вы демонстрируете свой священный символ, и вас окружает сумеречная сфера. " +
					"Сфера, наполненная тусклым светом, появляется с центром на вас, её радиус равен 30 футам. " +
					"Сфера перемещается вместе с вами и существует 1 минуту или до тех пор, пока вы не станете недееспособны или не умрёте. " +
					"Каждый раз, когда существо (включая вас) оканчивает свой ход в пределах сферы, вы можете даровать этому существу одно из описанных ниже преимуществ: <br>" +
					"* Даровать временные хиты в количестве, равном 1к6 + ваш уровень жреца. " +
					"* Окончить одно действие эффекта очарования или испуга.",
			},
			{
				Level: 6,
				Name:  "Шаги ночи",
				Description: "Вы можете использовать мистическую силу ночи, чтобы подняться в воздух. " +
					"Бонусным действием, когда вы находитесь в тусклом свете или темноте, " +
					"вы можете магическим образом даровать себе скорость полёта, равную вашей скорости ходьбы, на 1 минуту. " +
					"Вы можете использовать бонусное действие подобным образом количество раз, равное вашему бонусу мастерства. " +
					"Вы восстанавливаете все потраченные использования после окончания продолжительного отдыха.",
			},
			{
				Level: 8,
				Name:  "Божественный удар",
				Description: "Вы получаете способность наделять удары своего оружия божественной энергией. " +
					"Один раз в каждый свой ход, когда вы попадаете по существу атакой оружием, " +
					"вы можете нанести цели дополнительно 1к8 урона излучением. " +
					"Когда вы достигаете 14-го уровня, этот урон увеличивается до 2к8.",
			},
		}...)
	}

	for _, ability := range clericClassAbilities {
		if lvl >= ability.Level {
			clericClassAbilities = append(clericClassAbilities, ability)
		}
	}

	return clericClassAbilities, proficiencies
}

func GetArmorInfo(equip []classes.Item, classArchetypeName string, lvl int) classes.Armor {
	var armorData = classes.GetArmorFormDB()
	var armorName string

	if lvl >= 6 && classArchetypeName == "Домен кузни" {
		for _, item := range equip {
			if item.ItemType == "Armor" {
				armorName = item.Name
			}
		}

		for _, armor := range armorData {
			if armor.Name == armorName {
				armor.ArmorClassCount += 1
				return armor
			}
		}

	}

	for _, item := range equip {
		if item.ItemType == "Armor" {
			armorName = item.Name
		}
	}

	for _, armor := range armorData {
		if armor.Name == armorName {
			return armor
		}
	}
	return classes.Armor{}
}

func GetWeaponInfo(equip []classes.Item) []classes.Weapon {
	var weaponData = classes.GetWeaponFormDB()
	var weaponNameList []string
	var weaponAnswer []classes.Weapon

	for _, item := range equip {
		if item.ItemType == "Weapon" {
			weaponNameList = append(weaponNameList, item.Name)
		}
	}

	for _, weaponInfo := range weaponData {
		for _, equipWeaponName := range weaponNameList {
			if weaponInfo.Name == equipWeaponName {
				weaponAnswer = append(weaponAnswer, weaponInfo)
			}
		}
	}
	return weaponAnswer
}

func getSpells(mod classes.AbilityModifier, lvl int) []spells.SpellsJSON {
	var proficiencyBonus = classes.GetProficiencyBonus(lvl)
	spellList = []spells.SpellsJSON{}
	var spellCastingInfo = classes.GetClassSpellBasicCharacteristic("Жрец", lvl, mod, proficiencyBonus)
	for i := 0; i < 5; i++ {
		var spellCount int
		switch i {
		case 0:
			spellCount = spellCastingInfo.ZeroLevelSpellsKnownCount
		case 1:
			spellCount = spellCastingInfo.OneLevelSpellsKnownCount
		case 2:
			spellCount = spellCastingInfo.TwoLevelSpellsKnownCount
		case 3:
			spellCount = spellCastingInfo.TreeLevelSpellsKnownCount
		case 4:
			spellCount = spellCastingInfo.FourLevelSpellsKnownCount
		}
		spellList = append(spellList,
			spells.GetRandomSpellForClass("Бард", i, spellCount)...)

	}
	return spellList
}
