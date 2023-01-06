package elapsing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	assert := assert.New(t)

	on := time.Now()

	p := &point{
		name: "ABCD",
		on:   on,
	}

	assert.Equal("ABCD", p.name)
	assert.Equal(StepTypePoint, p.Type())
	assert.Equal(on, p.On())
}
