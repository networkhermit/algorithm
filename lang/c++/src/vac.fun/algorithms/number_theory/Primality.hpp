#ifndef VAC_FUN_ALGORITHMS_NUMBER_THEORY_PRIMALITY_HPP
#define VAC_FUN_ALGORITHMS_NUMBER_THEORY_PRIMALITY_HPP 1

#include <cmath>

namespace Primality {

    bool isPrime(long n) {
        if (n < 2) {
            return false;
        }
        if ((n & 1) == 0 && n != 2) {
            return false;
        }

        for (int i = 3, bound = static_cast<int>(std::sqrt(static_cast<double>(n))) + 1; i < bound; i += 2) {
            if (n % i == 0) {
                return false;
            }
        }

        return true;
    }

    bool isComposite(long n) {
        return n > 1 && !isPrime(n);
    }
}

#endif