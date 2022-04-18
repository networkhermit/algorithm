package muse.algorithms.sorting;

public final class InsertionSort {

  private InsertionSort() {}

  public static <T extends Comparable<T>> void sort(T[] arr) {
    T target;

    int cursor;

    for (int i = 1, length = arr.length; i < length; i++) {
      target = arr[i];
      for (cursor = i; cursor > 0; cursor--) {
        if (arr[cursor - 1].compareTo(target) <= 0) {
          break;
        }
        arr[cursor] = arr[cursor - 1];
      }
      arr[cursor] = target;
    }
  }
}
