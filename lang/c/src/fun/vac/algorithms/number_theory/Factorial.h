#ifndef FUN_VAC_ALGORITHMS_NUMBER_THEORY_FACTORIAL_H
#define FUN_VAC_ALGORITHMS_NUMBER_THEORY_FACTORIAL_H 1

long Factorial_iterativeProcedure(long n) {
    if (n < 0) {
        return 0;
    }

    long result = 1;
    for (long i = 1; i <= n; i++) {
        result *= i;
    }

    return result;
}

long Factorial_recursiveProcedure(long n) {
    if (n < 0) {
        return 0;
    }

    if (n == 0) {
        return 1;
    }
    return Factorial_recursiveProcedure(n - 1) * n;
}

#endif
