package elapsing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	assert := assert.New(t)

	on := time.Now()

	p := &Point{
		Name: "ABCD",
		on:   on,
	}

	assert.Equal("ABCD", p.Name)
	assert.Equal(StepTypePoint, p.Type())
	assert.Equal(on, p.On())
}
