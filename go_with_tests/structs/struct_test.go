package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Test rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		assertCorrect(t, got, want)
	})
	t.Run("Test circle perimeter", func(t *testing.T) {
		circle := Circle{1.00}
		got := circle.Perimeter()
		want := 2 * math.Pi
		assertCorrect(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("Test rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		assertCorrect(t, got, want)
	})
	t.Run("Test circle area", func(t *testing.T) {
		circle := Circle{1.00}
		got := circle.Area()
		want := math.Pi
		assertCorrect(t, got, want)
	})
}

func assertCorrect(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
