package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
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
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{name: "circle", shape: Circle{Radius: 1.00}, hasArea: math.Pi},
		{name: "triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g expected area %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func assertCorrect(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
