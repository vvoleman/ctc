package station

import (
	"fmt"
	"simulator/src/utils"
	"simulator/src/vehicle"
	"sync"
	"time"
)

type CashRegisterManager struct {
	CashRegisters []*CashRegister
}

func (crm *CashRegisterManager) MoveToRegister(v *vehicle.Vehicle) {
	// find register with the shortest queue
	var shortestQueue *CashRegister
	for _, cr := range crm.CashRegisters {
		if shortestQueue == nil || len(shortestQueue.Queue) > len(cr.Queue) {
			shortestQueue = cr
		}
	}

	fmt.Printf("Vehicle %d joined register %d\n", v.ID, shortestQueue.ID)
	shortestQueue.JoinQueue(v)
}

func (crm *CashRegisterManager) Close() {
	for _, cr := range crm.CashRegisters {
		close(cr.Queue)
	}
}

func (crm *CashRegisterManager) StartRegisters() {
	for _, cr := range crm.CashRegisters {
		go cr.Start()
	}
}

type CashRegister struct {
	ID         int
	Queue      chan *vehicle.Vehicle
	wg         *sync.WaitGroup
	lowerBound float64
	upperBound float64
}

func (cr *CashRegister) JoinQueue(v *vehicle.Vehicle) {
	v.RegisterJoinedAt = time.Now()

	cr.Queue <- v
}

func (cr *CashRegister) GetDuration() time.Duration {
	return utils.GetRandomDuration(cr.lowerBound, cr.upperBound)
}

func (cr *CashRegister) Start() {
	for v := range cr.Queue {
		v.RegisterStartAt = time.Now()
		time.Sleep(cr.GetDuration())
		v.LeftAt = time.Now()
		diff := v.LeftAt.Sub(v.RegisterJoinedAt)
		fmt.Printf("Vehicle %d left cash register after %fs\n", v.ID, utils.DurationToVirtualSeconds(diff))
		cr.wg.Done()
	}
}
