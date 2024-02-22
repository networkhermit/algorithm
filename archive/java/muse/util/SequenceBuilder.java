package muse.util;

public final class SequenceBuilder {

  private SequenceBuilder() {}

  public static void packIncreasing(Integer[] arr) {
    if (arr.length == 0) {
      return;
    }
    arr[0] = RandomFactory.genIntN(1, 3);
    for (int i = 1, length = arr.length; i < length; i++) {
      arr[i] = arr[i - 1] + RandomFactory.genIntN(1, 3);
    }
  }

  public static void packRandom(Integer[] arr) {
    for (int i = 0, length = arr.length; i < length; i++) {
      arr[i] = RandomFactory.genInt();
    }
  }

  public static void packDecreasing(Integer[] arr) {
    packIncreasing(arr);
    Sequences.reverse(arr);
  }
}
