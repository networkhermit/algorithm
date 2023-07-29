package least_common_multiple

import "muse/algorithms/number_theory/greatest_common_divisor"

func ReduceToBinaryGCD(m, n int64) int64 {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	if m == 0 || n == 0 {
		return 0
	}
	return m / greatest_common_divisor.IterativeBinaryGCD(m, n) * n
}

func ReduceToEuclidean(m, n int64) int64 {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}

	if m == 0 || n == 0 {
		return 0
	}
	return m / greatest_common_divisor.IterativeEuclidean(m, n) * n
}
