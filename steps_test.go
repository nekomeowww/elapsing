package elapsing

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndexes(t *testing.T) {
	assert := assert.New(t)

	emptyStep := Steps{}
	indexes, maxLength := emptyStep.Indexes()
	assert.Empty(indexes)
	assert.Zero(maxLength)

	step := Steps{
		Point{},
		Point{},
		&Elapsing{},
	}

	indexes, maxLength = step.Indexes()
	assert.Len(indexes, 3)
	assert.Equal(1, maxLength)
	assert.ElementsMatch([]string{"1", "2", "3"}, indexes)

	expectedIndexes := make([]string, 100)
	aLotOfSteps := Steps{}
	for i := 0; i < 100; i++ {
		aLotOfSteps = append(aLotOfSteps, &Point{})
		expectedIndexes[i] = strconv.FormatInt(int64(i+1), 10)
	}

	indexes, maxLength = aLotOfSteps.Indexes()
	assert.Len(indexes, 100)
	assert.Equal(3, maxLength)
	assert.ElementsMatch(expectedIndexes, indexes)
}

func TestNames(t *testing.T) {
	assert := assert.New(t)

	emptyStep := Steps{}
	names, maxLength := emptyStep.Names()
	assert.Empty(names)
	assert.Zero(maxLength)

	step := Steps{
		Point{Name: "A"},
		Point{Name: "B"},
		&Elapsing{Name: "C"},
	}

	names, maxLength = step.Names()
	assert.Len(names, 3)
	assert.Equal(1, maxLength)
	assert.ElementsMatch([]string{"A", "B", ""}, names)
}

func TestLasts(t *testing.T) {
	assert := assert.New(t)

	emptyStep := Steps{}
	lasts, maxLength := emptyStep.Lasts()
	assert.Empty(lasts)
	assert.Zero(maxLength)

	on := time.Now()

	last1 := time.Now().Add(time.Second).Sub(on)
	last2 := time.Now().Add(2 * time.Second).Sub(on)

	step := Steps{
		Point{Name: "A", SinceLast: last1},
		Point{Name: "B", SinceLast: last2},
		&Elapsing{Name: "C"},
	}

	lasts, maxLength = step.Lasts()
	assert.Len(lasts, 3)
	assert.NotZero(maxLength)
	assert.ElementsMatch([]string{last1.String(), last2.String(), "0s"}, lasts)
}

func TestTotals(t *testing.T) {
	assert := assert.New(t)

	emptyStep := Steps{}
	totals, maxLength := emptyStep.Totals()
	assert.Empty(totals)
	assert.Zero(maxLength)

	on := time.Now()

	total1 := time.Now().Add(time.Second).Sub(on)
	total2 := time.Now().Add(2 * time.Second).Sub(on)

	step := Steps{
		Point{Name: "A", SinceInitial: total1},
		Point{Name: "B", SinceInitial: total2},
		&Elapsing{Name: "C"},
	}

	totals, maxLength = step.Totals()
	assert.Len(totals, 3)
	assert.NotZero(maxLength)
	assert.ElementsMatch([]string{total1.String(), total2.String(), "0s"}, totals)
}
