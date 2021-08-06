package muse.util;

public final class RandomFactory {

    private RandomFactory() {}

    public static void seed() {
        // preserve for consistent interface
    }

    public static int genIntN(int min, int max) {
        return min + (int) (Math.random() * (max - min));
    }

    public static int genInt() {
        return genIntN(0, 2_147_483_647);
    }

    public static int genEven() {
        return genInt() >>> 1 << 1;
    }

    public static int genOdd() {
        return genInt() | 1;
    }
}
