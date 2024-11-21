package array_and_slice

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}
		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("expected %d got %d, given %v", expected, sum, numbers)			
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("slices of slices", func (t *testing.T) {
		slices := [][]int{{1,2}, {3,4}}
		got := SumAll(slices)
		expected := []int{3,7}

		if got != expected {
			t.Errorf("expected %v got %v, given %v", expected, got, slices)
		}
	})
}