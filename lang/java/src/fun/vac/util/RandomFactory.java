package fun.vac.util;

public class RandomFactory {

    private RandomFactory() {}

    public static void launch() {
        // preserve for consistent interface
    }

    public static int integerN(int min, int max) {
        return min + (int) (Math.random() * (max - min));
    }

    public static int generateInteger() {
        return integerN(0, 2_147_483_647);
    }

    public static int generateEven() {
        return generateInteger() >>> 1 << 1;
    }

    public static int generateOdd() {
        return generateInteger() | 1;
    }
}
