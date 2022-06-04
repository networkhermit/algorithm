package coprimality

import "muse/algorithms/number_theory/greatest_common_divisor"

func ReduceToBinaryGCD(m int64, n int64) bool {
	return greatest_common_divisor.IterativeBinaryGCD(m, n) == 1
}

func ReduceToEuclidean(m int64, n int64) bool {
	return greatest_common_divisor.IterativeEuclidean(m, n) == 1
}
