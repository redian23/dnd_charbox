package skills

type Skills struct {
	Acrobatics     skill
	AnimalHandling skill
	Arcana         skill
	Athletics      skill
	Deception      skill
	History        skill
	Insight        skill
	Intimidation   skill
	Investigation  skill
	Medicine       skill
	Nature         skill
	Perception     skill
	Performance    skill
	Persuasion     skill
	Religion       skill
	SleightOfHand  skill
	Stealth        skill
	Survival       skill
}
type skill struct {
	Point          int
	Predisposition string
	Mastery        bool
}
