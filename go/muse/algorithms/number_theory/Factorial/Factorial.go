package Factorial

func IterativeProcedure(n int64) int64 {
	if n < 0 {
		return 0
	}

	result := int64(1)
	for i := int64(1); i <= n; i++ {
		result *= i
	}

	return result
}

func RecursiveProcedure(n int64) int64 {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}
	return RecursiveProcedure(n-1) * n
}
