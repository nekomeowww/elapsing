package elapsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForFunc(t *testing.T) {
	assert := assert.New(t)

	e := empty()
	fc := e.ForFunc()
	assert.Equal(elapsingTypeFunc, fc.elapsingType)
	assert.Equal(1, len(e.steps))
	assert.Equal(fc.Elapsing, e.steps[0])
}

func testObtainFunctionName(fc *FuncCall) {
	fc.obtainFunctionName()
}

func TestObtainFunctionName(t *testing.T) {
	assert := assert.New(t)
	fc := &FuncCall{
		Elapsing: empty(),
	}

	testObtainFunctionName(fc)
	assert.Equal("elapsing.TestObtainFunctionName", fc.name)
}

func TestStepEnds(t *testing.T) {
	assert := assert.New(t)

	e := empty()
	fc := e.ForFunc()
	fc.StepEnds()
	assert.Equal(1, len(fc.steps))
	assert.Equal(1, len(e.steps))
	assert.Equal("elapsing.TestStepEnds", fc.name)
}

func TestReturn(t *testing.T) {
	assert := assert.New(t)

	e := empty()
	fc := e.ForFunc()
	fc.Return()
	assert.Empty(fc.steps)
	assert.Equal("elapsing.TestReturn", fc.name)
}
