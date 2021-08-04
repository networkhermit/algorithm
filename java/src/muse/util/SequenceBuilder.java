package muse.util;

public final class SequenceBuilder {

    private SequenceBuilder() {}

    public static void packIncreasing(Integer[] arr) {
        RandomFactory.seed();
        arr[0] = RandomFactory.integerN(1, 3);
        for (int i = 1, length = arr.length; i < length; i++) {
            arr[i] = arr[i - 1] + RandomFactory.integerN(1, 3);
        }
    }

    public static void packRandom(Integer[] arr) {
        RandomFactory.seed();
        for (int i = 0, length = arr.length; i < length; i++) {
            arr[i] = RandomFactory.generateInteger();
        }
    }

    public static void packDecreasing(Integer[] arr) {
        packIncreasing(arr);
        Sequences.reverse(arr);
    }
}
