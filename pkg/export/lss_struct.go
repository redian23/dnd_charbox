package export

type ExportToLss struct {
	JSONType      string        `json:"jsonType"`
	Template      string        `json:"template"`
	Name          Name          `json:"name"`
	HiddenName    string        `json:"hiddenName"`
	Info          Info          `json:"info"`
	SubInfo       SubInfo       `json:"subInfo"`
	SpellsInfo    SpellsInfo    `json:"spellsInfo"`
	Spells        Spells        `json:"spells"`
	Proficiency   int           `json:"proficiency"`
	Stats         Stats         `json:"stats"`
	Saves         Saves         `json:"saves"`
	Skills        Skills        `json:"skills"`
	Vitality      Vitality      `json:"vitality"`
	WeaponsList   []WeaponsList `json:"weaponsList"`
	Weapons       Weapons       `json:"weapons"`
	Text          Text          `json:"text"`
	Coins         Coins         `json:"coins"`
	Resources     Resources     `json:"resources"`
	BonusesSkills BonusesSkills `json:"bonusesSkills"`
	BonusesStats  BonusesStats  `json:"bonusesStats"`
	Conditions    []interface{} `json:"conditions"`
	CasterClass   CasterClass   `json:"casterClass"`
	Avatar        Avatar        `json:"avatar"`
}
type Avatar struct {
	Jpeg string `json:"jpeg"`
	Webp string `json:"webp"`
}

type Name struct {
	Value string `json:"value"`
}
type CharClass struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Level struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value int    `json:"value"`
}
type BackgroundInfo struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type PlayerName struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Race struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Alignment struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Experience struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Info struct {
	CharClass  CharClass      `json:"charClass"`
	Level      Level          `json:"level"`
	Background BackgroundInfo `json:"background"`
	PlayerName PlayerName     `json:"playerName"`
	Race       Race           `json:"race"`
	Alignment  Alignment      `json:"alignment"`
	Experience Experience     `json:"experience"`
}
type Age struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Height struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Weight struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Eyes struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Skin struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Hair struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type SubInfo struct {
	Age    Age    `json:"age"`
	Height Height `json:"height"`
	Weight Weight `json:"weight"`
	Eyes   Eyes   `json:"eyes"`
	Skin   Skin   `json:"skin"`
	Hair   Hair   `json:"hair"`
}
type Base struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Save struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type Mod struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type SpellsInfo struct {
	Base Base `json:"base"`
	Save Save `json:"save"`
	Mod  Mod  `json:"mod"`
}
type Slots1 struct {
	Value string `json:"value"`
}
type Spell10 struct {
	IsReady bool `json:"isReady"`
}
type Spell11 struct {
	IsReady bool `json:"isReady"`
}
type Spells struct {
	Slots1  Slots1  `json:"slots-1"`
	Spell10 Spell10 `json:"spell-1-0"`
	Spell11 Spell11 `json:"spell-1-1"`
}
type StrStat struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}
type DexStat struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}
type ConStat struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}
type IntStat struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}
type WisStat struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}
type ChaStat struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}
type Stats struct {
	Str StrStat `json:"str"`
	Dex DexStat `json:"dex"`
	Con ConStat `json:"con"`
	Int IntStat `json:"int"`
	Wis WisStat `json:"wis"`
	Cha ChaStat `json:"cha"`
}
type Str struct {
	Name   string `json:"name"`
	IsProf bool   `json:"isProf"`
}
type Dex struct {
	Name   string `json:"name"`
	IsProf bool   `json:"isProf"`
}
type Con struct {
	Name   string `json:"name"`
	IsProf bool   `json:"isProf"`
}
type Int struct {
	Name   string `json:"name"`
	IsProf bool   `json:"isProf"`
}
type Wis struct {
	Name   string `json:"name"`
	IsProf bool   `json:"isProf"`
}
type Cha struct {
	Name   string `json:"name"`
	IsProf bool   `json:"isProf"`
}
type Saves struct {
	Str Str `json:"str"`
	Dex Dex `json:"dex"`
	Con Con `json:"con"`
	Int Int `json:"int"`
	Wis Wis `json:"wis"`
	Cha Cha `json:"cha"`
}
type skill struct {
	BaseStat string      `json:"baseSta0t"`
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	IsProf   interface{} `json:"isProf"`
}

type Skills struct {
	Acrobatics     skill `json:"acrobatics"`
	Investigation  skill `json:"investigation"`
	Athletics      skill `json:"athletics"`
	Perception     skill `json:"perception"`
	Survival       skill `json:"survival"`
	Performance    skill `json:"performance"`
	Intimidation   skill `json:"intimidation"`
	History        skill `json:"history"`
	SleightOfHand  skill `json:"sleight of hand"`
	Arcana         skill `json:"arcana"`
	Medicine       skill `json:"medicine"`
	Deception      skill `json:"deception"`
	Nature         skill `json:"nature"`
	Insight        skill `json:"insight"`
	Religion       skill `json:"religion"`
	Stealth        skill `json:"stealth"`
	Persuasion     skill `json:"persuasion"`
	AnimalHandling skill `json:"animal handling"`
}
type HpDiceCurrent struct {
	Value int `json:"value"`
}
type HpDiceMulti struct {
}
type Speed struct {
	Value string `json:"value"`
}
type HpMax struct {
	Value string `json:"value"`
}
type Ac struct {
	Value string `json:"value"`
}
type HpCurrent struct {
	Value string `json:"value"`
}
type HitDie struct {
	Value string `json:"value"`
}
type Vitality struct {
	HpDiceCurrent HpDiceCurrent `json:"hp-dice-current"`
	HitDie        HitDie        `json:"hit-die"`
	HpCurrent     HpCurrent     `json:"hp-current"`
	Speed         Speed         `json:"speed"`
	HpMax         HpMax         `json:"hp-max"`
	Ac            Ac            `json:"ac"`
}
type Modif struct {
	Value string `json:"value"`
}
type Dmg struct {
	Value string `json:"value"`
}

type WeaponName0 struct {
	Value string `json:"value"`
}
type WeaponMod0 struct {
	Value string `json:"value"`
}
type WeaponDmg0 struct {
	Value string `json:"value"`
}
type WeaponName1 struct {
	Value string `json:"value"`
}
type WeaponMod1 struct {
	Value string `json:"value"`
}
type WeaponDmg1 struct {
	Value string `json:"value"`
}
type Weapons struct {
	WeaponName0 WeaponName0 `json:"weaponName-0"`
	WeaponMod0  WeaponMod0  `json:"weaponMod-0"`
	WeaponDmg0  WeaponDmg0  `json:"weaponDmg-0"`
	WeaponName1 WeaponName1 `json:"weaponName-1"`
	WeaponMod1  WeaponMod1  `json:"weaponMod-1"`
	WeaponDmg1  WeaponDmg1  `json:"weaponDmg-1"`
}
type WeaponsList struct {
	ID   string `json:"id"`
	Name Name   `json:"name"`
	Mod  Modif  `json:"mod"`
	Dmg  Dmg    `json:"dmg"`
}

type Value struct {
	Data string `json:"data"`
}
type Attacks struct {
	Value Value `json:"value"`
}
type Equipment struct {
	Value Value `json:"value"`
}
type Prof struct {
	Value Value `json:"value"`
}
type Traits struct {
	Value Value `json:"value"`
}
type Allies struct {
	Value Value `json:"value"`
}
type Features struct {
	Value Value `json:"value"`
}
type Personality struct {
	Value Value `json:"value"`
}
type Ideals struct {
	Value Value `json:"value"`
}
type Bonds struct {
	Value Value `json:"value"`
}
type Flaws struct {
	Value Value `json:"value"`
}
type Background struct {
	Value Value `json:"value"`
}
type Quests struct {
	Value Value `json:"value"`
	Size  int   `json:"size"`
}
type SpellsLevel1 struct {
	Value Value `json:"value"`
}
type SpellsLevel0 struct {
	Value Value `json:"value"`
}
type Text struct {
	Attacks      Attacks      `json:"attacks"`
	Equipment    Equipment    `json:"equipment"`
	Prof         Prof         `json:"prof"`
	Traits       Traits       `json:"traits"`
	Allies       Allies       `json:"allies"`
	Features     Features     `json:"features"`
	Personality  Personality  `json:"personality"`
	Ideals       Ideals       `json:"ideals"`
	Bonds        Bonds        `json:"bonds"`
	Flaws        Flaws        `json:"flaws"`
	Background   Background   `json:"background"`
	Quests       Quests       `json:"quests"`
	SpellsLevel1 SpellsLevel1 `json:"spells-level-1"`
	SpellsLevel0 SpellsLevel0 `json:"spells-level-0"`
}

type Gp struct {
	Value int `json:"value"`
}
type Total struct {
	Value float64 `json:"value"`
}
type Sp struct {
	Value int `json:"value"`
}
type Cp struct {
	Value int `json:"value"`
}
type Pp struct {
	Value int `json:"value"`
}
type Ep struct {
	Value int `json:"value"`
}
type Coins struct {
	Gp    Gp    `json:"gp"`
	Total Total `json:"total"`
	Sp    Sp    `json:"sp"`
	Cp    Cp    `json:"cp"`
	Pp    Pp    `json:"pp"`
	Ep    Ep    `json:"ep"`
}
type Resources struct {
}
type BonusesSkills struct {
}
type BonusesStats struct {
}
type Ability struct {
	BaseStat string      `json:"baseStat"`
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	IsProf   interface{} `json:"isProf"`
}

type Abilities struct {
	Acrobatics     Ability `json:"acrobatics"`
	Investigation  Ability `json:"investigation"`
	Athletics      Ability `json:"athletics"`
	Perception     Ability `json:"perception"`
	Survival       Ability `json:"survival"`
	Performance    Ability `json:"performance"`
	Intimidation   Ability `json:"intimidation"`
	History        Ability `json:"history"`
	SleightOfHand  Ability `json:"sleight of hand"`
	Arcana         Ability `json:"arcana"`
	Medicine       Ability `json:"medicine"`
	Deception      Ability `json:"deception"`
	Nature         Ability `json:"nature"`
	Insight        Ability `json:"insight"`
	Religion       Ability `json:"religion"`
	Stealth        Ability `json:"stealth"`
	Persuasion     Ability `json:"persuasion"`
	AnimalHandling Ability `json:"animal handling"`
}
type CasterClass struct {
	Value string `json:"value"`
}
