package utils

import (
	"math/rand"
	"simulator/src/constants"
	"time"
)

func GetRandomDuration(lowerBound, upperBound float64) time.Duration {
	lowerBound *= constants.SecondMultiplier
	upperBound *= constants.SecondMultiplier

	delta := upperBound - lowerBound
	r := lowerBound + rand.Float64()*(delta)
	return time.Duration(r * float64(time.Second))
}

func DurationToVirtualSeconds(d time.Duration) float64 {
	return d.Seconds() / constants.SecondMultiplier
}
