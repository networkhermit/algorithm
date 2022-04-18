import muse.algorithms.sorting.InsertionSort;
import muse.util.SequenceBuilder;
import muse.util.Sequences;
import muse.util.TestRunner;

public class InsertionSortTest {

  public static boolean testInsertionSort() {
    int size = 32768;

    Integer[] arr = new Integer[size];
    SequenceBuilder.packRandom(arr);

    int checksum = Sequences.parityChecksum(arr);

    InsertionSort.sort(arr);

    if (Sequences.parityChecksum(arr) != checksum) {
      return false;
    }

    return Sequences.isSorted(arr);
  }

  public static void main(String[] args) {
    TestRunner.pick(InsertionSortTest::testInsertionSort);
  }
}
