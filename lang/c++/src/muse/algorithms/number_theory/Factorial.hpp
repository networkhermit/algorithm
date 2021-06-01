#ifndef MUSE_ALGORITHMS_NUMBER_THEORY_FACTORIAL_HPP
#define MUSE_ALGORITHMS_NUMBER_THEORY_FACTORIAL_HPP 1

namespace Factorial {

    long iterativeProcedure(long n) {
        if (n < 0) {
            return 0;
        }

        long result = 1;
        for (long i = 1; i <= n; i++) {
            result *= i;
        }

        return result;
    }

    long recursiveProcedure(long n) {
        if (n < 0) {
            return 0;
        }

        if (n == 0) {
            return 1;
        }
        return recursiveProcedure(n - 1) * n;
    }
}

#endif
