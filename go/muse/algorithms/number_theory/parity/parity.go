package parity

func ModuloIsEven(n int64) bool {
	return n%2 == 0
}

func ModuloIsOdd(n int64) bool {
	return n%2 != 0
}

func BitwiseIsEven(n int64) bool {
	return (n & 1) == 0
}

func BitwiseIsOdd(n int64) bool {
	return (n & 1) != 0
}
