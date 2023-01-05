package elapsing

import (
	"time"
)

type Point struct {
	Name         string
	SinceInitial time.Duration
	SinceLast    time.Duration

	on time.Time
}

func (p Point) Type() StepType {
	return StepTypePoint
}

func (p Point) On() time.Time {
	return p.on
}
