package arithmetic

import "math"

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Circumference() float64 {
	return s.Side * 4
}

type Circle struct {
	Diameter float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius(), 2)
}

func (c Circle) Circumference() float64 {
	return math.Pi * c.Diameter
}

func (c Circle) Radius() float64 {
	return c.Diameter / 2
}
