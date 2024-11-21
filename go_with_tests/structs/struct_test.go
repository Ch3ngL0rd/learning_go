package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{1.00}, 2 * math.Pi},
	}
	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		assertCorrect(t, got, tt.want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12,6}, 72},
		{Circle{1.00}, math.Pi},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertCorrect(t, got, tt.want)
	}
}

func assertCorrect(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
