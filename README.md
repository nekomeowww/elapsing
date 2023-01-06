# elapsing

[![Go Report](https://goreportcard.com/badge/github.com/nekomeowww/elapsing)](https://goreportcard.com/report/github.com/nekomeowww/elapsing)
[![Testing](https://github.com/nekomeowww/elapsing/actions/workflows/ci.yml/badge.svg)](https://github.com/nekomeowww/elapsing/actions/workflows/ci.yml)

A simple Golang library for measuring elapsed time in application, function calls, and even for goroutine!

---

## Features

- Easy to use
- Pretty output with colors into terminal
- Able to measure elapsed time for individual steps
- Able to measure elapsed time for function calls
- Able to measure elapsed time for goroutines
- Concurrent safe

## Install

```shell
go get github.com/nekomeowww/elapsing
```

## Usage

It is just simple as:

```go
package main

import (
    "fmt"

    "github.com/nekomeowww/elapsing"
)

func main() {
    elapsing := elapsing.New()

    time.Sleep(100 * time.Millisecond)
    elapsing.StepEnds()
    fmt.Println(elapsing.Stats())
}
```

If you would like to measure the elapsed time of a function call, you can use `elapsing.ForFunc()` to get a new `FuncCall` instance to measure individually:

```go
package main

import (
    "fmt"

    "github.com/nekomeowww/elapsing"
)

func func1(elapsingFunc *elapsing.FuncCall) {
    defer elapsingFunc.Return() // remember to call this!

    time.Sleep(100 * time.Millisecond)
    elapsingFunc.StepEnds()
}

func main() {
    elapsing := elapsing.New()

    time.Sleep(100 * time.Millisecond)
    elapsing.StepEnds()

    func1(elapsing.ForFunc())

    fmt.Println(elapsing.Stats())
}
```

## Performance

Benchmark result as below.

### Benchmark on MacBook Pro (14-inches, M1 Pro, 2022)

```text
goos: darwin
goarch: arm64
pkg: github.com/nekomeowww/elapsing
BenchmarkStepEnds
BenchmarkStepEnds-10                     7178438        177.1 ns/op      180 B/op        2 allocs/op
BenchmarkStepEndsWithName
BenchmarkStepEndsWithName-10            11180500        122.8 ns/op      156 B/op        1 allocs/op
BenchmarkStepEndsWithTime
BenchmarkStepEndsWithTime-10             8067342        150.5 ns/op      169 B/op        2 allocs/op
BenchmarkForFunc
BenchmarkForFunc-10                       495000         4119 ns/op    16858 B/op        6 allocs/op
BenchmarkForFuncStepEnds
BenchmarkForFuncStepEnds-10              6566698        164.5 ns/op      168 B/op        2 allocs/op
BenchmarkForFuncStepEndsWithName
BenchmarkForFuncStepEndsWithName-10     10403928        113.7 ns/op      163 B/op        1 allocs/op
BenchmarkForFuncStepEndsWithTime
BenchmarkForFuncStepEndsWithTime-10      7167212        170.7 ns/op      180 B/op        2 allocs/op
```

## Example

Here's some example codes:

```go
func func1(elapsingFunc *FuncCall) {
    defer elapsingFunc.Return()

    time.Sleep(50 * time.Millisecond)
    elapsingFunc.StepEnds(WithName("Func1 step 1"))

    func2(elapsingFunc.ForFunc())

    time.Sleep(50 * time.Millisecond)
    elapsingFunc.StepEnds(WithName("Func1 step 2"))
}

func func2(elapsingFunc *FuncCall) {
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

    func1(elapsing.ForFunc())

    time.Sleep(50 * time.Millisecond)
    elapsing.StepEnds()

    require.NotPanics(func() {
        fmt.Println(elapsing.Stats())
    })
}
```

It will output the following result with colors:

```shell
── elapsing.TestStats
   ├─ #1 Step 0 [50.290209ms ( 50.290209ms total)]
   ├─ #2 elapsing.func1
   │  ├─ #1 Func1 step 1 [51.041583ms ( 51.041583ms total)]
   │  ├─ #2 elapsing.func2
   │  │  ├─ #1 Func2 step 1 [51.119917ms ( 51.119917ms total)]
   │  │  └─ #2 Func2 step 2 [51.095666ms (102.215583ms total)]
   │  └─ #3 Func1 step 2 [51.070458ms (204.339041ms total)]
   └─ #3 Step 2 [51.074333ms (305.816375ms total)]
```

Screenshot with colors:

![screenshot with colors](https://raw.githubusercontent.com/nekomeowww/elapsing/main/docs/screenshots-01.png)
