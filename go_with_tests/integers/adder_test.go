package integers

import (
	"testing"
	"fmt"
)

func ExampleAdd() {
	sum := Add(9, 10)
	fmt.Println(sum)
	// Output: 19
}

func TestAdd(t *testing.T) {
	t.Run("Testing standard add", func(t *testing.T) {
		sum := Add(4, 5)
		expected := 9
		assertCorrectSum(t, expected, sum)
	})
	t.Run("Adding two negative numbers", func(t *testing.T) {
		sum := Add(-4, -9)
		expected := -13
		assertCorrectSum(t, expected, sum)
	})
}

func assertCorrectSum(t testing.TB, expected, sum int) {
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}