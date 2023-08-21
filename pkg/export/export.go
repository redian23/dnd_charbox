package export

import (
	"pregen/pkg/characterCore"
	"strconv"
)

//func getHiddenName() {
//	var adj = []string{"адский", "буйный", "бурный", "веский", "волчий", "высший", "глухой", "грубый", "густой", "дивный", "дикий", "добрый", "жаркий", "жгучий", "живой", "жуткий", "злой", "крутой", "лютый", "мощный", "наглый", "немой", "острый", "первый", "полный", "прямой", "пущий", "редкий", "резкий", "рьяный", "седой", "слепой", "сущий", "тяжкий", "черный", "чертов", "чистый", "шумный", "щедрый", "яркий", "ярый"}
//	var noun = []string{"год", "раз", "друг", "глаз", "дом", "мир", "вид", "час", "бог", "стол", "суд", "свет", "язык", "век", "труд", "сын", "счет", "ряд", "план", "ход", "срок", "опыт", "лес", "зал", "шаг", "брат", "рост", "тип", "край", "пол", "банк", "союз", "врач", "факт", "угол", "двор", "род", "цвет", "круг", "поэт", "дух", "сон", "ум", "удар", "нос", "сад", "курс", "рот", "фонд", "лист", "снег"}
//
//}

func RunExportToLSS(lc characterCore.Character) *ExportToLss {

	return &ExportToLss{
		JSONType:   "character",
		Template:   "default",
		Name:       Name{Value: lc.Race.CharacterName}, //var
		HiddenName: "Крутой Язык_18647742",             //var
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
			Base: Base{Name: "base", Label: "Базовая характеристика заклинаний", Value: ""}, //var
			Save: Save{Name: "save", Label: "Сложность спасброска", Value: ""},              //var
			Mod:  Mod{Name: "mod", Label: "Бонус атаки заклинанием", Value: ""},             //var
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
			Acrobatics:     skill{BaseStat: "dex", Name: "acrobatics", Label: "Акробатика"},
			Investigation:  skill{BaseStat: "dex", Name: "investigation", Label: "Анализ"},
			Athletics:      skill{BaseStat: "dex", Name: "athletics", Label: "Атлетика"},
			Perception:     skill{BaseStat: "dex", Name: "perception", Label: "Восприятие"},
			Survival:       skill{BaseStat: "dex", Name: "survival", Label: "Выживание"},
			Performance:    skill{BaseStat: "dex", Name: "performance", Label: "Выступление"},
			Intimidation:   skill{BaseStat: "dex", Name: "intimidation", Label: "Запугивание"},
			History:        skill{BaseStat: "dex", Name: "history", Label: "История"},
			SleightOfHand:  skill{BaseStat: "dex", Name: "sleight of hand", Label: "Ловкость рук"},
			Arcana:         skill{BaseStat: "dex", Name: "arcana", Label: "Магия"},
			Medicine:       skill{BaseStat: "dex", Name: "medicine", Label: "Медицина"},
			Deception:      skill{BaseStat: "dex", Name: "deception", Label: "Обман"},
			Nature:         skill{BaseStat: "dex", Name: "nature", Label: "Природа"},
			Insight:        skill{BaseStat: "dex", Name: "insight", Label: "Проницательность"},
			Religion:       skill{BaseStat: "dex", Name: "religion", Label: "Религия"},
			Stealth:        skill{BaseStat: "dex", Name: "stealth", Label: "Скрытность"},
			Persuasion:     skill{BaseStat: "dex", Name: "persuasion", Label: "Убеждение"},
			AnimalHandling: skill{BaseStat: "dex", Name: "animal handling", Label: "Уход за животными"},
		},
		Abilities: Abilities{
			Acrobatics:     Ability{BaseStat: "dex", Name: "acrobatics", Label: "Акробатика", IsProf: lc.Skills.Acrobatics.Proficiency},
			Investigation:  Ability{BaseStat: "dex", Name: "investigation", Label: "Анализ", IsProf: lc.Skills.Investigation.Proficiency},
			Athletics:      Ability{BaseStat: "dex", Name: "athletics", Label: "Атлетика", IsProf: lc.Skills.Athletics.Proficiency},
			Perception:     Ability{BaseStat: "dex", Name: "perception", Label: "Восприятие", IsProf: lc.Skills.Perception.Proficiency},
			Survival:       Ability{BaseStat: "dex", Name: "survival", Label: "Выживание", IsProf: lc.Skills.Survival.Proficiency},
			Performance:    Ability{BaseStat: "dex", Name: "performance", Label: "Выступление", IsProf: lc.Skills.Performance.Proficiency},
			Intimidation:   Ability{BaseStat: "dex", Name: "intimidation", Label: "Запугивание", IsProf: lc.Skills.Intimidation.Proficiency},
			History:        Ability{BaseStat: "dex", Name: "history", Label: "История", IsProf: lc.Skills.History.Proficiency},
			SleightOfHand:  Ability{BaseStat: "dex", Name: "sleight of hand", Label: "Ловкость рук", IsProf: lc.Skills.SleightOfHand.Proficiency},
			Arcana:         Ability{BaseStat: "dex", Name: "arcana", Label: "Магия", IsProf: lc.Skills.Arcana.Proficiency},
			Medicine:       Ability{BaseStat: "dex", Name: "medicine", Label: "Медицина", IsProf: lc.Skills.Medicine.Proficiency},
			Deception:      Ability{BaseStat: "dex", Name: "deception", Label: "Обман", IsProf: lc.Skills.Deception.Proficiency},
			Nature:         Ability{BaseStat: "dex", Name: "nature", Label: "Природа", IsProf: lc.Skills.Nature.Proficiency},
			Insight:        Ability{BaseStat: "dex", Name: "insight", Label: "Проницательность", IsProf: lc.Skills.Insight.Proficiency},
			Religion:       Ability{BaseStat: "dex", Name: "religion", Label: "Религия", IsProf: lc.Skills.Religion.Proficiency},
			Stealth:        Ability{BaseStat: "dex", Name: "stealth", Label: "Скрытность", IsProf: lc.Skills.Stealth.Proficiency},
			Persuasion:     Ability{BaseStat: "dex", Name: "persuasion", Label: "Убеждение", IsProf: lc.Skills.Persuasion.Proficiency},
			AnimalHandling: Ability{BaseStat: "dex", Name: "animal handling", Label: "Уход за животными", IsProf: lc.Skills.AnimalHandling.Proficiency},
		},
		Vitality: Vitality{
			HpMax:         HpMax{Value: strconv.Itoa(lc.Class.Hits.HitCount)},
			HpCurrent:     HpCurrent{Value: strconv.Itoa(lc.Class.Hits.HitCount)},
			HitDie:        HitDie{Value: lc.Class.Hits.HitDice},
			HpDiceCurrent: HpDiceCurrent{Value: 1},
			Ac:            Ac{Value: "12"},
			Speed:         Speed{Value: strconv.Itoa(lc.Race.Speed)},
		},
		WeaponsList: []WeaponsList{
			{ID: "", Name: Name{Value: ""}, Mod: Modif{Value: ""}, Dmg: Dmg{Value: ""}},
			{ID: "", Name: Name{Value: ""}, Mod: Modif{Value: ""}, Dmg: Dmg{Value: ""}},
		},
		Weapons: Weapons{
			WeaponName0: WeaponName0{Value: "Кинжал"},
			WeaponMod0:  WeaponMod0{Value: "+5"},
			WeaponDmg0:  WeaponDmg0{Value: "1к4 + 3 колющий"},
		},
		Text: Text{
			Attacks:      Attacks{Value: Value{Data: ""}},
			Equipment:    Equipment{Value: Value{Data: ""}},
			Prof:         Prof{Value: Value{Data: ""}},
			Traits:       Traits{Value: Value{Data: ""}},
			Allies:       Allies{Value: Value{Data: ""}},
			Features:     Features{Value: Value{Data: ""}},
			Personality:  Personality{Value: Value{Data: ""}},
			Ideals:       Ideals{Value: Value{Data: ""}},
			Bonds:        Bonds{Value: Value{Data: ""}},
			Flaws:        Flaws{Value: Value{Data: ""}},
			Background:   Background{Value: Value{Data: ""}},
			Quests:       Quests{Value: Value{Data: ""}},
			SpellsLevel0: SpellsLevel0{Value: Value{Data: ""}},
			SpellsLevel1: SpellsLevel1{Value: Value{Data: ""}},
		},
		Coins:         Coins{},
		Resources:     Resources{},
		BonusesSkills: BonusesSkills{},
		BonusesStats:  BonusesStats{},
		Conditions:    nil,
		CasterClass:   CasterClass{Value: lc.Class.ClassNameRU},
	}

}
