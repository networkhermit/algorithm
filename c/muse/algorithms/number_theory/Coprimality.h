#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_COPRIMALITY_H
#define MUSE_ALGORITHMS_NUMBER_THEORY_COPRIMALITY_H 1

#include <stdbool.h>

#include <muse/algorithms/number_theory/GreatestCommonDivisor.h>

bool Coprimality_reduceToBinaryGCD(long m, long n) {
    return GreatestCommonDivisor_iterativeBinaryGCD(m, n) == 1;
}

bool Coprimality_reduceToEuclidean(long m, long n) {
    return GreatestCommonDivisor_iterativeEuclidean(m, n) == 1;
}

#endif
