package dice

import (
	"github.com/mazen160/go-random"
)

type Dice struct {
	Name     string
	MinRange int
	MaxRange int
}

var (
	D4   = Dice{Name: "D4", MinRange: 1, MaxRange: 4}
	D6   = Dice{Name: "D6", MinRange: 1, MaxRange: 6}
	D8   = Dice{Name: "D8", MinRange: 1, MaxRange: 8}
	D10  = Dice{Name: "D10", MinRange: 1, MaxRange: 10}
	D12  = Dice{Name: "D12", MinRange: 1, MaxRange: 12}
	D20  = Dice{Name: "D20", MinRange: 1, MaxRange: 20}
	D100 = Dice{Name: "D100", MinRange: 1, MaxRange: 100}
)

func (d Dice) RollDice() int {
	count, _ := random.IntRange(d.MinRange, d.MaxRange+1)
	return count
}

func (d Dice) GetDiceName() string {
	return d.Name
}

func (d Dice) GetMinRange() int {
	return d.MinRange
}

func (d Dice) GetMaxRange() int {
	return d.MaxRange
}
