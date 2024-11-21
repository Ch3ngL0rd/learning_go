package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Test rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		assertCorrect(t, got, want)
	})
	t.Run("Test circle perimeter", func(t *testing.T) {
		circle := Circle{1.00}
		got := Perimeter(circle)
		want := 2 * math.Pi
		assertCorrect(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("Test rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		want := 100.0
		assertCorrect(t, got, want)
	})
	t.Run("Test circle area", func(t *testing.T) {
		circle := Circle{1.00}
		got := Area(circle)
		want := math.Pi
		assertCorrect(t, got, want)
	})
}

func assertCorrect(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
