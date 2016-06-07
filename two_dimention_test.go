package arithmetic

import "testing"

func TestSquare(t *testing.T) {
	var calculate2D Calculate2D

	calculate2D = Square{10.0}

	if calculate2D.Area() != 100.0 {
		t.Error("Squre Area should be 100.0")
	}
	if calculate2D.Circumference() != 40.0 {
		t.Error("Squre Circumference should be 40.0")
	}
}

func TestCircle(t *testing.T) {
	var calculate2D Calculate2D

	calculate2D = Circle{14}

	if int(calculate2D.Area()) != 153 {
		t.Error("Circle Area should be 153.86")
	}

	if int(calculate2D.Circumference()) != 43 {
		t.Error("Circle Circumference should be 43.96")
	}

	if calculate2D.(Circle).Radius() != 7.0 {
		t.Error("Circle Radius should be 7.0")
	}
}
