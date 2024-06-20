# Starter

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/starter.svg)](https://pkg.go.dev/github.com/akramarenkov/starter)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/starter)](https://goreportcard.com/report/github.com/akramarenkov/starter)
[![codecov](https://codecov.io/gh/akramarenkov/starter/branch/master/graph/badge.svg?token=wQDw9CowDp)](https://codecov.io/gh/akramarenkov/starter)

## Purpose

Library that allows you to control the start of work of multiple goroutines

## Usage

Example:

```go
package main

import (
    "sync"
    "time"

    "github.com/akramarenkov/starter"
)

func main() {
    const (
        quantity = 5
    )

    wg := &sync.WaitGroup{}
    defer wg.Wait()

    starter := starter.New()

    for range quantity {
        wg.Add(1)
        starter.Ready()

        go func() {
            defer wg.Done()

            // Preparing for main work
            time.Sleep(time.Second)

            starter.Set()

            // Main work
            time.Sleep(time.Second)
        }()
    }

    starter.Go()

    // Output:
    //
}
```
