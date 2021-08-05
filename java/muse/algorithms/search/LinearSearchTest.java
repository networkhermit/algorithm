import muse.algorithms.search.LinearSearch;
import muse.util.SequenceBuilder;
import muse.util.TestRunner;

public class LinearSearchTest {

    public static boolean testLinearSearch() {
        int size = 32768;

        Integer[] arr = new Integer[size];
        SequenceBuilder.packIncreasing(arr);

        if (LinearSearch.find(arr, -1) != size) {
            return false;
        }

        if (LinearSearch.find(arr, 2_147_483_647) != size) {
            return false;
        }

        for (int i = 0; i < size; i++) {
            if (LinearSearch.find(arr, arr[i]) != i) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {
        TestRunner.pick(LinearSearchTest::testLinearSearch);
    }
}
