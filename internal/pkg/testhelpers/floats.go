package testhelpers

import (
	"math/rand"
	"time"
)

func GetRandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}
