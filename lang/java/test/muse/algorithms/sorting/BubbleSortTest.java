import muse.algorithms.sorting.BubbleSort;
import muse.util.SequenceBuilder;
import muse.util.Sequences;
import muse.util.TestRunner;

public class BubbleSortTest {

    public static boolean testBubbleSort() {
        int size = 32768;

        Integer[] arr = new Integer[size];
        SequenceBuilder.packRandom(arr);

        int checksum = Sequences.parityChecksum(arr);

        BubbleSort.sort(arr);

        if (Sequences.parityChecksum(arr) != checksum) {
            return false;
        }

        return Sequences.isSorted(arr);
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testBubbleSort());
    }
}
