package elapsing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithName(t *testing.T) {
	assert := assert.New(t)

	callOpts := WithName("ABCD")
	appliedOpts := applyOptions([]StepCallOptions{callOpts})
	assert.Equal("ABCD", appliedOpts.name)
}

func TestWithTime(t *testing.T) {
	assert := assert.New(t)

	on := time.Now()
	callOpts := WithTime(on)
	appliedOpts := applyOptions([]StepCallOptions{callOpts})
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

		opts := applyOptions(make([]StepCallOptions, 0))
		assert.Empty(opts.name)
		assert.True(opts.on.IsZero())
	})
}
