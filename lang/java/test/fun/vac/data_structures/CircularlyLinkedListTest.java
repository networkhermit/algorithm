import fun.vac.data_structures.CircularlyLinkedList;
import fun.vac.util.TestRunner;

public class CircularlyLinkedListTest {

    public static boolean testCircularlyLinkedList() {
        int size = 8192;

        CircularlyLinkedList<Integer> list = new CircularlyLinkedList<>();

        for (int i = 1; i <= size; i++) {
            list.append(i);
        }

        if (list.size() != size) {
            return false;
        }

        for (int i = 0; i < size; i++) {
            if (list.get(i) != i + 1) {
                return false;
            }
        }

        for (int i = 0; i < size; i++) {
            list.set(i, size - i);
        }

        for (int i = 0; i < size; i++) {
            if (list.get(i) != size - i) {
                return false;
            }
        }

        int x;
        int y;

        for (int i = 0, j = size - 1; i < j; i++, j--) {
            x = list.get(i);
            y = list.get(j);

            list.remove(i);
            list.insert(i, y);
            list.remove(j);
            list.insert(j, x);
        }

        for (int i = 0; i < size; i++) {
            if (list.get(i) != i + 1) {
                return false;
            }
        }

        for (int i = size; i >= 1; i--) {
            if (list.back() != i) {
                return false;
            }
            list.eject();
        }

        return list.isEmpty();
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testCircularlyLinkedList());
    }
}
