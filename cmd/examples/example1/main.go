package main

import (
	"fmt"
	"time"

	"github.com/nekomeowww/elapsing"
)

func func1(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("Func1 step 1"))

	func2(elapsingFunc.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("Func1 step 2"))
}

func func2(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("Func2 step 1"))
}

func main() {
	elapsing := elapsing.New()

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	func1(elapsing.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	fmt.Println(elapsing.Stats())
}
