#ifndef FUN_VAC_ALGORITHMS_NUMBER_THEORY_LEAST_COMMON_MULTIPLE_H
#define FUN_VAC_ALGORITHMS_NUMBER_THEORY_LEAST_COMMON_MULTIPLE_H 1

#include <fun/vac/algorithms/number_theory/GreatestCommonDivisor.h>

long LeastCommonMultiple_reduceToBinaryGCD(long m, long n) {
    if (m < 0) {
        m = -m;
    }
    if (n < 0) {
        n = -n;
    }

    if (m == 0 || n == 0) {
        return 0;
    } else {
        return m / GreatestCommonDivisor_iterativeBinaryGCD(m, n) * n;
    }
}

long LeastCommonMultiple_reduceToEuclidean(long m, long n) {
    if (m < 0) {
        m = -m;
    }
    if (n < 0) {
        n = -n;
    }

    if (m == 0 || n == 0) {
        return 0;
    } else {
        return m / GreatestCommonDivisor_iterativeEuclidean(m, n) * n;
    }
}

#endif
