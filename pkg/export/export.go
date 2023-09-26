package export

import (
	"github.com/mazen160/go-random"
	"pregen/pkg/characterCore"
	"strconv"
)

func getHiddenName() string {
	var adj = []string{"Адский", "Буйный", "Бурный", "Веский", "Волчий", "Высший", "Глухой", "грубый", "Густой", "Дивный", "Дикий", "Добрый", "Жаркий", "Жгучий", "Живой", "Жуткий", "Злой", "Крутой", "Лютый", "Мощный", "Наглый", "Немой", "Острый", "Первый", "Полный", "Прямой", "Пущий", "Редкий", "Резкий", "Рьяный", "Седой", "Слепой", "Несущий", "Тяжкий", "Черный", "Чертов", "Чистый", "Чумный", "Щедрый", "Яркий", "Ярый"}
	var noun = []string{"Год", "Раз", "Друг", "Глаз", "Дом", "Мир", "Вид", "Час", "Бог", "Стол", "Суд", "Свет", "Язык", "Век", "Труд", "Сын", "Счет", "Ряд", "План", "Ход", "срок", "Опыт", "Лес", "Зал", "Шаг", "Брат", "Рост", "Тип", "Край", "Пол", "Банк", "Союз", "Врач", "Факт", "Угол", "Двор", "Род", "Цвет", "Круг", "Поэт", "Дух", "Сон", "Ум", "Удар", "Нос", "Сад", "Курс", "Рот", "Фонд", "Лист", "Снег"}
	var randNum, _ = random.IntRange(10000000, 99999999)

	randAdj, _ := random.IntRange(0, len(adj))
	randNoun, _ := random.IntRange(0, len(noun))

	return adj[randAdj] + " " + noun[randNoun] + "_" + strconv.Itoa(randNum)
}

func RunExportToLSS(lc characterCore.Character) *ExportToLss {

	return &ExportToLss{
		JSONType:   "character",
		Template:   "default",
		Name:       Name{Value: lc.Race.CharacterName}, //var
		HiddenName: getHiddenName(),                    //var
		Info: Info{
			CharClass:  CharClass{Name: "charClass", Label: "класс и уровень", Value: lc.Class.ClassNameRU},             //var
			Level:      Level{Name: "level", Label: "уровень", Value: 1},                                                //var
			Background: BackgroundInfo{Name: "background", Label: "предыстория", Value: lc.Background.BackgroundNameRu}, //var
			PlayerName: PlayerName{Name: "playerName", Label: "имя игрока", Value: "Generated"},                         //var
			Race:       Race{Name: "race", Label: "раса", Value: lc.Race.RaceNameRu},                                    //var
			Alignment:  Alignment{Name: "alignment", Label: "мировоззрение", Value: lc.Background.Worldview},            //var
			Experience: Experience{Name: "experience", Label: "опыт", Value: "0"},                                       //var
		},
		SubInfo: SubInfo{
			Age:    Age{Name: "age", Label: "возраст", Value: strconv.Itoa(lc.Race.Age)},        //var
			Height: Height{Name: "height", Label: "рост", Value: lc.Race.HeightFt},              //var
			Weight: Weight{Name: "weight", Label: "вес", Value: strconv.Itoa(lc.Race.WeightLb)}, //var
			Eyes:   Eyes{Name: "eyes", Label: "глаза", Value: lc.Race.Eyes},                     //var
			Skin:   Skin{Name: "skin", Label: "кожа", Value: ""},                                //FIX var
			Hair:   Hair{Name: "hair", Label: "волосы", Value: lc.Race.Hair},                    //var
		},
		SpellsInfo: SpellsInfo{
			Base: Base{Name: "base", Label: "Базовая характеристика заклинаний", Value: lc.Class.Spellcasting.BasicSpellCharacteristics}, //var
			Save: Save{Name: "save", Label: "Сложность спасброска", Value: strconv.Itoa(lc.Class.Spellcasting.SavingThrowDifficulty)},    //var
			Mod:  Mod{Name: "mod", Label: "Бонус атаки заклинанием", Value: strconv.Itoa(lc.Class.Spellcasting.SpellDamageModifier)},     //var                                               //var
		},
		Spells: Spells{
			Slots1:  Slots1{Value: "2"},
			Spell10: Spell10{IsReady: true},
			Spell11: Spell11{IsReady: false},
		},
		Proficiency: 2,
		Stats: Stats{
			Str: StrStat{Name: "str", Label: "Сила", Score: lc.Class.Ability.Strength, Modifier: lc.Class.Modifier.Strength},                     //var var
			Dex: DexStat{Name: "dex", Label: "Ловкость", Score: lc.Class.Ability.Dexterity, Modifier: lc.Class.Modifier.Dexterity},               //var var
			Con: ConStat{Name: "con", Label: "Телосложение", Score: lc.Class.Ability.BodyDifficulty, Modifier: lc.Class.Modifier.BodyDifficulty}, //var var
			Int: IntStat{Name: "int", Label: "Интеллект", Score: lc.Class.Ability.Intelligence, Modifier: lc.Class.Modifier.Intelligence},        //var var
			Wis: WisStat{Name: "wis", Label: "Мудрость", Score: lc.Class.Ability.Wisdom, Modifier: lc.Class.Modifier.Wisdom},                     //var var
			Cha: ChaStat{Name: "cha", Label: "Харизма", Score: lc.Class.Ability.Charisma, Modifier: lc.Class.Modifier.Charisma},                  //var var
		},
		Saves: Saves{
			Str: Str{Name: "str", IsProf: lc.Class.SavingThrows.Strength.Mastery},       //var
			Dex: Dex{Name: "dex", IsProf: lc.Class.SavingThrows.Dexterity.Mastery},      //var
			Con: Con{Name: "con", IsProf: lc.Class.SavingThrows.BodyDifficulty.Mastery}, //var
			Int: Int{Name: "int", IsProf: lc.Class.SavingThrows.Intelligence.Mastery},   //var
			Wis: Wis{Name: "wis", IsProf: lc.Class.SavingThrows.Wisdom.Mastery},         //var
			Cha: Cha{Name: "cha", IsProf: lc.Class.SavingThrows.Charisma.Mastery},       //var
		},
		Skills: Skills{
			Acrobatics:     skill{BaseStat: "dex", Name: "acrobatics", Label: "Акробатика", IsProf: lc.Skills.Acrobatics.ProfLSS},
			Investigation:  skill{BaseStat: "dex", Name: "investigation", Label: "Анализ", IsProf: lc.Skills.Investigation.ProfLSS},
			Athletics:      skill{BaseStat: "dex", Name: "athletics", Label: "Атлетика", IsProf: lc.Skills.Athletics.ProfLSS},
			Perception:     skill{BaseStat: "dex", Name: "perception", Label: "Восприятие", IsProf: lc.Skills.Perception.ProfLSS},
			Survival:       skill{BaseStat: "dex", Name: "survival", Label: "Выживание", IsProf: lc.Skills.Survival.ProfLSS},
			Performance:    skill{BaseStat: "dex", Name: "performance", Label: "Выступление", IsProf: lc.Skills.Performance.ProfLSS},
			Intimidation:   skill{BaseStat: "dex", Name: "intimidation", Label: "Запугивание", IsProf: lc.Skills.Intimidation.ProfLSS},
			History:        skill{BaseStat: "dex", Name: "history", Label: "История", IsProf: lc.Skills.History.ProfLSS},
			SleightOfHand:  skill{BaseStat: "dex", Name: "sleight of hand", Label: "Ловкость рук", IsProf: lc.Skills.SleightOfHand.ProfLSS},
			Arcana:         skill{BaseStat: "dex", Name: "arcana", Label: "Магия", IsProf: lc.Skills.Arcana.ProfLSS},
			Medicine:       skill{BaseStat: "dex", Name: "medicine", Label: "Медицина", IsProf: lc.Skills.Medicine.ProfLSS},
			Deception:      skill{BaseStat: "dex", Name: "deception", Label: "Обман", IsProf: lc.Skills.Deception.ProfLSS},
			Nature:         skill{BaseStat: "dex", Name: "nature", Label: "Природа", IsProf: lc.Skills.Nature.ProfLSS},
			Insight:        skill{BaseStat: "dex", Name: "insight", Label: "Проницательность", IsProf: lc.Skills.Insight.ProfLSS},
			Religion:       skill{BaseStat: "dex", Name: "religion", Label: "Религия", IsProf: lc.Skills.Religion.ProfLSS},
			Stealth:        skill{BaseStat: "dex", Name: "stealth", Label: "Скрытность", IsProf: lc.Skills.Stealth.ProfLSS},
			Persuasion:     skill{BaseStat: "dex", Name: "persuasion", Label: "Убеждение", IsProf: lc.Skills.Persuasion.ProfLSS},
			AnimalHandling: skill{BaseStat: "dex", Name: "animal handling", Label: "Уход за животными", IsProf: lc.Skills.AnimalHandling.ProfLSS},
		},
		Vitality: Vitality{
			HpMax:         HpMax{Value: strconv.Itoa(lc.Class.Hits.HitCount)},
			HpCurrent:     HpCurrent{Value: strconv.Itoa(lc.Class.Hits.HitCount)},
			HitDie:        HitDie{Value: lc.Class.Hits.HitDice},
			HpDiceCurrent: HpDiceCurrent{Value: 1},
			Ac:            Ac{Value: strconv.Itoa(lc.Class.Armor[0].ArmorClassCount)},
			Speed:         Speed{Value: strconv.Itoa(lc.Race.Speed)},
		},
		WeaponsList: getWeaponList(lc),
		Text: Text{
			Attacks:   Attacks{Value: Value{Data: ""}}, //spells
			Equipment: Equipment{Value: Value{getEquipString(lc)}},
			Prof:      Prof{Value: Value{getProfs(lc)}},
			Traits: Traits{Value: Value{Data: "<strong>Умения от класса:</strong> " + getClassAbil(lc) +
				"<p><strong>Умения от расы:</strong> " + getRaceAbil(lc) + "</p>"}},
			Personality: Personality{Value: Value{Data: "Черта от расы: " + lc.Background.CharacterTrait}},
			Ideals:      Ideals{Value: Value{Data: lc.Background.Ideal}},
			Bonds:       Bonds{Value: Value{Data: lc.Background.Affection}},
			Flaws:       Flaws{Value: Value{Data: lc.Background.Weakness}},
			Background: Background{Value: Value{Data: lc.Background.Personalization + " <p>" +
				lc.Background.Description + "</p>" +
				"<strong>Совет:</strong> " + lc.Background.Advice}},
			Features: Features{Value: Value{Data: "<p><strong>" + lc.Background.BackgroundAbility.AbilityName + "</strong>: " +
				lc.Background.BackgroundAbility.Description}},
			SpellsLevel0: SpellsLevel{Value: Value{Data: getZeroLvlSpells(lc)}},
			SpellsLevel1: SpellsLevel{Value: Value{Data: ""}},
		},
		Avatar: Avatar{
			Jpeg: "https://charbox.swn.by/imgs/" + lc.Race.RacePhoto.Path + lc.Race.RacePhoto.FileName,
			Webp: "",
		},
		Coins: Coins{
			Total: Total{Value: float64(lc.Background.Gold)},
			Gp:    Gp{Value: lc.Background.Gold},
		},
	}
}

func getProfs(lc characterCore.Character) string {
	var profString string
	var langs string
	for i, lng := range lc.Langs {
		if i == len(lc.Langs)-1 {
			langs += lng
			break
		}
		langs += lng + ", "
	}

	var profTools string
	for i, tool := range lc.Background.MasteryOfTools {
		if i == len(lc.Background.MasteryOfTools)-1 {
			profTools += tool
			break
		}
		profTools += tool + ", "
	}
	profString += "<p><strong>Знание языков:</strong> " + langs + "</p>" +
		"<p><strong>Сопротивление к:</strong> " + lc.Race.Resist + "</p>" +
		"<p><strong>Доспехи:</strong> " + lc.Class.Proficiencies.Armor + "</p>" +
		"<p><strong>Оружие:</strong> " + lc.Class.Proficiencies.Weapons + "</p>" +
		"<p><strong>Инструменты от класса: </strong> " + lc.Class.Proficiencies.Instruments + "</p>" +
		"<p><strong>Инструменты от предыстории: </strong> " + profTools + "</p>"

	return profString
}

func getEquipString(lc characterCore.Character) string {
	var equipString string
	//кастыли для интеграции с LSS (самому больно, но приходится)
	for i, equipment := range lc.Class.ClassEquipment {
		if equipment.Count > 1 {
			if i == len(lc.Class.ClassEquipment)-1 {
				equipString += equipment.ItemName + " (" + strconv.Itoa(equipment.Count) + ")"
				break
			}
			equipString += equipment.ItemName + " (" + strconv.Itoa(equipment.Count) + "), "
		} else {
			if i == len(lc.Class.ClassEquipment)-1 {
				equipString += equipment.ItemName
				break
			}
			equipString += equipment.ItemName + ", "
		}
	}

	var backgroundEquipString string
	for i, equipment := range lc.Background.Equipment {
		if i == len(lc.Class.ClassEquipment)-1 {
			backgroundEquipString += equipment
			break
		}
		backgroundEquipString += equipment + ", "
	}
	return equipString + ", " + backgroundEquipString
}

func getWeaponList(lc characterCore.Character) []WeaponsList {
	var wplist []WeaponsList
	var wp WeaponsList

	for _, weapon := range lc.Class.Weapon {
		wp.Name.Value = weapon.WeaponName
		wp.Mod.Value = strconv.Itoa(lc.Class.Modifier.Strength)
		wp.Dmg.Value = weapon.Damage
		wplist = append(wplist, wp)
	}
	return wplist
}

func getClassAbil(lc characterCore.Character) string {
	var classAbilities string
	for _, ability := range lc.Class.ClassAbilities {
		classAbilities += "<p><strong>" + ability.Name + ":</strong> " + ability.Description + "</p>"
	}
	if classAbilities == "" {
		classAbilities = "Нет"
	}
	return classAbilities
}

func getRaceAbil(lc characterCore.Character) string {
	var raseAbilities string
	for _, ability := range lc.Race.RaceAbilities {
		raseAbilities += "<p><strong>" + ability.AbilityName + ":</strong> " + ability.Description + "</p>"
	}
	if raseAbilities == "" {
		raseAbilities = "Нет"
	}
	return raseAbilities
}

func getZeroLvlSpells(lc characterCore.Character) string {
	var zeroSpellList string

	if len(lc.SpellsList.ZeroLevelSpells) == 0 {
		zeroSpellList = "Не владеет заговорами."
		return zeroSpellList
	}
	for _, sp := range lc.SpellsList.ZeroLevelSpells {
		zeroSpellList += "<p>" + sp + "</p>"

	}
	return zeroSpellList
}
