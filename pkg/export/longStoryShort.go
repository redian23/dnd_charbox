package export

//
//type ExportLSS struct {
//	JSONType      string        `json:"jsonType"`
//	Template      string        `json:"template"`
//	Name          Name          `json:"name"`
//	HiddenName    string        `json:"hiddenName"`
//	Info          Info          `json:"info"`
//	SubInfo       SubInfo       `json:"subInfo"`
//	SpellsInfo    SpellsInfo    `json:"spellsInfo"`
//	Spells        Spells        `json:"spells"`
//	Proficiency   int           `json:"proficiency"`
//	Stats         Stats         `json:"stats"`
//	Saves         Saves         `json:"saves"`
//	Skills        Skills        `json:"skills"`
//	Vitality      Vitality      `json:"vitality"`
//	WeaponsList   []WeaponsList `json:"weaponsList"`
//	Weapons       Weapons       `json:"weapons"`
//	Text          Text          `json:"text"`
//	Coins         Coins         `json:"coins"`
//	Resources     Resources     `json:"resources"`
//	BonusesSkills BonusesSkills `json:"bonusesSkills"`
//	BonusesStats  BonusesStats  `json:"bonusesStats"`
//	Conditions    []interface{} `json:"conditions"`
//	ID            string        `json:"id"`
//	Avatar        Avatar        `json:"avatar"`
//	Inspiration   bool          `json:"inspiration"`
//}
//type Name struct {
//	Value string `json:"value"`
//}
//type CharClass struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Level struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value int    `json:"value"`
//}
//type Background struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type PlayerName struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Race struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Alignment struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Experience struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value int    `json:"value"`
//}
//type Info struct {
//	CharClass  CharClass  `json:"charClass"`
//	Level      Level      `json:"level"`
//	Background Background `json:"background"`
//	PlayerName PlayerName `json:"playerName"`
//	Race       Race       `json:"race"`
//	Alignment  Alignment  `json:"alignment"`
//	Experience Experience `json:"experience"`
//}
//type Age struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value int    `json:"value"`
//}
//type Height struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Weight struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Eyes struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Skin struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Hair struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type SubInfo struct {
//	Age    Age    `json:"age"`
//	Height Height `json:"height"`
//	Weight Weight `json:"weight"`
//	Eyes   Eyes   `json:"eyes"`
//	Skin   Skin   `json:"skin"`
//	Hair   Hair   `json:"hair"`
//}
//type Base struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Save struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type Mod struct {
//	Name  string `json:"name"`
//	Label string `json:"label"`
//	Value string `json:"value"`
//}
//type SpellsInfo struct {
//	Base Base `json:"base"`
//	Save Save `json:"save"`
//	Mod  Mod  `json:"mod"`
//}
//type Spells struct {
//}
//type Str struct {
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	Score    int    `json:"score"`
//	Modifier int    `json:"modifier"`
//}
//type Dex struct {
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	Score    int    `json:"score"`
//	Modifier int    `json:"modifier"`
//}
//type Con struct {
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	Score    int    `json:"score"`
//	Modifier int    `json:"modifier"`
//}
//type Int struct {
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	Score    int    `json:"score"`
//	Modifier int    `json:"modifier"`
//}
//type Wis struct {
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	Score    int    `json:"score"`
//	Modifier int    `json:"modifier"`
//}
//type Cha struct {
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	Score    int    `json:"score"`
//	Modifier int    `json:"modifier"`
//}
//type Stats struct {
//	Str Str `json:"str"`
//	Dex Dex `json:"dex"`
//	Con Con `json:"con"`
//	Int Int `json:"int"`
//	Wis Wis `json:"wis"`
//	Cha Cha `json:"cha"`
//}
//type SaveStr struct {
//	Name   string `json:"name"`
//	IsProf bool   `json:"isProf"`
//}
//type SaveDex struct {
//	Name   string `json:"name"`
//	IsProf bool   `json:"isProf"`
//}
//type SaveCon struct {
//	Name   string `json:"name"`
//	IsProf bool   `json:"isProf"`
//}
//type SaveInt struct {
//	Name   string `json:"name"`
//	IsProf bool   `json:"isProf"`
//}
//type SaveWis struct {
//	Name   string `json:"name"`
//	IsProf bool   `json:"isProf"`
//}
//type SaveCha struct {
//	Name   string `json:"name"`
//	IsProf bool   `json:"isProf"`
//}
//type Saves struct {
//	Str SaveStr `json:"str"`
//	Dex SaveDex `json:"dex"`
//	Con SaveCon `json:"con"`
//	Int SaveInt `json:"int"`
//	Wis SaveWis `json:"wis"`
//	Cha SaveCha `json:"cha"`
//}
//type Acrobatics struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Investigation struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Athletics struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Perception struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Survival struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Performance struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Intimidation struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type History struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type SleightOfHand struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Arcana struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Medicine struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Deception struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Nature struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Insight struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Religion struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Stealth struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Persuasion struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type AnimalHandling struct {
//	BaseStat string `json:"baseStat"`
//	Name     string `json:"name"`
//	Label    string `json:"label"`
//	IsProf   int    `json:"isProf"`
//}
//type Skills struct {
//	Acrobatics     Acrobatics     `json:"acrobatics"`
//	Investigation  Investigation  `json:"investigation"`
//	Athletics      Athletics      `json:"athletics"`
//	Perception     Perception     `json:"perception"`
//	Survival       Survival       `json:"survival"`
//	Performance    Performance    `json:"performance"`
//	Intimidation   Intimidation   `json:"intimidation"`
//	History        History        `json:"history"`
//	SleightOfHand  SleightOfHand  `json:"sleight of hand"`
//	Arcana         Arcana         `json:"arcana"`
//	Medicine       Medicine       `json:"medicine"`
//	Deception      Deception      `json:"deception"`
//	Nature         Nature         `json:"nature"`
//	Insight        Insight        `json:"insight"`
//	Religion       Religion       `json:"religion"`
//	Stealth        Stealth        `json:"stealth"`
//	Persuasion     Persuasion     `json:"persuasion"`
//	AnimalHandling AnimalHandling `json:"animal handling"`
//}
//type HpDiceCurrent struct {
//	Value int `json:"value"`
//}
//type HpDiceMulti struct {
//	Value int `json:"value"`
//}
//type Initiative struct {
//	Value interface{} `json:"value"`
//}
//type Speed struct {
//	Value string `json:"value"`
//}
//type Ac struct {
//	Value int `json:"value"`
//}
//type HpMax struct {
//	Value int `json:"value"`
//}
//type HpCurrent struct {
//	Value int `json:"value"`
//}
//type HpTemp struct {
//	Value int `json:"value"`
//}
//type Vitality struct {
//	HpDiceCurrent  HpDiceCurrent `json:"hp-dice-current"`
//	HpDiceMulti    HpDiceMulti   `json:"hp-dice-multi"`
//	Initiative     Initiative    `json:"initiative"`
//	Speed          Speed         `json:"speed"`
//	Ac             Ac            `json:"ac"`
//	HpMax          HpMax         `json:"hp-max"`
//	HpCurrent      HpCurrent     `json:"hp-current"`
//	IsDying        bool          `json:"isDying"`
//	DeathFails     int           `json:"deathFails"`
//	DeathSuccesses int           `json:"deathSuccesses"`
//	HpTemp         HpTemp        `json:"hp-temp"`
//}
//type Modif struct {
//	Value string `json:"value"`
//}
//type Dmg struct {
//	Value string `json:"value"`
//}
//type ModBonus struct {
//	Value int `json:"value"`
//}
//type WeaponsList struct {
//	ID       string   `json:"id"`
//	Name     Name     `json:"name"`
//	Modif    Modif    `json:"mod"`
//	Dmg      Dmg      `json:"dmg"`
//	Ability  string   `json:"ability"`
//	ModBonus ModBonus `json:"modBonus"`
//}
//type Weapons struct {
//}
//type Attrs struct {
//	Href   string `json:"href"`
//	Target string `json:"target"`
//	Class  string `json:"class"`
//}
//type Marks struct {
//	Type  string `json:"type"`
//	Attrs Attrs  `json:"attrs,omitempty"`
//}
//type Content struct {
//	Type  string  `json:"type"`
//	Marks []Marks `json:"marks"`
//	Text  string  `json:"text"`
//}
//type ContentText struct {
//	Type    string    `json:"type"`
//	Content []Content `json:"content"`
//}
//type Data struct {
//	Type        string        `json:"type"`
//	ContentText []ContentText `json:"content"`
//}
//type Value struct {
//	Data Data `json:"data"`
//}
//type Traits struct {
//	Value Value `json:"value"`
//}
//type Equipment struct {
//	Value Value `json:"value"`
//}
//type Prof struct {
//	Value Value `json:"value"`
//}
//type Attacks struct {
//	Value Value `json:"value"`
//	Size  int   `json:"size"`
//}
//type Allies struct {
//	Value Value `json:"value"`
//}
//type BackgroundFullText struct {
//	Value Value `json:"value"`
//	Size  int   `json:"size"`
//}
//type Features struct {
//	Value Value `json:"value"`
//	Size  int   `json:"size"`
//}
//type Quests struct {
//	Value Value `json:"value"`
//}
//type Flaws struct {
//	Value Value `json:"value"`
//	Size  int   `json:"size"`
//}
//type Personality struct {
//	Value Value `json:"value"`
//}
//type Bonds struct {
//	Value Value `json:"value"`
//}
//type Ideals struct {
//	Value Value `json:"value"`
//}
//type Items struct {
//	Value Value `json:"value"`
//}
//type Text struct {
//	Traits      Traits             `json:"traits"`
//	Equipment   Equipment          `json:"equipment"`
//	Prof        Prof               `json:"prof"`
//	Attacks     Attacks            `json:"attacks"`
//	Allies      Allies             `json:"allies"`
//	Background  BackgroundFullText `json:"background"`
//	Features    Features           `json:"features"`
//	Quests      Quests             `json:"quests"`
//	Flaws       Flaws              `json:"flaws"`
//	Personality Personality        `json:"personality"`
//	Bonds       Bonds              `json:"bonds"`
//	Ideals      Ideals             `json:"ideals"`
//	Items       Items              `json:"items"`
//}
//
//type Cp struct {
//	Value int `json:"value"`
//}
//type Sp struct {
//	Value int `json:"value"`
//}
//type Ep struct {
//	Value int `json:"value"`
//}
//type Gp struct {
//	Value int `json:"value"`
//}
//type Pp struct {
//	Value int `json:"value"`
//}
//type Coins struct {
//	Cp Cp `json:"cp"`
//	Sp Sp `json:"sp"`
//	Ep Ep `json:"ep"`
//	Gp Gp `json:"gp"`
//	Pp Pp `json:"pp"`
//}
//type Resources struct {
//}
//type BonusesSkills struct {
//}
//type BonusesStats struct {
//}
//type Avatar struct {
//	Jpeg string `json:"jpeg"`
//	Webp string `json:"webp"`
//}
