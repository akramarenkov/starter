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

func TestReadyN(_ *testing.T) {
	testReadyN(5)
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
		starter.Ready()

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

func testReadyN(quantity int) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	starter := New()

	wg.Add(quantity)
	starter.ReadyN(quantity)

	starter.ReadyN(-2 * quantity)

	for range quantity {
		go func() {
			defer wg.Done()

			starter.Set()
		}()
	}

	starter.Go()
}
