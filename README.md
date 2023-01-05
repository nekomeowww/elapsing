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

## Example

Here's some example codes:

```go
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
```

It will output the following result with colors:

```shell
── elapsing.TestStats
   ├─ #1 Step 1 [ 51.164875ms (51.164875ms total)]
   ├─ #2 elapsing.Func1
   │  ├─ #1 Func1 step 1 [   51.1435ms ( 51.1435ms total)]
   │  ├─ #2 elapsing.Func2
   │  │  ├─ #1 Func2 step 1 [51.094625ms ( 51.094625ms total)]
   │  │  └─ #2 Func2 step 2 [51.302833ms (102.397458ms total)]
   │  └─ #3 Func1 step 2 [152.680417ms (203.8375ms total)]
   └─ #3 Step 3 [254.866125ms (  306.036ms total)]
```

Screenshot with colors:

![](https://raw.githubusercontent.com/nekomeowww/elapsing/main/docs/screenshots-01.png)
