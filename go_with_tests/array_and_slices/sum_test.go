package array_and_slice

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1,2,3,4,5}

		sum := Sum(numbers)
		expected := 15
	
		if sum != expected {
			t.Errorf("expected %d got %d, given %v", expected, sum, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}
		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("expected %d got %d, given %v", expected, sum, numbers)			
		}
	})
}