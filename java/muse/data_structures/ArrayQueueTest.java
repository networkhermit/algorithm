import muse.data_structures.ArrayQueue;
import muse.util.TestRunner;

public class ArrayQueueTest {

    public static boolean testArrayQueue() {
        int size = 8192;

        ArrayQueue<Integer> queue = new ArrayQueue<>(0);

        for (int i = 1; i <= size; i++) {
            queue.enqueue(i);
        }

        queue.shrink();

        if (queue.size() != size) {
            return false;
        }

        if (queue.capacity() != size) {
            return false;
        }

        for (int i = 1; i <= size; i++) {
            if (queue.peek() != i) {
                return false;
            }
            queue.dequeue();
        }

        queue.shrink();

        if (!queue.isEmpty()) {
            return false;
        }

        return queue.capacity() == 0;
    }

    public static void main(String[] args) {
        TestRunner.pick(ArrayQueueTest::testArrayQueue);
    }
}
