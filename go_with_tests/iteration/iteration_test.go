package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("Testing standard repeat", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Testing variable amount of repeats", func(t *testing.T) {
		repeated := Repeat("abc",4)
		expected := "abcabcabcabc"

		if repeated != expected {
			t.Errorf("expected %q but got %q",expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}