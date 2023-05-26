package classes

type ClassJson struct {
	Classes []Class `json:"classes"`
}
type Class struct {
	ClassName      string `json:"class_name"`
	Description    string `json:"description"`
	Strength       int    `json:"strength"`
	Dexterity      int    `json:"dexterity"`
	BodyDifficulty int    `json:"body_difficulty"`
	Intelligence   int    `json:"intelligence"`
	Wisdom         int    `json:"wisdom"`
	Charisma       int    `json:"charisma"`
}

// start point for classes
var (
	Mage = Class{ClassName: "Mage",
		Strength:       10,
		Dexterity:      8,
		BodyDifficulty: 13,
		Intelligence:   15,
		Wisdom:         14,
		Charisma:       12,
	}

	Knight = Class{ClassName: "Knight",
		Strength:       10,
		Dexterity:      8,
		BodyDifficulty: 13,
		Intelligence:   15,
		Wisdom:         14,
		Charisma:       12,
	}
	Thief = Class{ClassName: "Thief",
		Strength:       10,
		Dexterity:      8,
		BodyDifficulty: 13,
		Intelligence:   15,
		Wisdom:         14,
		Charisma:       12,
	}
	Warrior = Class{ClassName: "Warrior",
		Strength:       10,
		Dexterity:      8,
		BodyDifficulty: 13,
		Intelligence:   15,
		Wisdom:         14,
		Charisma:       12,
	}

	Lucky = Class{ClassName: "Lucky",
		Strength:       RandomRollPoints(),
		Dexterity:      RandomRollPoints(),
		BodyDifficulty: RandomRollPoints(),
		Intelligence:   RandomRollPoints(),
		Wisdom:         RandomRollPoints(),
		Charisma:       RandomRollPoints(),
	}

	ClassArray     = []Class{Mage, Knight, Thief, Warrior, Lucky}
	ClassArrayName = []string{Mage.ClassName, Knight.ClassName, Thief.ClassName,
		Warrior.ClassName, Lucky.ClassName}
)

func (c *Class) ModifyDexterity(value int) {
	c.Dexterity = value
}
func (c *Class) ModifyBodyDifficulty(value int) {
	c.BodyDifficulty = value
}
func (c *Class) ModifyIntelligence(value int) {
	c.Intelligence = value
}
func (c *Class) ModifyWisdom(value int) {
	c.Wisdom = value
}
func (c *Class) ModifyCharisma(value int) {
	c.Charisma = value
}
func (c *Class) ModifyStrength(value int) {
	c.Strength = value
}
