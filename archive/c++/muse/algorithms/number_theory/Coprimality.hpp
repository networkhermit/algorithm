#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_COPRIMALITY_HPP
#define MUSE_ALGORITHMS_NUMBER_THEORY_COPRIMALITY_HPP

#include <muse/algorithms/number_theory/GreatestCommonDivisor.hpp>

namespace Coprimality {

bool reduceToBinaryGCD(long m, long n) {
  return GreatestCommonDivisor::iterativeBinaryGCD(m, n) == 1;
}

bool reduceToEuclidean(long m, long n) {
  return GreatestCommonDivisor::iterativeEuclidean(m, n) == 1;
}
} // namespace Coprimality

#endif
