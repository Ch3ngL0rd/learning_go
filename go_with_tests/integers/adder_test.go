package main

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Testing standard add", func(t *testing.T) {
		sum := Add(4, 5)
		expected := 9
		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}
