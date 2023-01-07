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

func func1(elapsingFunc *FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("func1 step 1"))

	func2(elapsingFunc.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("func1 call step 2"))
}

func func2(elapsingFunc *FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("func2 step 1"))

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(WithName("func2 call step 2"))
}

func TestStats(t *testing.T) {
	require := require.New(t)

	elapsing := New()

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	func1(elapsing.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	require.NotPanics(func() {
		fmt.Println(elapsing.Stats())
	})
}

func TestTotalElapsed(t *testing.T) {
	assert := assert.New(t)

	e := empty()
	assert.GreaterOrEqual(int64(time.Millisecond), int64(e.TotalElapsed()))
}
