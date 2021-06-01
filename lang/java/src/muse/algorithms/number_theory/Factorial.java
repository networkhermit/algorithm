package muse.algorithms.number_theory;

public final class Factorial {

    private Factorial() {}

    public static long iterativeProcedure(long n) {
        if (n < 0) {
            return 0;
        }

        long result = 1;
        for (long i = 1; i <= n; i++) {
            result *= i;
        }

        return result;
    }

    public static long recursiveProcedure(long n) {
        if (n < 0) {
            return 0;
        }

        if (n == 0) {
            return 1;
        }
        return recursiveProcedure(n - 1) * n;
    }
}
