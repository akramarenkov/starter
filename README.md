# Starter

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/starter.svg)](https://pkg.go.dev/github.com/akramarenkov/starter)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/starter)](https://goreportcard.com/report/github.com/akramarenkov/starter)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/starter/badge.svg)](https://coveralls.io/github/akramarenkov/starter)

## Purpose

Library that provides to control the start of work of multiple goroutines

## Usage

Example:

```go
package main

import (
    "fmt"
    "sync"
    "time"

    "github.com/akramarenkov/starter"
)

func main() {
    const workersQuantity = 5

    wg := &sync.WaitGroup{}
    actuator := starter.New()

    wg.Add(workersQuantity)
    actuator.ReadyN(workersQuantity)

    for range workersQuantity {
        go func() {
            defer wg.Done()

            // Preparing for main work
            time.Sleep(time.Second)

            actuator.Set()

            // Main work
            time.Sleep(time.Second)
        }()
    }

    actuator.Go()

    wg.Wait()
    // Output:
}
```
