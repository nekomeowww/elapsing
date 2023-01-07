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

## Examples

### Example 1 [[Try it]](https://go.dev/play/p/gpZziPOabFI)

Example 1 contains the use of `elapsing.ForFunc()` to measure the elapsed time of a function call.

It will output the following result with colors:

![screenshot with colors](https://raw.githubusercontent.com/nekomeowww/elapsing/main/docs/screenshots-01.png)

[[Source Code: cmd/examples/example1]](https://github.com/nekomeowww/elapsing/tree/main/cmd/examples/example1)

### Example 2 [[Try it]](https://go.dev/play/p/Cbpxdnc-q-m)

Example 2 contains the use of CJK names and the output will auto padding the spaces to make the output looks better.

It will output the following result with colors:

![screenshot with colors](https://raw.githubusercontent.com/nekomeowww/elapsing/main/docs/screenshots-02.png)

[[Source Code: cmd/examples/example2]](https://github.com/nekomeowww/elapsing/tree/main/cmd/examples/example2)

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
