package time

import (
	"math"
	"time"
)

// Clock represents a monotonic clock with a tick frequency and limit.
type Clock struct {
	// Tick exposes ticks received by the underlying ticker.
	Tick chan time.Time
	// Done exposes values received by the underlying done and stop channels.
	Done chan bool
	// ticker is the underlying monotonic clock.
	ticker *time.Ticker
	// done receives true when tick reaches limit or the clock is reset.
	done chan bool
	// stop receives true when the clock is stopped.
	stop chan bool
	// tick is a count of the number of ticks the clock has seen
	// in its current run.
	tick uint64
	// frequency represents the time interval between each clock tick.
	frequency time.Duration
	// limit represents the maximum value tick can reach,
	// which is math.MaxUint64 if limit is 0.
	limit uint64
}

// NewClock returns a clock initialized with the provided
// frequency and limit. The clock doesn't start running
// until it is explicitly started.
func NewClock(frequency time.Duration, limit uint64) Clock {
	if limit == 0 {
		limit = math.MaxUint64
	}
	return Clock{
		Tick:      make(chan time.Time),
		Done:      make(chan bool),
		ticker:    nil,
		done:      make(chan bool),
		stop:      make(chan bool),
		tick:      0,
		frequency: frequency,
		limit:     limit,
	}
}

// Start initializes and runs the clock's underlying time.Ticker.
func (c *Clock) Start() {
	if c.ticker == nil {
		// Create a new time.Ticker and start it.
		c.ticker = time.NewTicker(c.frequency)
	}

	for {
		select {
		case <-c.done:
			c.ticker.Stop()
			c.ticker = nil
			c.tick = 0
			c.Done <- true
			return
		case <-c.stop:
			c.Done <- true
			return
		case t := <-c.ticker.C:
			if c.tick < c.limit {
				c.tick++
				c.Tick <- t
			} else {
				c.Done <- true
			}
		}
	}
}

func (c *Clock) Stop() {
	c.stop <- true
}

func (c *Clock) Reset() {
	c.done <- true
}
