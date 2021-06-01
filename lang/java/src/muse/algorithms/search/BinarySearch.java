package muse.algorithms.search;

public final class BinarySearch {

    private BinarySearch() {}

    public static <T extends Comparable<T>> int find(T[] arr, T key) {
        int lo = 0;
        int hi = arr.length;

        int mid;

        while (lo < hi) {
            mid = lo + ((hi - lo) >>> 1);
            if (key.compareTo(arr[mid]) < 0) {
                hi = mid;
            } else if (key.compareTo(arr[mid]) > 0) {
                lo = mid + 1;
            } else {
                return mid;
            }
        }

        return arr.length;
    }
}
