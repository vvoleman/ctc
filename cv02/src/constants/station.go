package constants

// TotalNumberOfVehicles How many vehicles will be created
const TotalNumberOfVehicles = 500

const (
	GasChance      = 0.35
	DieselChance   = 0.35
	LpgChance      = 0.2
	ElectricChance = 0.1
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
	GasPumpQueueSize      = 1
	DieselPumpQueueSize   = 1
	LpgPumpQueueSize      = 1
	ElectricPumpQueueSize = 1
)
