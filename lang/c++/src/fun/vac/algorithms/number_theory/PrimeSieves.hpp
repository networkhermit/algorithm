#ifndef FUN_VAC_ALGORITHMS_NUMBER_THEORY_PRIME_SIEVES_HPP
#define FUN_VAC_ALGORITHMS_NUMBER_THEORY_PRIME_SIEVES_HPP 1

#include <cmath>
#include <cstddef>
#include <cstring>
#include <vector>

namespace PrimeSieves {

    std::vector<std::size_t> sieveOfEratosthenes(std::size_t n) {
        if (n < 2) {
            return std::vector<std::size_t>(0);
        }

        std::size_t size = (n + 1) >> 1;

        bool arr[size];
        std::memset(arr, false, size);

        std::size_t numPrimes = size;
        for (std::size_t i = 3, bound = static_cast<std::size_t>(std::sqrt(static_cast<double>(n))) + 1; i < bound; i += 2) {
            if (arr[i >> 1]) {
                continue;
            }
            for (std::size_t j = i * i; j <= n; j += i << 1) {
                if (!arr[j >> 1]) {
                    arr[j >> 1] = true;
                    numPrimes--;
                }
            }
        }

        std::vector<std::size_t> primes(numPrimes);

        primes[0] = 2;

        std::size_t curr = 1;
        for (std::size_t i = 3, bound = n + 1; i < bound; i += 2) {
            if (!arr[i >> 1]) {
                primes[curr] = i;
                curr++;
            }
        }

        return primes;
    }
}

#endif
