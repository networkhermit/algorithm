import fun.vac.data_structures.LinkedQueue;
import fun.vac.util.TestRunner;

public class LinkedQueueTest {

    public static boolean testLinkedQueue() {
        int size = 8192;

        LinkedQueue<Integer> queue = new LinkedQueue<>();

        for (int i = 1; i <= size; i++) {
            queue.enqueue(i);
        }

        if (queue.size() != size) {
            return false;
        }

        for (int i = 1; i <= size; i++) {
            if (queue.peek() != i) {
                return false;
            }
            queue.dequeue();
        }

        return queue.isEmpty();
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testLinkedQueue());
    }
}
