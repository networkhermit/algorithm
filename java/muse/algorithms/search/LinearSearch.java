package muse.algorithms.search;

public final class LinearSearch {

  private LinearSearch() {}

  public static <T extends Comparable<T>> int find(T[] arr, T key) {
    for (int i = 0, length = arr.length; i < length; i++) {
      if (arr[i].compareTo(key) == 0) {
        return i;
      }
    }

    return arr.length;
  }
}
