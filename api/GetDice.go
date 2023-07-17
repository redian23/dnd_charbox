package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pregen/backend/dice"
)

type multiDiceRoll struct {
	DiceName string `form:"dice_name"`
	Count    int    `form:"count"`
}

type diceRoll struct {
	RollNumber int `json:"roll_number"`
	Result     int `json:"result"`
}

type rollAnswer struct {
	DiceName  string     `json:"dice_name"`
	RollSteps []diceRoll `json:"roll_steps"`
	Total     int        `json:"total"`
}

func roll(name string) diceRoll {
	var result diceRoll
	switch name {
	case "D4":
		result.Result = dice.D4.RollDice()
	case "D6":
		result.Result = dice.D6.RollDice()
	case "D8":
		result.Result = dice.D8.RollDice()
	case "D10":
		result.Result = dice.D10.RollDice()
	case "D12":
		result.Result = dice.D12.RollDice()
	case "D20":
		result.Result = dice.D20.RollDice()
	case "D100":
		result.Result = dice.D100.RollDice()
	}

	return result
}

func GetMultiRoll(c *gin.Context) {
	var d multiDiceRoll
	c.Bind(&d)
	//ответ в консоль на запрос
	//проверка отправляемых данных
	//c.JSON(200, gin.H{
	//	"Name":  d.DiceName,
	//	"Count": d.Count,
	//})

	if d.DiceName == "D4" || d.DiceName == "D6" || d.DiceName == "D8" || d.DiceName == "D10" ||
		d.DiceName == "D12" || d.DiceName == "D20" || d.DiceName == "D100" {
		//skip
	} else {
		c.Data(http.StatusNotFound, "text/plain; charset=utf-8", []byte("Dice don't exist."+"\n"))
		return
	}
	//max count
	if d.Count > 100 {
		d.Count = 100
	}

	var answ rollAnswer
	var rollStepsArray []diceRoll

	answ.DiceName = d.DiceName

	for i := 0; i < d.Count; i++ {
		tmpDice := roll(d.DiceName)
		tmpDice.RollNumber = i + 1
		answ.Total = answ.Total + tmpDice.Result

		rollStepsArray = append(rollStepsArray, tmpDice)
	}
	answ.RollSteps = rollStepsArray
	c.JSONP(http.StatusOK, answ)
}
