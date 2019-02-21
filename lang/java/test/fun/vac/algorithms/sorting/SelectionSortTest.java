import fun.vac.algorithms.sorting.SelectionSort;
import fun.vac.util.SequenceBuilder;
import fun.vac.util.Sequences;
import fun.vac.util.TestRunner;

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

        if (!Sequences.isSorted(arr)) {
            return false;
        }

        return true;
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testSelectionSort());
    }
}
