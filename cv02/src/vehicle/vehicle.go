package vehicle

import (
	"math/rand"
	"simulator/src/constants"
	"simulator/src/utils"
	"time"
)

type Fuel struct {
	Name       constants.FuelName
	LowerBound float64
	UpperBound float64
}

type Vehicle struct {
	ID               int
	Fuel             Fuel
	QueuedAt         time.Time
	StartFuelAt      time.Time
	EndFuelAt        time.Time
	RegisterJoinedAt time.Time
	RegisterStartAt  time.Time
	LeftAt           time.Time
}

func (f *Fuel) getDuration() time.Duration {
	return utils.GetRandomDuration(f.LowerBound, f.UpperBound)
}

func (v *Vehicle) StartRefueling() {
	dur := v.Fuel.getDuration()
	time.Sleep(dur)
}

func InitVehicles(count int) []*Vehicle {
	var vehicles []*Vehicle
	for i := 0; i < count; i++ {
		fuel := getRandomFuelType()
		vehicles = append(vehicles, &Vehicle{
			ID:   i,
			Fuel: fuel,
		})
	}

	return vehicles
}

func getRandomFuelType() Fuel {
	// get number between 0 and 1
	r := rand.Float64()
	gas := constants.GasChance
	diesel := constants.DieselChance
	lpg := constants.LpgChance

	if r < gas {
		return Fuel{
			Name:       constants.Gas,
			LowerBound: constants.GasLowerBoundTime,
			UpperBound: constants.GasUpperBoundTime,
		}
	}
	r -= gas

	if r < diesel {
		return Fuel{
			Name:       constants.Diesel,
			LowerBound: constants.DieselLowerBoundTime,
			UpperBound: constants.DieselUpperBoundTime,
		}
	}
	r -= diesel

	if r < lpg {
		return Fuel{
			Name:       constants.LPG,
			LowerBound: constants.LpgLowerBoundTime,
			UpperBound: constants.LpgUpperBoundTime,
		}
	}

	return Fuel{
		Name:       constants.Electric,
		LowerBound: constants.ElectricLowerBoundTime,
		UpperBound: constants.ElectricUpperBoundTime,
	}
}
