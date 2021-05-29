#ifndef VAC_FUN_ALGORITHMS_NUMBER_THEORY_FIBONACCI_NUMBER_HPP
#define VAC_FUN_ALGORITHMS_NUMBER_THEORY_FIBONACCI_NUMBER_HPP 1

namespace FibonacciNumber {

    long iterativeProcedure(long n) {
        long sign = 1;

        if (n < 0) {
            if ((n & 1) == 0) {
                sign = -1;
            }
            n = -n;
        }

        if (n < 2) {
            return n;
        }

        long prev = 0;
        long curr = 1;

        long next;
        for (long i = 2; i <= n; i++) {
            next = prev + curr;
            prev = curr;
            curr = next;
        }

        return sign * curr;
    }

    long recursiveProcedure(long n) {
        if (n < 0) {
            if ((n & 1) == 0) {
                return -recursiveProcedure(-n);
            }
            return recursiveProcedure(-n);
        }
        if (n < 2) {
            return n;
        }
        return recursiveProcedure(n - 2) + recursiveProcedure(n - 1);
    }
}

#endif
