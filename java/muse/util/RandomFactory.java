package muse.util;

import java.security.SecureRandom;

public final class RandomFactory {

    private static SecureRandom random = new SecureRandom();

    private RandomFactory() {}

    public static int genIntN(int min, int max) {
        return min + random.nextInt(max - min + 1);
    }

    public static int genInt() {
        return genIntN(1, 2_147_483_647);
    }

    public static int genEven() {
        return genInt() >>> 1 << 1;
    }

    public static int genOdd() {
        return genInt() | 1;
    }
}
