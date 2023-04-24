package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MULTIPLIER = 1
)

const (
	GAS      = "gas"
	DIESEL   = "diesel"
	LPG      = "lpg"
	ELECTRIC = "electric"
)

type VehicleFuel struct {
	name       string
	lowerBound time.Time
	upperBound time.Time
}

type Vehicle struct {
	ID       int
	fuel     VehicleFuel
	joinedAt time.Time
	leftAt   time.Time
}

func (v *Vehicle) getDuration() time.Duration {
	var delta time.Duration

	if !v.joinedAt.IsZero() && !v.leftAt.IsZero() {
		delta = v.leftAt.Sub(v.joinedAt)
	}

	return delta
}

type WaitQueue struct {
	Min       float64
	Max       float64
	m         sync.Mutex
	QueueSize int32
	waiting   int32
}

func (queue *WaitQueue) getWaitTime() float64 {
	r := queue.Min + rand.Float64()*(queue.Max-queue.Min)*MULTIPLIER
	return math.Round(r*100) / 100
}

func (queue *WaitQueue) join(stats chan Vehicle, wg *sync.WaitGroup, c *sync.Cond, v Vehicle) {
	defer func() {
		stats <- v
		wg.Done()
	}()

	fmt.Printf("Vehicle %d arrives\n", v.ID)

	waiting := atomic.LoadInt32(&queue.waiting)
	fmt.Printf("Waiting: %d\n", waiting)
	if waiting >= queue.QueueSize {
		fmt.Printf("Vehicle %d leaves, pump is full\n", v.ID)
		return
	}

	v.joinedAt = time.Now()
	fmt.Printf("Vehicle %d joins queue\n", v.ID)

	atomic.AddInt32(&queue.waiting, 1)
	c.L.Lock()
	defer func() {
		atomic.AddInt32(&queue.waiting, -1)
		v.leftAt = time.Now()
		fmt.Printf("Vehicle %d finished in %fs\n", v.ID, v.getDuration().Seconds())
	}()

	for atomic.LoadInt32(&queue.waiting) != 0 {
		c.Wait()
	}

	dur := time.Duration(queue.getWaitTime()) * time.Second
	fmt.Printf("Vehicle %d starts pumping for %fs\n", v.ID, dur.Seconds())
	time.Sleep(dur)
	c.L.Unlock()
}

func getVehicle(id int, fuel) Vehicle {
	return Vehicle{ID: id, fuel: fuel}
}

func main() {
	gas := WaitQueue{Min: 5, Max: 15, QueueSize: 5}
	c := sync.NewCond(&gas.m)

	pumpsWg := sync.WaitGroup{}
	limit := 10
	stats := make(chan Vehicle, limit)
	pumpsWg.Add(limit)
	for i := 0; i < limit; i++ {
		r := rand.Float64() * MULTIPLIER
		go gas.join(stats, &pumpsWg, c, getVehicle(i, GAS))
		time.Sleep(time.Duration(r) * time.Second)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for el := range stats {
			fmt.Println(fmt.Sprintf("Vehicle %v: %v", el.ID, el.leftAt.Sub(el.joinedAt).Seconds()))
		}
	}()

	pumpsWg.Wait()
	close(stats)

	fmt.Println("done")

}
