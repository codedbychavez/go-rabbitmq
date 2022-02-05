package random

import (
	"math/rand"
	"time"
)

func GetRandomWord() string {
	randomWords := [5]string {
		"books",
		"pens",
		"mouse",
		"laptop",
		"keyboard",
	}

	// Create a seed -> a value that chnages to produce different numbers on each call
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	// Picks a random word from the list and returns it
	return (randomWords[random.Intn(5)])
}