package elapsing

import "time"

type stepOptions struct {
	name string
	on   time.Time
}

// StepCallOption
type StepCallOption struct {
	applyFunc func(opt *stepOptions)
}

func applyOptions(callOptions []StepCallOption) *stepOptions {
	if len(callOptions) == 0 {
		return &stepOptions{}
	}

	optCopy := &stepOptions{}
	for _, f := range callOptions {
		f.applyFunc(optCopy)
	}

	return optCopy
}

func WithName(name string) StepCallOption {
	return StepCallOption{
		applyFunc: func(opt *stepOptions) {
			opt.name = name
		},
	}
}

func WithTime(t time.Time) StepCallOption {
	return StepCallOption{
		applyFunc: func(opt *stepOptions) {
			opt.on = t
		},
	}
}
