package elapsing

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Func1(elapsingFunc *FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("Func1 step 1"))

	Func2(elapsingFunc.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("Func1 step 2"))
}

func Func2(elapsingFunc *FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("Func2 step 1"))

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("Func2 step 2"))
}

func TestStats(t *testing.T) {
	require := require.New(t)

	elapsing := New()

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	Func1(elapsing.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	require.NotPanics(func() {
		fmt.Println(elapsing.Stats())
	})
}

func TestTotalElapsed(t *testing.T) {
	assert := assert.New(t)

	e := Empty()
	assert.GreaterOrEqual(1000*int64(time.Nanosecond), int64(e.TotalElapsed()))
}
