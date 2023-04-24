package statistics

import (
	"fmt"
	"math"
	"simulator/src/constants"
	"simulator/src/vehicle"
	"time"
)

func Run(vehicles []*vehicle.Vehicle) {
	// Get count of vehicles per fuel type
	fuelTypeCount := make(map[constants.FuelName]int)
	for _, v := range vehicles {
		fuelTypeCount[v.Fuel.Name]++
	}
	for fuelType, count := range fuelTypeCount {
		fmt.Printf("Vehicles for %s: %d (%.2f%%)\n", fuelType, count, getPercentage(count, len(vehicles)))
	}
	println("............................................")

	// Get average time per fuel type
	fuelTypeTimes := make(map[constants.FuelName][]float64)
	for _, v := range vehicles {
		if v.QueuedAt.IsZero() {
			continue
		}
		fuelType := v.Fuel.Name
		fuelTypeTimes[fuelType] = append(fuelTypeTimes[fuelType], v.EndFuelAt.Sub(v.StartFuelAt).Seconds())
	}

	for fuelType, times := range fuelTypeTimes {
		avg := getAverage(times)
		fmt.Printf("Average time for %s is %fs\n", fuelType, avg.Seconds())
	}
	if len(fuelTypeTimes) > 0 {
		println("..........................................")
	}

	// Get how many vehicles left per fuel type
	leavingVehicles := calculateLeavingVehicles(vehicles)
	if len(leavingVehicles) == 0 {
		fmt.Println("All vehicles were served")
		return
	} else {
		for fuelType, count := range leavingVehicles {
			fmt.Printf("Vehicles unserverd for %s: %d (%.2f%%)\n", fuelType, count, getPercentage(count, fuelTypeCount[fuelType]))
		}
	}
	println("............................................")
}

func getPercentage(count, total int) float64 {
	return math.Round(float64(count)/float64(total)*10000) / 100
}

func getAverage(times []float64) time.Duration {
	var sum float64
	for _, t := range times {
		sum += t
	}

	val := sum / float64(len(times))
	return time.Duration(val*1000) * time.Millisecond
}

func calculateLeavingVehicles(vehicles []*vehicle.Vehicle) map[constants.FuelName]int {
	leavingVehicles := make(map[constants.FuelName]int)
	for _, v := range vehicles {
		if v.QueuedAt.IsZero() {
			leavingVehicles[v.Fuel.Name]++
		}
	}

	return leavingVehicles
}
