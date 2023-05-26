package classes

import "github.com/mazen160/go-random"

func CheckValidClass() bool {
	var sum int
	for _, class := range ClassArray {
		sum = class.BodyDifficulty + class.Dexterity + class.Intelligence +
			class.Charisma + class.Strength + class.Wisdom
		if sum > 72 {
			if class.ClassName == "Lucky" {
				for true {
					if sum < 67 || sum > 78 {
						Lucky = RerollLuckyClassStats()
						sum = Lucky.BodyDifficulty + Lucky.Dexterity + Lucky.Intelligence +
							Lucky.Charisma + Lucky.Strength + Lucky.Wisdom
						//fmt.Println("Перерол на", sum)
					} else {
						break
					}
				}
			}
		}
		//fmt.Println(class.ClassName, sum)
	}
	return true
}

func GetRandomClass() Class {
	return ClassArray[random.GetIntInsecure(5)]
}
