package utils

import (
	"math/rand"
	"time"
)

func RandomInRange(min, max int) int {
	// Create a new random source
	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number in the range [min, max]
	return src.Intn(max-min+1) + min
}
