package muse.algorithms.sorting;

public final class MergeSort {

    private MergeSort() {}

    public static <T extends Comparable<T>> void merge(T[] arr, int lo, int mid, int hi) {
        if (lo == mid) {
            return;
        }

        merge(arr, lo, (lo + mid) >>> 1, mid);
        merge(arr, mid, (mid + hi) >>> 1, hi);

        int m = lo;
        int n = mid;

        T[] sorted = java.util.Arrays.copyOfRange(arr, lo, hi);

        for (int i = 0, length = sorted.length; i < length; i++) {
            if (m != mid && (n == hi || arr[m].compareTo(arr[n]) < 0)) {
                sorted[i] = arr[m];
                m++;
            } else {
                sorted[i] = arr[n];
                n++;
            }
        }

        int cursor = 0;
        for (int i = lo; i < hi; i++) {
            arr[i] = sorted[cursor];
            cursor++;
        }
    }

    public static <T extends Comparable<T>> void sort(T[] arr) {
        merge(arr, 0, arr.length >>> 1, arr.length);
    }
}
