package Coprimality

import "muse/algorithms/number_theory/GreatestCommonDivisor"

func ReduceToBinaryGCD(m int64, n int64) bool {
    return GreatestCommonDivisor.IterativeBinaryGCD(m, n) == 1
}

func ReduceToEuclidean(m int64, n int64) bool {
    return GreatestCommonDivisor.IterativeEuclidean(m, n) == 1
}
