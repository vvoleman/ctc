package station

import (
	"fmt"
	"simulator/src/constants"
	"simulator/src/utils"
	"simulator/src/vehicle"
	"sync"
	"time"
)

type Pump struct {
	ID        int
	m         sync.Mutex
	Fuel      constants.FuelName
	QueueSize int32
	Queue     chan *vehicle.Vehicle
}

func (p *Pump) CanJoin() bool {

	return len(p.Queue) < int(p.QueueSize)
}

func (p *Pump) Join(v *vehicle.Vehicle) bool {
	for {
		p.m.Lock()

		if !p.CanJoin() {
			p.m.Unlock()
			return false
		}

		select {
		case p.Queue <- v:
			v.QueuedAt = time.Now()
			p.m.Unlock()
			return true
		default:
			p.m.Unlock()
			return false
		}
	}
}

func (p *Pump) Start(crm *CashRegisterManager) {
	for v := range p.Queue {
		go func(v *vehicle.Vehicle) {
			v.StartFuelAt = time.Now()
			fmt.Printf("Vehicle %d started fueling %s\n", v.ID, p.Fuel)
			v.StartRefueling()
			v.EndFuelAt = time.Now()
			diff := v.EndFuelAt.Sub(v.StartFuelAt)
			fmt.Printf("Vehicle %d ended fueling after (%fs)\n", v.ID, utils.DurationToVirtualSeconds(diff))
			crm.MoveToRegister(v)
		}(v)
	}
}

func InitPumps() map[constants.FuelName][]*Pump {
	pumps := make(map[constants.FuelName][]*Pump)

	// Gas
	for i := 0; i < constants.GasPumpCount; i++ {
		pumps[constants.Gas] = append(pumps[constants.Gas], &Pump{
			ID:        i,
			QueueSize: constants.GasPumpQueueSize,
			Fuel:      constants.Gas,
			Queue:     make(chan *vehicle.Vehicle, constants.GasPumpQueueSize),
		})
	}

	// Diesel
	for i := 0; i < constants.DieselPumpCount; i++ {
		pumps[constants.Diesel] = append(pumps[constants.Diesel], &Pump{
			ID:        i,
			QueueSize: constants.DieselPumpQueueSize,
			Fuel:      constants.Diesel,
			Queue:     make(chan *vehicle.Vehicle, constants.DieselPumpQueueSize),
		})
	}

	// LPG
	for i := 0; i < constants.LpgPumpCount; i++ {
		pumps[constants.LPG] = append(pumps[constants.LPG], &Pump{
			ID:        i,
			QueueSize: constants.LpgPumpQueueSize,
			Fuel:      constants.LPG,
			Queue:     make(chan *vehicle.Vehicle, constants.LpgPumpQueueSize),
		})
	}

	// Electric
	for i := 0; i < constants.ElectricPumpCount; i++ {
		pumps[constants.Electric] = append(pumps[constants.Electric], &Pump{
			ID:        i,
			QueueSize: constants.ElectricPumpQueueSize,
			Fuel:      constants.Electric,
			Queue:     make(chan *vehicle.Vehicle, constants.ElectricPumpQueueSize),
		})
	}

	return pumps
}
