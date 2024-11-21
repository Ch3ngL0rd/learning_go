package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertCorrect(t, got, want)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0
	assertCorrect(t, got, want)
}

func assertCorrect(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
