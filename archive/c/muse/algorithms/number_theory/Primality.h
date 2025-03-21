#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_PRIMALITY_H
#define MUSE_ALGORITHMS_NUMBER_THEORY_PRIMALITY_H

#include <math.h>

bool Primality_isPrime(long n) {
  if (n < 2) {
    return false;
  }
  if ((n & 1) == 0 && n != 2) {
    return false;
  }

  for (int i = 3, bound = (int)sqrt((double)n) + 1; i < bound; i += 2) {
    if (n % i == 0) {
      return false;
    }
  }

  return true;
}

bool Primality_isComposite(long n) { return n > 1 && !Primality_isPrime(n); }

#endif
