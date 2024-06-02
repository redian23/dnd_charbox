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
	skillCount    = 3

	musicalInstruments = []string{
		"Барабан", "Виола", "Волынка", "Лира", "Лютня",
		"Рожок", "Свирель", "Флейта", "Цимбалы", "Шалмей",
	}
)

func GetClass(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int, classArchetypeName string) *classes.Class {

	var statsPriority = []string{"Charisma", "Dexterity"}
	var abilitiesScore = general.GetClassAbilitiesScore(statsPriority, raceInfo, lvl)
	var modifier = classes.GetModifiersForClass(abilitiesScore)
	var savingThrows = classes.GetSaveThrowsForClass(modifier, statsPriority)
	var spellCastingInfo = classes.GetClassSpellBasicCharacteristic("Бард", lvl, modifier)
	var equip = getBardEquipment()

	return &classes.Class{
		ClassName:          "Bard",
		ClassNameRU:        "Бард",
		ClassArchetypeName: classArchetypeName,
		ClassAbilities:     getBardClassAbilitiesWithLevel(raceInfo, backgrInfo, lvl, classArchetypeName),
		AbilityScore:       abilitiesScore,
		AbilityModifier:    modifier,
		SavingThrows:       savingThrows,
		Inspiration:        false,
		Proficiencies:      *bardProf, //need to fix
		ProficiencyBonus:   classes.ProficiencyBonus,
		Hits:               getBardHits(modifier, lvl),
		Caster:             true,
		SpellCasting:       spellCastingInfo,
		SpellsList:         getBardSpells(modifier, lvl),
		Equipment:          equip,
		ArmorInfo:          GetArmorInfo(equip),
		WeaponInfo:         GetWeaponInfo(equip),
	}
}

func getBardHits(modifier classes.AbilityModifier, lvl int) classes.Hits {
	var hitCount int

	for i := 1; i < lvl; i++ {
		if i == 1 {
			hitCount = general.D6.GetMaxRange() + modifier.BodyDifficulty
		} else {
			hitCount += general.D6.RollDice() + modifier.BodyDifficulty
		}
	}

	return classes.Hits{
		HitDice:  general.D6.GetDiceName(),
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

	var availableSkillList = []string{"Acrobatics", "Animal Handling", "Arcana", "Athletics",
		"Deception", "History", "Insight", "Intimidation", "Investigation",
		"Medicine", "Nature", "Perception", "Performance", "Persuasion",
		"Religion", "Sleight Of Hand", "Stealth", "Survival"}

	classSkills := classes.GetClassSkillsArray(
		raceInfo.RaceSkill,
		backgrInfo.BackgroundSkills,
		availableSkillList,
		skillCount,
	)

	return &classes.Proficiencies{
		Armor:         []string{"Лёгкие доспехи"},
		Weapons:       []string{"Простое оружие", "Длинные мечи", "Короткие мечи", "Рапиры", "Ручные арбалеты"},
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

func getBardClassAbilities(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int, classArchetypeName string) []classes.ClassAbility {
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
		switch classArchetypeName {
		case "Коллегия доблести":
			bardProf = getBardProficiencies(raceInfo, backgrInfo) //новый екзмпляр
			bardProf.Armor = append(bardProf.Armor,
				[]string{"Средние доспехи", "Щит", "Воинское оружие"}...)

			bardClassAbilities = append(bardClassAbilities, []classes.ClassAbility{
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
			}...)
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
			randNum, _ := random.IntRange(0, len(classArray))

			tempSpellList := spells.GetAllSpellForClass(classArray[randNum], maxSpellLevel)

			if lvl >= 6 {
				for i := 0; i < 3; i++ {
					randNum, _ = random.IntRange(0, len(tempSpellList))
					bardSpellList = append(bardSpellList, tempSpellList[randNum])
				}
			}

			bardClassAbilities = append(bardClassAbilities, []classes.ClassAbility{
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
			}...)
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
			bardClassAbilities = append(bardClassAbilities, []classes.ClassAbility{
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
			}...)
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
		case "Коллегия шёпотов":
			bardClassAbilities = []classes.ClassAbility{
				{
					Level: 3,
					Name:  "Психические клинки",
					Description: "При вступлении в коллегию шёпотов вы получаете возможность сделать свою атаку оружием магически токсичной для разума существа. " +
						"Когда вы попадаете по существу атакой оружием, вы можете потратить одно «Вдохновение барда» и нанести цели дополнительные 2к6 урона психической энергией. " +
						"Вы можете использовать это умение один раз за раунд в свой ход.<br>" +
						"Урон психической энергией увеличивается, когда вы получаете определенный уровень в этом классе, " +
						"увеличиваясь до 3к6 на 5-м уровне, 5к6 на 10-м уровне.",
				},
				{
					Level: 3,
					Name:  "Слова ужаса",
					Description: "Вы учитесь наполнять кажущиеся безобидными слова коварной магией, которая вызывает ужас. " +
						"Если вы разговариваете с Гуманоидом наедине, в течение по крайней мере 1 минуты, то можете попытаться посеять семя паранойи в его разуме.<br>" +
						"В конце беседы цель должна преуспеть в спасброске Мудрости против вашей Сл спасброска от ваших заклинаний, " +
						"иначе будет напугана вами или другим существом по вашему выбору. Цель напугана на 1 час таким образом, " +
						"пока не будет атакована или не получит урон, или пока она не заметит, что её союзники атакованы или получили урон. " +
						"Если цель преуспевает в спасброске, то она не замечает никаких признаков того, что вы пытались ее испугать.<br>" +
						"Использовав это умение, вы не можете использовать его повторно до окончания короткого или продолжительного отдыха.",
				},
				{
					Level: 6,
					Name:  "Мантия шёпотов",
					Description: "Вы получаете возможность принимать личность Гуманоида. " +
						"Когда Гуманоид умирает в пределах 30 футов от вас, вы можете реакцией магическим образом поймать его тень. " +
						"Вы удерживаете эту тень, пока не используете её, или до окончания короткого или продолжительного отдыха.<br>" +
						"Действием вы можете использовать тень. " +
						"Когда вы это делаете, она исчезает, магически превращаясь в маскировку, которая появляется на вас. " +
						"Теперь вы выглядите как этот мёртвый Гуманоид, но кажетесь живым и здоровым. " +
						"Маскировка длится 1 час или пока вы не закончите её бонусным действием. " +
						"Находясь под этой маскировкой, вы получаете доступ ко всей информации, " +
						"которой это существо могло бы свободно поделиться со случайным знакомым. " +
						"Эта информация включает в себя общие данные о его биографии и личной жизни, но не его секреты. " +
						"Этой информации достаточно для того, чтобы вы смогли выдать себя за эту личность, используя его воспоминания. " +
						"Другие существа могут обнаружить вашу истинную сущность, в состязании их Мудрости (Проницательность) " +
						"против вашей Харизмы (Обман), при этом вы получаете бонус +5 к своей проверке.<br>" +
						"Когда вы ловите тень этим умением, вы не можете поймать другую тень, до окончания короткого или продолжительного отдыха.",
				},
			}
		case "Коллегия красноречия":
			bardClassAbilities = []classes.ClassAbility{
				{
					Level: 3,
					Name:  "Златоуст",
					Description: "Вы мастер говорить нужные вещи в нужное время. " +
						"Когда вы совершаете проверку харизмы (убеждение или обман), вы можете при выпадении на к20 результата «1–9» считать, что выпало «10».",
				},
				{
					Level: 3,
					Name:  "Тревожащие слова",
					Description: "Вместе с магией вы можете сплетать слова, которые выбивают существо из колеи и заставляют его сомневаться в себе. " +
						"Бонусным действием вы можете потратить одно использование «вдохновения барда» и выбрать существо, которое вы можете видеть в пределах 60 футов от вас. " +
						"Совершите бросок кости бардовского вдохновения. " +
						"Существо должно вычесть выпавший результат из следующего спасброска, который оно совершит до начала вашего следующего хода.",
				},
				{
					Level: 6,
					Name:  "Неисчерпаемое вдохновение",
					Description: "Ваши вдохновляющие слова настолько убедительны, что другие чувствуют стремление к успеху. " +
						"Когда существо добавляет вашу кость бардовского вдохновения к проверке характеристики, " +
						"броску атаки или спасброску, но бросок оказывается провален, то существо не теряет кость бардовского вдохновения.",
				},
				{
					Level: 6,
					Name:  "Неисчерпаемое вдохновение",
					Description: "Вы приобретаете способность сделать свою речь понятной любому существу. " +
						"Действием выберите одно или несколько существ в радиусе 60 футов от вас, количество которых не превышает ваш модификатор харизмы (минимум одно существо). " +
						"Выбранные существа могут в течение 1 часа магическим образом понимать вас независимо от языка, на котором вы говорите." +
						"Как только вы воспользуетесь этим умением, вы не можете использовать его снова, " +
						"пока не закончите продолжительный отдых или не израсходуете ячейку заклинания любого уровня, чтобы использовать его снова.",
				},
			}
		case "Коллегия созидания":
			bardClassAbilities = []classes.ClassAbility{
				{
					Level:       3,
					Name:        "Частица потенциала",
					Description: "",
				},
				{
					Level: 3,
					Name:  "Созидательное выступление",
					Description: "Действием вы можете направить магию Песни Созидания для создания одного немагического предмета по вашему выбору в свободном пространстве в пределах 10 футов от вас. " +
						"Предмет должен находиться на поверхности или в жидкости, которая может поддерживать его на плаву. Стоимость предмета в зм не может быть более чем в 20 раз больше вашего уровня барда, а размер должен быть Средним или меньше. " +
						"Предмет слегка мерцает, а коснувшись его, можно услышать слабую музыку. Созданный предмет исчезает через количество часов, равное вашему бонусу мастерства. " +
						"Примеры предметов, которые вы можете создать, смотрите в главе «Снаряжение» «Книги игрока».<br>" +
						"После того, как вы создадите предмет этим умением, вы не сможете сделать это снова, пока не закончите продолжительный отдых или не израсходуете ячейку заклинания 2-го уровня или выше, чтобы применить это умение снова. " +
						"У вас может быть только один предмет, созданный этим умением, одновременно: если вы используете это действие снова, то предыдущий предмет немедленно исчезает. " +
						"Максимально возможный размер предмета, который вы можете создать с помощью этого умения, увеличивается на одну категорию, когда вы достигаете 6-го уровня (Большой) и 14-го уровня (Огромный).",
				},
				{
					Level: 6,
					Name:  "Оживляющее выступление",
					Description: "Действием вы можете нацелиться на немагический предмет не больше Большого размера, который вы можете видеть в пределах 30 футов от вас, который никто не держит и не несёт, и оживить его. " +
						"Пробуждённый предмет использует блок статистики танцующий предмет [dancing Item], зависящий от вашего бонуса мастерства (БМ). Этот предмет дружелюбен к вам и вашим спутникам и подчиняется вашим командам. " +
						"Он оживает либо на 1 час, либо до тех пор, пока его хиты не упадут до 0, либо пока вы не умрете.<br>" +
						"В бою предмет получает вашу инициативу и ходит сразу же после вас. Он может перемещаться и использовать свою реакцию самостоятельно, но единственное действие, " +
						"которое он может совершать в свой ход — это Уклонение, если только вы не используете своё бонусное действие, чтобы приказать ему сделать что-то другое. " +
						"Это действие может быть одним из действий из его блока статистики или любым другим действием. Если вы недееспособны, то предмет может использовать любое действие по своему выбору, а не только Уклонение.<br>" +
						"Когда вы используете умение «Вдохновение барда», частью того же бонусного действия вы можете отдавать приказы этому оживленному предмету.<br>" +
						"После того, как вы оживляете предмет с помощью этого умения, вы не можете сделать это снова, пока не закончите продолжительный отдых или не израсходуете ячейку заклинания 3-го уровня или выше, " +
						"чтобы снова воспользоваться этим умением. Вы можете одновременно иметь только один предмет, пробуждённый этим умением. Если вы используете это действие и уже имеете танцующий предмет от этого умения, то первый из них мгновенно становится неодушевленным.",
				},
			}
		case "Коллегия духов":
			bardClassAbilities = []classes.ClassAbility{
				{
					Level:       3,
					Name:        "Направляющий шёпот",
					Description: "Вы можете обратиться к духам, чтобы они направили вас и других. Вы изучаете заговор указание [guidance], который не учитывается при подсчёте известных вам заговоров барда. Для вас он имеет дистанцию 60 футов.",
				},
				{
					Level: 3,
					Name:  "Спиритическая фокусировка",
					Description: "Вы используете инструменты, которые помогают вам обращаться к духам, будь то исторические личности или вымышленные архетипы. " +
						"Вы можете использовать следующие предметы в качестве фокусировки для заклинаний барда: свечу, хрустальный шар, череп, спиритическую доску или колоду Тарокка.<br>" +
						"Начиная с 6-го уровня, когда вы с помощью спиритической фокусировки накладываете заклинание барда, " +
						"которое наносит урон или восстанавливает хиты, бросьте к6, и вы получите бонус к одному броску урона или исцеления этого заклинания, равный выпавшему результату.",
				},
				{
					Level: 3,
					Name:  "Истории с того света",
					Description: "Вы обращаетесь к духам, которые рассказывают свои истории через вас. " +
						"Пока вы держите свою спиритическую фокусировку, вы можете бонусным действием потратить одно использование вашего «Вдохновения барда» " +
						"и совершить бросок по таблице <a href=\"https://dnd.su/class/88-bard\">«Истории духов»</a>, используя кость «Вдохновения барда», чтобы узнать историю, которую духи поведают вам. " +
						"Вы сохраняете историю в памяти до тех пор, пока не используете её эффект или не закончите короткий или продолжительный отдых. <br>" +
						"Действием вы можете выбрать одно существо, которое вы видите в пределах 30 футов от себя (вы можете выбрать себя), которое станет целью эффекта истории. " +
						"Сделав это, вы не можете использовать этот же эффект, пока не совершите повторный бросок. <br>" +
						"Одновременно вы можете держать в голове только одну из этих историй, повторный бросок по таблице «Истории духов» " +
						"немедленно прекращает действие эффекта предыдущего броска. ",
				},
				{
					Level: 6,
					Name:  "Спиритический сеанс",
					Description: "Духи дарят вам сверхъестественные озарения. " +
						"Вы можете провести часовой ритуал обращения к духам (который может быть выполнен во время короткого или продолжительного отдыха), " +
						"используя свою спиритическую фокусировку. " +
						"Вы можете провести ритуал вместе с несколькими согласными существами, количество которых равно вашему бонусу мастерства (включая вас). " +
						"В конце ритуала вы временно изучаете одно заклинание по вашему выбору из списка заклинаний любого класса. <br>" +
						"Выбранное заклинание должно иметь уровень не выше количества существ, проводивших ритуал, должно быть того уровня, который вы способны накладывать, " +
						"и должно принадлежать к школе Прорицания или Некромантии. Выбранное заклинание считается для вас заклинанием барда, " +
						"но не учитываются при подсчёте известных вам заклинаний барда. <br>После проведения ритуала вы не можете провести его снова, " +
						"пока не начнёте продолжительный отдых, и вы помните выбранное заклинание, пока не начнёте продолжительный отдых.",
				},
			}
		}
	}

	return bardClassAbilities
}

func getBardClassAbilitiesWithLevel(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int, classArchetypeName string) []classes.ClassAbility {
	var abilitiesList = getBardClassAbilities(raceInfo, backgrInfo, lvl, classArchetypeName)
	var abilitiesAnswer []classes.ClassAbility

	for _, ability := range abilitiesList {
		if lvl >= ability.Level {
			abilitiesAnswer = append(abilitiesAnswer, ability)
		}
	}
	return abilitiesAnswer
}

func GetArmorInfo(equip []classes.Item) classes.Armor {
	var armorData = classes.GetArmorFormDB()
	var armorName string

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

func getBardSpells(mod classes.AbilityModifier, lvl int) []spells.SpellsJSON {

	bardSpellList = []spells.SpellsJSON{}
	var bardSpellCastingInfo = classes.GetClassSpellBasicCharacteristic("Бард", lvl, mod)
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
