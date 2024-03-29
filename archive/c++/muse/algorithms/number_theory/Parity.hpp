#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_PARITY_HPP
#define MUSE_ALGORITHMS_NUMBER_THEORY_PARITY_HPP

namespace Parity {

bool moduloIsEven(long n) { return n % 2 == 0; }

bool moduloIsOdd(long n) { return n % 2 != 0; }

bool bitwiseIsEven(long n) { return (n & 1) == 0; }

bool bitwiseIsOdd(long n) { return (n & 1) != 0; }
} // namespace Parity

#endif
