package throws

type SavingThrows struct {
	Strength       savingThrows
	Dexterity      savingThrows
	BodyDifficulty savingThrows
	Intelligence   savingThrows
	Wisdom         savingThrows
	Charisma       savingThrows
}

type savingThrows struct {
	Point   int
	Mastery bool
}
