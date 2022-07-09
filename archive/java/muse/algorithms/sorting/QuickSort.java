package muse.algorithms.sorting;

public final class QuickSort {

  private QuickSort() {}

  public static <T extends Comparable<T>> void partition(T[] arr, int lo, int hi) {
    if (lo == hi) {
      return;
    }

    T pivot = arr[lo];

    int left = lo;
    int right = hi - 1;

    while (left != right) {
      for (int i = right; i > left; i--) {
        if (arr[i].compareTo(pivot) < 0) {
          arr[left] = arr[i];
          arr[i] = pivot;
          break;
        }
        right--;
      }
      for (int i = left; i < right; i++) {
        if (arr[i].compareTo(pivot) > 0) {
          arr[right] = arr[i];
          arr[i] = pivot;
          break;
        }
        left++;
      }
    }

    partition(arr, lo, left);
    partition(arr, left + 1, hi);
  }

  public static <T extends Comparable<T>> void sort(T[] arr) {
    partition(arr, 0, arr.length);
  }
}
