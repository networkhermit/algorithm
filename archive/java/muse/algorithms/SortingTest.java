import muse.algorithms.Sorting;
import muse.util.SequenceBuilder;
import muse.util.Sequences;
import muse.util.TestRunner;

public class SortingTest {

  public static boolean testBubbleSort() {
    int size = 32768;

    Integer[] arr = new Integer[size];
    SequenceBuilder.packRandom(arr);

    int checksum = Sequences.parityChecksum(arr);

    Sorting.bubbleSort(arr);

    if (Sequences.parityChecksum(arr) != checksum) {
      return false;
    }

    return Sequences.isSorted(arr);
  }

  public static boolean testInsertionSort() {
    int size = 32768;

    Integer[] arr = new Integer[size];
    SequenceBuilder.packRandom(arr);

    int checksum = Sequences.parityChecksum(arr);

    Sorting.insertionSort(arr);

    if (Sequences.parityChecksum(arr) != checksum) {
      return false;
    }

    return Sequences.isSorted(arr);
  }

  public static boolean testMergeSort() {
    int size = 32768;

    Integer[] arr = new Integer[size];
    SequenceBuilder.packRandom(arr);

    int checksum = Sequences.parityChecksum(arr);

    Sorting.mergeSort(arr);

    if (Sequences.parityChecksum(arr) != checksum) {
      return false;
    }

    return Sequences.isSorted(arr);
  }

  public static boolean testQuickSort() {
    int size = 32768;

    Integer[] arr = new Integer[size];
    SequenceBuilder.packRandom(arr);

    int checksum = Sequences.parityChecksum(arr);

    Sorting.quickSort(arr);

    if (Sequences.parityChecksum(arr) != checksum) {
      return false;
    }

    return Sequences.isSorted(arr);
  }

  public static boolean testSelectionSort() {
    int size = 32768;

    Integer[] arr = new Integer[size];
    SequenceBuilder.packRandom(arr);

    int checksum = Sequences.parityChecksum(arr);

    Sorting.selectionSort(arr);

    if (Sequences.parityChecksum(arr) != checksum) {
      return false;
    }

    return Sequences.isSorted(arr);
  }

  public static void main(String[] args) {
    TestRunner.pick(SortingTest::testBubbleSort);

    TestRunner.pick(SortingTest::testInsertionSort);

    TestRunner.pick(SortingTest::testMergeSort);

    TestRunner.pick(SortingTest::testQuickSort);

    TestRunner.pick(SortingTest::testSelectionSort);
  }
}
