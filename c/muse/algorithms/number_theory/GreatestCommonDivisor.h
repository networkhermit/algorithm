#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_GREATEST_COMMON_DIVISOR_H
#define MUSE_ALGORITHMS_NUMBER_THEORY_GREATEST_COMMON_DIVISOR_H

#include <stdbool.h>

long GreatestCommonDivisor_iterativeBinaryGCD(long m, long n) {
    if (m < 0) {
        m = -m;
    }
    if (n < 0) {
        n = -n;
    }

    int shift = 0;

    while (true) {
        if (m == n) {
            return m << shift;
        }
        if (m == 0) {
            return n << shift;
        }
        if (n == 0) {
            return m << shift;
        }

        if ((m & 1) == 0) {
            m >>= 1;
            if ((n & 1) == 0) {
                n >>= 1;
                shift++;
            }
        } else if ((n & 1) == 0) {
            n >>= 1;
        } else if (m > n) {
            m = (m - n) >> 1;
        } else {
            n = (n - m) >> 1;
        }
    }
}

long GreatestCommonDivisor_recursiveBinaryGCD(long m, long n) {
    if (m < 0) {
        m = -m;
    }
    if (n < 0) {
        n = -n;
    }

    if (m == n) {
        return m;
    }
    if (m == 0) {
        return n;
    }
    if (n == 0) {
        return m;
    }

    if ((m & 1) == 0) {
        if ((n & 1) == 0) {
            return GreatestCommonDivisor_recursiveBinaryGCD(m >> 1, n >> 1) << 1;
        }
        return GreatestCommonDivisor_recursiveBinaryGCD(m >> 1, n);
    }
    if ((n & 1) == 0) {
        return GreatestCommonDivisor_recursiveBinaryGCD(m, n >> 1);
    }
    if (m > n) {
        return GreatestCommonDivisor_recursiveBinaryGCD((m - n) >> 1, n);
    }
    return GreatestCommonDivisor_recursiveBinaryGCD(m, (n - m) >> 1);
}

long GreatestCommonDivisor_iterativeEuclidean(long m, long n) {
    if (m < 0) {
        m = -m;
    }
    if (n < 0) {
        n = -n;
    }

    long r;
    while (n != 0) {
        r = m % n;
        m = n;
        n = r;
    }

    return m;
}

long GreatestCommonDivisor_recursiveEuclidean(long m, long n) {
    if (m < 0) {
        m = -m;
    }
    if (n < 0) {
        n = -n;
    }

    if (n == 0) {
        return m;
    }
    return GreatestCommonDivisor_recursiveEuclidean(n, m % n);
}

#endif
