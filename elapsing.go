package elapsing

import (
	"strconv"
	"sync"
	"time"

	"github.com/samber/lo"

	"github.com/nekomeowww/elapsing/internal/utils"
)

const (
	defaultPointNamePrefix     = "Step "
	defaultUnknownName         = "(unknown name)"
	defaultUnknownFunctionName = "(unknown function name)"
)

type elapsingType int

const (
	_                elapsingType = iota
	elapsingTypeBase              // Elapsing of base
	elapsingTypeFunc              // Elapsing for a function
)

type Elapsing struct {
	name         string
	steps        steps
	elapsingType elapsingType

	on            time.Time
	lastStepOn    time.Time
	lastStepIndex int64

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
	e := empty()
	e.name = utils.FunctionNameOfCaller(2)
	e.name = lo.Ternary(e.name == "", defaultUnknownName, e.name)
	return e
}

// empty creates a new empty elapsing, the only difference with New is that
// the name of the elapsing is empty
func empty() *Elapsing {
	return &Elapsing{
		on:           time.Now(),
		elapsingType: elapsingTypeBase,
		steps:        make(steps, 0, 1024),
	}
}

func (e *Elapsing) newStepEnds(name string, on time.Time) {
	e.stepsLock.Lock()
	defer e.stepsLock.Unlock()

	point := point{
		name:         name,
		on:           on,
		sinceInitial: on.Sub(e.on),
	}
	if len(e.steps) > 0 {
		point.sinceLast = on.Sub(e.lastStepOn)
	} else {
		point.sinceLast = point.sinceInitial
	}

	e.lastStepOn = on
	e.lastStepIndex += 1
	e.steps = append(e.steps, point)
}

// StepEnds adds a new step to the elapsing.
func (e *Elapsing) StepEnds(callOpts ...StepCallOptions) {
	opts := applyOptions(callOpts)
	if opts.name == "" {
		opts.name = defaultPointNamePrefix + strconv.FormatInt(e.lastStepIndex, 10)
	}
	if opts.on.IsZero() {
		opts.on = time.Now()
	}

	e.newStepEnds(opts.name, opts.on)
}

// TotalElapsed returns the total elapsed time since the elapsing is created
func (e *Elapsing) TotalElapsed() time.Duration {
	return time.Since(e.on)
}
