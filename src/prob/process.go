package prob

import (
	"context"
	"time"
	"tradesim/src/time/clock"
)

// Process represents a discrete-time stochastic process.
type Process struct {
	// Event receives a clock tick when the success event
	// of the probability distribution is satisfied.
	Event chan time.Time
	// distribution represents the probability distribution of the process.
	distribution Distribution
	// clock represents the discrete-time index set of the process.
	clock clock.Clock
}

func NewProcess(distribution Distribution, clock clock.Clock) *Process {
	return &Process{
		Event:        make(chan time.Time, 8),
		distribution: distribution,
		clock:        clock,
	}
}

func (p *Process) Start(ctx context.Context) error {
	go p.clock.Start(ctx)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-p.clock.Done:
			return nil
		case t := <-p.clock.Tick:
			if ok := p.distribution.Indicate(); ok {
				p.Event <- t
			}
		}
	}
}
