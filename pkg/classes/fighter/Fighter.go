package fighter

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/backgrounds"
	"pregen/pkg/classes"
	"pregen/pkg/general"
	"pregen/pkg/spells"
)

var (
	fighterProf      *classes.Proficiencies
	skillCount       int
	fighterSpellList []spells.SpellsJSON
)

var classAbilityPriority []string

func getHits(raceInfo *races.Race, lvl int) classes.Hits {
	var hitCount int
	var modifier = getAbilityModifier(raceInfo, lvl)
	for i := 1; i < lvl; i++ {
		if i == 1 {
			hitCount = general.D10.GetMaxRange() + modifier.BodyDifficulty
		} else {
			hitCount += general.D10.RollDice() + modifier.BodyDifficulty
		}
	}

	return classes.Hits{
		HitDice:  general.D10.GetDiceName(),
		HitCount: hitCount,
	}
}

func getProficiencies(raceInfo *races.Race, backgrInfo *backgrounds.Background) *classes.Proficiencies {
	var availableSkillList = []string{"Acrobatics", "Animal Handling", "Athletics", "Perception",
		"Survival", "Intimidation", "History", "Insight"}

	skillCount = 2
	classSkills := classes.GetClassSkillsArray(
		raceInfo.RaceSkill,
		backgrInfo.BackgroundSkills,
		availableSkillList,
		skillCount)

	return &classes.Proficiencies{
		Weapons:       []string{"Все доспехи", "Щиты"},
		Armor:         []string{"Простое оружие", "Воинское оружие"},
		Tools:         []string{},
		SavingThrow:   []string{"Strength", "BodyDifficulty"},
		SkillsOfClass: classSkills,
	}
}

func getBardEquipment() []classes.Item {
	var equipment []classes.Item

	variant, _ := random.IntRange(0, 1)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Кольчуга",
				ItemType: "Armor",
				Count:    1,
			})
	case 1:
		equipment = append(equipment,
			[]classes.Item{
				{
					Name:     "Кожаный доспех",
					ItemType: "Weapon",
					Count:    1,
				},
				{
					Name:     "Длинный лук",
					ItemType: "Weapon",
					Count:    1,
				},
				{
					Name:     "Колчан и 20 стрел",
					ItemType: "Weapon",
					Count:    1,
				},
			}...,
		)
	}

	variant, _ = random.IntRange(0, 1)
	switch variant {
	case 0:
		equipment = append(equipment,
			classes.Item{
				Name:     "Набор исследователя подземелий",
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

	variant, _ = random.IntRange(0, 1)
	switch variant {
	case 0:
		var weaponsList []string
		for _, weapon := range classes.WeaponData {
			if weapon.WeaponType == "Воинское рукопашное оружие" {
				weaponsList = append(weaponsList, weapon.Name)
			}
		}
		randNum, _ := random.IntRange(0, len(weaponsList))
		equipment = append(equipment,
			[]classes.Item{
				{
					Name:     weaponsList[randNum],
					ItemType: "Weapon",
					Count:    1,
				},
				{
					Name:     "Щит",
					ItemType: "Armor",
					Count:    1,
				},
			}...)
	case 1:
		var weaponsList []string
		for _, weapon := range classes.WeaponData {
			if weapon.WeaponType == "Простое рукопашное оружие" ||
				weapon.WeaponType == "Простое дальнобойное оружие" ||
				weapon.Name != "Рапира" && weapon.Name != "Длинный меч" {
				weaponsList = append(weaponsList, weapon.Name)
			}
		}

		var randNum, randNum2 int
		for true {
			randNum, _ = random.IntRange(0, len(weaponsList))
			randNum2, _ = random.IntRange(0, len(weaponsList))
			if randNum != randNum2 {
				break
			}
		}

		equipment = append(equipment,
			[]classes.Item{
				{
					Name:     weaponsList[randNum],
					ItemType: "Weapon",
					Count:    1,
				},
				{
					Name:     weaponsList[randNum2],
					ItemType: "Weapon",
					Count:    1,
				},
			}...)
	}
	variant, _ = random.IntRange(0, 1)
	switch variant {
	case 0:
		equipment = append(equipment,
			[]classes.Item{
				{
					Name:     "Лёгкий арбалет",
					ItemType: "Weapon",
					Count:    1,
				},
				{
					Name:     "Колчан и 20 болтов",
					ItemType: "Weapon",
					Count:    1,
				},
			}...)
	case 1:
		equipment = append(equipment,
			classes.Item{
				Name:     "Ручной топор",
				ItemType: "Weapon",
				Count:    2,
			})
	}

	return equipment
}

func getAbilitiesScore(raceInfo *races.Race, lvl int) classes.AbilityScore {
	var abilitiesScore = general.GetClassAbilitiesScore(classAbilityPriority, raceInfo, lvl)
	return abilitiesScore
}

func getAbilityModifier(raceInfo *races.Race, lvl int) classes.AbilityModifier {
	return classes.GetModifiersForClass(getAbilitiesScore(raceInfo, lvl))
}

func getSavingThrows(raceInfo *races.Race, lvl int) *classes.SavingThrows {
	return classes.GetSaveThrowsForClass(getAbilityModifier(raceInfo, lvl), classAbilityPriority)
}

func getClassAbilities(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) []classes.ClassAbility {
	fighterProf = getProficiencies(raceInfo, backgrInfo)

	var basicBattleStylesMap = map[string]string{
		"Дуэлянт": "Пока вы держите рукопашное оружие в одной руке и не используете другого оружия, вы получаете бонус +2 к броскам урона этим оружием.",

		"Защита": "Если существо, которое вы видите, атакует не вас, а другое существо, находящееся в пределах 5 футов от вас, вы можете реакцией создать помеху его броску атаки." +
			"Для этого вы должны использовать щит.",

		"Оборона": "Пока вы носите доспехи, вы получаете бонус +1 к КД.",

		"Сражение большим оружием": "Если у вас выпало «1» или «2» на кости урона при атаке, которую вы совершали рукопашным оружием, удерживая его двумя руками, " +
			"то вы можете перебросить эту кость, и должны использовать новый результат, даже если снова выпало «1» или «2». " +
			"Чтобы воспользоваться этим преимуществом, ваше оружие должно иметь свойство «двуручное» или «универсальное».",

		"Сражение двумя оружиями": "Если вы сражаетесь двумя оружиями, вы можете добавить модификатор характеристики к урону от второй атаки.",

		"Стрельба": "Вы получаете бонус +2 к броску атаки, когда атакуете дальнобойным оружием.",
	}
	var basicBattleStylesArray = []string{
		"Дуэлянт", "Защита", "Оборона", "Сражение большим оружием", "Сражение двумя оружиями", "Стрельба",
	}

	randBattleStyle, _ := random.IntRange(0, len(basicBattleStylesArray))
	//var optionalBattleStylesMap = map[string]string{
	//	"Перехват": "Когда существо, которое вы можете видеть, попадает атакой по цели в пределах 5 футов от вас (но не по вам), " +
	//		"вы можете реакцией уменьшить урон, получаемый целью, на количество, равное 1к10 + ваш бонус мастерства (вплоть до 0 урона). " +
	//		"Чтобы использовать эту реакцию, вы должны держать в руках щит либо простое или воинское оружие.",
	//
	//	"Превосходная техника": "Вы изучаете один боевой приём  на ваш выбор из списка архетипа воина «Мастер боевых искусств». " +
	//		"Если приём требует, чтобы цель совершила спасбросок для сопротивления эффекту приёма, " +
	//		"то Сл спасброска равна 8 + ваш бонус мастерства + ваш модификатор Силы или Ловкости (по вашему выбору)." +
	//		"Вы получаете одну кость превосходства: к6. Эта кость добавляется к другим костям превосходства, полученным вами из других источников. " +
	//		"Эта кость используется для ваших приёмов и тратится при использовании. " +
	//		"Вы восстанавливаете все свои израсходованные кости превосходства после окончания короткого или продолжительного отдыха.",
	//
	//	"Сражение вслепую": "Вы получаете слепое зрение в пределах 10 футов. " +
	//		"В пределах этой дистанции вы можете видеть всё, что не находится за полным укрытием, даже если вы ослеплены или находитесь в темноте. " +
	//		"Более того, вы можете увидеть невидимое существо в пределах этой дистанции, если только оно не преуспело в попытке спрятаться от вас.",
	//
	//	"Сражение голыми руками": "Ваши безоружные удары могут наносить дробящий урон, равный 1к6 + ваш модификатор Силы, при попадании. " +
	//		"Если вы не держите в руках ни оружие, ни щит, то кость урона увеличивается с к6 до к8. " +
	//		"В начале каждого своего хода вы можете нанести 1к4 дробящего урона одному существу, которое вы держите в захвате.",
	//
	//	"Сражение метательным оружием": "Вы можете вытащить оружие со свойством «метательное» частью атаки, которую вы совершаете этим оружием. " +
	//		"Кроме того, когда вы попадаете по существу дальнобойной атакой, используя метательное оружие, вы получаете бонус +2 к броску урона.",
	//}

	var fighterClassAbilities = []classes.ClassAbility{
		{
			Level: 1,
			Name:  "Боевой стиль",
			Description: "Вы выбираете боевой стиль, соответствующий вашей специализации. " +
				"<strong>" + basicBattleStylesArray[randBattleStyle] + "</strong > : " + basicBattleStylesMap[basicBattleStylesArray[randBattleStyle]],
		},
		{
			Level: 1,
			Name:  "Второе дыхание",
			Description: "Вы обладаете ограниченным источником выносливости, которым можете воспользоваться, чтобы уберечь себя." +
				"В свой ход вы можете бонусным действием восстановить хиты в размере 1к10 + ваш уровень воина. " +
				"Использовав это умение, вы должны завершить короткий либо продолжительный отдых, чтобы получить возможность использовать его снова.",
		},
		{
			Level: 2,
			Name:  "Всплеск действий",
			Description: "Вы получаете возможность на мгновение преодолеть обычные возможности. " +
				"В свой ход вы можете совершить одно дополнительное действие помимо обычного и бонусного действий. " +
				"Использовав это умение, вы должны завершить короткий или продолжительный отдых, чтобы получить возможность использовать его снова.",
		},
		{
			Level:       3,
			Name:        "Воинский архетип",
			Description: "Вы выбираете архетип, отражающий стиль и технику, к которым вы стремитесь.",
		},
		{
			Level:       5,
			Name:        "Дополнительная атака",
			Description: "Если вы в свой ход совершаете действие Атака, вы можете совершить две атаки вместо одной.",
		},
	}
	if lvl >= 3 {
		//var archetypesList = []string{
		//	"Мастер боевых искусств",
		//	"Мистический рыцарь",
		//	"Чемпион",
		//	"Рыцарь пурпурного дракона",
		//	"Кавалерист",
		//	"Мистический лучник",
		//	"Самурай",
		//	"Рыцарь эха",
		//	"Пси-воин",
		//	"Рунный рыцарь",
		//}
		//randNum, _ := random.IntRange(0, len(archetypesList))
		archetypeName := "Мистический рыцарь"

		switch archetypeName {
		case "Мастер боевых искусств":
			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level: 3,
					Name:  "Боевое превосходство",
					Description: "Если вы выбираете этот архетип, вы изучаете приёмы, использующие специальные кости, называемые костями превосходства. <br>" +
						"<strong>Приёмы.</strong> Вы изучаете три приёма на ваш выбор. Большинство приёмов тем или иным образом усиливают атаку." +
						"Во время одной атаки вы можете использовать только один приём." +
						"Вы изучаете два дополнительных приёма при достижении 7-го, 10-го и 15-го уровней." +
						"Каждый раз, при изучении новых приёмов, вы можете также заменить один из известных вам приёмов на другой. <br>" +
						"<strong>Кости превосходства.</strong> У вас есть четыре кости превосходства. Это кости <strong>к8.</strong>" +
						"Кости превосходства тратятся при использовании." +
						"Вы восполняете все потраченные кости в конце короткого или продолжительного отдыха." +
						"Вы получаете ещё по одной кости превосходства на 7-м и 15-м уровнях. <br>" +
						"<strong>Спасброски.</strong> Некоторые из ваших приёмов требуют от цели спасброска, чтобы избежать эффекта приёма. <br>" +
						"Сложность такого спасброска рассчитывается следующим образом: <br>" +
						"<strong>Сложность спасброска приёма = 8 + ваш бонус мастерства + ваш модификатор Силы или Ловкости (на ваш выбор)</strong>",
				},
				{
					Level:       3,
					Name:        "Ученик войны",
					Description: "Вы осваиваете владением одним из ремесленных инструментов на ваш выбор.",
				},
				{
					Level:       7,
					Name:        "Познай своего врага",
					Description: "Если вы потратите как минимум 1 минуту, рассматривая, или по другому взаимодействуя с существом вне боя, вы можете узнать некоторую информацию о его способностях в сравнении с вашими. Мастер сообщит вам, равняется ли существо, превосходит или уступает вам в двух характеристиках на ваш выбор из перечисленных: <br>Значение Силы\nЗначение Ловкости\nЗначение Телосложения\nКласс Доспеха\nТекущие хиты\nОбщее количество уровней (если есть)\nКоличество уровней в классе Воин (если есть)",
				},
			}...)
		case "Мистический рыцарь":
			var cast classes.SpellCasting
			var modifier = getAbilityModifier(raceInfo, lvl)

			var ProficiencyBonus int
			if lvl <= 4 {
				ProficiencyBonus = 2
			}
			if lvl > 4 && lvl <= 8 {
				ProficiencyBonus = 3
			}

			cast.BasicSpellCharacteristics = "Интеллект"
			cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Charisma
			cast.SpellDamageModifier = ProficiencyBonus + modifier.Charisma
			switch lvl {
			case 3:
				cast.SpellsSlots = classes.SpellsSlots{OneLevel: 2}
				cast.ZeroLevelSpellsKnownCount = 2
				cast.OneLevelSpellsKnownCount = 3
				cast.TotalSpellCount = 5
			case 4:
				cast.SpellsSlots = classes.SpellsSlots{OneLevel: 3}
				cast.ZeroLevelSpellsKnownCount = 2
				cast.OneLevelSpellsKnownCount = 3
				cast.TotalSpellCount = 5
			case 5:
				cast.SpellsSlots = classes.SpellsSlots{OneLevel: 3}
				cast.ZeroLevelSpellsKnownCount = 2
				cast.OneLevelSpellsKnownCount = 4
				cast.TotalSpellCount = 6
			case 6:
				cast.SpellsSlots = classes.SpellsSlots{OneLevel: 3}
				cast.ZeroLevelSpellsKnownCount = 2
				cast.OneLevelSpellsKnownCount = 4
				cast.TotalSpellCount = 6
			case 7:
				cast.SpellsSlots = classes.SpellsSlots{OneLevel: 4, TwoLevel: 2}
				cast.ZeroLevelSpellsKnownCount = 2
				cast.OneLevelSpellsKnownCount = 5
				cast.TwoLevelSpellsKnownCount = 2
				cast.TotalSpellCount = 7
			case 8:
				cast.SpellsSlots = classes.SpellsSlots{OneLevel: 4, TwoLevel: 2}
				cast.ZeroLevelSpellsKnownCount = 2
				cast.OneLevelSpellsKnownCount = 6
				cast.TwoLevelSpellsKnownCount = 2
				cast.TotalSpellCount = 8
			}

			for i := 0; i < 5; i++ {
				var spellCount int
				switch i {
				case 0:
					spellCount = cast.ZeroLevelSpellsKnownCount
				case 1:
					spellCount = cast.OneLevelSpellsKnownCount
				case 2:
					spellCount = cast.TwoLevelSpellsKnownCount
				case 3:
					spellCount = cast.TreeLevelSpellsKnownCount
				case 4:
					spellCount = cast.FourLevelSpellsKnownCount
				}
				fighterSpellList = append(fighterSpellList,
					spells.GetRandomSpellForClass("Волшебник", i, spellCount)...)
			} //я ХЗ как дать ему спеллы  школ магии: Воплощения и Ограждения. Они все из ХомБрю

			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level: 3,
					Name:  "Использование заклинаний",
					Description: "К своим воинским талантам вы добавляете возможность накладывать заклинания. " +
						"Базовая характеристика заклинаний. Интеллект является базовой характеристикой для ваших заклинаний, " +
						"поскольку вы узнаёте заклинания посредством изучения и запоминания. " +
						"Когда заклинание требует использовать базовую характеристику, вы используете Интеллект. " +
						"Кроме того, вы используете модификатор Интеллекта при определении Сл спасбросков от ваших заклинаний волшебника и при броске атаки заклинаниями. <br>" +
						"Сл спасброска = 8 + ваш бонус мастерства + ваш модификатор Интеллекта <br>" +
						"Модификатор броска атаки = ваш бонус мастерства + ваш модификатор Интеллекта",
				},
				{
					Level: 3,
					Name:  "Связь с оружием",
					Description: "Вы узнаёте ритуал, позволяющий создать магическую связь между вами и одним оружием. " +
						"Вы выполняете ритуал в течение 1 часа, и он может быть совершён в течение короткого отдыха. " +
						"Оружие во время проведения ритуала должно находиться на доступном расстоянии от вас, и в конце вы должны прикоснуться к нему и создать связь. <br>" +
						"Как только вы привязали к себе оружие, вы не можете быть обезоруженным, пока не станете недееспособным. " +
						"Если оружие находится на одном плане существования с вами, вы можете в свой ход бонусным действием призвать его, телепортируя себе в руку. <br>" +
						"У вас может быть не более двух привязанных оружий одновременно, и бонусным действием вы призываете их по одному. " +
						"Если вы попытаетесь создать связь с третьим оружием, вам придётся разорвать связь с одним из первых двух.",
				},
				{
					Level:       7,
					Name:        "classAbilityPriority",
					Description: "classAbilityPriority",
				},
			}...)
		case "Чемпион":
			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level:       3,
					Name:        "Улучшенные критические попадания",
					Description: "Ваши атаки оружием совершают критическое попадание при выпадении «19» или «20» на кости атак",
				},
				{
					Level: 7,
					Name:  "Выдающийся атлет",
					Description: "Вы можете добавлять половину бонуса мастерства, округлённую в большую сторону, " +
						"ко всем проверкам Силы, Ловкости или Телосложения, куда этот бонус еще не включён.<br>" +
						"Кроме того, если вы совершаете прыжок в длину с разбега, " +
						"дальность прыжка увеличивается на количество футов, равное модификатору Силы.",
				},
			}...)
		case "Баннерет":
			fighterProf = getProficiencies(raceInfo, backgrInfo)

			//archetypeSkillsAvailable := []string{"Animal Handling", "Insight", "Intimidation", "Performance", "Persuasion"}

			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level: 3,
					Name:  "Сплочающий крик",
					Description: "Вы изучаете приёмы, позволяющие вдохновить ваших союзников на сражение, не смотря на былые ранения.\n\n" +
						"Когда вы используете «Второе дыхание», вы можете выбрать до трёх союзных существ в пределах 60 футах от вас. " +
						"Каждый из них восстанавливает количество хитов равное вашему уровню воина, при условии, что существо может вас видеть и слышать.",
				},
				{
					Level: 7,
					Name:  "Посланник короны",
					Description: "Рыцари Пурпурного Дракона служат посланниками короны Кормира. " +
						"Предполагается, что рыцари высокого положения будут вести себя достойно.\n\n" +
						"Вы получаете владение навыком Убеждение. " +
						"Если вы уже владеете этим навыком, выберите владение в одном из следующих навыков: " +
						"Уход за животными, Проницательность, Запугивание или Выступление.\n\n" +
						"Ваш бонус владения удваивается при любой проверке характеристики, использующей Убеждение. " +
						"Вы сохраняете выгоду от этого умения вне зависимости от навыка, которым вы овладели, используя это умение.\n\n",
				},
			}...)
		case "Кавалерист":
			//archetypeSkillsAvailable := []string{"Animal Handling", "History", "Intimidation", "Performance", "Persuasion"}

			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level:       3,
					Name:        "Дополнительные владения",
					Description: "",
				},
				{
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				},
			}...)
		case "Мистический лучник":
			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level:       3,
					Name:        "",
					Description: "",
				},
				{
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				},
			}...)
		case "Самурай":
			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level:       3,
					Name:        "",
					Description: "",
				},
				{
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				},
			}...)
		case "Рыцарь эха":
			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level:       3,
					Name:        "",
					Description: "",
				},
				{
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				},
			}...)
		case "Пси-воин":
			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level:       3,
					Name:        "",
					Description: "",
				},
				{
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				},
			}...)
		case "Рунный рыцарь":
			fighterClassAbilities = append(fighterClassAbilities, []classes.ClassAbility{
				{
					Level:       3,
					Name:        "",
					Description: "",
				},
				{
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				}, {
					Level:       3,
					Name:        "",
					Description: "",
				},
			}...)
		}

	}

	return fighterClassAbilities
}

func getBardClassAbilitiesWithLevel(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) []classes.ClassAbility {
	var abilitiesList = getClassAbilities(raceInfo, backgrInfo, lvl)
	var abilitiesAnswer []classes.ClassAbility

	for _, ability := range abilitiesList {
		if lvl >= ability.Level {
			abilitiesAnswer = append(abilitiesAnswer, ability)
		}
	}
	return abilitiesAnswer
}

func GetClass(raceInfo *races.Race, backgrInfo *backgrounds.Background, lvl int) *classes.Class {
	return &classes.Class{
		ClassName:       "Fighter",
		ClassNameRU:     "Воин",
		ClassAbilities:  getBardClassAbilitiesWithLevel(raceInfo, backgrInfo, lvl),
		AbilityScore:    getAbilitiesScore(raceInfo, lvl),
		AbilityModifier: getAbilityModifier(raceInfo, lvl),
		SavingThrows:    getSavingThrows(raceInfo, lvl),
		Inspiration:     false,
		Proficiencies:   *fighterProf, //need to fix
		Hits:            getHits(raceInfo, lvl),
		Caster:          false,
		SpellCasting:    classes.SpellCasting{},
		SpellsList:      fighterSpellList,
		Equipment:       getBardEquipment(),
	}
}
