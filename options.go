package elapsing

import "time"

type stepOptions struct {
	name string
	on   time.Time
}

var (
	defaultOptions = stepOptions{}
)

// StepCallOptions
type StepCallOptions struct {
	Name string
	On   time.Time
}

func applyOptions(callOptions []StepCallOptions) stepOptions {
	if len(callOptions) == 0 {
		return defaultOptions
	}

	opts := callOptions[0]
	optCopy := defaultOptions
	if opts.Name != "" {
		optCopy.name = opts.Name
	}
	if !opts.On.IsZero() {
		optCopy.on = opts.On
	}

	return optCopy
}

func WithName(name string) StepCallOptions {
	return StepCallOptions{
		Name: name,
	}
}

func WithTime(t time.Time) StepCallOptions {
	return StepCallOptions{
		On: t,
	}
}
