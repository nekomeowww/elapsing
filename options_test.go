package elapsing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithName(t *testing.T) {
	assert := assert.New(t)

	callOpts := WithName("ABCD")

	emptyOpts := new(stepOptions)
	callOpts.applyFunc(emptyOpts)
	assert.Equal("ABCD", emptyOpts.name)

	appliedOpts := applyOptions([]StepCallOption{callOpts})
	assert.Equal("ABCD", appliedOpts.name)
}

func TestWithTime(t *testing.T) {
	assert := assert.New(t)

	on := time.Now()
	callOpts := WithTime(on)

	emptyOpts := new(stepOptions)
	callOpts.applyFunc(emptyOpts)
	assert.NotZero(emptyOpts.on)
	assert.Equal(on, emptyOpts.on)

	appliedOpts := applyOptions([]StepCallOption{callOpts})
	assert.NotZero(appliedOpts.on)
}

func TestApplyOptions(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		assert := assert.New(t)

		opts := applyOptions(nil)
		assert.Empty(opts.name)
		assert.True(opts.on.IsZero())
	})

	t.Run("Empty", func(t *testing.T) {
		assert := assert.New(t)

		opts := applyOptions(make([]StepCallOption, 0))
		assert.Empty(opts.name)
		assert.True(opts.on.IsZero())
	})
}
