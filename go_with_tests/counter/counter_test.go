package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("0 increments", func (t *testing.T) {
		counter := Counter{}

		assertValue(t, counter, 0)
	})
	t.Run("3 increments", func (t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
	
		assertValue(t, counter, 3)
	})
}

func assertValue(t *testing.T, c Counter, count int) {
	t.Helper()

	if c.Value() != count {
		t.Errorf("got %d, want %d",c.Value(), count)
	}
}