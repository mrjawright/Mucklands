package dice

import (
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func rollDie(numberOfSides int) int {
	return rand.Intn(numberOfSides) + 1
}

func Roll(numberOfDie int, numberOfSides int, modifier int) int {
	var total int = 0
	for d := 0; d < numberOfDie; d++ {
		total += rollDie(numberOfSides)
	}
	return total + modifier
}

func GetRandomIndex(arrayLen int) int {
	return rand.Intn(arrayLen)
}
