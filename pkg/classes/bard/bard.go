package bard

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/general"
	"pregen/pkg/races"
	"pregen/pkg/spells"
)

var (
	bardSpellList []spells.SpellsJSON
	bardProf      *classes.Proficiencies
	skillCount    int

	musicalInstruments = []string{
		"Барабан", "Виола", "Волынка", "Лира", "Лютня",
		"Рожок", "Свирель", "Флейта", "Цимбалы", "Шалмей",
	}
)

func getBardHits(lvl int) classes.Hits {
	var hitCount int
	var D6 = general.D6

	for i := 1; i < lvl; i++ {
		if i == 1 {
			hitCount = D6.GetMaxRange()
		} else {
			hitCount += D6.RollDice()
		}
	}

	return classes.Hits{
		HitDice:  D6.GetDiceName(),
		HitCount: hitCount,
	}
}

func getBardProficiencies(raceInfo *races.Race, backgrInfo *backgrounds.Background) *classes.Proficiencies {
	var instrumentsArray []string
	var iter int
	for _, value := range musicalInstruments {
		instrumentsArray = append(instrumentsArray, value)
		iter++
		if iter == 3 {
			break
		}
	}
	skillCount = 3
	classSkills := classes.GetClassSkillsArray(
		raceInfo.RaceSkill,
		backgrInfo.BackgroundSkills,
		skillCount)

	return &classes.Proficiencies{
		Weapons:       []string{"Лёгкие доспехи"},
		Armor:         []string{"Простое оружие", "Длинные мечи", "Короткие мечи", "Рапиры", "Ручные арбалеты"},
		Tools:         instrumentsArray,
		SavingThrow:   []string{"Dexterity", "Charisma"},
		SkillsOfClass: classSkills,
	}
}

func getBardEquipment() []classes.Item {
	var equipment = []classes.Item{
		{
			Name:     "Кожаный доспех",
			ItemType: "Armor",
			Count:    1,
		},
		{
			Name:     "Кинжал",
			ItemType: "Weapon",
			Count:    1,
		},
	}

	variant, _ := random.IntRange(0, 1)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Набор дипломата",
				ItemType: "Tool",
				Count:    1,
			})
	case 1:
		equipment = append(equipment,
			classes.Item{
				Name:     "Набор артиста",
				ItemType: "Tool",
				Count:    1,
			})
	}

	variant, _ = random.IntRange(0, 2)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Рапира",
				ItemType: "Weapon",
				Count:    1,
			})
	case 1:
		equipment = append(equipment,
			classes.Item{
				Name:     "Длинный меч",
				ItemType: "Weapon",
				Count:    1,
			})
	case 2:
		var weaponsList []string
		for _, weapon := range classes.WeaponData {
			if weapon.WeaponType == "Простое рукопашное оружие" ||
				weapon.WeaponType == "Простое дальнобойное оружие" ||
				weapon.Name != "Рапира" && weapon.Name != "Длинный меч" {
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

	randNum, _ := random.IntRange(0, len(musicalInstruments))
	equipment = append(equipment,
		classes.Item{
			Name:     musicalInstruments[randNum],
			ItemType: "Tool",
			Count:    1,
		})

	return equipment
}

func getBardAbilitiesScore(raceInfo *races.Race, lvl int) classes.AbilityScore {
	var classAbilityPriority = []string{"Charisma", "Dexterity"}
	var abilitiesScore = general.GetClassAbilitiesScore(classAbilityPriority, raceInfo, lvl)
	return abilitiesScore
}

func getBardAbilityModifier(raceInfo *races.Race, lvl int) classes.AbilityModifier {
	return classes.GetModifiersForClass(getBardAbilitiesScore(raceInfo, lvl))
}

func getBardSavingThrows(raceInfo *races.Race, lvl int) *classes.SavingThrows {
	return classes.GetSaveThrowsForClass(getBardAbilityModifier(raceInfo, lvl), []string{"Dexterity", "Charisma"})
}

func getBardClassAbilities(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) []classes.ClassAbility {
	bardProf = getBardProficiencies(raceInfo, backgrInfo)
	var bardClassAbilities = []classes.ClassAbility{
		{
			Level: 1,
			Name:  "Использование заклинаний",
			Description: "Вы научились изменять ткань реальности в соответствии с вашими волей и музыкой. " +
				"Ваши заклинания являются частью вашего обширного репертуара; это магия, которой вы найдёте применение в любой ситуации.",
		},
		{
			Level: 1,
			Name:  "Вдохновение барда",
			Description: "Своими словами или музыкой вы можете вдохновлять других. " +
				"Для этого вы должны бонусным действием выбрать одно существо, отличное от вас, в пределах 60 футов, которое может вас слышать. " +
				"Это существо получает кость бардовского вдохновения — <strong>K6.</strong><br>" +
				"В течение следующих 10 минут это существо может один раз бросить эту кость и добавить результат к проверке характеристики, броску атаки или спасброску, который оно совершает. " +
				"Существо может принять решение о броске кости вдохновения уже после броска к20, но должно сделать это прежде, чем Мастер объявит результат броска. " +
				"Как только кость бардовского вдохновения брошена, она исчезает. Существо может иметь только одну такую кость одновременно.<br>" +
				"Вы можете использовать это умение количество раз, равное модификатору вашей Харизмы, но как минимум один раз. " +
				"Потраченные использования этого умения восстанавливаются после продолжительного отдыха.<br>" +
				"Ваша кость бардовского вдохновения изменяется с ростом вашего уровня в этом классе. Она становится к8 на 5-м уровне.",
		},
		{
			Level: 2,
			Name:  "Мастер на все руки",
			Description: "Вы можете добавлять половину бонуса мастерства, округлённую в меньшую сторону, " +
				"ко всем проверкам характеристики, куда этот бонус еще не включён.",
		},
		{
			Level: 2,
			Name:  "Песнь отдыха",
			Description: "Вы с помощью успокаивающей музыки или речей можете помочь своим раненым союзникам восстановить их силы во время короткого отдыха. " +
				"Если вы или любые союзные существа, способные слышать ваше исполнение, восстанавливаете хиты в конце короткого отдыха, " +
				"используя Кости Хитов, каждое потратившее Кость Хитов существо восстанавливает дополнительно <strong>1K6</strong> хитов.",
		},
		{
			Level: 3,
			Name:  "Коллегия бардов",
			Description: "Вы углубляетесь в традиции выбранной вами коллегии бардов. Все коллегии описаны ниже. " +
				"Этот выбор предоставляет вам умения на 3-м, 6-м и 14-м уровнях.",
		},
		{
			Level: 3,
			Name:  "Компетентность",
			Description: "Выберите 2 навыка из тех, которыми вы владеете (Навыки выбраны случайно). " +
				"Ваш бонус мастерства для этих навыков удваивается.<br>" +
				"На 10-м уровне вы можете выбрать еще 2 навыка и получить для них это преимущество.",
		},
		{
			Level:       5,
			Name:        "Источник вдохновения",
			Description: "Вы восстанавливаете истраченные вдохновения барда и после короткого и после продолжительного отдыха.",
		},
		{
			Level: 6,
			Name:  "Контрочарование",
			Description: "Вы получаете возможность использовать звуки или слова силы для разрушения воздействующих на разум эффектов. " +
				"Вы можете действием начать исполнение, которое продлится до конца вашего следующего хода. " +
				"В течение этого времени вы и все дружественные существа в пределах 30 футов от вас совершают " +
				"спасброски от запугивания и очарования с преимуществом. Чтобы получить это преимущество, " +
				"существа должны слышать вас. Исполнение заканчивается преждевременно, если вы оказываетесь недееспособны, " +
				"теряете способность говорить, или прекращаете исполнение добровольно (на это не требуется действие).",
		},
	}
	if lvl >= 3 {
		var collegiumsList = []string{
			"Коллегия доблести",
			"Коллегия знаний",
			"Коллегия мечей",
			"Коллегия очарования",
			//"Коллегия шёпотов",
			//"Коллегия красноречия",
			//"Коллегия созидания",
			//"Коллегия духов",
		}
		randNum, _ := random.IntRange(0, len(collegiumsList))
		collegiumName := collegiumsList[randNum]

		switch collegiumName {
		case "Коллегия доблести":
			bardProf = getBardProficiencies(raceInfo, backgrInfo) //новый екзмпляр
			bardProf.Armor = append(bardProf.Armor,
				[]string{"Средние доспехи", "Щит", "Воинское оружие"}...)

			bardClassAbilities = []classes.ClassAbility{
				{
					Level:       3,
					Name:        "Дополнительные Навыки",
					Description: "Присоединяясь к коллегии доблести, вы получаете владение средними доспехами, щитами и воинским оружием.",
				},
				{
					Level: 3,
					Name:  "Боевое Вдохновение",
					Description: "Вы узнаёте, как вдохновлять других в бою. " +
						"Существо, получившее от вас кость бардовского вдохновения, может бросить эту кость и добавить результат к своему броску урона оружием. " +
						"В качестве альтернативы, если существо атаковано, оно может реакцией совершить бросок кости вдохновения и добавить результат к своему КД от этой атаки. " +
						"Оно может сделать это после броска атаки, но до того, как узнает, попали ли по нему.",
				},
				{
					Level:       6,
					Name:        "Дополнительная Атака",
					Description: "Если вы в свой ход совершаете действие Атака, вы можете совершить две атаки вместо одной.",
				},
			}
		case "Коллегия знаний":
			skillCount = 3 + 3
			bardProf = getBardProficiencies(raceInfo, backgrInfo)

			var maxSpellLevel int
			if lvl == 6 {
				maxSpellLevel = 3
			} else if lvl > 6 {
				maxSpellLevel = 4
			}

			classArray := []string{"Изобретатель", "Жрец", "Друид", "Паладин", "Следопыт", "Чародей", "Колдун", "Волшебник"}
			randNum, _ = random.IntRange(0, len(classArray))

			tempSpellList := spells.GetAllSpellForClass(classArray[randNum], maxSpellLevel)

			if lvl >= 6 {
				for i := 0; i < 3; i++ {
					randNum, _ = random.IntRange(0, len(tempSpellList))
					bardSpellList = append(bardSpellList, tempSpellList[randNum])
				}
			}

			bardClassAbilities = []classes.ClassAbility{
				{
					Level:       3,
					Name:        "Дополнительные Навыки",
					Description: "Если вы присоединяетесь к коллегии знаний, вы овладеваете тремя навыками на ваш выбор.",
				},
				{
					Level: 3,
					Name:  "Острое словцо",
					Description: "Вы узнаёте, как использовать собственное остроумие, чтобы отвлечь, смутить или по-другому подорвать способности и уверенность противников. " +
						"Если существо, которое вы можете видеть, в пределах 60 футов от вас совершает бросок атаки, урона или проверку характеристики, " +
						"вы можете реакцией потратить одну из ваших костей бардовского вдохновения, и вычесть результат броска этой кости из броска этого существа. " +
						"Вы можете принять решение об использовании этого умения после броска существа, но до того момента, когда Мастер объявит результат броска или проверки. " +
						"Существо не подвержено этому умению, если не может слышать вас, или обладает иммунитетом к очарованию.",
				},
				{
					Level: 6,
					Name:  "Дополнительные Тайны Магии",
					Description: "Вы можете выучить 2 заклинания из доступных любому классу на свой выбор. " +
						"Их уровень не должен превышать уровня заклинаний, которые вы можете использовать на этом уровне, как показано в таблице Барда. " +
						"Они также могут быть заговорами. Выбранные заклинания теперь считаются для вас заклинаниями барда, " +
						"но они не учитываются в общем количестве известных вам заклинаний барда.",
				},
			}

		case "Коллегия мечей":
			battleStyle := []string{"<strong>Дуэлянт.</strong> Пока вы держите рукопашное оружие в одной руке и не используете другого оружия, вы получаете бонус +2 к броскам урона этим оружием.",
				"<strong>Сражение двумя оружиями.</strong> Если вы сражаетесь двумя оружиями, вы можете добавить модификатор характеристики к урону от второй атаки."}
			randBattleStyleNum, _ := random.IntRange(0, len(battleStyle))

			bardProf = getBardProficiencies(raceInfo, backgrInfo) //новый екзмпляр
			bardProf.Armor = append(bardProf.Armor,
				[]string{"Средние доспехи", "Скимитар"}...)

			daggerLine := []string{
				"<strong>Оборонительный росчерк.</strong> Вы можете потратить одно свое «Вдохновение барда», чтобы оружие нанесло дополнительный урон по цели. " +
					"Урон равен числу, выпавшему на кости бардовского вдохновения. Вы также добавляете это число к своему КД до начала своего следующего хода.",
				"<strong>Режущий росчерк.</strong> Вы можете потратить одно свое «Вдохновение барда», чтобы оружие нанесло дополнительный урон по цели и по любому другому существу по вашему выбору, " +
					"которое вы можете видеть в 5 футах от себя. Урон равен числу, выпавшему на кости бардовского вдохновения.",
				"<strong>Мобильный росчерк.</strong> Вы можете потратить одно свое «Вдохновение барда», чтобы оружие нанесло дополнительный урон по цели. " +
					"Урон равен числу, выпавшему на кости бардовского вдохновения. Вы также можете оттолкнуть цель на 5 футов от себя, плюс количество футов, выпавшее на кости бардовского вдохновения. После этого вы незамедлительно можете реакцией переместиться на расстояние не большее своей скорости в незанятое место в 5 футах от цели.",
			}
			randDaggerLineNum, _ := random.IntRange(0, len(battleStyle))
			bardClassAbilities = []classes.ClassAbility{
				{
					Level: 3,
					Name:  "Боевой стиль",
					Description: "Вы выбираете боевой стиль, соответствующий вашей специализации. <br>" +
						"Боевой стиль:" + battleStyle[randBattleStyleNum],
				},
				{
					Level: 3,
					Name:  "Дополнительные владения",
					Description: "Когда вы вступаете в Коллегию Мечей, вы получаете владение средними доспехами и скимитарами. " +
						"Если вы владеете простым или воинским оружием ближнего боя, то можете использовать его в качестве фокусировки для ваших заклинаний барда.",
				},
				{
					Level: 3,
					Name:  "Росчерк клинка",
					Description: "Вы учитесь исполнять впечатляющие демонстрации боевого мастерства и проворства. " +
						"Когда вы совершаете действие Атака в свой ход, ваша скорость ходьбы увеличивается на 10 футов до конца хода, " +
						"а если атака оружием, которую вы совершаете частью этого действия, попадает по существу, " +
						"вы можете использовать " + daggerLine[randDaggerLineNum] + "Вы можете использовать только один вариант «Росчерка клинка» за ход.",
				},
				{
					Level:       6,
					Name:        "Дополнительная атака",
					Description: "Если вы в свой ход совершаете действие Атака, вы можете совершить две атаки вместо одной.",
				},
			}
		case "Коллегия очарования":
			bardClassAbilities = []classes.ClassAbility{
				{
					Level:       3,
					Name:        "Завораживающее выступление",
					Description: "Вы можете наполнить свое выступление притягательными чарами фей. Если вы выступаете как минимум в течение 1 минуты, то можете попытаться очаровать слушателей своим пением, чтением стихотворения или танцем. В конце выступления выберите такое число Гуманоидов в пределах 60 футов от вас, которые смотрели и слушали все выступление, количество которых не превышает ваш модификатор Харизмы (минимум 1).\n\nКаждая цель должна преуспеть в спасброске Мудрости против Сл спасброска от ваших заклинаний, иначе будет очарована вами. Очарованные таким образом существа боготворят вас и говорят о вас крайне восторженно с любым заговорившим с ними. Они мешают вашим недругам, хотя при этом избегают насилия, если только к этому моменту уже и так не были настроены драться за вас.\n\nЭффект на цели заканчивается через 1 час, если цель получает любой урон, если вы атакуете её или если она замечает, как вы атакуете или наносите урон кому-то из её союзников. Если цель преуспевает в спасброске, то она не замечает никаких признаков того, что вы пытались её очаровать.\n\nИспользовав это умение, вы не можете использовать его повторно до окончания короткого или продолжительного отдыха.",
				},
				{
					Level:       3,
					Name:        "Мантия вдохновения",
					Description: "Вступая в коллегию очарования, вы получаете возможность слагать песни, сплетённые с чарами фей, которые наполняют ваших союзников отвагой и проворством.\n\nВы можете бонусным действием потратить одно использование своего «Вдохновения барда», чтобы придать себе поразительный внеземной облик. Когда вы делаете это, выберите несколько союзников, которых вы можете видеть, и которые могут видеть вас в пределах 60 футов от вас. Количество выбранных целей не должно превышать ваш модификатор Харизмы (минимум 1).\n\nКаждая выбранная цель получает 5 временных хитов. Когда цель получает временные хиты от этого умения, она также немедленно может реакцией переместиться на расстояние, не превышающее своей скорости, не вызывая провоцированных атак.\n\nКоличество временных хитов увеличивается, когда вы достигаете определенных уровней в этом классе: до 8 на 5-м уровне, 11 на 10-м уровне и 14 на 15-м уровне.",
				},
				{
					Level:       6,
					Name:        "Мантия величия",
					Description: "Вы получаете возможность укутаться в чары фей, заставляющие других хотеть служить вам.\n\nБонусным действием вы можете сотворить заклинание <strong>приказ [command]</strong> без траты ячейки заклинания и принимаете облик неземной красоты на 1 минуту или до тех пор, пока не завершится ваша концентрация (как если бы вы концентрировались на заклинании). В течение этого времени вы можете накладывать заклинание <strong>приказ [command]</strong> бонусным действием в каждый свой ход, не тратя ячейку заклинания.\n\nЛюбое существо, очарованное вами, автоматически проваливает спасбросок против заклинания <strong>приказ [command]</strong>, которое вы накладываете с помощью этого умения.\nИспользовав это умение, вы не можете использовать его повторно до окончания продолжительного отдыха.",
				},
			}
			//case "Коллегия шёпотов":
			//	bardClassAbilities = []classes.ClassAbility{
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//	}
			//case "Коллегия красноречия":
			//	bardClassAbilities = []classes.ClassAbility{
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//	}
			//case "Коллегия созидания":
			//	bardClassAbilities = []classes.ClassAbility{
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//	}
			//case "Коллегия духов":
			//	bardClassAbilities = []classes.ClassAbility{
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//		{
			//			Level:       3,
			//			Name:        "",
			//			Description: "",
			//		},
			//	}
		}

	}

	return bardClassAbilities
}

func getBardClassAbilitiesWithLevel(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) []classes.ClassAbility {
	var abilitiesList = getBardClassAbilities(raceInfo, backgrInfo, lvl)
	var abilitiesAnswer []classes.ClassAbility

	for _, ability := range abilitiesList {
		if lvl >= ability.Level {
			abilitiesAnswer = append(abilitiesAnswer, ability)
		}
	}
	return abilitiesAnswer
}

func getBardSpells(raceInfo *races.Race, lvl int) []spells.SpellsJSON {
	var bardSpellCastingInfo = classes.GetClassSpellBasicCharacteristic("Бард", lvl, getBardAbilityModifier(raceInfo, lvl))
	for i := 0; i < 5; i++ {
		var spellCount int
		switch i {
		case 0:
			spellCount = bardSpellCastingInfo.ZeroLevelSpellsKnownCount
		case 1:
			spellCount = bardSpellCastingInfo.OneLevelSpellsKnownCount
		case 2:
			spellCount = bardSpellCastingInfo.TwoLevelSpellsKnownCount
		case 3:
			spellCount = bardSpellCastingInfo.TreeLevelSpellsKnownCount
		case 4:
			spellCount = bardSpellCastingInfo.FourLevelSpellsKnownCount
		}
		bardSpellList = append(bardSpellList,
			spells.GetRandomSpellForClass("Бард", i, spellCount)...)
	}
	return bardSpellList
}

func GetBardClass(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) *classes.Class {

	return &classes.Class{
		ClassName:       "Bard",
		ClassNameRU:     "Бард",
		ClassAbilities:  getBardClassAbilitiesWithLevel(raceInfo, backgrInfo, lvl),
		AbilityScore:    getBardAbilitiesScore(raceInfo, lvl),
		AbilityModifier: getBardAbilityModifier(raceInfo, lvl),
		SavingThrows:    getBardSavingThrows(raceInfo, lvl),
		Inspiration:     false,
		Proficiencies:   *bardProf, //need to fix
		Hits:            getBardHits(lvl),
		Initiative:      "",
		Caster:          true,
		SpellCasting:    classes.GetClassSpellBasicCharacteristic("Бард", lvl, getBardAbilityModifier(raceInfo, lvl)),
		SpellsList:      getBardSpells(raceInfo, lvl),
		Equipment:       getBardEquipment(),
	}
}
