package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := newCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertValue(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantendCount := 1000
		counter := newCounter()

		var wg sync.WaitGroup
		wg.Add(wantendCount)

		for i := 0; i < wantendCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertValue(t, counter, wantendCount)
	})
}

func assertValue(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func newCounter() *Counter {
	return &Counter{}
}
