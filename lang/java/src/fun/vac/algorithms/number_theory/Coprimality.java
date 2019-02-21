package fun.vac.algorithms.number_theory;

import fun.vac.algorithms.number_theory.GreatestCommonDivisor;

public class Coprimality {

    private Coprimality() {}

    public static boolean reduceToBinaryGCD(long m, long n) {
        return GreatestCommonDivisor.iterativeBinaryGCD(m, n) == 1;
    }

    public static boolean reduceToEuclidean(long m, long n) {
        return GreatestCommonDivisor.iterativeEuclidean(m, n) == 1;
    }
}
