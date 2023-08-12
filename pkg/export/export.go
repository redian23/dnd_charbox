package export

//
//import (
//	"encoding/json"
//	"pregen/pkg/characterCore"
//	"strconv"
//)
//
//func RunExportToLSS() {
//	var lc characterCore.Character
//
//	var exp = ExportLSS{JSONType: "character", Template: "default",
//		Name:       Name{Value: lc.CharacterName}, //var
//		HiddenName: "234567",                      //var??
//		Info: Info{
//			CharClass:  CharClass{Name: "charClass", Label: "класс и уровень", Value: lc.Class.ClassNameRU},         //var
//			Level:      Level{Name: "level", Label: "уровень", Value: 1},                                            //var
//			Background: Background{Name: "background", Label: "предыстория", Value: lc.Background.BackgroundNameRu}, //var
//			PlayerName: PlayerName{Name: "playerName", Label: "имя игрока", Value: "Generated"},                     //FIX var
//			Race:       Race{Name: "race", Label: "раса", Value: lc.Race.RaceNameRu},                                //var
//			Alignment:  Alignment{Name: "alignment", Label: "мировоззрение", Value: lc.Background.Worldview},        //var
//			Experience: Experience{Name: "experience", Label: "опыт", Value: 0},                                     //var
//		},
//		SubInfo: SubInfo{
//			Age:    Age{Name: "age", Label: "возраст", Value: lc.Race.Age},                      //var
//			Height: Height{Name: "height", Label: "рост", Value: lc.Race.HeightFt},              //var
//			Weight: Weight{Name: "weight", Label: "вес", Value: strconv.Itoa(lc.Race.WeightLb)}, //var
//			Eyes:   Eyes{Name: "eyes", Label: "глаза", Value: lc.Race.Eyes},                     //var
//			Skin:   Skin{Name: "skin", Label: "кожа", Value: ""},                                //FIX var
//			Hair:   Hair{Name: "hair", Label: "волосы", Value: lc.Race.Hair},                    //var
//		},
//		SpellsInfo: SpellsInfo{
//			Base: Base{Name: "base", Label: "Базовая характеристика заклинаний", Value: ""}, //var
//			Save: Save{Name: "save", Label: "Сложность спасброска", Value: ""},              //var
//			Mod:  Mod{Name: "mod", Label: "Бонус атаки заклинанием", Value: ""},             //var
//		},
//		Spells:      Spells{},
//		Proficiency: 2,
//		Stats: Stats{
//			Str: Str{Name: "str", Label: "Сила", Score: lc.Class.Ability.Strength, Modifier: lc.Class.Modifier.Strength},                     //var var
//			Dex: Dex{Name: "dex", Label: "Ловкость", Score: lc.Class.Ability.Dexterity, Modifier: lc.Class.Modifier.Dexterity},               //var var
//			Con: Con{Name: "con", Label: "Телосложение", Score: lc.Class.Ability.BodyDifficulty, Modifier: lc.Class.Modifier.BodyDifficulty}, //var var
//			Int: Int{Name: "int", Label: "Интеллект", Score: lc.Class.Ability.Intelligence, Modifier: lc.Class.Modifier.Intelligence},        //var var
//			Wis: Wis{Name: "wis", Label: "Мудрость", Score: lc.Class.Ability.Wisdom, Modifier: lc.Class.Modifier.Wisdom},                     //var var
//			Cha: Cha{Name: "cha", Label: "Харизма", Score: lc.Class.Ability.Charisma, Modifier: lc.Class.Modifier.Charisma},                  //var var
//		},
//		Saves: Saves{
//			Str: SaveStr{Name: "str", IsProf: lc.Class.SavingThrows.Strength.Mastery},       //var
//			Dex: SaveDex{Name: "dex", IsProf: lc.Class.SavingThrows.Dexterity.Mastery},      //var
//			Con: SaveCon{Name: "con", IsProf: lc.Class.SavingThrows.BodyDifficulty.Mastery}, //var
//			Int: SaveInt{Name: "int", IsProf: lc.Class.SavingThrows.Intelligence.Mastery},   //var
//			Wis: SaveWis{Name: "wis", IsProf: lc.Class.SavingThrows.Wisdom.Mastery},         //var
//			Cha: SaveCha{Name: "cha", IsProf: lc.Class.SavingThrows.Charisma.Mastery},       //var
//		},
//		Skills: Skills{
//			Acrobatics:     Acrobatics{BaseStat: "dex", Name: "", Label: "", IsProf: 0},     //var
//			Investigation:  Investigation{BaseStat: "int", Name: "", Label: "", IsProf: 0},  //var
//			Athletics:      Athletics{BaseStat: "str", Name: "", Label: "", IsProf: 0},      //var
//			Perception:     Perception{BaseStat: "wis", Name: "", Label: "", IsProf: 0},     //var
//			Survival:       Survival{BaseStat: "wis", Name: "", Label: "", IsProf: 0},       //var
//			Performance:    Performance{BaseStat: "cha", Name: "", Label: "", IsProf: 0},    //var
//			Intimidation:   Intimidation{BaseStat: "cha", Name: "", Label: "", IsProf: 0},   //var
//			History:        History{BaseStat: "int", Name: "", Label: "", IsProf: 0},        //var
//			SleightOfHand:  SleightOfHand{BaseStat: "dex", Name: "", Label: "", IsProf: 0},  //var
//			Arcana:         Arcana{BaseStat: "int", Name: "", Label: "", IsProf: 0},         //var
//			Medicine:       Medicine{BaseStat: "wis", Name: "", Label: "", IsProf: 0},       //var
//			Deception:      Deception{BaseStat: "cha", Name: "", Label: "", IsProf: 0},      //var
//			Nature:         Nature{BaseStat: "int", Name: "", Label: "", IsProf: 0},         //var
//			Insight:        Insight{BaseStat: "wis", Name: "", Label: "", IsProf: 0},        //var
//			Religion:       Religion{BaseStat: "int", Name: "", Label: "", IsProf: 0},       //var
//			Stealth:        Stealth{BaseStat: "dex", Name: "", Label: "", IsProf: 0},        //var
//			Persuasion:     Persuasion{BaseStat: "cha", Name: "", Label: "", IsProf: 0},     //var
//			AnimalHandling: AnimalHandling{BaseStat: "wis", Name: "", Label: "", IsProf: 0}, //var
//		},
//		Vitality: Vitality{
//			HpDiceCurrent:  HpDiceCurrent{Value: 1},
//			HpDiceMulti:    HpDiceMulti{Value: 0},
//			Initiative:     Initiative{Value: 0},
//			Speed:          Speed{Value: strconv.Itoa(lc.Race.Speed)},
//			Ac:             Ac{Value: 0},
//			HpMax:          HpMax{Value: 0},
//			HpCurrent:      HpCurrent{Value: 0},
//			IsDying:        false,
//			DeathFails:     0,
//			DeathSuccesses: 0,
//			HpTemp:         HpTemp{Value: 0},
//		},
//		WeaponsList: []WeaponsList{{}},
//		Weapons:     Weapons{},
//		Text: Text{
//			Traits: Traits{},
//			Equipment: Equipment{Value: Value{Data: Data{
//				Type: "doc",
//				ContentText: []ContentText{{
//					Type: "paragraph",
//					Content: []Content{{
//						Type:  "text",
//						Marks: []Marks{{Type: "bold"}},
//						Text:  "Контейнер для свитков битком набитый вашими изысканиями",
//					}}},
//					{
//						Type: "paragraph",
//						Content: []Content{{
//							Type:  "text",
//							Marks: []Marks{{Type: "bold"}},
//							Text:  "тёплое одеяло,",
//						}}},
//					{
//						Type: "paragraph",
//						Content: []Content{{
//							Type:  "text",
//							Marks: []Marks{{Type: "bold"}},
//							Text:  "Кожаный доспех,",
//						}}}},
//			}}},
//			Prof: Prof{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//			Attacks: Attacks{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//			Allies: Allies{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//			Background: BackgroundFullText{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//			Features: Features{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}, Size: 0},
//			Quests: Quests{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//			Flaws: Flaws{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}, Size: 0},
//			Personality: Personality{Value: Value{Data: Data{Type: "", ContentText: nil}}},
//			Bonds: Bonds{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//			Ideals: Ideals{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//			Items: Items{Value: Value{Data: Data{
//				Type:        "",
//				ContentText: nil,
//			}}},
//		},
//		Coins: Coins{
//			Cp: Cp{Value: 0},
//			Sp: Sp{Value: 0},
//			Ep: Ep{Value: 0},
//			Gp: Gp{Value: 0},
//			Pp: Pp{Value: 0}},
//		Resources:     Resources{},
//		BonusesSkills: BonusesSkills{},
//		BonusesStats:  BonusesStats{},
//		Conditions:    nil,
//		ID:            "",
//		Avatar: Avatar{
//			Jpeg: "",
//			Webp: "",
//		},
//		Inspiration: false,
//	}
//
//	json.Marshal(exp)
//	//fmt.Println(string(js))
//}
