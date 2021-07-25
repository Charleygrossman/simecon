package mkt

import (
	"math"
	"time"
)

// Clock represents a monotonic clock with a tick frequency and limit.
type Clock struct {
	// Ticker is the underlying monotonic clock.
	Ticker *time.Ticker
	// Frequency represents the time interval between each clock tick.
	Frequency time.Duration
	// Limit represents the maximum value Tick can reach,
	// which is math.MaxUint64 if Limit is nil.
	Limit *uint64
	// Tick is a count of the number of ticks the clock has seen
	// in its current run.
	Tick uint64
	// Done receives true when the clock has reached its limit
	// or is reset.
	Done chan bool
}

// NewClock returns a clock initialized with the provided
// frequency and limit. The clock doesn't start running
// until it is explicitly started.
func NewClock(frequency time.Duration, limit uint64) *Clock {
	l := limit
	return &Clock{
		Ticker:    nil,
		Frequency: frequency,
		Limit:     &l,
		Tick:      0,
		Done:      make(chan bool),
	}
}

// Start initializes and runs the clock's underlying time.Ticker.
func (c *Clock) Start() {
	// Create a new time.Ticker and start it.
	c.Ticker = time.NewTicker(c.Frequency)
	go func() {
		// Run the clock until Limit is reached or Tick
		// reaches system limits (maximum integer value).
		for (c.Limit == nil || c.Tick < *c.Limit) && c.Tick < math.MaxUint64 {
			time.Sleep(c.Frequency)
			c.Tick++
		}
		c.Done <- true
	}()
}

// Reset enables the next start of
// the clock to run its full duration.
func (c *Clock) Reset() {
	c.Ticker.Stop()
	c.Ticker = nil
	c.Tick = 0
	c.Done = make(chan bool)
}
