package starter

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStarter(t *testing.T) {
	testStarter(t, 5, 200*time.Millisecond, 10*time.Millisecond)
}

func testStarter(
	t *testing.T,
	quantity int,
	delay time.Duration,
	diffLimit time.Duration,
) {
	wg := &sync.WaitGroup{}

	starter := New()

	starts := make(chan time.Duration, quantity)

	beginAt := time.Now()

	for id := range quantity {
		wg.Add(1)

		if quantity%2 == 0 {
			starter.Ready()
		} else {
			starter.ReadyN(1)
		}

		go func(id int) {
			defer wg.Done()

			// Preparing for main work
			time.Sleep(time.Duration(id+1) * delay)

			starter.Set()

			starts <- time.Since(beginAt)

			// Main work
			time.Sleep(time.Second)
		}(id)
	}

	starter.Go()

	wg.Wait()

	close(starts)

	require.InDelta(t,
		time.Duration(quantity)*delay,
		starter.StartedAt().Sub(beginAt),
		float64(diffLimit),
	)

	for started := range starts {
		require.InDelta(t,
			time.Duration(quantity)*delay,
			started,
			float64(diffLimit),
		)
	}
}

func TestInvalidReadyN(t *testing.T) {
	starter := New()

	require.Panics(t, func() { starter.ReadyN(-1) })
	require.Panics(t, func() { starter.ReadyN(0) })
}
