package fun.vac.algorithms.number_theory;

import fun.vac.algorithms.number_theory.GreatestCommonDivisor;

public class LeastCommonMultiple {

    private LeastCommonMultiple() {}

    public static long reduceToBinaryGCD(long m, long n) {
        if (m < 0) {
            m = -m;
        }
        if (n < 0) {
            n = -n;
        }

        if (m == 0 || n == 0) {
            return 0;
        } else {
            return m / GreatestCommonDivisor.iterativeBinaryGCD(m, n) * n;
        }
    }

    public static long reduceToEuclidean(long m, long n) {
        if (m < 0) {
            m = -m;
        }
        if (n < 0) {
            n = -n;
        }

        if (m == 0 || n == 0) {
            return 0;
        } else {
            return m / GreatestCommonDivisor.iterativeEuclidean(m, n) * n;
        }
    }
}
