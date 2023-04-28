package main

import (
	"fmt"
	"simulator/src/constants"
	"simulator/src/station"
	"simulator/src/statistics"
	"simulator/src/utils"
	"simulator/src/vehicle"
	"time"
)

func main() {
	vehicles := vehicle.InitVehicles(constants.TotalNumberOfVehicles)
	gs := station.InitStation()

	startTime := time.Now()
	fmt.Printf("Starting simulation for %d vehicles\n", len(vehicles))

	gs.Start(vehicles)

	endTime := time.Now()
	total_time := endTime.Sub(startTime)
	dur := time.Duration(utils.DurationToVirtualSeconds(total_time) * float64(time.Second))
	fmt.Printf("===== Simulation finished in %s =====\n", dur)

	statistics.Run(vehicles)
}
