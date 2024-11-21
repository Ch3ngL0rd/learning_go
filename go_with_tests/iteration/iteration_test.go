package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Testing standard repeat", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Testing variable amount of repeats", func(t *testing.T) {
		repeated := Repeat("abc", 4)
		expected := "abcabcabcabc"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("abc", 2)
	fmt.Println(repeated)
	// Output: abcabc
}

func BenchmarkRepeat(b *testing.B) {
	for repeatCount := 1; repeatCount < 100_000; repeatCount *= 10 {
		for i := 0; i < b.N; i++ {
			Repeat("a", repeatCount)
		}
	}
}
