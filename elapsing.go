package elapsing

import (
	"fmt"
	"sync"
	"time"

	"github.com/samber/lo"

	"github.com/nekomeowww/elapsing/pkg/utils"
)

type ElapsingType int

const (
	_                ElapsingType = iota
	ElapsingTypeBase              // Elapsing of base
	ElapsingTypeFunc              // Elapsing for a function
)

type Elapsing struct {
	Name         string
	Steps        Steps
	ElapsingType ElapsingType

	on time.Time

	stepsLock sync.Mutex
}

func (e *Elapsing) Type() StepType {
	return StepTypeElapsing
}

func (e *Elapsing) On() time.Time {
	return e.on
}

// New creates a new elapsing with the name of the caller function
func New() *Elapsing {
	e := Empty()
	e.Name = utils.FunctionNameOfCaller(2)
	e.Name = lo.Ternary(e.Name == "", "(unknown name)", e.Name)
	return e
}

// Empty creates a new empty elapsing, the only difference with New is that
// the name of the elapsing is empty
func Empty() *Elapsing {
	return &Elapsing{
		on:           time.Now(),
		ElapsingType: ElapsingTypeBase,
	}
}

func (e *Elapsing) newStepEnds(name string, on time.Time) {
	e.stepsLock.Lock()
	defer e.stepsLock.Unlock()

	// create a new point to store the step name and time infos
	point := Point{
		Name:         name,
		on:           on,
		SinceInitial: on.Sub(e.on),
	}
	if len(e.Steps) > 0 {
		// calculate the time elapsed since the last step
		point.SinceLast = on.Sub(e.Steps[len(e.Steps)-1].On())
	} else {
		// calculate the time elapsed since the initial elapsing time
		point.SinceLast = on.Sub(e.on)
	}

	e.Steps = append(e.Steps, point)
}

// StepEnds adds a new step to the elapsing.
func (e *Elapsing) StepEnds(callOpts ...StepCallOption) {
	// apply the options
	opts := applyOptions(callOpts)

	// if the name is empty, create a default name
	if opts.name == "" {
		opts.name = fmt.Sprintf("Step %d", len(e.Steps)+1)
	}
	// if the time is zero, use the current time
	if opts.on.IsZero() {
		opts.on = time.Now()
	}

	e.newStepEnds(opts.name, opts.on)
}

// TotalElapsed returns the total elapsed time since the elapsing is created
func (e *Elapsing) TotalElapsed() time.Duration {
	return time.Since(e.on)
}
