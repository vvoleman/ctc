package main

import (
	"fmt"
	"simulator/src/constants"
	"simulator/src/station"
	"simulator/src/statistics"
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
	fmt.Printf("===== Simulation finished in %s =====\n", endTime.Sub(startTime))

	statistics.Run(vehicles)
}
