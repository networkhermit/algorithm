#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_PARITY_H
#define MUSE_ALGORITHMS_NUMBER_THEORY_PARITY_H

bool Parity_moduloIsEven(long n) { return n % 2 == 0; }

bool Parity_moduloIsOdd(long n) { return n % 2 != 0; }

bool Parity_bitwiseIsEven(long n) { return (n & 1) == 0; }

bool Parity_bitwiseIsOdd(long n) { return (n & 1) != 0; }

#endif
