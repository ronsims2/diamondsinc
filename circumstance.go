package main

import (
	"math/rand"
)

//Circumstance - the data model for the game chance engine
type Circumstance struct {
	influence float64
	goal      int
}

//GenerateCircumstances is the generator for the circumstances
func GenerateCircumstances(rounds int) []Circumstance {
	chances := []Circumstance{}

	for i := 0; i < rounds; i++ {
		rand.Seed(int64(i))
		flip := rand.Intn(rounds)
		goal := rand.Intn(1000)

		infl := rand.Float64() + 0.7

		if infl < 1 {
			infl = 1
		}

		if flip%4 == 0 {
			infl = infl * -1
		} else if flip%6 == 0 {
			infl = infl * -2
		} else if flip%10 == 0 {
			infl = infl * -3
		}

		chance := Circumstance{
			influence: infl,
			goal:      goal,
		}

		chances = append(chances, chance)
	}

	return chances
}
