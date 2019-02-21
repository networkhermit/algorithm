import fun.vac.algorithms.meta.Search;
import fun.vac.util.SequenceBuilder;
import fun.vac.util.TestRunner;

public class SearchTest {

    public static boolean testBinarySearch() {
        int size = 32768;

        Integer[] arr = new Integer[size];
        SequenceBuilder.packIncreasing(arr);

        if (Search.binarySearch(arr, -1) != size) {
            return false;
        }

        if (Search.binarySearch(arr, 2_147_483_647) != size) {
            return false;
        }

        for (int i = 0; i < size; i++) {
            if (Search.binarySearch(arr, arr[i]) != i) {
                return false;
            }
        }

        return true;
    }

    public static boolean testLinearSearch() {
        int size = 32768;

        Integer[] arr = new Integer[size];
        SequenceBuilder.packIncreasing(arr);

        if (Search.linearSearch(arr, -1) != size) {
            return false;
        }

        if (Search.linearSearch(arr, 2_147_483_647) != size) {
            return false;
        }

        for (int i = 0; i < size; i++) {
            if (Search.linearSearch(arr, arr[i]) != i) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {

        TestRunner.parseTest(testBinarySearch());

        TestRunner.parseTest(testLinearSearch());
    }
}
