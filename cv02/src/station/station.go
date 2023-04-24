package station

import (
	"fmt"
	"simulator/src/constants"
	"simulator/src/utils"
	"simulator/src/vehicle"
	"sync"
	"time"
)

type Station struct {
	Pumps               map[constants.FuelName][]*Pump
	CashRegisterManager *CashRegisterManager
	wg                  *sync.WaitGroup
}

func (s *Station) startPumps() {
	for _, pumps := range s.Pumps {
		for _, pump := range pumps {
			go pump.Start(s.CashRegisterManager)
		}
	}
}

func (s *Station) startVehicles(vehicles []*vehicle.Vehicle) {
	for _, v := range vehicles {
		s.wg.Add(1)
		go s.Join(v)
		waitDuration := utils.GetRandomDuration(constants.ArrivalLowerBoundTime, constants.ArrivalUpperBoundTime)
		time.Sleep(waitDuration)
	}
}

func (s *Station) Start(vehicles []*vehicle.Vehicle) {
	go s.CashRegisterManager.StartRegisters()
	go s.startPumps()
	s.startVehicles(vehicles)
	s.wg.Wait()
}

func (s *Station) Join(v *vehicle.Vehicle) {
	fuelType := v.Fuel.Name
	pumps := s.Pumps[fuelType]
	if len(pumps) == 0 {
		fmt.Printf("No pumps for fuel type %s\n", fuelType)
		s.wg.Done()
		return
	}

	for _, pump := range pumps {
		if pump.CanJoin() {
			pump.Join(v)
			return
		}
	}

	fmt.Printf("Vehicle %d leaves, no pumps for %s available\n", v.ID, fuelType)
	s.wg.Done()
}

func InitStation() *Station {
	pumps := InitPumps()

	wg := sync.WaitGroup{}

	registers := []*CashRegister{
		{ID: 1, lowerBound: constants.RegisterLowerBoundTime, upperBound: constants.RegisterLowerBoundTime, wg: &wg, Queue: make(chan *vehicle.Vehicle)},
		{ID: 2, lowerBound: constants.RegisterLowerBoundTime, upperBound: constants.RegisterLowerBoundTime, wg: &wg, Queue: make(chan *vehicle.Vehicle)},
	}

	crm := CashRegisterManager{
		CashRegisters: registers,
	}

	return &Station{
		Pumps:               pumps,
		CashRegisterManager: &crm,
		wg:                  &wg,
	}
}
