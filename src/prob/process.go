package prob

import (
	"context"
	libTime "time"
	"tradesim/src/time"
)

// Process represents a discrete-time stochastic process.
type Process struct {
	// Event receives a clock tick when the success event
	// of the probability distribution is satisfied.
	Event chan libTime.Time
	// distribution represents the probability distribution of the process.
	distribution Distribution
	// clock represents the discrete-time index set of the process.
	clock time.Clock
}

func NewProcess(distribution Distribution, clock time.Clock) Process {
	return Process{
		Event:        make(chan libTime.Time),
		distribution: distribution,
		clock:        clock,
	}
}

func (p Process) Start(ctx context.Context) error {
	go p.clock.Start()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-p.clock.Done:
			return nil
		case t := <-p.clock.Tick:
			if ok := p.distribution.Indicate(); ok {
				p.Event <- t
			}
		}
	}
}
