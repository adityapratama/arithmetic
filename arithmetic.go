package arithmetic

func Add(i ...int) int {
	var x int
	for _, v := range i {
		x += v
	}

	return x
}

func Sub(i ...int) int {
	var x int

	for k, v := range i {
		if k == 0 {
			x += v
		} else {
			x -= v
		}
	}

	return x
}

func Mult(i ...int) int {
	var x = 1
	for _, v := range i {
		x *= v
	}

	return x
}

func Div(i ...int) int {
	var x = 1
	for k, v := range i {
		if k == 0 {
			x *= v
		} else {
			x /= v
		}
	}

	return x
}
