package elapsing

import (
	"github.com/nekomeowww/elapsing/internal/utils"
	"github.com/samber/lo"
)

type FuncCall struct {
	*Elapsing
}

func (e *Elapsing) ForFunc() *FuncCall {
	e.stepsLock.Lock()
	defer e.stepsLock.Unlock()

	fc := &FuncCall{
		Elapsing: empty(),
	}

	fc.Elapsing.elapsingType = elapsingTypeFunc

	e.lastStepOn = fc.on
	e.lastStepIndex += 1
	e.steps = append(e.steps, fc.Elapsing)
	return fc
}

func (f *FuncCall) obtainFunctionName() {
	functionName := utils.FunctionNameOfCaller(3)
	functionName = lo.Ternary(functionName == "", "(unknown function name)", functionName)
	f.name = functionName
}

func (f *FuncCall) StepEnds(callOpts ...StepCallOptions) {
	if f.name == "" {
		f.obtainFunctionName()
	}

	f.Elapsing.StepEnds(callOpts...)
}

func (f *FuncCall) Return() {
	if f.name == "" {
		f.obtainFunctionName()
	}
}
