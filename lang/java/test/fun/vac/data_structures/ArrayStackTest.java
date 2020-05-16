import fun.vac.data_structures.ArrayStack;
import fun.vac.util.TestRunner;

public class ArrayStackTest {

    public static boolean testArrayStack() {
        int size = 8192;

        ArrayStack<Integer> stack = new ArrayStack<>(0);

        for (int i = 1; i <= size; i++) {
            stack.push(i);
        }

        stack.shrink();

        if (stack.size() != size) {
            return false;
        }

        if (stack.capacity() != size) {
            return false;
        }

        for (int i = size; i > 0; i--) {
            if (stack.peek() != i) {
                return false;
            }
            stack.pop();
        }

        stack.shrink();

        if (!stack.isEmpty()) {
            return false;
        }

        return stack.capacity() == 0;
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testArrayStack());
    }
}
