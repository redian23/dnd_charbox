package export

import (
	"fmt"
)

func runExport() {
	var exp = ExportLSS{JSONType: "character", Template: "default",
		Name:       Name{Value: ""}, //var
		HiddenName: "234567",        //var??
		Info: Info{
			CharClass:  CharClass{Name: "charClass", Label: "класс и уровень", Value: "sdf"}, //var
			Level:      Level{Name: "level", Label: "уровень", Value: 1},                     //var
			Background: Background{Name: "background", Label: "предыстория", Value: "sdf"},   //var
			PlayerName: PlayerName{Name: "playerName", Label: "имя игрока", Value: "sdf"},    //var
			Race:       Race{Name: "race", Label: "раса", Value: "dfsd"},                     //var
			Alignment:  Alignment{Name: "alignment", Label: "мировоззрение", Value: "dfsd"},  //var
			Experience: Experience{Name: "experience", Label: "опыт", Value: 0},              //var
		},
		SubInfo: SubInfo{
			Age:    Age{Name: "age", Label: "возраст", Value: 0},     //var
			Height: Height{Name: "height", Label: "рост", Value: ""}, //var
			Weight: Weight{Name: "weight", Label: "вес", Value: ""},  //var
			Eyes:   Eyes{Name: "eyes", Label: "глаза", Value: ""},    //var
			Skin:   Skin{Name: "skin", Label: "кожа", Value: ""},     //var
			Hair:   Hair{Name: "hair", Label: "волосы", Value: ""},   //var
		},
		SpellsInfo: SpellsInfo{
			Base: Base{Name: "base", Label: "Базовая характеристика заклинаний", Value: ""}, //var
			Save: Save{Name: "save", Label: "Сложность спасброска", Value: ""},              //var
			Mod:  Mod{Name: "mod", Label: "Бонус атаки заклинанием", Value: ""},             //var
		},
		Spells:      Spells{},
		Proficiency: 2,
		Stats: Stats{
			Str: Str{Name: "str", Label: "Сила", Score: 0, Modifier: 0},         //var var
			Dex: Dex{Name: "dex", Label: "Ловкость", Score: 0, Modifier: 0},     //var var
			Con: Con{Name: "con", Label: "Телосложение", Score: 0, Modifier: 0}, //var var
			Int: Int{Name: "int", Label: "Интеллект", Score: 0, Modifier: 0},    //var var
			Wis: Wis{Name: "wis", Label: "Мудрость", Score: 0, Modifier: 0},     //var var
			Cha: Cha{Name: "cha", Label: "Харизма", Score: 0, Modifier: 0},      //var var
		},
		Saves: Saves{
			Str: SaveStr{Name: "str", IsProf: false}, //var
			Dex: SaveDex{Name: "dex", IsProf: false}, //var
			Con: SaveCon{Name: "con", IsProf: false}, //var
			Int: SaveInt{Name: "int", IsProf: false}, //var
			Wis: SaveWis{Name: "wis", IsProf: false}, //var
			Cha: SaveCha{Name: "cha", IsProf: false}, //var
		},
		Skills: Skills{
			Acrobatics:     Acrobatics{BaseStat: "dex", Name: "", Label: "", IsProf: 0},     //var
			Investigation:  Investigation{BaseStat: "int", Name: "", Label: "", IsProf: 0},  //var
			Athletics:      Athletics{BaseStat: "str", Name: "", Label: "", IsProf: 0},      //var
			Perception:     Perception{BaseStat: "wis", Name: "", Label: "", IsProf: 0},     //var
			Survival:       Survival{BaseStat: "wis", Name: "", Label: "", IsProf: 0},       //var
			Performance:    Performance{BaseStat: "cha", Name: "", Label: "", IsProf: 0},    //var
			Intimidation:   Intimidation{BaseStat: "cha", Name: "", Label: "", IsProf: 0},   //var
			History:        History{BaseStat: "int", Name: "", Label: "", IsProf: 0},        //var
			SleightOfHand:  SleightOfHand{BaseStat: "dex", Name: "", Label: "", IsProf: 0},  //var
			Arcana:         Arcana{BaseStat: "int", Name: "", Label: "", IsProf: 0},         //var
			Medicine:       Medicine{BaseStat: "wis", Name: "", Label: "", IsProf: 0},       //var
			Deception:      Deception{BaseStat: "cha", Name: "", Label: "", IsProf: 0},      //var
			Nature:         Nature{BaseStat: "int", Name: "", Label: "", IsProf: 0},         //var
			Insight:        Insight{BaseStat: "wis", Name: "", Label: "", IsProf: 0},        //var
			Religion:       Religion{BaseStat: "int", Name: "", Label: "", IsProf: 0},       //var
			Stealth:        Stealth{BaseStat: "dex", Name: "", Label: "", IsProf: 0},        //var
			Persuasion:     Persuasion{BaseStat: "cha", Name: "", Label: "", IsProf: 0},     //var
			AnimalHandling: AnimalHandling{BaseStat: "wis", Name: "", Label: "", IsProf: 0}, //var
		},
		Vitality: Vitality{
			HpDiceCurrent:  HpDiceCurrent{Value: 1},
			HpDiceMulti:    HpDiceMulti{Value: 0},
			Initiative:     Initiative{Value: 0},
			Speed:          Speed{Value: ""},
			Ac:             Ac{Value: 0},
			HpMax:          HpMax{Value: 0},
			HpCurrent:      HpCurrent{Value: 0},
			IsDying:        false,
			DeathFails:     0,
			DeathSuccesses: 0,
			HpTemp:         HpTemp{Value: 0},
		},
		WeaponsList: []WeaponsList{{}},
		Weapons:     Weapons{},
		Text: Text{
			Traits: Traits{},
			Equipment: Equipment{Value: Value{Data: Data{
				Type: "doc",
				ContentText: []ContentText{{
					Type: "paragraph",
					Content: []Content{{
						Type:  "text",
						Marks: []Marks{{Type: "bold"}},
						Text:  "Контейнер для свитков битком набитый вашими изысканиями",
					}}},
					{
						Type: "paragraph",
						Content: []Content{{
							Type:  "text",
							Marks: []Marks{{Type: "bold"}},
							Text:  "тёплое одеяло,",
						}}},
					{
						Type: "paragraph",
						Content: []Content{{
							Type:  "text",
							Marks: []Marks{{Type: "bold"}},
							Text:  "Кожаный доспех,",
						}}}},
			}}},
			Prof: Prof{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
			Attacks: Attacks{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
			Allies: Allies{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
			Background: BackgroundFullText{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
			Features: Features{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}, Size: 0},
			Quests: Quests{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
			Flaws: Flaws{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}, Size: 0},
			Personality: Personality{Value: Value{Data: Data{Type: "", ContentText: nil}}},
			Bonds: Bonds{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
			Ideals: Ideals{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
			Items: Items{Value: Value{Data: Data{
				Type:        "",
				ContentText: nil,
			}}},
		},
		Coins: Coins{
			Cp: Cp{Value: 0},
			Sp: Sp{Value: 0},
			Ep: Ep{Value: 0},
			Gp: Gp{Value: 0},
			Pp: Pp{Value: 0}},
		Resources:     Resources{},
		BonusesSkills: BonusesSkills{},
		BonusesStats:  BonusesStats{},
		Conditions:    nil,
		ID:            "",
		Avatar: Avatar{
			Jpeg: "",
			Webp: "",
		},
		Inspiration: false,
	}

	fmt.Println(exp)
}
