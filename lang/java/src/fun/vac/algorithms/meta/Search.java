package fun.vac.algorithms.meta;

import fun.vac.algorithms.search.BinarySearch;
import fun.vac.algorithms.search.LinearSearch;

public final class Search {

    private Search() {}

    public static <T extends Comparable<T>> int binarySearch(T[] arr, T key) {
        return BinarySearch.find(arr, key);
    }

    public static <T extends Comparable<T>> int linearSearch(T[] arr, T key) {
        return LinearSearch.find(arr, key);
    }
}
