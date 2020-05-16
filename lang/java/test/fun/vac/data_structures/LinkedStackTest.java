import fun.vac.data_structures.LinkedStack;
import fun.vac.util.TestRunner;

public class LinkedStackTest {

    public static boolean testLinkedStack() {
        int size = 8192;

        LinkedStack<Integer> stack = new LinkedStack<>();

        for (int i = 1; i <= size; i++) {
            stack.push(i);
        }

        if (stack.size() != size) {
            return false;
        }

        for (int i = size; i > 0; i--) {
            if (stack.peek() != i) {
                return false;
            }
            stack.pop();
        }

        return stack.isEmpty();
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testLinkedStack());
    }
}
