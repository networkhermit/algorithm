#ifndef FUN_VAC_ALGORITHMS_NUMBER_THEORY_COPRIMALITY_HPP
#define FUN_VAC_ALGORITHMS_NUMBER_THEORY_COPRIMALITY_HPP 1

#include <fun/vac/algorithms/number_theory/GreatestCommonDivisor.hpp>

namespace Coprimality {

    bool reduceToBinaryGCD(long m, long n) {
        return GreatestCommonDivisor::iterativeBinaryGCD(m, n) == 1;
    }

    bool reduceToEuclidean(long m, long n) {
        return GreatestCommonDivisor::iterativeEuclidean(m, n) == 1;
    }
}

#endif
