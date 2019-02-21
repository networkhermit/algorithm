package fun.vac.algorithms.meta;

import fun.vac.algorithms.number_theory.Coprimality;
import fun.vac.algorithms.number_theory.Factorial;
import fun.vac.algorithms.number_theory.FibonacciNumber;
import fun.vac.algorithms.number_theory.GreatestCommonDivisor;
import fun.vac.algorithms.number_theory.LeastCommonMultiple;
import fun.vac.algorithms.number_theory.Parity;
import fun.vac.algorithms.number_theory.Primality;
import fun.vac.algorithms.number_theory.PrimeSieves;

public class NumberTheory {

    private NumberTheory() {}

    public static boolean isCoprime(long m, long n) {
        return Coprimality.reduceToBinaryGCD(m, n);
    }

    public static long factorial(long n) {
        return Factorial.iterativeProcedure(n);
    }

    public static long fibonacci(long n) {
        return FibonacciNumber.iterativeProcedure(n);
    }

    public static long gcd(long m, long n) {
        return GreatestCommonDivisor.iterativeBinaryGCD(m, n);
    }

    public static long lcm(long m, long n) {
        return LeastCommonMultiple.reduceToBinaryGCD(m, n);
    }

    public static boolean isEven(long n) {
        return Parity.bitwiseIsEven(n);
    }

    public static boolean isOdd(long n) {
        return Parity.bitwiseIsOdd(n);
    }

    public static boolean isPrime(long n) {
        return Primality.isPrime(n);
    }

    public static boolean isComposite(long n) {
        return Primality.isComposite(n);
    }

    public static int[] sieveOfPrimes(int n) {
        return PrimeSieves.sieveOfEratosthenes(n);
    }
}
