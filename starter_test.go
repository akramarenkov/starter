package starter

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStarter(t *testing.T) {
	const (
		workersQuantity = 5
		baseWorkDelay   = 200 * time.Millisecond
		startDiffLimit  = 10 * time.Millisecond
	)

	wg := &sync.WaitGroup{}
	starter := New()
	starts := make(chan time.Duration, workersQuantity)

	beginAt := time.Now()

	for id := range workersQuantity {
		wg.Add(1)

		if id%2 == 0 {
			starter.Ready()
		} else {
			starter.ReadyN(1)
		}

		go func() {
			defer wg.Done()

			// Preparing for main work
			time.Sleep(time.Duration(id+1) * baseWorkDelay)

			starter.Set()

			starts <- time.Since(beginAt)

			// Main work
			time.Sleep(time.Second)
		}()
	}

	starter.Go()

	wg.Wait()

	close(starts)

	require.InDelta(
		t,
		time.Duration(workersQuantity)*baseWorkDelay,
		starter.StartedAt().Sub(beginAt),
		float64(startDiffLimit),
	)

	for started := range starts {
		require.InDelta(
			t,
			time.Duration(workersQuantity)*baseWorkDelay,
			started,
			float64(startDiffLimit),
		)
	}
}

func TestReadyNInvalid(t *testing.T) {
	starter := New()

	require.Panics(t, func() { starter.ReadyN(-1) })
	require.Panics(t, func() { starter.ReadyN(0) })
}
