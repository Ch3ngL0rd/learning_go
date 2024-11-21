package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("Testing standard repeat", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}