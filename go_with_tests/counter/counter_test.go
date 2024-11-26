package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("0 increments", func (t *testing.T) {
		counter := Counter{}

		if counter.Value() != 0 {
			t.Errorf("got %d, wanted %d", counter.Value(), 0)
		}
	})
	t.Run("3 increments", func (t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
	
		if counter.Value() != 3 {
			t.Errorf("got %d, wanted %d", counter.Value(), 3)
		}
	})
}