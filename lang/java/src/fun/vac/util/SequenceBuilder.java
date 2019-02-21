package fun.vac.util;

import fun.vac.util.RandomFactory;
import fun.vac.util.Sequences;

public class SequenceBuilder {

    private SequenceBuilder() {}

    public static void packIncreasing(Integer[] arr) {
        RandomFactory.launch();
        arr[0] = RandomFactory.integerN(1, 4);
        for (int i = 1, length = arr.length; i < length; i++) {
            arr[i] = arr[i - 1] + RandomFactory.integerN(1, 4);
        }
    }

    public static void packRandom(Integer[] arr) {
        RandomFactory.launch();
        for (int i = 0, length = arr.length; i < length; i++) {
            arr[i] = RandomFactory.generateInteger();
        }
    }

    public static void packDecreasing(Integer[] arr) {
        packIncreasing(arr);
        Sequences.reverse(arr);
    }
}
