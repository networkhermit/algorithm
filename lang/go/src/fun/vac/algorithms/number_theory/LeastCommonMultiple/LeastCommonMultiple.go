package LeastCommonMultiple

import "fun/vac/algorithms/number_theory/GreatestCommonDivisor"

func ReduceToBinaryGCD(m int64, n int64) int64 {
    if m < 0 {
        m = -m
    }
    if n < 0 {
        n = -n
    }

    if m == 0 || n == 0 {
        return 0
    }
    return m / GreatestCommonDivisor.IterativeBinaryGCD(m, n) * n
}

func ReduceToEuclidean(m int64, n int64) int64 {
    if m < 0 {
        m = -m
    }
    if n < 0 {
        n = -n
    }

    if m == 0 || n == 0 {
        return 0
    }
    return m / GreatestCommonDivisor.IterativeEuclidean(m, n) * n
}
