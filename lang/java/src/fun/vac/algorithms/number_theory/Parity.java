package fun.vac.algorithms.number_theory;

public final class Parity {

    private Parity() {}

    public static boolean moduloIsEven(long n) {
        return n % 2 == 0;
    }

    public static boolean moduloIsOdd(long n) {
        return n % 2 != 0;
    }

    public static boolean bitwiseIsEven(long n) {
        return (n & 1) == 0;
    }

    public static boolean bitwiseIsOdd(long n) {
        return (n & 1) != 0;
    }
}
