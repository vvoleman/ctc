package station

type Pump struct {
	ID        int
	Fuel      string
	QueueSize int32
	Queue
}
