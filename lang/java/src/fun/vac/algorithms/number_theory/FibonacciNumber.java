package fun.vac.algorithms.number_theory;

public class FibonacciNumber {

    private FibonacciNumber() {}

    public static long iterativeProcedure(long n) {
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

    public static long recursiveProcedure(long n) {
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
