package counter

import (
	"testing"
	"sync"
)

func TestCounter(t *testing.T) {
	t.Run("0 increments", func (t *testing.T) {
		counter := NewCounter()

		assertValue(t, counter, 0)
	})
	t.Run("3 increments", func (t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
	
		assertValue(t, counter, 3)
	})
	t.Run("Test concurrent increments", func(t *testing.T) {
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

		assertValue(t, counter, wantedCount)
	})
}

func assertValue(t *testing.T, c *Counter, count int) {
	t.Helper()

	if c.Value() != count {
		t.Errorf("got %d, want %d",c.Value(), count)
	}
}