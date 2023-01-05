package elapsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForFunc(t *testing.T) {
	assert := assert.New(t)

	e := Empty()
	fc := e.ForFunc()
	assert.Equal(ElapsingTypeFunc, fc.ElapsingType)
	assert.Equal(1, len(e.Steps))
	assert.Equal(fc.Elapsing, e.Steps[0])
}

func testObtainFunctionName(fc *FuncCall) {
	fc.obtainFunctionName()
}

func TestObtainFunctionName(t *testing.T) {
	assert := assert.New(t)
	fc := &FuncCall{
		Elapsing: Empty(),
	}

	testObtainFunctionName(fc)
	assert.Equal("elapsing.TestObtainFunctionName", fc.Name)
}

func TestStepEnds(t *testing.T) {
	assert := assert.New(t)

	e := Empty()
	fc := e.ForFunc()
	fc.StepEnds()
	assert.Equal(1, len(fc.Steps))
	assert.Equal(1, len(e.Steps))
	assert.Equal("elapsing.TestStepEnds", fc.Name)
}

func TestReturn(t *testing.T) {
	assert := assert.New(t)

	e := Empty()
	fc := e.ForFunc()
	fc.Return()
	assert.Empty(fc.Steps)
	assert.Equal("elapsing.TestReturn", fc.Name)
}
