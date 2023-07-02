package characterCore

import (
	"fmt"
	"github.com/mazen160/go-random"
)

func GenerateAppearance() Appearance {
	var appearance Appearance

	appearance.Gender = setGender()
	appearance.Age = setAge()
	appearance.Eyes = setEyesColor()
	appearance.Height = setHeight()
	appearance.Weight = setWeight(appearance.Height)
	appearance.Hair = setHairColor()

	return appearance
}

func setGender() string {
	var gender string
	count, _ := random.IntRange(0, 10)
	if count%2 == 0 {
		gender = "Man"
	} else {
		gender = "Woman"
	}
	return gender
}

func setAge() int {
	age, _ := random.IntRange(18, 45)
	return age
}

func setEyesColor() string {
	var colors = []string{"Yellow", "Blue", "Red", "Green", "Brown", "Azure"}

	color, err := random.Choice(colors)
	if err != nil {
		fmt.Println(err)
	}

	return color
}

func setHairColor() string {
	var colors = []string{"Yellow", "Blond", "Black", "Blue", "Red", "Green", "Brown", "Azure"}

	color, err := random.Choice(colors)
	if err != nil {
		fmt.Println(err)
	}

	return color
}

// Высчитывать от значений BodyDifficulty
func setHeight() int {
	height, _ := random.IntRange(150, 210)
	return height
}

func setWeight(height int) int {
	minus, _ := random.IntRange(5, 20)
	weight := (height - 100) - minus

	return weight
}
