package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup

		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(tb testing.TB, got *Counter, want int) {
	tb.Helper()

	if got.Value() != want {
		tb.Errorf("got %d, want %d", got.Value(), want)
	}
}
