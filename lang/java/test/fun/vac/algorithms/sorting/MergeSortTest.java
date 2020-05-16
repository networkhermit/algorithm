import fun.vac.algorithms.sorting.MergeSort;
import fun.vac.util.SequenceBuilder;
import fun.vac.util.Sequences;
import fun.vac.util.TestRunner;

public class MergeSortTest {

    public static boolean testMergeSort() {
        int size = 32768;

        Integer[] arr = new Integer[size];
        SequenceBuilder.packRandom(arr);

        int checksum = Sequences.parityChecksum(arr);

        MergeSort.sort(arr);

        if (Sequences.parityChecksum(arr) != checksum) {
            return false;
        }

        return Sequences.isSorted(arr);
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testMergeSort());
    }
}
