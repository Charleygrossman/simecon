package clock

import (
	"math"
	"time"
)

// Clock represents a monotonic clock with a tick frequency and limit.
type Clock struct {
	// Tick exposes ticks received by the underlying ticker.
	Tick chan time.Time
	// Done receives true when the underlying ticker stops.
	Done chan bool
	// ticker is the underlying monotonic clock.
	ticker *time.Ticker
	// stop receives true when the clock is stopped.
	stop chan bool
	// reset receives true when the clock is reset.
	reset chan bool
	// tick is a count of the number of ticks the clock has seen
	// in its current run.
	tick uint64
	// frequency represents the time interval between each clock tick.
	frequency time.Duration
	// limit represents the maximum value tick can reach.
	limit uint64
}

// NewClock returns a clock initialized with the provided
// frequency and limit. If the provided limit is 0,
// the clock limit will be set to math.MaxUint64.
//
// The clock doesn't start running until it is explicitly started.
func NewClock(frequency time.Duration, limit uint64) Clock {
	if limit == 0 {
		limit = math.MaxUint64
	}
	return Clock{
		Tick:      make(chan time.Time),
		Done:      make(chan bool),
		ticker:    nil,
		stop:      make(chan bool),
		reset:     make(chan bool),
		tick:      0,
		frequency: frequency,
		limit:     limit,
	}
}

// Start initializes and runs the clock's underlying ticker.
func (c *Clock) Start() {
	// Create a new ticker and start it.
	c.ticker = time.NewTicker(c.frequency)
	for {
		select {
		case <-c.stop:
			c.stopTicker()
			c.Done <- true
			return
		case <-c.reset:
			c.stopTicker()
			c.tick = 0
			c.Done <- true
			return
		case t := <-c.ticker.C:
			if c.tick >= c.limit {
				c.stopTicker()
				c.Done <- true
				return
			}
			c.tick++
			c.Tick <- t
		}
	}
}

// Stop stops the clock and leaves its progress
// toward its tick limit in its current state.
func (c *Clock) Stop() {
	c.stop <- true
}

// Reset stops the clock and sets its progress
// toward its tick limit back to the beginning.
func (c *Clock) Reset() {
	c.reset <- true
}

func (c *Clock) stopTicker() {
	c.ticker.Stop()
	c.ticker = nil
}
