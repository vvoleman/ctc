package constants

// TotalNumberOfVehicles How many vehicles will be created
const TotalNumberOfVehicles = 1000

const (
	GasChance      = 0.4
	DieselChance   = 0.1
	LpgChance      = 0.1
	ElectricChance = 0.4
)

const (
	RegisterLowerBoundTime = 0.5
	RegisterUpperBoundTime = 2
)

// Number of pumps
const (
	GasPumpCount      = 4
	DieselPumpCount   = 4
	LpgPumpCount      = 1
	ElectricPumpCount = 8
)

// Time to refuel
const (
	GasLowerBoundTime      = 1
	GasUpperBoundTime      = 5
	DieselLowerBoundTime   = 1
	DieselUpperBoundTime   = 5
	LpgLowerBoundTime      = 1
	LpgUpperBoundTime      = 5
	ElectricLowerBoundTime = 3
	ElectricUpperBoundTime = 10
)

const (
	GasPumpQueueSize      = 3
	DieselPumpQueueSize   = 3
	LpgPumpQueueSize      = 3
	ElectricPumpQueueSize = 1
)
