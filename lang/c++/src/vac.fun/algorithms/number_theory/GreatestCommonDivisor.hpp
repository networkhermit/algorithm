#ifndef VAC_FUN_ALGORITHMS_NUMBER_THEORY_GREATEST_COMMON_DIVISOR_HPP
#define VAC_FUN_ALGORITHMS_NUMBER_THEORY_GREATEST_COMMON_DIVISOR_HPP 1

namespace GreatestCommonDivisor {

    long iterativeBinaryGCD(long m, long n) {
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

    long recursiveBinaryGCD(long m, long n) {
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
                return recursiveBinaryGCD(m >> 1, n >> 1) << 1;
            }
            return recursiveBinaryGCD(m >> 1, n);
        }
        if ((n & 1) == 0) {
            return recursiveBinaryGCD(m, n >> 1);
        }
        if (m > n) {
            return recursiveBinaryGCD((m - n) >> 1, n);
        }
        return recursiveBinaryGCD(m, (n - m) >> 1);
    }

    long iterativeEuclidean(long m, long n) {
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

    long recursiveEuclidean(long m, long n) {
        if (m < 0) {
            m = -m;
        }
        if (n < 0) {
            n = -n;
        }

        if (n == 0) {
            return m;
        }
        return recursiveEuclidean(n, m % n);
    }
}

#endif
