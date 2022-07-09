import muse.data_structures.LinkedList;
import muse.util.TestRunner;

public class LinkedListTest {

  public static boolean testLinkedList() {
    int size = 8192;

    LinkedList<Integer> list = new LinkedList<>();

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

    for (int i = 0, j = size - 1; i < j; i++, j--) {
      int x = list.get(i);
      int y = list.get(j);

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
    TestRunner.pick(LinkedListTest::testLinkedList);
  }
}
