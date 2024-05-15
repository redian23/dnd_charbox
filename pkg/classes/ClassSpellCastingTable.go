package classes

var ProficiencyBonus int

func GetClassSpellBasicCharacteristic(ClassName string, ClassLevel int, modifier AbilityModifier) SpellCasting {
	var cast SpellCasting

	if ClassLevel <= 4 {
		ProficiencyBonus = 2
	} else if ClassLevel > 4 && ClassLevel <= 8 {
		ProficiencyBonus = 3
	} else {
		ProficiencyBonus = 4
	}

	switch ClassName {
	case "Бард":
		cast.BasicSpellCharacteristics = "Харизма"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Charisma
		cast.SpellDamageModifier = ProficiencyBonus + modifier.Charisma
		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 4
			cast.TotalSpellCount = 4
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 5
			cast.TotalSpellCount = 5
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 6
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 7
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 8
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 9
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 10
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 11
		}
	case "Волшебник":
		cast.BasicSpellCharacteristics = "Интеллект"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Intelligence
		cast.SpellDamageModifier = ProficiencyBonus + modifier.Intelligence

		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 6
			cast.TotalSpellCount = 6
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 7
			cast.TotalSpellCount = 7
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 8
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 9
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 10
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 11
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 12
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 13
		}
	case "Друид":
		cast.BasicSpellCharacteristics = "Мудрость"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Wisdom
		cast.SpellDamageModifier = ProficiencyBonus + modifier.Wisdom

		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Wisdom + ClassLevel
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Wisdom + ClassLevel
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 4
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 6
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		}
	case "Жрец":
		cast.BasicSpellCharacteristics = "Мудрость"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Wisdom
		cast.SpellDamageModifier = ProficiencyBonus + modifier.Wisdom

		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Wisdom + ClassLevel
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Wisdom + ClassLevel
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 4
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (modifier.Wisdom + ClassLevel) - 6
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Wisdom + ClassLevel
		}
	case "Изобретатель":
		cast.BasicSpellCharacteristics = "Интеллект"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Intelligence
		cast.SpellDamageModifier = ProficiencyBonus + modifier.Intelligence

		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2)
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2)
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2)
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2)
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = modifier.Intelligence + (ClassLevel / 2) - 3
			cast.TwoLevelSpellsKnownCount = 3
			cast.TotalSpellCount = modifier.Intelligence + (ClassLevel / 2)
		}
	case "Колдун":
		cast.BasicSpellCharacteristics = "Харизма"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Charisma
		cast.SpellDamageModifier = ProficiencyBonus + modifier.Charisma

		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 1, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 2
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 4
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 5
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 6
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 6
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 7
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 8
		}
	case "Паладин":
		if ClassLevel > 1 {
			cast.BasicSpellCharacteristics = "Харизма"
			cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Charisma
			cast.SpellDamageModifier = ProficiencyBonus + modifier.Charisma
		}
		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 0, TwoLevel: 0}

			cast.TotalSpellCount = 0
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}

			cast.OneLevelSpellsKnownCount = modifier.Charisma + (ClassLevel / 2)
			cast.TotalSpellCount = modifier.Charisma + (ClassLevel / 2)
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}

			cast.OneLevelSpellsKnownCount = modifier.Charisma + (ClassLevel / 2)
			cast.TotalSpellCount = modifier.Charisma + (ClassLevel / 2)
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}

			cast.OneLevelSpellsKnownCount = modifier.Charisma + (ClassLevel / 2)
			cast.TotalSpellCount = modifier.Charisma + (ClassLevel / 2)
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.OneLevelSpellsKnownCount = modifier.Charisma + (ClassLevel / 2) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = modifier.Charisma + (ClassLevel / 2)
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.OneLevelSpellsKnownCount = modifier.Charisma + (ClassLevel / 2) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = modifier.Charisma + (ClassLevel / 2)
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = modifier.Charisma + (ClassLevel / 2) - 3
			cast.TwoLevelSpellsKnownCount = 3
			cast.TotalSpellCount = modifier.Charisma + (ClassLevel / 2)
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = modifier.Charisma + (ClassLevel / 2) - 3
			cast.TwoLevelSpellsKnownCount = 3
			cast.TotalSpellCount = modifier.Charisma + (ClassLevel / 2)
		}
	case "Следопыт":
		if ClassLevel > 1 {
			cast.BasicSpellCharacteristics = "Мудрость"
			cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Wisdom
			cast.SpellDamageModifier = ProficiencyBonus + modifier.Wisdom
		}
		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 0, TwoLevel: 0}
			cast.TotalSpellCount = 0
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.OneLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 2
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 4
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}

			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 4
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 5
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 5
		}
	case "Чародей":
		cast.BasicSpellCharacteristics = "Харизма"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + modifier.Charisma
		cast.SpellDamageModifier = ProficiencyBonus + modifier.Charisma
		switch ClassLevel {
		case 1:
			cast.SpellsSlots = SpellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 2
		case 2:
			cast.SpellsSlots = SpellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 3:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 4
		case 4:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 5
		case 5:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 6
		case 6:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 7
		case 7:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 8
		case 8:
			cast.SpellsSlots = SpellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 9
		}
	}

	return cast
}
