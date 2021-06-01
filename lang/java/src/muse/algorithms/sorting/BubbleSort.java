package muse.algorithms.sorting;

public final class BubbleSort {

    private BubbleSort() {}

    public static <T extends Comparable<T>> void sort(T[] arr) {
        T temp;

        int margin;
        int unsorted = arr.length;

        while (unsorted > 1) {
            margin = 0;
            for (int i = 1; i < unsorted; i++) {
                if (arr[i - 1].compareTo(arr[i]) > 0) {
                    temp = arr[i - 1];
                    arr[i - 1] = arr[i];
                    arr[i] = temp;
                    margin = i;
                }
            }
            unsorted = margin;
        }
    }
}
