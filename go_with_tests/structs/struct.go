package structs

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return 0
}

func (r Rectangle) Perimeter() float64 {
	return 0
}

func (c Circle) Area() float64 {
	return 0
}

func (c Circle) Perimeter() float64 {
	return 0
}
