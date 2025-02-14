package starter_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/akramarenkov/starter"
)

func ExampleStarter() {
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

func ExampleStarter_complicated() {
	const (
		workersQuantity = 5
		baseWorkDelay   = 200 * time.Millisecond
		startDiffLimit  = 10 * time.Millisecond
	)

	wg := &sync.WaitGroup{}
	actuator := starter.New()

	wg.Add(workersQuantity)
	actuator.ReadyN(workersQuantity)

	beginAt := time.Now()

	for id := range workersQuantity {
		go func() {
			defer wg.Done()

			// Preparing for main work
			time.Sleep(time.Duration(id+1) * baseWorkDelay)

			actuator.Set()

			startDiff := time.Since(beginAt) - workersQuantity*baseWorkDelay

			fmt.Println(startDiff < startDiffLimit && startDiff > -startDiffLimit)

			// Main work
			time.Sleep(time.Second)
		}()
	}

	actuator.Go()

	wg.Wait()

	startDiff := actuator.StartedAt().Sub(beginAt) - workersQuantity*baseWorkDelay

	fmt.Println(startDiff < startDiffLimit && startDiff > -startDiffLimit)
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
