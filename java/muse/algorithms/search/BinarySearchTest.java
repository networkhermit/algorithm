import muse.algorithms.search.BinarySearch;
import muse.util.SequenceBuilder;
import muse.util.TestRunner;

public class BinarySearchTest {

    public static boolean testBinarySearch() {
        int size = 32768;

        Integer[] arr = new Integer[size];
        SequenceBuilder.packIncreasing(arr);

        if (BinarySearch.find(arr, -1) != size) {
            return false;
        }

        if (BinarySearch.find(arr, 2_147_483_647) != size) {
            return false;
        }

        for (int i = 0; i < size; i++) {
            if (BinarySearch.find(arr, arr[i]) != i) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {
        TestRunner.pick(BinarySearchTest::testBinarySearch);
    }
}
