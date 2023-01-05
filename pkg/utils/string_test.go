package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsMaxLength(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(20, StringsMaxLength([]string{strings.Repeat("A", 10), strings.Repeat("A", 20)}))
	assert.Equal(1, StringsMaxLength([]string{"A", "B", "C"}))
	assert.Equal(1, StringsMaxLength([]string{"啊", "😯", "C"}))
}

func TestStringPadStart(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("      A", StringPadStart("A", 7))
	assert.Equal("     A😯", StringPadStart("A😯", 7))
}

func TestStringPadEnd(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("A      ", StringPadEnd("A", 7))
	assert.Equal("A😯     ", StringPadEnd("A😯", 7))
}
