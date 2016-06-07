package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, 2, 3, 4)

	if res != 10 {
		t.Error("res should be 10")
	}
}

func TestSub(t *testing.T) {
	res := Sub(10, 5, 3)

	if res != 2 {
		t.Error("res should be 2")
	}
}

func TestMult(t *testing.T) {
	res := Mult(2, 3, 4)

	if res != 24 {
		t.Error("res should be 24")
	}
}

func TestDiv(t *testing.T) {
	res := Div(10, 5, 2)

	if res != 1 {
		t.Error("res should be 1")
	}
}
