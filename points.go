package elapsing

import (
	"time"
)

type point struct {
	name         string
	on           time.Time
	sinceInitial time.Duration
	sinceLast    time.Duration
}

func (p point) Type() StepType {
	return StepTypePoint
}

func (p point) On() time.Time {
	return p.on
}
