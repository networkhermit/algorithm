#ifndef FUN_VAC_ALGORITHMS_NUMBER_THEORY_PRIME_SIEVES_H
#define FUN_VAC_ALGORITHMS_NUMBER_THEORY_PRIME_SIEVES_H 1

#include <math.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

size_t *PrimeSieves_sieveOfEratosthenes(size_t n, size_t *length) {
    if (n < 2) {
        *length = 0;
        return (size_t *) malloc(0);
    }

    size_t size = (n + 1) >> 1;

    bool arr[size];
    memset(arr, false, size);

    size_t numPrimes = size;
    for (size_t i = 3, bound = (size_t) sqrt((double) n) + 1; i < bound; i += 2) {
        if (arr[i >> 1]) {
            continue;
        }
        for (size_t j = i * i; j <= n; j += i << 1) {
            if (!arr[j >> 1]) {
                arr[j >> 1] = true;
                numPrimes--;
            }
        }
    }

    *length = numPrimes;
    size_t *primes = (size_t *) malloc(*length * sizeof(size_t));

    primes[0] = 2;

    size_t curr = 1;
    for (size_t i = 3, bound = n + 1; i < bound; i += 2) {
        if (!arr[i >> 1]) {
            primes[curr] = i;
            curr++;
        }
    }

    return primes;
}

#endif
