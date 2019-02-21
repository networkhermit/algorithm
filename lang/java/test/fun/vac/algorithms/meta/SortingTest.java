import fun.vac.algorithms.meta.Sorting;
import fun.vac.util.SequenceBuilder;
import fun.vac.util.Sequences;
import fun.vac.util.TestRunner;

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

        if (!Sequences.isSorted(arr)) {
            return false;
        }

        return true;
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

        if (!Sequences.isSorted(arr)) {
            return false;
        }

        return true;
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

        if (!Sequences.isSorted(arr)) {
            return false;
        }

        return true;
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

        if (!Sequences.isSorted(arr)) {
            return false;
        }

        return true;
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

        if (!Sequences.isSorted(arr)) {
            return false;
        }

        return true;
    }

    public static void main(String[] args) {

        TestRunner.parseTest(testBubbleSort());

        TestRunner.parseTest(testInsertionSort());

        TestRunner.parseTest(testMergeSort());

        TestRunner.parseTest(testQuickSort());

        TestRunner.parseTest(testSelectionSort());
    }
}
