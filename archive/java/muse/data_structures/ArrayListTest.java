import muse.data_structures.ArrayList;
import muse.util.TestRunner;

public class ArrayListTest {

  public static boolean testArrayList() {
    int size = 8192;

    ArrayList<Integer> list = new ArrayList<>(0);

    for (int i = 1; i <= size; i++) {
      list.append(i);
    }

    list.shrink();

    if (list.size() != size) {
      return false;
    }

    if (list.capacity() != size) {
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

    list.shrink();

    if (!list.isEmpty()) {
      return false;
    }

    return list.capacity() == 0;
  }

  public static void main(String[] args) {
    TestRunner.pick(ArrayListTest::testArrayList);
  }
}
