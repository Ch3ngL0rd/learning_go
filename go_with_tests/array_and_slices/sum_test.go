package array_and_slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("expected %d got %d, given %v", expected, sum, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("slices of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		expected := []int{3, 7}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v got %v", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Testing standard case", func (t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{9,9})
		want := []int{2, 9}
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	})
	t.Run("Testing empty case", func (t *testing.T) {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	})
}