package main

import (
	"fmt"
	"time"

	"github.com/nekomeowww/elapsing"
)

func func1(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("func1 step 1"))

	func2(elapsingFunc.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("func1 call step 2"))
}

func func2(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("func2 step 1"))

	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("func2 call step 2"))
}

func funcWithSimplifiedChineseNames(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	// Chinese
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("简体中文 函数3 步骤 1"))

	// Chinese longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("简体中文 函数3 调用步骤 2"))
}

func funcWithTraditionalChineseNames(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	// Chinese
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("繁體中文 函數3 步驟 1"))

	// Chinese longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("繁體中文 函數3 調用步驟 2"))
}

func funcWithJapaneseNames(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	// Japanese
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("日本語 関数3 ステップ 1"))

	// Japanese longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("日本語 関数3 呼び出しステップ 2"))
}

func funcWithKoreanNames(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	// Korean
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("한국어 함수3 단계 1"))

	// Korean longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("한국어 함수3 호출 단계 2"))
}

func funcWithCJKNamesAllTogether(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	// Simplified Chinese
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("简体中文 函数3 步骤 1"))

	// Simplified Chinese longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("简体中文 函数3 调用步骤 2"))

	// Traditional Chinese
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("繁體中文 函數3 步驟 1"))

	// Traditional Chinese longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("繁體中文 函數3 調用步驟 2"))

	// Japanese
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("日本語 関数3 ステップ 1"))

	// Japanese longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("日本語 関数3 呼び出しステップ 2"))

	// Korean
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("한국어 함수3 단계 1"))

	// Korean longer
	time.Sleep(50 * time.Millisecond)
	elapsingFunc.StepEnds(elapsing.WithName("한국어 함수3 호출 단계 2"))
}

func func3WithCJKNames(elapsingFunc *elapsing.FuncCall) {
	defer elapsingFunc.Return()

	funcWithSimplifiedChineseNames(elapsingFunc.ForFunc())
	funcWithTraditionalChineseNames(elapsingFunc.ForFunc())
	funcWithJapaneseNames(elapsingFunc.ForFunc())
	funcWithKoreanNames(elapsingFunc.ForFunc())
	funcWithCJKNamesAllTogether(elapsingFunc.ForFunc())
}

func main() {
	elapsing := elapsing.New()

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	func1(elapsing.ForFunc())

	time.Sleep(50 * time.Millisecond)
	elapsing.StepEnds()

	func3WithCJKNames(elapsing.ForFunc())

	fmt.Println(elapsing.Stats())
}
