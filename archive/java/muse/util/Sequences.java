package muse.util;

import muse.algorithms.sorting.QuickSort;

public final class Sequences {

  private Sequences() {}

  public static void inspect(Integer[] arr) {
    System.out.println("[");
    for (int i = 0, length = arr.length; i < length; i++) {
      System.out.printf("\t#%04X  -->  %d%n", i, arr[i]);
    }
    System.out.println("]");
  }

  public static boolean isSorted(Integer[] arr) {
    for (int i = 1, length = arr.length; i < length; i++) {
      if (arr[i] < arr[i - 1]) {
        return false;
      }
    }

    return true;
  }

  public static int parityChecksum(Integer[] arr) {
    int checksum = 0;

    for (Integer v : arr) {
      checksum ^= v;
    }

    return checksum;
  }

  public static void reverse(Integer[] arr) {
    int k;
    Integer temp;

    for (int i = 0, bound = arr.length >>> 1, length = arr.length; i < bound; i++) {
      k = length - i - 1;
      temp = arr[i];
      arr[i] = arr[k];
      arr[k] = temp;
    }
  }

  public static void shuffle(Integer[] arr) {
    int k;
    Integer temp;

    for (int i = 0, length = arr.length; i < length; i++) {
      k = RandomFactory.genIntN(i, length - 1);
      temp = arr[i];
      arr[i] = arr[k];
      arr[k] = temp;
    }
  }

  public static void sort(Integer[] arr) {
    QuickSort.sort(arr);
  }
}
