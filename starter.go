// Library that provides to control the start of work of multiple goroutines.
package starter

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNegativeReadyValue = errors.New("negative ready value is specified")
	ErrZeroReadyValue     = errors.New("zero ready value is specified")
)

// Starts of work of multiple goroutines at the same time.
type Starter struct {
	startedAt time.Time
	trigger   chan struct{}
	wg        *sync.WaitGroup
}

// Creates Starter instance.
func New() *Starter {
	str := &Starter{
		trigger: make(chan struct{}),
		wg:      &sync.WaitGroup{},
	}

	return str
}

// Increases counter of controlled goroutines.
//
// It must be called in a control goroutine.
func (str *Starter) Ready() {
	str.wg.Add(1)
}

// Increases counter of controlled goroutines to specified value.
//
// It must be called in a control goroutine.
//
// Negative and zero values is not supported and causes panic.
func (str *Starter) ReadyN(value int) {
	if value < 0 {
		panic(ErrNegativeReadyValue)
	}

	if value == 0 {
		panic(ErrZeroReadyValue)
	}

	str.wg.Add(value)
}

// Marks the controlled goroutine as up for a start.
//
// It must be called in a controlled goroutine before doing the work.
//
// Goroutine execution will continue after calling the Go() method.
func (str *Starter) Set() {
	str.wg.Done()

	<-str.trigger
}

// Waits for the goroutines to be ready to start and continues their work.
//
// It must be called in a control goroutine.
func (str *Starter) Go() {
	str.wg.Wait()

	str.startedAt = time.Now()

	close(str.trigger)
}

// Returns time of goroutines start.
//
// Thread-unsafe in parallel with Go method.
func (str *Starter) StartedAt() time.Time {
	return str.startedAt
}
