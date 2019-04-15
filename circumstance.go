package main

import (
	"math/rand"
)

//Circumstance - the data model for the game chance engine
type Circumstance struct {
	influence   float64
	description string
}

//GenerateCircumstances is the generator for the circumstances
func GenerateCircumstances() []Circumstance {
	chances := []Circumstance{}

	for i := 0; i < 100; i++ {
		rand.Seed(int64(i))
		flip := rand.Intn(100)

		infl := rand.Float64() + 0.7
		desc := "Pretty good sales today."

		if flip%4 == 0 {
			infl = infl * -1
			desc = "Sales where a little soft today."
		} else if flip%6 == 0 {
			infl = infl * -2
			desc = "The weather was terrible today, if this keeps up, we will be in trouble."
		} else if flip%10 == 0 {
			infl = infl * -3
			desc = "Someone's negative post went viral on social media, sales were disasterous!"
		}

		chance := Circumstance{
			influence:   infl,
			description: desc,
		}

		chances = append(chances, chance)
	}

	return chances
}
