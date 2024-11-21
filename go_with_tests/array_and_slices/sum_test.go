package array_and_slice

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1,2,3,4,5}

	sum := Sum(numbers)
	expected := 15

	if sum != expected {
		t.Errorf("expected %d got %d, given %v", expected, sum, numbers)
	}
}