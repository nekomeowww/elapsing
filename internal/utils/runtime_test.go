package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFunctionNameOfCaller() string {
	return FunctionNameOfCaller(1)
}

func testFunctionNameOfCaller2() string {
	return FunctionNameOfCaller(2)
}

func TestFunctionNameOfCaller(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("utils.testFunctionNameOfCaller", testFunctionNameOfCaller())
	assert.Equal("utils.TestFunctionNameOfCaller", testFunctionNameOfCaller2())
	assert.Equal("utils.TestFunctionNameOfCaller", FunctionNameOfCaller(1))
}
