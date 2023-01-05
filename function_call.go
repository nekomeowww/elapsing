package elapsing

import (
	"github.com/nekomeowww/elapsing/pkg/utils"
	"github.com/samber/lo"
)

type FuncCall struct {
	*Elapsing
}

func (e *Elapsing) ForFunc() *FuncCall {
	e.stepsLock.Lock()
	defer e.stepsLock.Unlock()

	fc := &FuncCall{
		Elapsing: Empty(),
	}

	fc.Elapsing.ElapsingType = ElapsingTypeFunc
	e.Steps = append(e.Steps, fc.Elapsing)
	return fc
}

func (f *FuncCall) obtainFunctionName() {
	functionName := utils.FunctionNameOfCaller(3)
	functionName = lo.Ternary(functionName == "", "(unknown function name)", functionName)
	f.Name = functionName
}

func (f *FuncCall) StepEnds(callOpts ...StepCallOption) {
	if f.Name == "" {
		f.obtainFunctionName()
	}

	f.Elapsing.StepEnds(callOpts...)
}

func (f *FuncCall) Return() {
	if f.Name == "" {
		f.obtainFunctionName()
	}
}
