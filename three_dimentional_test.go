package arithmetic

import "testing"

func TestCube(t *testing.T) {
	var calculate Calculate

	calculate = &Cube{10}

	if calculate.Area() != 600 {
		t.Error("Cube Area should be 600")
	}
	if calculate.Circumference() != 120 {
		t.Error("Cube Circumference should be 120")
	}
	if calculate.Volume() != 1000 {
		t.Error("Cube Volume should be 1000")
	}
}
