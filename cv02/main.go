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
	MULTIPLIER = 0.1
)

const (
	GAS      = "gas"
	DIESEL   = "diesel"
	LPG      = "lpg"
	ELECTRIC = "electric"
)

type Vehicle struct {
	ID       int
	fuel     string
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
	r := queue.Min + rand.Float64()*(queue.Max-queue.Min)
	return math.Round(r*100) / 100
}

func (queue *WaitQueue) join(stats chan Vehicle, c *sync.Cond, v Vehicle) {
	defer func() {
		stats <- v
	}()

	waiting := atomic.LoadInt32(&queue.waiting)
	fmt.Printf("Waiting: %d\n", waiting)
	if waiting >= queue.QueueSize {
		fmt.Printf("Vehicle %d leaves, pump is full\n", v.ID)
		return
	}

	v.joinedAt = time.Now()
	fmt.Printf("Vehicle %d joins queue\n", v.ID)

	c.L.Lock()
	atomic.AddInt32(&queue.waiting, 1)
	defer func() {
		atomic.AddInt32(&queue.waiting, -1)
		v.leftAt = time.Now()
		fmt.Printf("Vehicle %d finished in %fs\n", v.ID, v.getDuration().Seconds())
	}()

	for atomic.LoadInt32(&queue.waiting) > 1 {
		c.Wait()
	}

	time.Sleep(time.Duration(queue.getWaitTime()*MULTIPLIER) * time.Second)

	c.L.Unlock()
}

func getVehicle(id int, fuel string) Vehicle {
	return Vehicle{ID: id, fuel: fuel}
}

func main() {
	gas := WaitQueue{Min: 1, Max: 5, QueueSize: 5}
	c := sync.NewCond(&gas.m)
	stats := make(chan Vehicle)

	for i := 0; i < 100; i++ {
		r := rand.Float64() * MULTIPLIER
		fmt.Printf("Vehicle %d arriving at pump\n", i)
		go gas.join(stats, c, getVehicle(i, GAS))
		time.Sleep(time.Duration(r) * time.Second)
	}

	for {

	}

}
