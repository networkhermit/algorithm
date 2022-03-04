#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_LEAST_COMMON_MULTIPLE_HPP
#define MUSE_ALGORITHMS_NUMBER_THEORY_LEAST_COMMON_MULTIPLE_HPP

#include <muse/algorithms/number_theory/GreatestCommonDivisor.hpp>

namespace LeastCommonMultiple {

    long reduceToBinaryGCD(long m, long n) {
        if (m < 0) {
            m = -m;
        }
        if (n < 0) {
            n = -n;
        }

        if (m == 0 || n == 0) {
            return 0;
        }
        return m / GreatestCommonDivisor::iterativeBinaryGCD(m, n) * n;
    }

    long reduceToEuclidean(long m, long n) {
        if (m < 0) {
            m = -m;
        }
        if (n < 0) {
            n = -n;
        }

        if (m == 0 || n == 0) {
            return 0;
        }
        return m / GreatestCommonDivisor::iterativeEuclidean(m, n) * n;
    }
}

#endif
