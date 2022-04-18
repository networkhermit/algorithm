package muse.algorithms.sorting;

public final class SelectionSort {

  private SelectionSort() {}

  public static <T extends Comparable<T>> void sort(T[] arr) {
    if (arr.length == 0) {
      return;
    }

    T temp;

    int iMin;

    for (int i = 0, bound = arr.length - 1; i < bound; i++) {
      iMin = i;
      for (int j = i + 1; j <= bound; j++) {
        if (arr[j].compareTo(arr[iMin]) < 0) {
          iMin = j;
        }
      }
      if (iMin != i) {
        temp = arr[i];
        arr[i] = arr[iMin];
        arr[iMin] = temp;
      }
    }
  }
}
