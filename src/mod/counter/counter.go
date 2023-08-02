package counter

import "sync/atomic"

type Counter interface {
	Increase()
	Decrease()
	Count() uint64
	Clear()
}

type counter struct {
	n uint64
}

func (c *counter) Increase() {
	atomic.AddUint64(&c.n, 1)
}

func (c *counter) Decrease() {
	atomic.AddUint64(&c.n, ^uint64(0))
}

func (c *counter) Count() uint64 {
	return atomic.LoadUint64(&c.n)
}

func (c *counter) Clear() {
	atomic.StoreUint64(&c.n, 0)
}

func New() Counter {
	return &counter{}
}
