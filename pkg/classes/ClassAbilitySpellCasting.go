package classes

var CharacterLevelGlobal int
var ProficiencyBonus int

func GetClassSpellBasicCharacteristic() Spellcasting {
	var cast Spellcasting

	switch ClassNameGlobalRu {
	case "Бард":
		cast.BasicSpellCharacteristics = "Харизма"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Charisma
		cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Charisma
		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 4
			cast.TotalSpellCount = 4
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 5
			cast.TotalSpellCount = 5
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 6
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 7
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 8
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 9
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 10
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 11
		}
	case "Волшебник":
		cast.BasicSpellCharacteristics = "Интеллект"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Intelligence
		cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Intelligence

		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 6
			cast.TotalSpellCount = 6
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 7
			cast.TotalSpellCount = 7
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 8
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 9
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 10
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 11
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 12
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 7
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 13
		}
	case "Друид":
		cast.BasicSpellCharacteristics = "Мудрость"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Wisdom
		cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Wisdom

		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 4
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 6
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		}
	case "Жрец":
		cast.BasicSpellCharacteristics = "Мудрость"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Wisdom
		cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Wisdom

		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 4
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 5
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = (ModifierGlobal.Wisdom + CharacterLevelGlobal) - 6
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Wisdom + CharacterLevelGlobal
		}
	case "Изобретатель":
		cast.BasicSpellCharacteristics = "Интеллект"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Intelligence
		cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Intelligence

		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2) - 3
			cast.TwoLevelSpellsKnownCount = 3
			cast.TotalSpellCount = ModifierGlobal.Intelligence + (CharacterLevelGlobal / 2)
		}
	case "Колдун":
		cast.BasicSpellCharacteristics = "Харизма"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Charisma
		cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Charisma

		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 1, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 2
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 2
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 4
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 5
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 6
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 6
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 7
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 2, TreeLevel: 2, FourLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 3
			cast.ZeroLevelSpellsKnownCount = 3
			cast.OneLevelSpellsKnownCount = 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 8
		}
	case "Паладин":
		if CharacterLevelGlobal > 1 {
			cast.BasicSpellCharacteristics = "Харизма"
			cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Charisma
			cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Charisma
		}
		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 0, TwoLevel: 0}

			cast.TotalSpellCount = 0
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}

			cast.OneLevelSpellsKnownCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
			cast.TotalSpellCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}

			cast.OneLevelSpellsKnownCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
			cast.TotalSpellCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}

			cast.OneLevelSpellsKnownCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
			cast.TotalSpellCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2) - 1
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2) - 2
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2) - 3
			cast.TwoLevelSpellsKnownCount = 3
			cast.TotalSpellCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2) - 3
			cast.TwoLevelSpellsKnownCount = 3
			cast.TotalSpellCount = ModifierGlobal.Charisma + (CharacterLevelGlobal / 2)
		}
	case "Следопыт":
		if CharacterLevelGlobal > 1 {
			cast.BasicSpellCharacteristics = "Мудрость"
			cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Wisdom
			cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Wisdom
		}
		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 0, TwoLevel: 0}
			cast.TotalSpellCount = 0
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.OneLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 2
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 4
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}

			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 4
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 5
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 5
		}
	case "Чародей":
		cast.BasicSpellCharacteristics = "Харизма"
		cast.SavingThrowDifficulty = 8 + ProficiencyBonus + ModifierGlobal.Charisma
		cast.SpellDamageModifier = ProficiencyBonus + ModifierGlobal.Charisma
		switch CharacterLevelGlobal {
		case 1:
			cast.SpellsSlots = spellsSlots{OneLevel: 2, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 2
		case 2:
			cast.SpellsSlots = spellsSlots{OneLevel: 3, TwoLevel: 0}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 3
			cast.TotalSpellCount = 3
		case 3:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 4
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 4
		case 4:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 5
		case 5:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 2}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 6
		case 6:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.TotalSpellCount = 7
		case 7:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 1}
			cast.ZeroLevelSpellsKnownCount = 5
			cast.OneLevelSpellsKnownCount = 3
			cast.TwoLevelSpellsKnownCount = 2
			cast.TreeLevelSpellsKnownCount = 2
			cast.FourLevelSpellsKnownCount = 1
			cast.TotalSpellCount = 8
		case 8:
			cast.SpellsSlots = spellsSlots{OneLevel: 4, TwoLevel: 3, TreeLevel: 3, FourLevel: 2}
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
