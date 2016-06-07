package arithmetic

type Calculate2D interface {
	Area() float64
	Circumference() float64
}

type Calculate3D interface {
	Volume() float64
}

type Calculate interface {
	Calculate2D
	Calculate3D
}
