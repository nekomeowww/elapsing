package elapsing

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndexes(t *testing.T) {
	assert := assert.New(t)

	emptyStep := steps{}
	indexes, maxLength := emptyStep.Indexes()
	assert.Empty(indexes)
	assert.Zero(maxLength)

	step := steps{
		point{},
		point{},
		&Elapsing{},
	}

	indexes, maxLength = step.Indexes()
	assert.Len(indexes, 3)
	assert.NotZero(maxLength)
	assert.ElementsMatch([]string{"1", "2", "3"}, indexes)

	expectedIndexes := make([]string, 100)
	aLotOfSteps := steps{}
	for i := 0; i < 100; i++ {
		aLotOfSteps = append(aLotOfSteps, &point{})
		expectedIndexes[i] = strconv.FormatInt(int64(i+1), 10)
	}

	indexes, maxLength = aLotOfSteps.Indexes()
	assert.Len(indexes, 100)
	assert.NotZero(maxLength)
	assert.ElementsMatch(expectedIndexes, indexes)
}

func TestNames(t *testing.T) {
	assert := assert.New(t)

	emptyStep := steps{}
	names, maxLength := emptyStep.Names()
	assert.Empty(names)
	assert.Zero(maxLength)

	step := steps{
		point{name: "A"},
		point{name: "B"},
		&Elapsing{name: "C"},
	}

	names, maxLength = step.Names()
	assert.Len(names, 3)
	assert.NotZero(maxLength)
	assert.ElementsMatch([]string{"A", "B", ""}, names)
}

func TestLasts(t *testing.T) {
	assert := assert.New(t)

	emptyStep := steps{}
	lasts, maxLength := emptyStep.Lasts()
	assert.Empty(lasts)
	assert.Zero(maxLength)

	on := time.Now()

	last1 := time.Now().Add(time.Second).Sub(on)
	last2 := time.Now().Add(2 * time.Second).Sub(on)

	step := steps{
		point{name: "A", sinceLast: last1},
		point{name: "B", sinceLast: last2},
		&Elapsing{name: "C"},
	}

	lasts, maxLength = step.Lasts()
	assert.Len(lasts, 3)
	assert.NotZero(maxLength)
	assert.ElementsMatch([]string{last1.String(), last2.String(), "0s"}, lasts)
}

func TestTotals(t *testing.T) {
	assert := assert.New(t)

	emptyStep := steps{}
	totals, maxLength := emptyStep.Totals()
	assert.Empty(totals)
	assert.Zero(maxLength)

	on := time.Now()

	total1 := time.Now().Add(time.Second).Sub(on)
	total2 := time.Now().Add(2 * time.Second).Sub(on)

	step := steps{
		point{name: "A", sinceInitial: total1},
		point{name: "B", sinceInitial: total2},
		&Elapsing{name: "C"},
	}

	totals, maxLength = step.Totals()
	assert.Len(totals, 3)
	assert.NotZero(maxLength)
	assert.ElementsMatch([]string{total1.String(), total2.String(), "0s"}, totals)
}
