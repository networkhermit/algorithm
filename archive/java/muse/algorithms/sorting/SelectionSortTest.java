import muse.algorithms.sorting.SelectionSort;
import muse.util.SequenceBuilder;
import muse.util.Sequences;
import muse.util.TestRunner;

public class SelectionSortTest {

  public static boolean testSelectionSort() {
    int size = 32768;

    Integer[] arr = new Integer[size];
    SequenceBuilder.packRandom(arr);

    int checksum = Sequences.parityChecksum(arr);

    SelectionSort.sort(arr);

    if (Sequences.parityChecksum(arr) != checksum) {
      return false;
    }

    return Sequences.isSorted(arr);
  }

  public static void main(String[] args) {
    TestRunner.pick(SelectionSortTest::testSelectionSort);
  }
}
