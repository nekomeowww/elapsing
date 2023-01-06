package elapsing

import (
	"testing"
	"time"
)

func BenchmarkStepEnds(b *testing.B) {
	elapsing := empty()
	for i := 0; i < b.N; i++ {
		elapsing.StepEnds()
	}
}

func BenchmarkStepEndsWithName(b *testing.B) {
	elapsing := empty()
	for i := 0; i < b.N; i++ {
		elapsing.StepEnds(WithName("ABCD"))
	}
}

func BenchmarkStepEndsWithTime(b *testing.B) {
	elapsing := empty()
	for i := 0; i < b.N; i++ {
		elapsing.StepEnds(WithTime(time.Now()))
	}
}

func benchmarkForFunc(fc *FuncCall) {
	defer fc.Return()
}

func BenchmarkForFunc(b *testing.B) {
	elapsing := empty()
	for i := 0; i < b.N; i++ {
		benchmarkForFunc(elapsing.ForFunc())
	}
}

func benchmarkForFuncStepEnds(fc *FuncCall) {
	defer fc.Return()
	fc.StepEnds()
}

func BenchmarkForFuncStepEnds(b *testing.B) {
	elapsing := empty()
	fc := elapsing.ForFunc()

	for i := 0; i < b.N; i++ {
		benchmarkForFuncStepEnds(fc)
	}
}

func benchmarkForFuncStepEndsWithName(fc *FuncCall) {
	defer fc.Return()
	fc.StepEnds(WithName("ABCD"))
}

func BenchmarkForFuncStepEndsWithName(b *testing.B) {
	elapsing := empty()
	fc := elapsing.ForFunc()

	for i := 0; i < b.N; i++ {
		benchmarkForFuncStepEndsWithName(fc)
	}
}

func benchmarkForFuncStepEndsWithTime(fc *FuncCall) {
	defer fc.Return()
	fc.StepEnds(WithTime(time.Now()))
}

func BenchmarkForFuncStepEndsWithTime(b *testing.B) {
	elapsing := empty()
	fc := elapsing.ForFunc()

	for i := 0; i < b.N; i++ {
		benchmarkForFuncStepEndsWithTime(fc)
	}
}
