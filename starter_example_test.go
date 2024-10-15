package starter_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/akramarenkov/starter"
)

func ExampleStarter() {
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
}

func ExampleStarter_complicated() {
	const (
		quantity  = 5
		delay     = 200 * time.Millisecond
		diffLimit = 10 * time.Millisecond
	)

	wg := &sync.WaitGroup{}

	starter := starter.New()

	beginAt := time.Now()

	for id := range quantity {
		wg.Add(1)
		starter.Ready()

		go func(id int) {
			defer wg.Done()

			// Preparing for main work
			time.Sleep(time.Duration(id+1) * delay)

			starter.Set()

			diff := time.Since(beginAt) - quantity*delay

			fmt.Println(diff < diffLimit && diff > -diffLimit)

			// Main work
			time.Sleep(time.Second)
		}(id)
	}

	starter.Go()

	wg.Wait()

	diff := starter.StartedAt().Sub(beginAt) - quantity*delay

	fmt.Println(diff < diffLimit && diff > -diffLimit)
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
