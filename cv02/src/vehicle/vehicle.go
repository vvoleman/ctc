package vehicle

import (
	"math/rand"
	"simulator/src/constants"
	"time"
)

type Fuel struct {
	name       string
	LowerBound float64
	UpperBound float64
}

type Vehicle struct {
	ID        int
	Fuel      Fuel
	QueuedAt  time.Time
	StartedAt time.Time
	LeftAt    time.Time
}

func (f *Fuel) getDuration() time.Duration {
	// generate random duration from lowerBound to upperBound
	delta := f.UpperBound - f.LowerBound
	r := f.LowerBound + rand.Float64()*(delta)*constants.SecondMultiplier
	duration := time.Duration(r)
	return duration
}

func GetVehicle(id int, fuel Fuel) Vehicle {
	return Vehicle{ID: id, Fuel: fuel}
}
